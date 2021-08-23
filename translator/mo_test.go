package translator

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestReadMo(t *testing.T) {
	fMo, err := os.Open("../testdata/zh_CN.mo")
	if err != nil {
		t.Fatalf("can not open test mo file: err=%+v", err)
	}
	mo := read(t, fMo)
	t.Logf("mo=%#v", mo)
	w, err := os.OpenFile("../testdata/zh_CN_save.mo", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o666)
	if err != nil {
		t.Fatalf("can not open save file: err=%+v", err)
	}
	defer w.Close()
	mo.SaveAsMo(w)

	w.Seek(0, io.SeekStart)
	mo2 := read(t, w)
	if !reflect.DeepEqual(mo.entries, mo2.entries) {
		t.Errorf("entries:\n origin=%#v\n read=%#v\n", mo, mo2)
	}
	if !reflect.DeepEqual(mo.headers, mo2.headers) {
		t.Errorf("headers:\n origin=%#v\n read=%#v\n", mo, mo2)
	}
	if !reflect.DeepEqual(mo.plural.totalForms, mo2.plural.totalForms) {
		t.Errorf("plural.totalForms:\n origin=%#v\n read=%#v\n", mo, mo2)
	}
	if !reflect.DeepEqual(mo.plural.expression, mo2.plural.expression) {
		t.Errorf("plural.expression:\n origin=%#v\n read=%#v\n", mo, mo2)
	}
}

func read(t *testing.T, r io.Reader) *File {
	t.Helper()
	content, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("can not read test mo file: err=%+v", err)
	}
	mo, err := ReadMo(content)
	if err != nil {
		t.Fatalf("read mo fail: err=%+v", err)
	}
	return mo
}
