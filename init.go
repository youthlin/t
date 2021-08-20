package t

// global is a global translations instance
var global = NewTranslations()

func Global() *Translations {
	return global.copy()
}

func ResetGlobal() {
	global = NewTranslations()
}
