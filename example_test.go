package t_test

import (
	"embed"
	"fmt"
	"sync"

	"github.com/youthlin/t"
)

var mu sync.Mutex

func markSeq() func() {
	mu.Lock() // 顺序执行
	t.SetGlobal(t.NewTranslations())
	return func() { mu.Unlock() }
}

// Example_init_path Load 绑定 path 到默认文本域
func Example_init_path() {
	defer markSeq()()
	t.Load("testdata")
	// empty means system default locale
	// 设置为使用系统语言这一步骤可以省略，因为初始化时就是使用系统语言
	t.SetLocale("")
	// 为了能在其他环境测试通过，所以指定 zh_CN
	t.SetLocale("zh_CN")
	fmt.Println(t.T("Hello, World"))
	// Output:
	// 你好，世界
}

// Example_init_file Load 绑定 path 到默认文本域, path 可以是单个文件
func Example_init_file() {
	defer markSeq()()
	t.Load("testdata/zh_CN.po")
	// will normalize to ll_CC form => zh_CN
	// 会格式化为 ll_CC 的形式
	t.SetLocale("zh_hans")
	fmt.Println(t.T("Hello, World"))
	// Output:
	// 你好，世界
}

//go:embed testdata
var pathFS embed.FS

//go:embed testdata/zh_CN.po
var fileFS embed.FS

// Example_initFS LoadFS 绑定文件系统到默认文本域
func Example_initFS() {
	defer markSeq()()
	t.LoadFS(pathFS)
	t.SetLocale("zh")
	fmt.Println(t.T("Hello, World"))
	// Output:
	// 你好，世界
}

// Example_init_fileFS embed.FS 可以是单个文件或文件夹
func Example_init_fileFS() {
	defer markSeq()()
	t.LoadFS(fileFS)
	t.SetLocale("zh_hans")
	fmt.Println("Current locale =", t.Locale()) // zh_CN
	fmt.Println(t.T("Hello, World"))

	// 设置不支持的语言，会原样返回
	t.SetLocale("zh_hant")
	fmt.Println("Current locale =", t.Locale()) // zh_TW
	fmt.Println(t.T("Hello, World"))
	// Output:
	// Current locale = zh_CN
	// 你好，世界
	// Current locale = zh_TW
	// Hello, World
}

// Example_locale 语言设置示例
func Example_locale() {
	defer markSeq()()
	t.LoadFS(pathFS)
	t.SetLocale("zh")
	fmt.Println("UsedLocale =", t.UsedLocale()) // zh_CN
	t.SetLocale("zh_TW")
	fmt.Println("UsedLocale =", t.UsedLocale())             // en_US
	fmt.Println("SourceCodeLocale =", t.SourceCodeLocale()) // en_US

	ts := t.NewTranslations()
	ts.SetSourceCodeLocale("zh_CN")
	ts.SetLocale("en_US")
	fmt.Println("empty ts SourceCodeLocale =", ts.SourceCodeLocale()) // zh_CN
	fmt.Println("empty ts Locale =", ts.Locale())                     // en_US
	fmt.Println("empty ts UsedLocale =", ts.UsedLocale())             // zh_CN
	fmt.Println(ts.T("你好，世界"))
	// Output:
	// UsedLocale = zh_CN
	// UsedLocale = en_US
	// SourceCodeLocale = en_US
	// empty ts SourceCodeLocale = zh_CN
	// empty ts Locale = en_US
	// empty ts UsedLocale = zh_CN
	// 你好，世界
}

// Example_bindDomain 绑定到指定文本域
func Example_bindDomain() {
	defer markSeq()()
	t.SetLocale("zh")
	t.Bind("main", "testdata/zh_CN.po")
	fmt.Println("HasDomain(main) =", t.HasDomain("main"))
	fmt.Println("Domains =", t.Domains())
	fmt.Println(t.T("Hello, World"))
	t.SetDomain("main")
	fmt.Println(t.T("Hello, World"))
	// Output:
	// HasDomain(main) = true
	// Domains = [main]
	// Hello, World
	// 你好，世界
}

func Example_gettext() {
	defer markSeq()()
	t.Load("testdata")
	t.SetLocale("zh_CN")
	fmt.Println(t.T("Hello, World"))                                 // 你好，世界
	fmt.Println(t.T("Hello, %s", "Tom"))                             // 你好，世界
	fmt.Println(t.N("One apple", "%d apples", 1))                    // %d 个苹果
	fmt.Println(t.N("One apple", "%d apples", 1, 1))                 // 1 个苹果
	fmt.Println(t.N("One apple", "%d apples", 2))                    // %d 个苹果
	fmt.Println(t.N("One apple", "%d apples", 2, 2))                 // 2 个苹果
	fmt.Println(t.N64("One apple", "%d apples", 200, 200))           // 200 个苹果
	fmt.Println(t.X("File|", "Open"))                                // 打开文件
	fmt.Println(t.X("Project|", "Open"))                             // 打开工程
	fmt.Println(t.XN("File|", "Open One", "Open %d", 1, 1))          // 打开 1 个文件
	fmt.Println(t.XN("Project|", "Open One", "Open %d", 1))          // 打开 %d 个工程
	fmt.Println(t.XN("Project|", "Open One", "Open %d", 1, 1))       // 打开 1 个工程
	fmt.Println(t.XN64("Project|", "Open One", "Open %d", 100, 100)) // 打开 100 个工程
	// Output:
	// 你好，世界
	// 你好，Tom
	// %d 个苹果
	// 1 个苹果
	// %d 个苹果
	// 2 个苹果
	// 200 个苹果
	// 打开文件
	// 打开工程
	// 打开 1 个文件
	// 打开 %d 个工程
	// 打开 1 个工程
	// 打开 100 个工程
}
