# AGENTS.md — 项目架构指南

## 项目概述

`t` 是一个受 GNU gettext 启发的 Go 国际化（i18n）库，提供完整的翻译基础设施：

- 加载 `.po` / `.mo` 翻译文件
- 基于文本域（domain）的翻译管理
- 语言检测与匹配（系统语言、HTTP Accept-Language）
- 复数形式求值（ANTLR4 语法解析）
- CLI 工具 `xtemplate`：从 Go `html/template` 文件中提取可翻译字符串生成 `.pot`

**模块路径：** `github.com/youthlin/t`  
**Go 版本：** 1.23.0（主模块），1.25.0（workspace）  
**许可证：** MIT

## 架构概览

```
┌─────────────────────────────────────────────────────┐
│  公开 API（根包 t）                                 │
│  T(), N(), N1(), N64(), X(), XN(), XN64()           │
│  D(), L(), Load(), Bind(), SetLocale(), SetDomain() │
│  Mark（noop 标记，辅助 xgettext 提取）              │
│  GetUserLang()（HTTP 语言检测）                     │
│  WithContext()（context.Context 支持）              │
├─────────────────────────────────────────────────────┤
│  Translations（全局单例）                           │
│  ├── locale, domain, sourceCodeLocale               │
│  └── domains: map[string]*Translation               │
│       └── Translation（每个 domain 一个）           │
│            └── langs: map[string]Translator         │
│                 └── translator.File（po/mo 文件）   │
│                      ├── entries: map[string]*Entry │
│                      ├── headers                    │
│                      └── plural（Plural-Forms）     │
├─────────────────────────────────────────────────────┤
│  子包                                               │
│  ├── translator/  PO/MO/POT 读写、Entry、File       │
│  ├── locale/      系统语言检测与标准化              │
│  ├── f/           安全的 Sprintf、DefaultPlural     │
│  ├── errors/      错误包装辅助                      │
│  └── plurals/     复数表达式求值（ANTLR4）          │
├─────────────────────────────────────────────────────┤
│  cmd/xtemplate/   CLI：从模板提取字符串             │
│  └── internal/    AST 遍历、关键字解析、POT 生成    │
└─────────────────────────────────────────────────────┘
```

## 核心类型

| 类型 | 文件 | 说明 |
|------|------|------|
| `Translator` 接口 | `translator/translator.go:4` | 核心翻译接口：`Lang()`, `X()`, `XN64()` |
| `File` | `translator/file.go:22` | 实现 `Translator`，表示单个 po/mo 文件 |
| `Entry` | `translator/entry.go:10` | 单条翻译条目（msgctxt, msgid, msgid_plural, msgstr） |
| `Translation` | `translaton.go:22` | 一个 domain 的多语言翻译集合 |
| `Translations` | `translations.go:18` | 顶层管理器，持有多个 domain，跟踪当前语言和域 |
| `noop` | `mark.go:8` | 标记类型，方法返回参数原值，辅助 xgettext 提取 |

## 关键设计

1. **全局单例模式：** 根包使用 `global` 变量（`*Translations`），所有顶层函数委托给它。`D()` 和 `L()` 返回修改了 domain/locale 的克隆副本，支持链式调用：`t.D("main").L("zh_CN").T("Hello")`。

2. **PO/MO 对称：** `translator` 包可读写 `.po`、`.mo`、`.pot` 文件。MO 文件使用自定义二进制格式，标志字符串为 `ThisFileIsGenerateBy:github.com/youthlin/t`。

3. **复数形式：** 使用 ANTLR4 语法（`plurals/plural.g4`）解析任意 C 风格复数表达式。同时内置 11 种常见复数公式的硬编码映射用于快速求值。

4. **模板提取：** `xtemplate` 遍历 Go `text/template/parse` AST，处理 `if`/`range`/`with`/`template` 节点和管道参数。自动从模板源码检测函数名，无需手动指定 `-f`。

5. **安全格式化：** `f.Format()` 通过追加哑参数和使用索引占位符，避免 `fmt.Sprintf` 的 `%!(EXTRA)` 和 `%!(MISSING)` 错误。

## 构建与命令

```bash
# 运行测试（含覆盖率、竞态检测）
bash unittest.sh

# 等价于：
mkdir -p output
gofmt -w .
go mod tidy
go test -gcflags all=-l -cover -race -coverprofile=output/cover.txt ./... \
  && go tool cover -func=output/cover.txt \
  && go tool cover -html=output/cover.txt

# 构建 xtemplate 工具
cd cmd/xtemplate && bash build.sh
# ENV=dev 可构建 debug 版本（禁用优化和内联）

# 安装 xtemplate
go install github.com/youthlin/t/cmd/xtemplate@latest
```

## 代码风格

- **格式化：** 使用 `gofmt`（`unittest.sh` 中自动执行）
- **命名：** 遵循 Go 惯例。公开 API 使用短名称（`T`, `N`, `X`, `D`, `L`）模仿 gettext 风格
- **类型参数：** 使用 `any` 而非 `any`（Go 1.18+）
- **注释：** 公开函数和类型需有中文注释

## 测试

- **框架：** `github.com/smartystreets/goconvey`（主要）+ 标准 `testing`
- **测试文件：** 17 个测试文件，覆盖所有包
- **测试数据：** `testdata/` 目录包含 `.po`、`.mo`、`.pot` 示例文件
- **CI：** GitHub Actions 在 push/PR 时运行 `go test -race -coverprofile=coverage.txt -covermode=atomic ./...`，结果上传到 Codecov

## 配置

无需配置文件。库完全通过 Go API 配置：

```go
t.Load("path/to/dir")                    // 加载翻译文件
t.Bind("my-domain", "path")              // 绑定文本域
t.SetLocale("zh_CN")                     // 设置语言（空字符串=系统默认）
t.SetDomain("my-domain")                 // 设置当前文本域
t.SetSourceCodeLocale("en_US")           // 设置源码语言（默认 en_US）
t.GetUserLang(request)                   // 从 HTTP 请求检测语言
```

## xtemplate 用法

```
xtemplate -i input-pattern -k keywords [-f functions] [-o output]
  -d          调试模式
  -f string   模板函数名（逗号分隔，通常可自动检测）
  -i string   输入文件 glob 模式（必填）
  -k string   关键字定义，如 gettext;T:1;N:1,2;X:1c,2;XN:1c,2,3
  -left       左分隔符（默认 "{{"）
  -right      右分隔符（默认 "}}"）
  -o string   输出文件，- 表示 stdout
```

## 安全

- 本库不涉及用户输入的直接执行或网络通信
- 翻译文件加载时信任文件内容，应确保翻译文件来源可信
- `f.Format()` 防止了格式化字符串注入导致的 panic
