package NumToWords

import (
	"errors"
	"fmt"
)

type NumToWord struct{}

type Entry interface {
	Translate(int) string
	BootstrapLanguage()
	ZeroResponse() string
}

func Convert(Input int, Language string) (string, error) {
	//Check if language doesn't support
	if Language == "" {
		err := errors.New("language is not provided")
		panic(err)
		return "", err
	}
	//Initiate an entry based on the selected language
	Ent := InitEntry(Language)
	//Bootstrap the languages and localized entities
	Ent.BootstrapLanguage()

	if IsZero(Input) {
		//fmt.Println("Input")
		return Ent.ZeroResponse(), nil
	}

	//Translate to words
	return Ent.Translate(Input), nil
}

func InitEntry(Language string) Entry {
	var ent Entry

	switch Language {
	case "ar":
		ent = &EntryAr{}
	case "cs":
		ent = &EntryCs{}
	case "de":
		ent = &EntryDe{}
	case "es":
		ent = &EntryEs{}
	case "fr":
		ent = &EntryFr{}
	case "hu":
		ent = &EntryHu{}
	case "pl":
		ent = &EntryPl{}
	case "ru":
		ent = &EntryRu{}
	case "sk":
		ent = &EntrySk{}
	case "uk":
		ent = &EntryUk{}
	default:
		ent = &EntryEn{}
	}
	return ent
}

func IsZero(Input int) bool {
	return Input == 0
}
