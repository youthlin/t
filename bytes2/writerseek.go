package bytes2

import (
	"fmt"
	"io"
)

var ErrInvalidSeekPos = fmt.Errorf("invalid seek positon")

type WriteSeeker struct {
	buf []byte
	pos int
}

var _ io.WriteSeeker = (*WriteSeeker)(nil)

func (ws *WriteSeeker) Write(p []byte) (n int, err error) {
	n = len(p)
	if grow := ws.pos + n - len(ws.buf); grow > 0 { // 容量不够需要先扩容
		ws.buf = append(ws.buf, make([]byte, grow)...)
	}
	copy(ws.buf[ws.pos:], p)
	ws.pos += n
	return n, err
}

// Seek 根据参数设置偏移量
// offset=偏移量 whence=相对于哪里偏移
// 返回相对于起始位置和发生的错误
func (ws *WriteSeeker) Seek(offset int64, whence int) (int64, error) {
	pos := ws.pos + int(offset) // 默认相对于当前位置偏移
	switch whence {
	case io.SeekStart:
		pos = int(offset)
	case io.SeekEnd:
		pos = len(ws.buf) + int(offset)
	}
	if pos < 0 {
		return int64(ws.pos), ErrInvalidSeekPos
	}
	ws.pos = pos
	return int64(pos), nil
}

func (ws *WriteSeeker) Bytes() []byte {
	return ws.buf
}

func (ws *WriteSeeker) CurrentPos() int {
	return ws.pos
}
