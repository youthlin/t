package translator

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/cockroachdb/errors"
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
	nul        = "\x00"
	eot        = "\x04"
	moMagic    = 0x950412de
	moMagicBig = 0xde120495
	flag       = "ThisFileIsGenerateBy:github.com/youthlin/t" + nul
	flagLen    = len(flag)        // 43
	offsetO    = 28 + 4 + flagLen // 75: 28=??????header 4=uint32:flagLen 43=string:flag
)

var errReadMo = fmt.Errorf("read mo")

type header struct {
	Major      uint16 // ????????????????????????0???1
	Minor      uint16 // ????????????????????????0???1
	IDCount    uint32 // N msgID ??????
	OffsetID   uint32 // O msgID ?????????????????????
	OffsetStr  uint32 // T msgStr ????????????
	SizeHash   uint32 // S ????????? hash ?????????
	OffsetHash uint32 // H ????????? hash ???????????????
}

type lenOff struct {
	Length uint32 // ??????
	OffSet uint32 // ????????????
}

// ReadMo read mo from []byte content
func ReadMo(content []byte) (*File, error) {
	file := new(File)
	r := bytes.NewReader(content)
	// 1 ????????????
	var magic uint32
	if err := binary.Read(r, binary.LittleEndian, &magic); err != nil {
		return nil, errors.WithSecondaryError(errReadMo, errors.Wrapf(err, "failed to read magic number"))
	}
	var bo binary.ByteOrder
	switch magic {
	case moMagicBig:
		bo = binary.BigEndian
	case moMagic:
		bo = binary.LittleEndian
	default:
		return nil, errors.WithSecondaryError(errReadMo, errors.Errorf("invalid magic number: %v", magic))
	}

	// 2 ????????????????????????????????????
	var h header
	if err := binary.Read(r, bo, &h); err != nil {
		return nil, errors.WithSecondaryError(errReadMo, errors.Wrapf(err, "failed to read mo header"))
	}
	if h.Major != 0 && h.Major != 1 {
		return nil, errors.WithSecondaryError(errReadMo, errors.Errorf("unsopported major version: %v", h.Major))
	}
	if h.Minor != 0 && h.Minor != 1 {
		return nil, errors.WithSecondaryError(errReadMo, errors.Errorf("unsopported minor version: %v", h.Minor))
	}

	// 3 ????????? O ??? ?????? msgID ??????
	if _, err := r.Seek(int64(h.OffsetID), io.SeekStart); err != nil {
		return nil, errors.WithSecondaryError(errReadMo,
			errors.Wrapf(err, "failed to seek to O(message id table): %d", h.OffsetID))
	}
	var n = h.IDCount // ????????? n ??? msgID
	var msgIDMeta []lenOff
	for i := uint32(0); i < n; i++ {
		var lo lenOff
		if err := binary.Read(r, bo, &lo); err != nil {
			return nil, errors.WithSecondaryError(errReadMo,
				errors.Wrapf(err, "failed to read msg_id[%d] length & offset", i))
		}
		msgIDMeta = append(msgIDMeta, lo)
	}

	// 4 ????????? T ???????????? msgStr ??????
	if _, err := r.Seek(int64(h.OffsetStr), io.SeekStart); err != nil {
		return nil, errors.WithSecondaryError(errReadMo,
			errors.Wrapf(err, "failed to seek to O(message string table): %d", h.OffsetStr))
	}
	var msgStrMeta []lenOff
	for i := uint32(0); i < n; i++ {
		var lo lenOff
		if err := binary.Read(r, bo, &lo); err != nil {
			return nil, errors.WithSecondaryError(errReadMo,
				errors.Wrapf(err, "failed to read msg_str[%d] length & offset", i))
		}
		msgStrMeta = append(msgStrMeta, lo)
	}

	// 5 ????????????
	for i := uint32(0); i < n; i++ {
		// 5.1 ???????????? i ??? msg_id ?????????
		if _, err := r.Seek(int64(msgIDMeta[i].OffSet), io.SeekStart); err != nil {
			return nil, errors.WithSecondaryError(errReadMo,
				errors.Wrapf(err, "failed to seek to msg_id[%d]: %d", i, msgIDMeta[i].OffSet))
		}
		id := make([]byte, msgIDMeta[i].Length) // msg_id ?????????
		if err := binary.Read(r, bo, &id); err != nil {
			return nil, errors.WithSecondaryError(errReadMo,
				errors.Wrapf(err, "failed to read msg_id[%d]", i))
		}

		// 5.2 ???????????? msg_str
		if _, err := r.Seek(int64(msgStrMeta[i].OffSet), io.SeekStart); err != nil {
			return nil, errors.WithSecondaryError(errReadMo,
				errors.Wrapf(err, "failed to seek to msg_str[%d]: %d", i, msgStrMeta[i].OffSet))
		}
		str := make([]byte, msgStrMeta[i].Length) // msg_str ?????????
		if err := binary.Read(r, bo, &str); err != nil {
			return nil, errors.WithSecondaryError(errReadMo,
				errors.Wrapf(err, "failed to read msg_str[%d]", i))
		}

		// 5.3 as Entry
		var entry = &Entry{
			MsgID:  string(id),
			MsgStr: string(str),
		}
		// 0x04 ?????? msgCtxt ??? msgId
		if index := strings.Index(entry.MsgID, eot); index >= 0 {
			entry.MsgCtxt, entry.MsgID = entry.MsgID[:index], entry.MsgID[index+1:]
		}
		// 0x00 ?????? msgId ??? msgIdPlural
		if index := strings.Index(entry.MsgID, nul); index >= 0 {
			entry.MsgID, entry.MsgID2 = entry.MsgID[:index], entry.MsgID[index+1:]
			entry.MsgStrN = strings.Split(entry.MsgStr, nul)
			entry.MsgStr = ""
		}
		file.AddEntry(entry)
	}
	return file, nil
}

// SaveAsMo save as machine object file(.mo)
//  0: magic number = 0x950412de
//  4: version = 0
//  8: count = count
// 12: offset of origin string table = O = 75
// 16: offset of translation string table
// 20: size of hash table = 0
// 24: offset of hash table = 0
// 28: custom header entry: flag size = len(flag) = 43
// 32: custom header entry: flag
// 75: offsetO:   id table: (length & offset) * count
// xx: offsetT:   string table. xx=75+count*8
// aa: offsetID:  ids. aa=75+count*8*2
// bb: offsetStr: strs.
func (f *File) SaveAsMo(w io.Writer) error {
	count := len(f.entries)
	var ws = new(bytes.Buffer)
	writeMoHeader(ws, count)

	// map ??????????????????????????????????????????
	var entries = f.SortedEntry()

	// pos=O. from here is O.
	offsetID := offsetO + count*8*2 // 8=length(uint32)+offset(uint32) 2=id table + str table
	// length/offset of 0th string
	for _, entry := range entries {
		// length
		msgID := moMsgID(entry)
		msgIDLen := len(msgID)
		if err := binary.Write(ws, binary.LittleEndian, uint32(msgIDLen)); err != nil {
			return err
		}
		// offset ?????????
		if err := binary.Write(ws, binary.LittleEndian, uint32(offsetID)); err != nil {
			return err
		}
		offsetID += msgIDLen + 1 // +1: string end with null
	}

	// pos=T form here is T
	offsetStr := offsetID
	// length & offset 0th translation
	for _, entry := range entries {
		// length
		msgStr := moMsgStr(entry)
		msgStrLen := len(msgStr)
		if err := binary.Write(ws, binary.LittleEndian, uint32(msgStrLen)); err != nil {
			return err
		}
		// offset ?????????
		if err := binary.Write(ws, binary.LittleEndian, uint32(offsetStr)); err != nil {
			return err
		}
		offsetStr += msgStrLen + 1
	}

	// offsetH ignore

	// pos=offsetID
	// NUL terminated 0th string
	for _, entry := range entries {
		msgID := moMsgID(entry)
		if err := binary.Write(ws, binary.LittleEndian, []byte(msgID+nul)); err != nil {
			return err
		}
	}

	// pos=offsetStr
	// NUL terminated 0th translation
	for _, entry := range entries {
		msgStr := moMsgStr(entry)
		if err := binary.Write(ws, binary.LittleEndian, []byte(msgStr+nul)); err != nil {
			return err
		}
	}
	// ????????????????????????????????? ??????????????????
	_, err := ws.WriteTo(w)
	return err
}

func writeMoHeader(ws io.Writer, count int) error {
	var offsetT = offsetO + count*8 // string table is after id table
	// pos=0.  magic number
	if err := binary.Write(ws, binary.LittleEndian, uint32(moMagic)); err != nil {
		return err
	}
	// pos=4.  version=0
	if err := binary.Write(ws, binary.LittleEndian, uint32(0)); err != nil {
		return err
	}
	// pos=8.  N=number of strings
	if err := binary.Write(ws, binary.LittleEndian, uint32(count)); err != nil {
		return err
	}
	// pos=12. O=offset of ids table
	if err := binary.Write(ws, binary.LittleEndian, uint32(offsetO)); err != nil {
		return err
	}
	// pos=16. T=offset of translated str table
	if err := binary.Write(ws, binary.LittleEndian, uint32(offsetT)); err != nil {
		return err
	}
	// pos=20. S=0 size of hashtable
	if err := binary.Write(ws, binary.LittleEndian, uint32(0)); err != nil {
		return err
	}
	// pos=24. H=0 offset of hashtable
	if err := binary.Write(ws, binary.LittleEndian, uint32(0)); err != nil {
		return err
	}

	// ?????????????????????
	if err := binary.Write(ws, binary.LittleEndian, uint32(flagLen)); err != nil {
		return err
	}
	return binary.Write(ws, binary.LittleEndian, []byte(flag))
}

func moMsgID(entry *Entry) string {
	msgID := entry.MsgID
	if entry.MsgCtxt != "" {
		msgID = entry.MsgCtxt + eot + msgID
	}
	if entry.MsgID2 != "" {
		msgID += nul + entry.MsgID2
	}
	return msgID
}

func moMsgStr(entry *Entry) string {
	msgStr := entry.MsgStr
	if len(entry.MsgStrN) > 0 {
		msgStr = strings.Join(entry.MsgStrN, nul)
	}
	return msgStr
}
