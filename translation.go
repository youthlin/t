package t

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/youthlin/t/f"
	"github.com/youthlin/t/po"
)

// Translations holds a map of domain to Translation
type Translations struct {
	userLang      string
	currentDomain string
	trMap         map[string]*Translation // key is domain
}

func NewTranslations() *Translations {
	return &Translations{currentDomain: DefaultDomain, trMap: make(map[string]*Translation)}
}

func (ts *Translations) SetUserLang(lang string) {
	ts.userLang = lang
}
func (ts *Translations) UserLang() string {
	return ts.userLang
}

// Bind search .po/.mo file in path, and bind the result with domain
func (ts *Translations) Bind(domain, path string) {
	tr := NewTranslation(domain)
	tr.Bind(path)
	if len(tr.Files) > 0 {
		ts.trMap[domain] = tr
	}
}

// TextDomain if domain exists, set it as current domain, else set current domain to DefaultDomain
func (ts *Translations) TextDomain(domain string) string {
	if _, ok := ts.trMap[domain]; ok {
		return ts.setDomain(domain)
	}
	return ts.setDomain(DefaultDomain)
}
func (ts *Translations) setDomain(domain string) string {
	ts.currentDomain = domain
	return domain
}

func (ts *Translations) SupportLangs(domain string) (langs []string) {
	if tr, ok := ts.trMap[domain]; ok {
		langs = append(langs, tr.Langs...)
	}
	return
}

func (ts *Translations) Add(tr *Translation) {
	if ts.trMap == nil {
		ts.trMap = make(map[string]*Translation)
	}
	ts.trMap[tr.Domain] = tr
}
func (ts *Translations) Get(domain string) (*Translation, bool) {
	tr, ok := ts.trMap[domain]
	return tr, ok
}
func (ts *Translations) GetOrNoop(domain string) *Translation {
	tr, ok := ts.Get(domain)
	if !ok {
		return trNoop
	}
	return tr
}

func (ts *Translations) DLT(domain, lang, msgID string, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LT(lang, msgID, args...)
}
func (ts *Translations) DLN(domain, lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LN(lang, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) DLN64(domain, lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LN64(lang, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) DLX(domain, lang, msgCtxt, msgID string, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LX(lang, msgCtxt, msgID, args...)
}
func (ts *Translations) DLXN(domain, lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LXN(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) DLXN64(domain, lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LXN64(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}

// Translation holds a map of lang to po/mo file
type Translation struct {
	Domain string
	Files  map[string]File
	Langs  []string
}

var trNoop = &Translation{}

func NewTranslation(domain string, languages ...File) *Translation {
	tr := Translation{Domain: domain, Files: map[string]File{}}
	for _, file := range languages {
		tr.Add(file)
	}
	return &tr
}

// Add add a po file to current Translation
// if the language of this po file is already exist, then ignore this po file and return false
func (tr *Translation) Add(poFile File) bool {
	lang := poFile.Lang()
	if _, ok := tr.Files[lang]; ok {
		return false
	}
	tr.Files[lang] = poFile
	tr.Langs = append(tr.Langs, lang)
	return true
}

// AddOrReplace if the language of file exist, then replace
func (tr *Translation) AddOrReplace(poFile File) {
	lang := poFile.Lang()
	if _, ok := tr.Files[lang]; !ok {
		tr.Langs = append(tr.Langs, lang)
	}
	tr.Files[lang] = poFile
}

func (tr *Translation) Bind(path string) {
	of, err := os.Open(path)
	if err != nil {
		return
	}
	fInfo, err := of.Stat()
	if err != nil {
		return
	}
	if fInfo.IsDir() {
		sub, err := of.Readdir(0)
		if err != nil {
			return
		}
		for _, v := range sub {
			tr.Bind(filepath.Join(path, v.Name()))
		}
	} else {
		tr.AddFile(of)
	}
}
func (tr *Translation) AddFile(file *os.File) {
	if strings.HasSuffix(file.Name(), ".po") {
		b, err := io.ReadAll(file)
		if err != nil {
			return
		}
		poFile, err := po.Parse(string(b))
		if err != nil {
			return
		}
		tr.Add(poFile)
	}
}

func (tr *Translation) LT(lang, msgID string, args ...interface{}) string {
	file, ok := tr.Files[lang]
	if !ok {
		return f.Format(msgID, args...)
	}
	return file.T(msgID, args...)
}
func (tr *Translation) LN(lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	file, ok := tr.Files[lang]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, int64(n), args...)
	}
	return file.N(msgID, msgIDPlural, n, args...)
}
func (tr *Translation) LN64(lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	file, ok := tr.Files[lang]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	return file.N64(msgID, msgIDPlural, n, args...)
}
func (tr *Translation) LX(lang, msgCtxt, msgID string, args ...interface{}) string {
	file, ok := tr.Files[lang]
	if !ok {
		return f.Format(msgID, args...)
	}
	return file.X(msgCtxt, msgID, args...)
}
func (tr *Translation) LXN(lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	file, ok := tr.Files[lang]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, int64(n), args...)
	}
	return file.XN(msgCtxt, msgID, msgIDPlural, n, args...)
}
func (tr *Translation) LXN64(lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	file, ok := tr.Files[lang]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, n)
	}
	return file.XN64(msgCtxt, msgID, msgIDPlural, n, args...)
}
