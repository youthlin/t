module github.com/youthlin/t

go 1.23.0

// 发布流程：
// 1. 完成改动并测试通过
// 2. 如果同时发布 cmd/xtemplate，更新 cmd/xtemplate/go.mod 中 github.com/youthlin/t 的版本到本次版本
// 3. 在同一个 commit 上创建 tag：v0.1.10 与 cmd/xtemplate/v0.1.10
// 4. git push && git push --tags

require (
	github.com/Xuanwo/go-locale v1.1.0
	github.com/antlr4-go/antlr/v4 v4.13.0
	github.com/smartystreets/goconvey v1.8.1
	golang.org/x/text v0.23.0
)

require (
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/smarty/assertions v1.15.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
	golang.org/x/sys v0.31.0 // indirect
)
