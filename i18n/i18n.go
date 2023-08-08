package i18n

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry/jibber_jabber"
	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

type Locale struct {
	// Embedding methods from gotext's Locale
	*gotext.Locale

	// Language code of this locale
	lang string
}

var (
	locales       map[string]*Locale
	emptyLocale   = Locale{Locale: &gotext.Locale{}}
	defaultDomain string
	defaultLang   string

	// DefaultLocaleRoot is the root dir of locales, usually is
	// "{Prefix}/share/locale".
	DefaultLocaleRoot = "/opt/go-i18n/share/locale"
	localeRoot        string
)

// Setup rootDir and domain for locale, and will return
// locale for default language.
func Setup(rootDir, domain string) *Locale {
	if rootDir != "" {
		localeRoot = rootDir
	} else {
		localeRoot = DefaultLocaleRoot
	}
	defaultDomain = domain
	return GetLocale("")
}

// GetLocale returns locale for the specific language.
func GetLocale(lang string) *Locale {
	if defaultDomain == "" || localeRoot == "" {
		fmt.Fprintln(os.Stderr, "ERROR: has not run i18n.Setup() yet")
		return &emptyLocale
	}
	if lang == "" {
		lang = defaultLang
	} else {
		lang = strings.Replace(lang, "-", "_", -1)
	}
	lang = gotext.SimplifiedLocale(lang)
	if l, ok := locales[lang]; ok {
		return l
	}
	l := gotext.NewLocale(localeRoot, lang)
	l.AddDomain(defaultDomain)
	locales[lang] = &Locale{Locale: l, lang: lang}
	return locales[lang]
}

func (v Locale) Lang() string {
	return v.lang
}

func (v Locale) LocaleName() string {
	tag := language.Make(v.lang)
	return display.Self.Name(tag)
}

/*
 * Mark i18n strings using the following wrapper:
 *
 * L_(), L_D(), L_C(), L_DC(),
 * Q_(), Q_D(), Q_C(), Q_DC(),
 * and N_().
 */
func (v Locale) L_(str string) string {
	return v.Get(str)
}

func (v Locale) L_D(dom, str string) string {
	return v.GetD(dom, str)
}

func (v Locale) L_C(str, ctx string) string {
	return v.GetC(str, ctx)
}

func (v Locale) L_DC(dom, str, ctx string) string {
	return v.GetDC(dom, str, ctx)
}

func (v Locale) Q_(str, plural string, n int) string {
	return v.GetN(str, plural, n)
}

func (v Locale) Q_D(dom, str, plural string, n int) string {
	return v.GetND(dom, str, plural, n)
}

func (v Locale) Q_C(str, plural string, n int, ctx string) string {
	return v.GetNC(str, plural, n, ctx)
}

func (v Locale) Q_DC(dom, str, plural string, n int, ctx string) string {
	return v.GetNDC(dom, str, plural, n, ctx)
}

/*
 * Mark for translation, but not expand immediately. Usually save the
 * raw string in a variable, and expand the variable in runtime. E.g.
 *
 *     prompt := []string { i18n.N_("prompt1"), i18n.N_("prompt2"), ... }
 *     fmt.Print( i18n.L_(prompt[1]) )
 */
func (v Locale) N_(str string) string {
	return str
}

func init() {
	var err error

	locales = make(map[string]*Locale)
	defaultLang, err = jibber_jabber.DetectIETF()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: cannot detect language: %s\n", err)
		defaultLang = "en_US"
	}
	defaultLang = strings.Replace(defaultLang, "-", "_", -1)
}
