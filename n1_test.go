package t_test

import (
	"fmt"

	"github.com/youthlin/t"
)

// ExampleN1 单复数同形简写，无翻译时原样返回
func ExampleN1() {
	defer markSeq()()
	fmt.Println(t.N1("%d 个苹果", 1))
	fmt.Println(t.N1("%d 个苹果", 1, 1))
	fmt.Println(t.N1("%d 个苹果", 2))
	fmt.Println(t.N1("%d 个苹果", 2, 2))
	// Output:
	// %d 个苹果
	// 1 个苹果
	// %d 个苹果
	// 2 个苹果
}

// ExampleN1_withTranslation 加载翻译后 N1 正常工作
func ExampleN1_withTranslation() {
	defer markSeq()()
	t.Load("testdata")
	t.SetLocale("zh_CN")
	// zh_CN.po 中 "One apple" 的翻译是 "%d 个苹果"
	fmt.Println(t.N1("One apple", 1))
	fmt.Println(t.N1("One apple", 1, 1))
	fmt.Println(t.N1("One apple", 2))
	fmt.Println(t.N1("One apple", 2, 2))
	// Output:
	// %d 个苹果
	// 1 个苹果
	// %d 个苹果
	// 2 个苹果
}

// ExampleTranslations_N1 Translations 上的 N1 方法
func ExampleTranslations_N1() {
	defer markSeq()()
	t.Load("testdata")
	ts := t.D("default").L("zh_CN")
	fmt.Println(ts.N1("One apple", 1))
	fmt.Println(ts.N1("One apple", 1, 1))
	fmt.Println(ts.N1("One apple", 3))
	fmt.Println(ts.N1("One apple", 3, 3))
	// Output:
	// %d 个苹果
	// 1 个苹果
	// %d 个苹果
	// 3 个苹果
}
