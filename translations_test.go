package t

import (
	"testing"
)

func TestTranslationsRemoveDomain(t *testing.T) {
	ts := NewTranslations()
	ts.BindFS("main", asFS("testdata"))
	ts.BindFS("extra", asFS("testdata/zh_CN.po"))
	ts.SetDomain("main")

	if !ts.HasDomain("main") || !ts.HasDomain("extra") {
		t.Fatalf("expected domains to be loaded, got %v", ts.Domains())
	}
	if ok := ts.RemoveDomain(""); ok {
		t.Fatal("RemoveDomain(\"\") should return false")
	}
	if ok := ts.RemoveDomain("missing"); ok {
		t.Fatal("RemoveDomain(missing) should return false")
	}
	if ok := ts.RemoveDomain("main"); !ok {
		t.Fatal("RemoveDomain(main) should return true")
	}
	if ts.HasDomain("main") {
		t.Fatal("main domain should be removed")
	}
	if got := ts.Domain(); got != DefaultDomain {
		t.Fatalf("current domain should fall back to %q, got %q", DefaultDomain, got)
	}
	if !ts.HasDomain("extra") {
		t.Fatal("other domains should remain after RemoveDomain")
	}
}

func TestTranslationsClearDomains(t *testing.T) {
	ts := NewTranslations()
	ts.SetLocale("zh_CN")
	ts.SetSourceCodeLocale("zh_CN")
	ts.BindFS("main", asFS("testdata"))
	ts.BindFS("extra", asFS("testdata/zh_CN.po"))
	ts.SetDomain("main")

	ts.ClearDomains()

	if got := len(ts.Domains()); got != 0 {
		t.Fatalf("Domains() len = %d, want 0", got)
	}
	if got := ts.Domain(); got != DefaultDomain {
		t.Fatalf("Domain() = %q, want %q", got, DefaultDomain)
	}
	if got := ts.Locale(); got != "zh_CN" {
		t.Fatalf("Locale() = %q, want zh_CN", got)
	}
	if got := ts.SourceCodeLocale(); got != "zh_CN" {
		t.Fatalf("SourceCodeLocale() = %q, want zh_CN", got)
	}
	if got := ts.T("Hello, World"); got != "Hello, World" {
		t.Fatalf("T() after ClearDomains = %q, want source text", got)
	}
}
