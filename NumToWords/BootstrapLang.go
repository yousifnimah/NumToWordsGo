package NumToWords

import "github.com/yousifnimah/NumToWordsGo/NumToWords/locales"

func (Entry *Entry) BootstrapLanguage() {
	Entry.LocalizedEntity = Languages[Entry.Language]
}

var Languages = map[string]locales.Lang{
	"en": locales.EN,
	"ar": locales.AR,
}
