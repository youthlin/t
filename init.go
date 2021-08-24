package t

import (
	"io/fs"
)

var global = NewTranslations()

func Global() *Translations {
	return global
}
func SetGlobal(g *Translations) {
	global = g
}

// Load load translation from path to current domain
func Load(path string) {
	LoadFS(asFS(path))
}

// LoadFS load translation from file system to current domain
func LoadFS(fsys fs.FS) {
	BindFS(Domain(), fsys)
}

// Bind bind translation from path to specified domain
func Bind(domain, path string) {
	BindFS(domain, asFS(path))
}

// BindFS bind translation from file system to specified domain
func BindFS(domain string, fsys fs.FS) {
	global.BindFS(domain, fsys)
}

// Locale return current locale(it may not be used locale)
func Locale() string {
	return global.Locale()
}

// SetLocale set user language
func SetLocale(locale string) {
	global.SetLocale(locale)
}

// SourceCodeLocale 返回源代码中使用的语言
func SourceCodeLocale() string {
	return global.SourceCodeLocale()
}

// SetSourceCodeLocale 设置源代码的语言
func SetSourceCodeLocale(locale string) {
	global.SetSourceCodeLocale(locale)
}

// UsedLocale return the actually used locale
func UsedLocale() string {
	return global.UsedLocale()
}

// Domain return current domain
func Domain() string {
	return global.Domain()
}

// SetDomain set current domain
func SetDomain(domain string) {
	global.SetDomain(domain)
}

// HasDomain return if we have loaded the specified domain
func HasDomain(domain string) bool {
	return global.HasDomain(domain)
}

// Domains return all loaded domains
func Domains() []string {
	return global.Domains()
}

// T: shor name of gettext
func T(msgID string, args ...interface{}) string {
	return global.T(msgID, args...)
}
