module github.com/youthlin/t/cmd/xtemplate

go 1.25.0

// 发布流程（同仓库多模块，go.sum 无需包含 t 的 hash，Go 自动从同仓库解析）：
// 1. 修改此处 t 版本号为新版本（如 v0.1.7）
// 2. 确保已删除 replace 指令（本地开发请用 go.work，不要用 replace）
// 3. 回到仓库根目录，提交所有改动
// 4. git tag v0.1.7 && git tag cmd/xtemplate/v0.1.7
// 5. git push && git push --tags
// 注意：两个 tag 必须在同一个 commit 上。

require (
	github.com/cockroachdb/errors v1.13.0
	github.com/smartystreets/goconvey v1.8.1
	github.com/youthlin/t v0.1.9
)

require (
	github.com/Xuanwo/go-locale v1.1.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/getsentry/sentry-go v0.46.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/smarty/assertions v1.15.0 // indirect
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
