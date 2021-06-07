// package t
//
// C syntax, gettext:
//
//  setlocale(LC_MESSAGE, "");
//  bindtextdomain(PACKAGE, LOCALEDIR);
//  textdomain(PACKAGE);
//
// while using this package:
//
// 	BindDefaultDomain("/path/to/po_or_mo/dir")
// 	// BindTextDomain("my-domain", "path/to/dir/or/name.po")
// 	// langs := SupportLang()
// 	// supported := []language.Tag{ language.Make(lang) }
// 	// matcher := language.NewMatcher(supported)
// 	// bestMatch, index, confidence := matcher.Match("<user-accept>"...)
// 	// SetUserLang(langs[index])
// 	SetUserLang("zh-CN")
// 	fmt.Println(T("hello, world"))
// 	// plurals: N/XN/DN/DXN, the argument n is used to choose plural form
// 	// if you want format the number n, you should pass it to the additional args
// 	fmt.Println(N("One apple", "%d apples", 1, 1))   // One apple
// 	fmt.Println(N("One apple", "%d apples", 2, 2))   // 2 apples
// 	fmt.Println(N("One apple", "%d apples", 2))      // %d apples
// 	fmt.Println(N("One apple", "%d apples", 2, 200)) // 200 apples
// 	fmt.Println(N("One apple", "%d apples", 1, 200)) // One apple
package t

// globalTranslatins is a global translations struct
var globalTranslatins = NewTranslations()
