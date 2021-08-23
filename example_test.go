package t_test

import (
	"fmt"

	"github.com/youthlin/t"
)

func Example_init() {
	t.Load("testdata")
	fmt.Println(t.T("Hello, World"))
	// Output:
	// 你好，世界
}
