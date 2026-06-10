package parse

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	tree, err := ParseInput("++")
	t.Logf("tree=\n%v, err=%v", tree, err)
	tree, err = ParseInput("++1")
	t.Logf("tree=\n%v, err=%v", tree, err)
	tree, err = ParseInput("*")
	t.Logf("tree=\n%v, err=%v", tree, err)
	tree, err = ParseInput("(*")
	t.Logf("tree=\n%v, err=%v", tree, err)

	tree, err = ParseInput("1++")
	t.Logf("tree=\n%v, err=%v", tree, err)

	tree, err = ParseInput("1+")
	t.Logf("tree=\n%v, err=%v", tree, err)

	tree, err = ParseInput("1?:")
	t.Logf("tree=\n%v, err=%v", tree, err)

	tree, err = ParseInput("1?0:")
	t.Logf("tree=\n%v, err=%v", tree, err)

	tree, err = ParseInput("n==0?0:n==1?1:n==2?2:n%100>=3&&n%100<=10?3:n%100>=11?4:5")
	t.Logf("tree=\n%v, err=%v", tree, err)
}
