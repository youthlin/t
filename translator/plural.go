package translator

import (
	"context"
	"regexp"
	"strconv"

	"github.com/youthlin/t/plurals"
)

// plural 复数
type plural struct {
	totalForms int
	expression string
	fn         func(int64) int
}

var (
	rePlurals       = regexp.MustCompile(`^\s*nplurals\s*=\s*(\d)\s*;\s*plural\s*=\s*(.*)\s*;$`)
	invalidPluralFn = func(i int64) int { return -1 }
)

func parsePlural(forms string) *plural {
	var p plural
	p.fn = invalidPluralFn
	if forms == "" {
		return &p
	}
	find := rePlurals.FindAllStringSubmatch(forms, -1)
	if len(find) == 1 && len(find[0]) == 3 {
		n := find[0][1]
		exp := find[0][2]
		if total, err := strconv.ParseInt(n, 10, 64); err == nil {
			p.totalForms = int(total)
			p.expression = exp
			p.fn = func(i int64) int {
				index, err := plurals.Eval(context.Background(), exp, i)
				if err != nil {
					return -1
				}
				return int(index)
			}
		}
	}
	return &p
}
