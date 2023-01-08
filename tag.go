package t

import (
	"sync"

	"golang.org/x/text/language"
)

var cachedTag sync.Map

func Tag(locale string) language.Tag {
	if v, ok := cachedTag.Load(locale); ok {
		return v.(language.Tag)
	}
	tag := language.Make(locale)
	cachedTag.Store(locale, tag)
	return tag
}

func Tags(locales []string) (tags []language.Tag) {
	for _, locale := range locales {
		tags = append(tags, Tag(locale))
	}
	return
}
