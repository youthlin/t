package plurals

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var ctx = context.Background()

type args struct {
	ctx context.Context
	exp string
	n   int64
}
type tCase struct {
	name    string
	args    args
	want    int64
	wantErr bool
}

func okCase() []tCase {
	tests := []tCase{
		// https://www.gnu.org/software/gettext/manual/html_node/Plural-forms.html#index-plural-form-formulas

		// 中日韩越泰 等
		{"#0-only-one-form: 亚洲语言", args{ctx, "0", 0}, 0, false},
		{"#1-only-one-form: 亚洲语言", args{ctx, "0", 1}, 0, false},
		{"#2-only-one-form: 亚洲语言", args{ctx, "0", 2}, 0, false},
		// 英/德/丹麦/西班牙/意大利 等
		{"#0-two-forms: 单数只有1", args{ctx, "n != 1", 0}, 1, false},
		{"#1-two-forms: 单数只有1", args{ctx, "n != 1", 1}, 0, false},
		{"#2-two-forms: 单数只有1", args{ctx, "n != 1", 2}, 1, false},
		// 法
		{"#1-two-forms: 01是单数", args{ctx, "n > 1", 1}, 0, false},
		{"#2-two-forms: 01是单数", args{ctx, "n > 1", 2}, 1, false},
		{"#0-two-forms: 01是单数", args{ctx, "n > 1", 0}, 0, false},
		// 三种复数，1、2特殊 Gaeilge (Irish) 爱尔兰语
		{"#0-three-forms: 12特殊", args{ctx, "n==1 ? 0 : n==2 ? 1 : 2", 0}, 2, false},
		{"#1-three-forms: 12特殊", args{ctx, "n==1 ? 0 : n==2 ? 1 : 2", 1}, 0, false},
		{"#2-three-forms: 12特殊", args{ctx, "n==1 ? 0 : n==2 ? 1 : 2", 2}, 1, false},
		{"#3-three-forms: 12特殊", args{ctx, "n==1 ? 0 : n==2 ? 1 : 2", 3}, 2, false},
	}

	// 拉脱维亚语 Latvian 三种复数，0特殊
	for i := int64(0); i < 215; i++ {
		tests = append(tests,
			tCase{fmt.Sprintf("test-case-3-#%v", i),
				args{ctx, "n%10==1 && n%100!=11 ? 0 : n != 0 ? 1 : 2", i},
				func(n int64) int64 {
					// n%10==1 && n%100!=11 ? 0 : n != 0 ? 1 : 2
					if n%10 == 1 && n%100 != 11 {
						return 0
					}
					if n != 0 {
						return 1
					}
					return 2
				}(i), false},
		)
	}
	// 三种复数，00 [2-9][0-9] Romanian 罗马尼亚
	for i := int64(0); i < 215; i++ {
		tests = append(tests,
			tCase{fmt.Sprintf("test-case-4-#%v", i),
				args{ctx, "n==1 ? 0 : (n==0 || (n%100 > 0 && n%100 < 20)) ? 1 : 2", i},
				func(n int64) int64 {
					// n==1 ? 0 : (n==0 || (n%100 > 0 && n%100 < 20)) ? 1 : 2
					if n == 1 {
						return 0
					}
					if n == 0 || (n%100 > 0 && n%100 < 20) {
						return 1
					}
					return 2
				}(i), false},
		)
	}
	// 三种复数，1[2-9] 立陶宛 Lithuanian
	for i := int64(0); i < 215; i++ {
		tests = append(tests,
			tCase{fmt.Sprintf("test-case-5-#%v", i),
				args{ctx, "n%10==1 && n%100!=11 ? 0 : n%10>=2 && (n%100<10 || n%100>=20) ? 1 : 2", i},
				func(n int64) int64 {
					// n%10==1 && n%100!=11 ? 0 : n%10>=2 && (n%100<10 || n%100>=20) ? 1 : 2
					if n%10 == 1 && n%100 != 11 {
						return 0
					}
					if n%10 >= 2 && (n%100 < 10 || n%100 >= 20) {
						return 1
					}
					return 2
				}(i), false},
		)
	}
	// 三种复数，1234结尾的特殊，但不是 1[1-4] Russian, Ukrainian, Belarusian, Serbian, Croatian
	for i := int64(0); i < 215; i++ {
		tests = append(tests,
			tCase{fmt.Sprintf("test-case-6-#%v", i),
				args{ctx, "n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2", i},
				func(n int64) int64 {
					// n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2
					if n%10 == 1 && n%100 != 11 {
						return 0
					}
					if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
						return 1
					}
					return 2
				}(i), false},
		)
	}
	// 三种复数，1234结尾的特殊 Czech, Slovak
	for i := int64(0); i < 215; i++ {
		tests = append(tests,
			tCase{fmt.Sprintf("test-case-7-#%v", i),
				args{ctx, "(n==1) ? 0 : (n>=2 && n<=4) ? 1 : 2", i},
				func(n int64) int64 {
					// (n==1) ? 0 : (n>=2 && n<=4) ? 1 : 2
					if n == 1 {
						return 0
					}
					if n >= 2 && n <= 4 {
						return 1
					}
					return 2
				}(i), false},
		)
	}
	// 三种复数，Three forms, special case for one and some numbers ending in 2, 3, or 4
	// Polish
	for i := int64(0); i < 215; i++ {
		tests = append(tests,
			tCase{fmt.Sprintf("test-case-8-#%v", i),
				args{ctx, "n==1 ? 0 :  n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2", i},
				func(n int64) int64 {
					// n==1 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2
					if n == 1 {
						return 0
					}
					if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
						return 1
					}
					return 2
				}(i), false},
		)
	}
	// Four forms, special case for one and all numbers ending in 02, 03, or 04
	// Slovenian
	for i := int64(0); i < 215; i++ {
		tests = append(tests,
			tCase{fmt.Sprintf("test-case-9-#%v", i),
				args{ctx, "n%100==1 ? 0 : n%100==2 ? 1 : n%100==3 || n%100==4 ? 2 : 3", i},
				func(n int64) int64 {
					//n%100==1 ? 0 : n%100==2 ? 1 : n%100==3 || n%100==4 ? 2 : 3
					if n%100 == 1 {
						return 0
					}
					if n%100 == 2 {
						return 1
					}
					if n%100 == 3 || n%100 == 4 {
						return 2
					}
					return 3
				}(i), false},
		)
	}
	// Six forms, special cases for one, two, all numbers ending in 02, 03, … 10, all numbers ending in 11 … 99, and others
	// Arabic
	for i := int64(0); i < 215; i++ {
		tests = append(tests,
			tCase{fmt.Sprintf("test-case-10-#%v", i),
				args{ctx, " n==0 ? 0 : n==1 ? 1 : n==2 ? 2 : n%100>=3 && n%100<=10 ? 3  : n%100>=11 ? 4 : 5", i},
				func(n int64) int64 {
					if n == 0 {
						return 0
					}
					if n == 1 {
						return 1
					}
					if n == 2 {
						return 2
					}
					if n%100 >= 3 && n%100 <= 10 {
						return 3
					}
					if n%100 >= 11 {
						return 4
					}
					return 5
				}(i), false},
		)
	}
	return tests
}

func myCase() []tCase {
	return []tCase{
		{"custom-case0", args{ctx, "n++", 0}, 1, false},
		{"custom-case1", args{ctx, "n++", 1}, 2, false},
		{"custom-case2", args{ctx, "n--", 1}, 0, false},
		{"custom-case3", args{ctx, "+n", 1}, 1, false},
		{"custom-case4", args{ctx, "-n", 1}, -1, false},
		{"custom-case5", args{ctx, "++n", 1}, 2, false},
		{"custom-case6", args{ctx, "--n", 1}, 0, false},
		{"custom-case7", args{ctx, "~n", 1}, -2, false},
		{"custom-case8", args{ctx, "!n", 1}, 0, false},
		{"custom-case9", args{ctx, "n*1", 1}, 1, false},
		{"custom-case10", args{ctx, "n/1", 1}, 1, false},
		{"custom-case11", args{ctx, "n%1", 1}, 0, false},
		{"custom-case12", args{DebugCtx(ctx), "n+1", 0}, 1, false},
		{"custom-case13", args{ctx, "n-1", 1}, 0, false},
		{"custom-case14", args{ctx, "n>>1", 1}, 0, false},
		{"custom-case15", args{ctx, "n<<1", 1}, 2, false},
		{"custom-case16", args{ctx, "n<1", 1}, 0, false},
		{"custom-case17", args{ctx, "n>1", 1}, 0, false},
		{"custom-case18", args{ctx, "n>=1", 1}, 1, false},
		{"custom-case19", args{ctx, "n<=1", 1}, 1, false},
		{"custom-case20", args{ctx, "n==1", 1}, 1, false},
		{"custom-case21", args{ctx, "n!=1", 1}, 0, false},
		{"custom-case22", args{ctx, "n&1", 1}, 1, false},
		{"custom-case23", args{ctx, "n|1", 1}, 1, false},
		{"custom-case24", args{ctx, "n^1", 1}, 1, false},
		{"custom-case25", args{ctx, "n||1", 0}, 1, false},
		{"custom-case26", args{ctx, "n&&1", 0}, 0, false},
		{"custom-case27", args{ctx, "n&&1", 1}, 1, false},
		{"custom-case28", args{ctx, "n==1?1:0", 1}, 1, false},
		{"custom-case29", args{ctx, "n==1?1:0", 2}, 0, false},
		{"custom-case30", args{ctx, "0", 100}, 0, false},
		{"custom-case31", args{ctx, "n", 100}, 100, false},
		{"custom-case32", args{ctx, "(n)", 100}, 100, false},
		{"custom-case33", args{ctx, "!1", 100}, 0, false},
	}
}
func errCase() []tCase {
	return []tCase{
		{"err0", args{}, 0, true},              // SyntaxError: EOF
		{"err1", args{ctx, `nn`, 1}, 1, false}, // todo why 这个不太符合预期。应该返回错误才对
	}
}

func TestEval(t *testing.T) {
	tests := []tCase{}
	tests = append(tests, okCase()...)
	tests = append(tests, myCase()...)
	tests = append(tests, errCase()...)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Eval(tt.args.ctx, tt.args.exp, tt.args.n)
			// if err != nil {
			// 	t.Logf("err=%+v", err)
			// }
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEval2(t *testing.T) {
	Convey("Eval", t, func() {
		Convey("empty-exp", func() {
			_, err := Eval(DebugCtx(ctx), ``, 0)
			t.Logf("%v", err)
			So(err, ShouldNotBeNil)
		})
	})
}
