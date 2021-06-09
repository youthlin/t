package locale

import (
	"fmt"

	"github.com/Xuanwo/go-locale"
	"golang.org/x/text/language"
)

// GetDefault get system default language in the form of:
// ll or ll_CC, where 'll' an ISO 639 two-letter language code (lowercase)
// and ‘CC’ is an ISO 3166 two-letter country code (uppercase)
// https://www.gnu.org/software/gettext/manual/html_node/Header-Entry.html
// 获取系统默认语言，返回格式是 ll 或 ll_CC，其中 ll 是小写语言二字码，CC 是大写地区二字码
func GetDefault() string {
	tag, _ := locale.Detect()
	return NormalizeTag(tag)
}

// Normalize return ll_CC code
func Normalize(lang string) string {
	return NormalizeTag(language.Make(lang))
}

// NormalizeTag return ll_CC code
func NormalizeTag(tag language.Tag) string {
	base, _ := tag.Base()
	region, _ := tag.Region()
	return fmt.Sprintf("%v_%v", base, region)
}
