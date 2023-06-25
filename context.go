package t

import "context"

type (
	typeLang   struct{}
	typeDomain struct{}
)

var (
	keyLang   = typeLang{}
	keyDomain = typeDomain{}
)

func SetCtxLocale(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, keyLang, lang)
}

func SetCtxDomain(ctx context.Context, domain string) context.Context {
	return context.WithValue(ctx, keyDomain, domain)
}

func GetCtxLocale(ctx context.Context) (string, bool) {
	v := ctx.Value(keyLang)
	lang, ok := v.(string)
	return lang, ok
}

func GetCtxDomain(ctx context.Context) (string, bool) {
	v := ctx.Value(keyDomain)
	domain, ok := v.(string)
	return domain, ok
}

func WithContext(ctx context.Context) *Translations {
	t := Global()
	if lang, ok := GetCtxLocale(ctx); ok {
		t = t.L(lang)
	}
	if domain, ok := GetCtxDomain(ctx); ok {
		t = t.D(domain)
	}
	return t
}
