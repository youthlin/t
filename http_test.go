package t

import (
	"net/http"
	"testing"
)

func TestGetUserLang(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{Name: "language", Value: "zh_CN"})
	r.Header.Add(HTTPHeaderAcceptLanguage, "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	type args struct {
		request *http.Request
		opts    []getUserLangOpt
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "cookie", args: args{request: r, opts: []getUserLangOpt{WithCookieName("language")}}, want: "zh_CN"},
		{name: "header", args: args{request: r}, want: "en_US"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserLang(tt.args.request, tt.args.opts...); got != tt.want {
				t.Errorf("GetUserLang() = %v, want %v", got, tt.want)
			}
		})
	}
}
