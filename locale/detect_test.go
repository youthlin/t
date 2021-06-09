package locale

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/text/language"
)

func TestGetDefault(t *testing.T) {
	Convey("GetDefault", t, func() {
		locale := GetDefault()
		t.Logf("Default locale is: %v", locale)
		So(locale, ShouldContainSubstring, "_")
		So(locale, ShouldNotBeEmpty)
		So(language.Make(locale), ShouldNotResemble, language.Und)
	})
}

func TestNormalize(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{""}, "en_US"},
		{"", args{"en_US"}, "en_US"},
		{"", args{"en-US"}, "en_US"},
		{"", args{"en"}, "en_US"},
		{"", args{"en_GB"}, "en_GB"},
		{"", args{"en-GB"}, "en_GB"},
		{"", args{"_"}, "en_US"},

		{"", args{"zh"}, "zh_CN"},
		{"", args{"zh_CN"}, "zh_CN"},
		{"", args{"zh_CN.UTF-8"}, "zh_CN"},
		{"", args{"zh_CN.UTF8"}, "zh_CN"},
		{"", args{"zh-CN"}, "zh_CN"},
		{"", args{"zh-Hans"}, "zh_CN"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Normalize(tt.args.lang); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeTag(t *testing.T) {
	type args struct {
		tag language.Tag
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{language.English}, "en_US"},
		{"", args{language.AmericanEnglish}, "en_US"},
		{"", args{language.BritishEnglish}, "en_GB"},

		{"", args{language.Chinese}, "zh_CN"},
		{"", args{language.SimplifiedChinese}, "zh_CN"},
		{"", args{language.TraditionalChinese}, "zh_TW"},
		{"", args{language.Make("zh")}, "zh_CN"},
		{"", args{language.Make("zh-Hans")}, "zh_CN"},
		{"", args{language.Make("zh-Hant")}, "zh_TW"},
		{"", args{language.Make("zh-Hant-TW")}, "zh_TW"},
		{"", args{language.Make("zh-Hant-HK")}, "zh_HK"},
		{"", args{language.Make("zh-Hans-HK")}, "zh_HK"}, // 香港简体 怪怪的
		{"", args{language.Make("zh_MO")}, "zh_MO"},
		{"", args{language.Make("zh-TW")}, "zh_TW"},
		{"", args{language.Make("zh_HK")}, "zh_HK"},
		{"", args{language.Make("zh-SG")}, "zh_SG"},

		{"", args{language.Und}, "en_US"},
		{"", args{language.Make("CN")}, "en_US"},
		{"", args{language.Make("")}, "en_US"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeTag(tt.args.tag); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
