package t

import (
	"reflect"
	"testing"

	"golang.org/x/text/language"
)

func TestMatch(t *testing.T) {
	type args struct {
		supported  []string
		userAccept []string
	}
	tests := []struct {
		name      string
		args      args
		wantTag   language.Tag
		wantIndex int
		wantC     language.Confidence
	}{
		{
			name:      "exact", // 精确匹配
			args:      args{supported: []string{"zh-CN"}, userAccept: []string{"zh-CN"}},
			wantTag:   language.Make("zh-CN"),
			wantIndex: 0,
			wantC:     language.Exact,
		},
		{
			name:      "exact2", // 精确匹配
			args:      args{supported: []string{"en-US", "zh-CN"}, userAccept: []string{"zh-CN"}},
			wantTag:   language.Make("zh-CN"),
			wantIndex: 1,
			wantC:     language.Exact,
		},
		{
			name:      "exact3", // 精确匹配 zh 默认就是指 zh-CN
			args:      args{supported: []string{"zh"}, userAccept: []string{"zh-CN"}},
			wantTag:   language.Make("zh-u-rg-cnzzzz"), // 但返回的 tag 和 support / userAccept 都不同(带 -u 后缀指定了地区 rg=region?)
			wantIndex: 0,                               // 使用 supported 的第 0 个，即 zh
			wantC:     language.Exact,
		},
		{
			name:      "exact4", // 精确匹配 zh-Hant 繁体中文默认就是指 中国台湾的繁体中文
			args:      args{supported: []string{"zh-Hant"}, userAccept: []string{"zh-TW"}},
			wantTag:   language.Make("zh-Hant-u-rg-twzzzz"),
			wantIndex: 0,
			wantC:     language.Exact,
		},
		{
			name:      "exact5", // 精确匹配 zh-HK 中国香港 默认就是 Hant 繁体
			args:      args{supported: []string{"zh-Hant-HK"}, userAccept: []string{"zh-HK"}},
			wantTag:   language.Make("zh-Hant-HK"),
			wantIndex: 0,
			wantC:     language.Exact,
		},
		{
			name:      "exact6", // 精确匹配 en 默认就是指美国英语
			args:      args{supported: []string{"en"}, userAccept: []string{"en-US"}},
			wantTag:   language.Make("en-u-rg-uszzzz"),
			wantIndex: 0,
			wantC:     language.Exact,
		},
		{
			name:      "exact7", // 精确匹配
			args:      args{supported: []string{"en"}, userAccept: []string{"en"}},
			wantTag:   language.Make("en"),
			wantIndex: 0,
			wantC:     language.Exact,
		},

		{
			name:      "high", // 高度匹配 zh-HK中国香港也使用 zh-Hant 繁体中文（未指明地区时其实默认是指中国台湾）
			args:      args{supported: []string{"zh-Hant"}, userAccept: []string{"zh-HK"}},
			wantTag:   language.Make("zh-Hant-u-rg-hkzzzz"),
			wantIndex: 0,
			wantC:     language.High,
		},
		{
			name:      "high2", // 高度匹配 提供了 zh-CN中国大陆，zh-MO中国澳门， zh-HK中国香港 更匹配同样使用 Hant繁体 的澳门
			args:      args{supported: []string{"zh-CN", "zh-MO"}, userAccept: []string{"zh-HK"}},
			wantTag:   language.Make("zh-MO-u-rg-hkzzzz"),
			wantIndex: 1,
			wantC:     language.High,
		},
		{
			name:      "high3", // 高度匹配 en 默认指美国英语 en-GB 英国也是英语
			args:      args{supported: []string{"en"}, userAccept: []string{"en-GB"}},
			wantTag:   language.Make("en-u-rg-gbzzzz"),
			wantIndex: 0,
			wantC:     language.High,
		},

		{
			name:      "low", // 也算找到了最接近的
			args:      args{supported: []string{"en", "zh"}, userAccept: []string{"zh-HK"}},
			wantTag:   language.Make("zh-u-rg-hkzzzz"),
			wantIndex: 1,
			wantC:     language.Low,
		},
		{
			name:      "low2", // 也算找到了最接近的 zh-CN中国大陆，从 英文，繁体中文(中国台湾) 中选择 繁体中文 最接近
			args:      args{supported: []string{"en", "zh-TW"}, userAccept: []string{"zh-CN"}},
			wantTag:   language.Make("zh-TW-u-rg-cnzzzz"),
			wantIndex: 1,
			wantC:     language.Low,
		},

		{
			name:      "no", // 不匹配
			args:      args{supported: []string{"en"}, userAccept: []string{"zh-HK"}},
			wantTag:   language.Make("en-u-rg-hkzzzz"),
			wantIndex: 0,
			wantC:     language.No,
		},
		{
			name:      "no1", // 不匹配
			args:      args{supported: []string{"en-US"}, userAccept: []string{"zh-HK"}},
			wantTag:   language.Make("en-US-u-rg-hkzzzz"),
			wantIndex: 0,
			wantC:     language.No,
		},
		{
			name:      "no2", // 不匹配
			args:      args{supported: []string{"en"}, userAccept: []string{"zh"}},
			wantTag:   language.Make("en"),
			wantIndex: 0,
			wantC:     language.No,
		},
		{
			name:      "no3", // 不匹配
			args:      args{supported: []string{"en"}, userAccept: []string{"zh-Hans"}},
			wantTag:   language.Make("en"),
			wantIndex: 0,
			wantC:     language.No,
		},
		{
			name:      "no4", // 不匹配
			args:      args{supported: []string{"ja"}, userAccept: []string{"zh-CN"}},
			wantTag:   language.Make("ja-u-rg-cnzzzz"),
			wantIndex: 0,
			wantC:     language.No,
		},
		{
			name:      "no5", // 不匹配
			args:      args{supported: []string{"ja", "en"}, userAccept: []string{"zh"}},
			wantTag:   language.Make("ja"), // 不匹配时都是返回第 0 个 supported
			wantIndex: 0,
			wantC:     language.No,
		},
		{
			name:      "no6", // 不匹配
			args:      args{supported: []string{"en", "jp"}, userAccept: []string{"zh"}},
			wantTag:   language.Make("en"),
			wantIndex: 0,
			wantC:     language.No,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTag, gotIndex, gotC := Match(tt.args.supported, tt.args.userAccept)
			if !reflect.DeepEqual(gotTag, tt.wantTag) {
				t.Errorf("Match() gotTag = %v, want %v", gotTag, tt.wantTag)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("Match() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Match() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
