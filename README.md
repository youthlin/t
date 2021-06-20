# t
t: a translation util for go, inspired by GNU gettext  
t: GNU gettext 的 Go 语言实现，Go 程序国际化工具  
[![test](https://github.com/youthlin/t/actions/workflows/test.yaml/badge.svg)](https://github.com/youthlin/t/actions/workflows/test.yaml)
<!-- [![sync-to-gitee](https://github.com/youthlin/t/actions/workflows/gitee.yaml/badge.svg)](https://github.com/youthlin/t/actions/workflows/gitee.yaml) -->
[![codecov](https://codecov.io/gh/youthlin/t/branch/main/graph/badge.svg?token=6RyU5nb3YT)](https://codecov.io/gh/youthlin/t)
[![Go Report Card](https://goreportcard.com/badge/github.com/youthlin/t)](https://goreportcard.com/report/github.com/youthlin/t)
[![Go Reference](https://pkg.go.dev/badge/github.com/youthlin/t.svg)](https://pkg.go.dev/github.com/youthlin/t)

## Install 安装

```bash
go get -u github.com/youthlin/t
```

go.mod  
Gitee 镜像：[gitee.com/youthlin/t](gitee.com/youthlin/t)
```go
require (
    github.com/youthlin/t latest
)

// 使用 gitee 镜像
replace github.com/youthlin/t latest => gitee.com/youthlin/t latest
```

## Usage 使用
```go
path := "path/to/filename.po" // .po, .mo file
path = "path/to/po_mo/dir"    // or dir
// 1 bind domain 绑定翻译文件
t.BindTextDomain("my-domain", path)
// 2 set current domain 设置使用的文本域
t.TextDomain("my-domain")
// 3 set user language 设置用户语言
// t.SetLocale("zh_CN")
t.SetLocale("") // empty to use system default
// 4 use the gettext api 使用 gettext 翻译接口
fmt.Println(t.T("Hello, world"))
fmt.Println(t.T("Hello, %v", "Tom"))
fmt.Println(t.N("One apple", "%d apples", 1)) // One apple
fmt.Println(t.N("One apple", "%d apples", 2)) // %d apples
// args... supported, used to format string
// 支持传入 args... 参数用于格式化输出
fmt.Println(t.N("One apple", "%d apples", 2, 2)) // 2 apples
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

// T:  gettext
// N:  ngettext
// X:  pgettext
// XN: npgettext
// D:  domain
// L:  locale(language)

DT(domain, msgID, args...)
// and DN, DX, DXN, DN64, DXN64

LT(lang, msgID, args...)
// and LT, LX, LXN, LN64, LXN64

DLT(domain, lang, msgID, args...)
// and DLN, DLX, DLXN, DLN64, DLXN64
```

## Domain 文本域
```go
t.BindTextDomain(domain1, path1)
t.BindTextDomain(domain2, path2)
t.SetLocale("zh_CN")

t.T("msg_id")           // use default domain

t.TextDomain(domain1)
t.T("msg_id")           // use domain1

t.DT(domain2, "msg_id") // use domain2

t.DT("unknown-domain", "msg_id") // return "msg_id" directly

// or new domain
d := t.NewDomain(domain1)
d.T("msg_id")   // use domain1
```

## Language 指定语言
If you are building a web application, you may want each request use diffenrent language, the code below may help you:  
如果你写的是 web 应用而不是 cli 工具，你可能想让每个 request 使用不同的语言，请看：

```go
t.BindDefaultDomain(path)

// Specify a language 可以指定语言
t.LT("zh_CN", "msg_id")

// or Judging by the browser header（Accept-Language）
// 或者根据浏览器标头获取用户语言
langs := t.SupportLangs(t.DefaultDomain）
// golang.org/x/text/language
// EN: https://blog.golang.org/matchlang
// 中文: https://learnku.com/docs/go-blog/matchlang/6525
var supported []language.Tag
for _, lang =range langs{
    supported = append(supported, language.Make(lang))
}
matcher := language.NewMatcher(supported)
// or: userAccept := []language.Tag{ language.Make("lang-code-from-cookie") }
// 或从 cookie 中获取用户语言偏好
userAccept, q, err :=language.ParseAcceptLanguage("zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
matchedTag, index, confidence := matcher.Match(userAccept...)
// confidence may language.No, language.Low, language.High, language.Exact
// 这里 confidence 是指匹配度，可以根据你的实际需求决定是否使用匹配的语言。
// 如服务器支持 英文、简体中文，如果用户是繁体中文，那匹配度就不是 Exact，
// 这时根据实际需求决定是使用英文，还是简体中文。
userLang := langs[index]
t.LT(userLang, "msg_id")

// or NewLocale
l := t.NewLocale("zh_CN")
l.T("msg_id")

// with domain, language 同时指定文本域、用户语言
t.DLT(domain, userLang, "msg_id")
```

## How to extract string 提取翻译文本
```bash
# if you use PoEdit, add a extractor
# 如果你使用 PoEdit，在设置-提取器中新增一个提取器
# ‪xgettext -C --add-comments=TRANSLATORS: --force-po -o %o %C %K %F
# keywords: 关键字这样设置：
# T:1;N:1,2;N64:1,2;X:2,1c;XN:2,3,1c;XN64:2,3,1c;
# LT:2;LN:2,3;LN64:2,3;LX:3,2c;LXN:3,4,2c;LXN64:3,4,2c;
# DLT:3;DLN:3,4;DLN64:3,4;DLX:4,3c;DLXN:4,5,3c;DLXN64:4,5,3c
‪xgettext -C --add-comments=TRANSLATORS: --force-po ‪-kT -kN:1,2 -kX:2,1c -kXN:2,3,1c -k...  *.go
```

## Todo 待办
- [x] mo file 支持 mo 二进制文件  
- [x] extract from html templates 从模板文件中提取: [xtemplate](cmd/xtemplate/)  
```bash
go install github.com/youthlin/t/cmd/xtemplate
```

## Links 链接
- https://www.gnu.org/software/gettext/manual/html_node/index.html
- https://github.com/search?l=Go&q=gettext&type=Repositories
- https://github.com/antlr/antlr4/
- https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/
- https://xuanwo.io/2019/12/11/golang-i18n/ (中文)

