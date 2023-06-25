package t

import (
	"net/http"

	"golang.org/x/text/language"
)

type httpConfig struct {
	cookieName string
}
type getUserLangOpt func(*httpConfig)

func WithCookieName(cookieName string) getUserLangOpt {
	return func(hc *httpConfig) { hc.cookieName = cookieName }
}

const HTTPHeaderAcceptLanguage = "Accept-Language"

func GetUserLang(request *http.Request, opts ...getUserLangOpt) string {
	cfg := &httpConfig{cookieName: "lang"}
	for _, opt := range opts {
		opt(cfg)
	}
	if cookie, err := request.Cookie(cfg.cookieName); err == nil {
		return cookie.Value
	}

	langs := Locales()
	var supported []language.Tag // 转换为 Tag
	for _, lang := range langs {
		supported = append(supported, language.Make(lang))
	}
	matcher := language.NewMatcher(supported)                     // 匹配器
	acceptLangs := request.Header.Get(HTTPHeaderAcceptLanguage)   // 用户支持的语言
	userAccept, _, _ := language.ParseAcceptLanguage(acceptLangs) // 转为 Tag
	_, index, _ := matcher.Match(userAccept...)                   // 找到最匹配的语言
	return langs[index]
}
