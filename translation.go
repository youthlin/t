package t

import (
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/youthlin/t/f"
	"github.com/youthlin/t/translator"
)

// trNoop is a Noop-Translation 一个直接返回原文的翻译实例
var trNoop = &Translation{}

// Translation has a text domain, holds a map of lang to translator
// 翻译结构体，对应一个文本域，多个翻译文件。每个翻译文件对应一种语言。
type Translation struct {
	domain      string
	langs       []string
	translators map[string]Translator // key is language
}

// NewTranslation return a new Translation
func NewTranslation(domain string, translators ...Translator) *Translation {
	tr := Translation{domain: domain, translators: map[string]Translator{}}
	for _, translator := range translators {
		tr.Add(translator)
	}
	return &tr
}

// Add add a translator to current Translation
// if the language is already exist, then replace it and return the previous
func (tr *Translation) Add(translator Translator) Translator {
	lang := translator.Lang()
	pre, ok := tr.translators[lang]
	if lang == "" {
		return pre
	}
	if ok {
		tr.translators[lang] = translator
		return pre
	}
	tr.translators[lang] = translator
	tr.langs = append(tr.langs, lang)
	return nil
}

// AddOrReplace if the language exist, then replace
func (tr *Translation) AddOrReplace(translator Translator) {
	lang := translator.Lang()
	if _, ok := tr.translators[lang]; !ok {
		tr.langs = append(tr.langs, lang)
	}
	tr.translators[lang] = translator
}

// Load load a translator from path
func (tr *Translation) Load(path string) {
	tr.LoadFS(os.DirFS(path))
}

// LoadFS load a translator from file system
func (tr *Translation) LoadFS(f fs.FS) {
	fn := func(exts ...string) func(path string, d fs.DirEntry, err error) error {
		return func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d != nil && !d.IsDir() {
				var shouldVisit = false
				for _, ext := range exts {
					if strings.HasSuffix(path, ext) {
						shouldVisit = true
						break
					}
				}
				if shouldVisit {
					of, err := f.Open(path)
					if err == nil {
						defer of.Close()
						tr.LoadFile(of)
					}
				}
			}
			return nil
		}
	}
	fs.WalkDir(f, ".", fn(".po", ".mo"))
}

// LoadFile load a translator from a file
func (tr *Translation) LoadFile(file fs.File) error {
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	fileName := fi.Name()
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if strings.HasSuffix(fileName, ".po") {
		err = tr.LoadPo(content)
	} else if strings.HasSuffix(fileName, ".mo") {
		err = tr.LoadMo(content)
	}
	return err
}

func (tr *Translation) LoadPo(content []byte) error {
	poFile, err := translator.ReadPo(content)
	if err != nil {
		return err
	}
	tr.Add(poFile)
	return nil
}

func (tr *Translation) LoadMo(content []byte) error {
	moFile, err := translator.ReadMo(content)
	if err != nil {
		return err
	}
	tr.Add(moFile)
	return nil
}

func (tr *Translation) LT(lang, msgID string, args ...interface{}) string {
	translator, ok := tr.translators[lang]
	if !ok {
		return f.Format(msgID, args...)
	}
	return translator.T(msgID, args...)
}

func (tr *Translation) LN(lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	translator, ok := tr.translators[lang]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, int64(n), args...)
	}
	return translator.N(msgID, msgIDPlural, n, args...)
}

func (tr *Translation) LN64(lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	translator, ok := tr.translators[lang]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	return translator.N64(msgID, msgIDPlural, n, args...)
}

func (tr *Translation) LX(lang, msgCtxt, msgID string, args ...interface{}) string {
	translator, ok := tr.translators[lang]
	if !ok {
		return f.Format(msgID, args...)
	}
	return translator.X(msgCtxt, msgID, args...)
}

func (tr *Translation) LXN(lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	translator, ok := tr.translators[lang]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, int64(n), args...)
	}
	return translator.XN(msgCtxt, msgID, msgIDPlural, n, args...)
}

func (tr *Translation) LXN64(lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	translator, ok := tr.translators[lang]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	return translator.XN64(msgCtxt, msgID, msgIDPlural, n, args...)
}
