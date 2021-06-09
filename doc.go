// package t
//
// C syntax, gettext: (https://www.gnu.org/software/gettext/manual/html_node/Triggering.html#Triggering)
//
// 	setlocale (LC_ALL, "");	// use the system locale
//  setlocale(LC_MESSAGES, ""); // or only set LC_MESSAGES
// 	bindtextdomain (PACKAGE, LOCALEDIR);
// 	textdomain (PACKAGE);
//
// while using this package:
//
// 	// 1 bind(search po file)
// 	BindTextDomain("my-name","/path/to/po_or_mo/dir")
//
// 	// BindDefaultDomain("/path")
// 	// 2 set current domain
// 	TextDomain("my-name")
//
// 	// BindTextDomain("my-domain", "path/to/dir/or/name.po")
// 	// langs := SupportLang(domain)
// 	// supported := []language.Tag{ language.Make(lang) }
// 	// matcher := language.NewMatcher(supported)
// 	// bestMatch, index, confidence := matcher.Match("<user-accept>"...)
// 	// SetUserLang(langs[index])
//
// 	// 3 set user language
// 	SetLocale("") // set locale, if empty, use system default
//
// 	// 4 use T/N/X/XN to gettext
// 	fmt.Println(T("hello, world"))
// 	// plurals: N/XN/DN/DXN, the argument n is used to choose plural form
// 	// if you want format the number n, you should pass it to the additional args
// 	fmt.Println(N("One apple", "%d apples", 1, 1))   // One apple
// 	fmt.Println(N("One apple", "%d apples", 2, 2))   // 2 apples
// 	fmt.Println(N("One apple", "%d apples", 2))      // %d apples
// 	fmt.Println(N("One apple", "%d apples", 2, 200)) // 200 apples
// 	fmt.Println(N("One apple", "%d apples", 1, 200)) // One apple
package t

// global is a global translations instance
var global = NewTranslations()
