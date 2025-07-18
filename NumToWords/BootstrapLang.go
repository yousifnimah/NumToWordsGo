package NumToWords

import "github.com/yousifnimah/NumToWordsGo/NumToWords/locales"

func (Entry *EntryEn) BootstrapLanguage() {
	Entry.Language = "en"
	Entry.LocalizedEntity = locales.EN
}

func (Entry *EntryAr) BootstrapLanguage() {
	Entry.Language = "ar"
	Entry.LocalizedEntity = locales.AR
}

func (Entry *EntryCs) BootstrapLanguage() {
	Entry.Language = "cs"
	Entry.LocalizedEntity = locales.CS
}

func (Entry *EntryDe) BootstrapLanguage() {
	Entry.Language = "de"
	Entry.LocalizedEntity = locales.DE
}

func (Entry *EntryEs) BootstrapLanguage() {
	Entry.Language = "es"
	Entry.LocalizedEntity = locales.ES
}

func (Entry *EntryFr) BootstrapLanguage() {
	Entry.Language = "fr"
	Entry.LocalizedEntity = locales.FR
}

func (Entry *EntryHu) BootstrapLanguage() {
	Entry.Language = "hu"
	Entry.LocalizedEntity = locales.HU
}

func (Entry *EntryPl) BootstrapLanguage() {
	Entry.Language = "pl"
	Entry.LocalizedEntity = locales.PL
}

func (Entry *EntryRu) BootstrapLanguage() {
	Entry.Language = "ru"
	Entry.LocalizedEntity = locales.RU
}

func (Entry *EntrySk) BootstrapLanguage() {
	Entry.Language = "sk"
	Entry.LocalizedEntity = locales.SK
}

func (Entry *EntryUk) BootstrapLanguage() {
	Entry.Language = "uk"
	Entry.LocalizedEntity = locales.UK
}
