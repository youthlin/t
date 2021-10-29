# t
t: a translation util for go, inspired by GNU gettext  
t: GNU gettext 的 Go 语言实现，Go 程序的国际化工具  
[![sync-to-gitee](https://github.com/youthlin/t/actions/workflows/gitee.yaml/badge.svg)](https://github.com/youthlin/t/actions/workflows/gitee.yaml)
[![test](https://github.com/youthlin/t/actions/workflows/test.yaml/badge.svg)](https://github.com/youthlin/t/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/youthlin/t/branch/main/graph/badge.svg?token=6RyU5nb3YT)](https://codecov.io/gh/youthlin/t)
[![Go Report Card](https://goreportcard.com/badge/github.com/youthlin/t)](https://goreportcard.com/report/github.com/youthlin/t)
[![Go Reference](https://pkg.go.dev/badge/github.com/youthlin/t.svg)](https://pkg.go.dev/github.com/youthlin/t)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyouthlin%2Ft.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyouthlin%2Ft?ref=badge_shield)


## Install 安装

```bash
go get -u github.com/youthlin/t
```

go.mod
```go
require (
    github.com/youthlin/t latest
)
```

Gitee 镜像：[gitee.com/youthlin/gottext](gitee.com/youthlin/gottext) (gottext: go + gettext)
> 鸣谢仓库同步工具：https://github.com/Yikun/hub-mirror-action
```
// 使用 gitee 镜像
// go.mod:
replace github.com/youthlin/t latest => gitee.com/youthlin/gottext latest
```


## Usage 使用
```go
path := "path/to/filename.po" // .po, .mo file
path = "path/to/po_mo/dir"    // or dir.
// (mo po 同名的话，po 后加载，会覆盖 mo 文件，因为 po 是文本文件，方便修改生效)
// 1 bind domain 绑定翻译文件
t.Load(path)
t.Bind("my-domain", path)
// 2 set current domain 设置使用的文本域
t.SetDomain("my-domain")
// 3 set user language 设置用户语言
// t.SetLocale("zh_CN")
t.SetLocale("") // empty to use system default
// 4 use the gettext api 使用 gettext 翻译接口
fmt.Println(t.T("Hello, world"))
fmt.Println(t.T("Hello, %v", "Tom"))
fmt.Println(t.N("One apple", "%d apples", 1)) // One apple
fmt.Println(t.N("One apple", "%d apples", 2)) // %d apples
// t.N(single, plural, n, args...)
// n => used to choose single or plural
// args => to format
// args... supported, used to format string
// 支持传入 args... 参数用于格式化输出
fmt.Println(t.N("One apple", "%d apples", 2, 2)) // 2 apples
fmt.Println(t.N("%[2]s has one apple", "%[2]s has %[1]d apples", 2, 200, "Bob"))
// Bob has 200 apples
t.X("msg_context_text", "msg_id")
t.X("msg_context_text", "msg_id")
t.XN("msg_context_text", "msg_id", "msg_plural", n)
```

## API
```go
T(msgID, args...)
N(msgID, msgIDPlural, n, args...) // and N64
X(msgCTxt, msgID, args...)
XN(msgCTxt, msgID, msgIDPlural, n, args...) // and XN64
D(domain)
L(locale)
// T:  gettext
// N:  ngettext
// X:  pgettext
// XN: npgettext
// D:  with domain
// L:  with locale(language)
```

## Domain 文本域
```go
t.Bind(domain1, path1)
t.Bind(domain2, path2)
t.SetLocale("zh_CN")

t.T("msg_id")           // use default domain

t.SetDomain(domain1)
t.T("msg_id")            // use domain1
t.D(domain2).T("msg_id") // use domain2
t.D("unknown-domain").T("msg_id") // return "msg_id" directly

```

## Language 指定语言
If you are building a web application, you may want each request use diffenrent language, the code below may help you:  
如果你写的是 web 应用而不是 cli 工具，你可能想让每个 request 使用不同的语言，请看：

```go
t.Load(path)

// a) Specify a language 可以指定语言
t.L("zh_CN").T("msg_id")

// b) every one use his own language 每个用户使用他接受的语言
// b.1) server supports 第一步，服务器支持的语言
langs := t.Locales()
// golang.org/x/text/language
// EN: https://blog.golang.org/matchlang
// 中文: https://learnku.com/docs/go-blog/matchlang/6525
var supported []language.Tag
for _, lang =range langs{
    supported = append(supported, language.Make(lang))
}
matcher := language.NewMatcher(supported)
// b.2) user accept 第二步，用户接受的语言
// Judging by the browser header（Accept-Language）
// 根据浏览器标头获取用户语言
// or: userAccept := []language.Tag{ language.Make("lang-code-from-cookie") }
// 或从 cookie 中获取用户语言偏好
userAccept, q, err :=language.ParseAcceptLanguage("zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
// b.3) match 第三步，匹配出最合适的
matchedTag, index, confidence := matcher.Match(userAccept...)
// confidence may language.No, language.Low, language.High, language.Exact
// 这里 confidence 是指匹配度，可以根据你的实际需求决定是否使用匹配的语言。
// 如服务器支持 英文、简体中文，如果用户是繁体中文，那匹配度就不是 Exact，
// 这时根据实际需求决定是使用英文，还是简体中文。
userLang := langs[index]
t.L(userLang).T("msg_id")

// with domain, language 同时指定文本域、用户语言
t.D(domain).L(userLang).T("msg_id")
```

> more examples can be find at: [example_test.go](example_test.go)

## How to extract string 提取翻译文本
```bash
# if you use PoEdit, add a extractor
# 如果你使用 PoEdit，在设置-提取器中新增一个提取器
# ‪xgettext -C --add-comments=TRANSLATORS: --force-po -o %o %C %K %F
# keywords: 关键字这样设置：
# T:1;N:1,2;N64:1,2;X:2,1c;XN:2,3,1c;XN64:2,3,1c
‪xgettext -C --add-comments=TRANSLATORS: --force-po ‪-kT -kN:1,2 -kX:2,1c -kXN:2,3,1c  *.go
```

## Done 已完成
- ✅ mo file 支持 mo 二进制文件
- ✅ extract from html templates 从模板文件中提取: [xtemplate](cmd/xtemplate/)
```bash
go install github.com/youthlin/t/cmd/xtemplate@latest
```

## Links 链接
- https://www.gnu.org/software/gettext/manual/html_node/index.html
- https://github.com/search?l=Go&q=gettext&type=Repositories
- https://github.com/antlr/antlr4/
- https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/
- https://xuanwo.io/2019/12/11/golang-i18n/ (中文)



## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyouthlin%2Ft.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyouthlin%2Ft?ref=badge_large)