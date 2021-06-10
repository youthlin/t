package po

import (
	"io"

	"github.com/pkg/errors"
)

type reader struct {
	lines     []string
	lineNo    int
	totalLine int
}

func newReader(lines []string) *reader {
	return &reader{
		lines:     lines,
		lineNo:    -1,
		totalLine: len(lines),
	}
}

func (r *reader) currentLine() (string, error) {
	if r.lineNo < 0 {
		return "", errors.Errorf("you should call nextLine() first")
	}
	if r.lineNo >= r.totalLine {
		return "", io.EOF
	}
	return r.lines[r.lineNo], nil
}

func (r *reader) nextLine() (string, error) {
	r.lineNo++
	return r.currentLine()
}

func (r *reader) unGetLine() error {
	if r.lineNo < 0 {
		return errors.Errorf("already at the begining")
	}
	r.lineNo--
	return nil
}
