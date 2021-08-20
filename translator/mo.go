package translator

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
)

const (
	nul     = "\x00"
	eot     = "\x04"
	moMagic = 0x950412de
	flag    = "ThisFileIsGenerateBy:github.com/youthlin/t" + nul
	flagLen = len(flag)        // 43
	offsetO = 28 + 4 + flagLen // 75: 28=固定header 4=uint32:flagLen 43=string:flag
)

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

	// pos=O. from here is O.
	offsetID := offsetO + count*8*2 // 8=length(uint32)+offset(uint32) 2=id table + str table
	// length/offset of 0th string
	for _, entry := range f.entries {
		// length
		msgID := moMsgID(entry)
		msgIDLen := len(msgID)
		if err := binary.Write(ws, binary.LittleEndian, uint32(msgIDLen)); err != nil {
			return err
		}
		// offset 先占位
		if err := binary.Write(ws, binary.LittleEndian, uint32(offsetID)); err != nil {
			return err
		}
		offsetID += msgIDLen + 1 // +1: string end with null
	}

	// pos=T form here is T
	offsetStr := offsetID
	// length & offset 0th translation
	for _, entry := range f.entries {
		// length
		msgStr := moMsgStr(entry)
		msgStrLen := len(msgStr)
		if err := binary.Write(ws, binary.LittleEndian, uint32(msgStrLen)); err != nil {
			return err
		}
		// offset 先占位
		if err := binary.Write(ws, binary.LittleEndian, uint32(offsetStr)); err != nil {
			return err
		}
		offsetStr += msgStrLen + 1
	}

	// offsetH ignore

	// pos=offsetID
	// NUL terminated 0th string
	for _, entry := range f.entries {
		msgID := moMsgID(entry)
		if err := binary.Write(ws, binary.LittleEndian, []byte(msgID+nul)); err != nil {
			return err
		}
	}

	// pos=offsetStr
	// NUL terminated 0th translation
	for _, entry := range f.entries {
		msgStr := moMsgStr(entry)
		if err := binary.Write(ws, binary.LittleEndian, []byte(msgStr+nul)); err != nil {
			return err
		}
	}
	// 不必回去填充各个占位符 已经计算好了
	return nil
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

	// 随便多填点内容
	if err := binary.Write(ws, binary.LittleEndian, uint32(flagLen)); err != nil {
		return err
	}
	return binary.Write(ws, binary.LittleEndian, []byte(flag))
}

func moMsgID(entry *Entry) string {
	msgID := entry.msgID
	if entry.msgCtxt != "" {
		msgID = entry.msgCtxt + eot + msgID
	}
	if entry.msgID2 != "" {
		msgID += nul + entry.msgID2
	}
	return msgID
}

func moMsgStr(entry *Entry) string {
	msgStr := entry.msgStr
	if len(entry.msgStrN) > 0 {
		msgStr = strings.Join(entry.msgStrN, nul)
	}
	return msgStr
}
