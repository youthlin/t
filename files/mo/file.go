package mo

import (
	"encoding/binary"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/youthlin/t/files"
)

// https://www.gnu.org/software/gettext/manual/html_node/MO-Files.html#MO-Files
/*
        byte
             +------------------------------------------+
          0  | magic number = 0x950412de                |
             |                                          |
          4  | file format revision = 0                 |
             |                                          |
          8  | number of strings                        |  == N
             |                                          |
         12  | offset of table with original strings    |  == O
             |                                          |
         16  | offset of table with translation strings |  == T
             |                                          |
         20  | size of hashing table                    |  == S
             |                                          |
         24  | offset of hashing table                  |  == H
             |                                          |
             .                                          .
             .    (possibly more entries later)         .
             .                                          .
             |                                          |
          O  | length & offset 0th string  ----------------.
      O + 8  | length & offset 1st string  ------------------.
              ...                                    ...   | |
O + ((N-1)*8)| length & offset (N-1)th string           |  | |
             |                                          |  | |
          T  | length & offset 0th translation  ---------------.
      T + 8  | length & offset 1st translation  -----------------.
              ...                                    ...   | | | |
T + ((N-1)*8)| length & offset (N-1)th translation      |  | | | |
             |                                          |  | | | |
          H  | start hash table                         |  | | | |
              ...                                    ...   | | | |
  H + S * 4  | end hash table                           |  | | | |
             |                                          |  | | | |
             | NUL terminated 0th string  <----------------' | | |
             |                                          |    | | |
             | NUL terminated 1st string  <------------------' | |
             |                                          |      | |
              ...                                    ...       | |
             |                                          |      | |
             | NUL terminated 0th translation  <---------------' |
             |                                          |        |
             | NUL terminated 1st translation  <-----------------'
             |                                          |
              ...                                    ...
             |                                          |
             +------------------------------------------+
*/

const (
	magicLittleEndian = 0x950412de
	magicBigEndian    = 0xde120495

	// https://www.asciihex.com/ascii-table

	nul = "\x00"
	eot = "\x04"

	errPrefix = "read mo file"
)

type header struct {
	Major      uint16 // 主版本号，只能是0或1
	Minor      uint16 // 次版本好，只能是0或1
	IDCount    uint32 // N msgID 数量
	OffsetID   uint32 // O msgID 从哪里开始读取
	OffsetStr  uint32 // T msgStr 的偏移量
	SizeHash   uint32 // S 可忽略 hash 表大小
	OffsetHash uint32 // H 可忽略 hash 表偏移位置
}
type lenOff struct {
	Length uint32 // 长度
	OffSet uint32 // 偏移位置
}

// Read read a mo file
// see also https://github.com/chai2010/gettext-go/blob/master/mo/file.go
func Read(r *os.File) (*files.File, error) {
	// 0 从头开始读取
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return nil, errors.Wrapf(err, "%s: failed to seek file", errPrefix)
	}
	// 1 读取魔数
	var magic uint32
	if err := binary.Read(r, binary.LittleEndian, &magic); err != nil {
		return nil, errors.Wrapf(err, "%s: failed to read magic number", errPrefix)
	}
	var bo binary.ByteOrder
	switch magic {
	case magicBigEndian:
		bo = binary.BigEndian
	case magicLittleEndian:
		bo = binary.LittleEndian
	default:
		return nil, errors.Errorf("%s: invalid magic number", errPrefix)
	}

	// 2 读取魔数后的固定头部字段
	var h header
	if err := binary.Read(r, bo, &h); err != nil {
		return nil, errors.Wrapf(err, "%s: failed to read mo header", errPrefix)
	}
	if h.Major != 0 && h.Major != 1 {
		return nil, errors.Errorf("%s: unsopported major version|%v", errPrefix, h.Major)
	}
	if h.Minor != 0 && h.Minor != 1 {
		return nil, errors.Errorf("%s: unsopported minor version|%v", errPrefix, h.Minor)
	}

	// 3 跳转到 O 处 读取 msgID 信息
	if _, err := r.Seek(int64(h.OffsetID), io.SeekStart); err != nil {
		// Seek 第二个参数：0=相对于文件开头，1=相对于当前，2=相对于文件尾
		return nil, errors.Wrapf(err, "%s: failed to seek msg_id", errPrefix)
	}
	var n = h.IDCount // 一共有 n 条 msgID
	var msgIDMeta []lenOff
	for i := uint32(0); i < n; i++ {
		var lo lenOff
		if err := binary.Read(r, bo, &lo); err != nil {
			return nil, errors.Wrapf(err, "%s: failed to read msg_id[%d] length & offset", errPrefix, i)
		}
		msgIDMeta = append(msgIDMeta, lo)
	}

	// 4 跳转到 T 处，读取 msgStr 信息
	if _, err := r.Seek(int64(h.OffsetStr), io.SeekStart); err != nil {
		return nil, errors.Wrapf(err, "%s: failed to seek msg_str", errPrefix)
	}
	var msgStrMeta []lenOff
	for i := uint32(0); i < n; i++ {
		var lo lenOff
		if err := binary.Read(r, bo, &lo); err != nil {
			return nil, errors.Wrapf(err, "%s: failed to read msg_str[%d] length & offset", errPrefix, i)
		}
		msgStrMeta = append(msgStrMeta, lo)
	}

	var result = files.NewEmptyFile()
	// 5 开始读取
	for i := uint32(0); i < n; i++ {
		// 5.1 跳转到第 i 条 msg_id 偏移处
		if _, err := r.Seek(int64(msgIDMeta[i].OffSet), io.SeekStart); err != nil {
			return nil, errors.Wrapf(err, "%s: failed to seek msg_id[%d]", errPrefix, i)
		}
		id := make([]byte, msgIDMeta[i].Length) // msg_id 的长度
		if err := binary.Read(r, bo, &id); err != nil {
			return nil, errors.Wrapf(err, "%s: failed to read msg_id[%d]", errPrefix, i)
		}
		// 5.2 跳转读取 msg_str
		if _, err := r.Seek(int64(msgStrMeta[i].OffSet), io.SeekStart); err != nil {
			return nil, errors.Wrapf(err, "%s: failed to seek msg_str[%d]", errPrefix, i)
		}
		str := make([]byte, msgStrMeta[i].Length) // msg_str 的长度
		if err := binary.Read(r, bo, &str); err != nil {
			return nil, errors.Wrapf(err, "%s: failed to read msg_str[%d]", errPrefix, i)
		}
		var msg = &files.Message{
			MsgID:  string(id),
			MsgStr: string(str),
		}
		// 0x04 分割 msgCtxt 和 msgId
		if index := strings.Index(msg.MsgID, eot); index >= 0 {
			msg.MsgCtxt, msg.MsgID = msg.MsgID[:index], msg.MsgID[index+1:]
		}
		// 0x00 分割 msgId 和 msgIdPlural
		if index := strings.Index(msg.MsgID, nul); index >= 0 {
			msg.MsgID, msg.MsgID2 = msg.MsgID[:index], msg.MsgID[index+1:]
			msg.MsgStrN = strings.Split(msg.MsgStr, nul)
			msg.MsgStr = ""
		}
		if err := result.AddMessage(msg); err != nil {
			return nil, errors.Wrapf(err, errPrefix)
		}
	}
	return result, nil
}
