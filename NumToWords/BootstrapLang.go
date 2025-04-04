package NumToWords

import "github.com/neurlang/NumToWordsGo/NumToWords/locales"

func (Entry *EntryEn) BootstrapLanguage() {
	Entry.Language = "en"
	Entry.LocalizedEntity = locales.EN
}

func (Entry *EntryAr) BootstrapLanguage() {
	Entry.Language = "ar"
	Entry.LocalizedEntity = locales.AR
}
