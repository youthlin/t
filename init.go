package t

import (
	"io/fs"
	"os"
)

var global = NewTranslations()

// Load load translations from path to current domain
func Load(path string) {
	LoadFS(os.DirFS(path))
}

// LoadFS load translations from file system to current domain
func LoadFS(fsys fs.FS) {
	BindFS(Domain(), fsys)
}

// Bind load translations from path to specified domain
func Bind(domain, path string) {
	BindFS(domain, os.DirFS(path))
}

// BindFS load translations from file system to specified domain
func BindFS(domain string, fsys fs.FS) {
	global.BindFS(domain, fsys)
}

// SetLocale set user language
func SetLocale(locale string) {
	global.SetLocale(locale)
}

// Domain return current domain
func Domain() string {
	return global.Domain()
}

func SetDomain(domain string) {
	global.SetDomain(domain)
}

func T(msgID string, args ...interface{}) string {
	return global.T(msgID, args...)
}
