package bytes2_test

import (
	"io"
	"testing"

	"github.com/youthlin/t/bytes2"
)

func TestWrite(t *testing.T) {
	var ws = new(bytes2.WriteSeeker)
	// write, result string, want position
	checkWrite(t, ws, "Hello", "Hello", len("Hello"))
	checkWrite(t, ws, ", World", "Hello, World", len("Hello, World"))
}

func TestSeek(t *testing.T) {
	var ws = new(bytes2.WriteSeeker)
	// seek offset, whence, want positon, want err
	checkSeek(t, ws, -1, io.SeekStart, 0, true)

	checkWrite(t, ws, "Hello", "Hello", len("Hello"))

	checkSeek(t, ws, 1, io.SeekStart, 1, false)
	checkWrite(t, ws, "a", "Hallo", 2)

	checkSeek(t, ws, 1, io.SeekEnd, len("Hallo")+1, false)
	checkWrite(t, ws, "a", "Hallo\x00a", len("Hallo\x00a"))

	checkSeek(t, ws, -2, io.SeekCurrent, len("Hallo"), false)
	checkWrite(t, ws, "A", "HalloAa", len("HalloA"))
	checkWrite(t, ws, "B", "HalloAB", len("HalloAB"))
}

func checkWrite(t *testing.T, ws *bytes2.WriteSeeker, write, exp string, wantPos int) {
	t.Helper()
	n, err := ws.Write([]byte(write))
	if err != nil {
		t.Fatalf("Write(%q) failed: err=%+v", write, err)
	}
	if n != len(write) {
		t.Fatalf("Write(%q) write %d bytes, want %d", write, n, len(write))
	}
	if result := string(ws.Bytes()); result != exp {
		t.Fatalf("Write(%q) result in %q, want %q", write, result, exp)
	}
	if ws.CurrentPos() != wantPos {
		t.Fatalf("Write(%q) current pos=%d, want %d", write, ws.CurrentPos(), wantPos)
	}
}

func checkSeek(t *testing.T, ws *bytes2.WriteSeeker, offset int64, whence, exp int, wanterr bool) {
	t.Helper()
	pos, err := ws.Seek(offset, whence)
	if (err != nil) != wanterr {
		t.Fatalf("Seek(%d, %d) failed: err=%+v", offset, whence, err)
	}
	if int(pos) != exp {
		t.Fatalf("Seek(%d, %d) new positon=%d want=%d", offset, whence, pos, exp)
	}
}
