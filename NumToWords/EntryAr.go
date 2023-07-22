package NumToWords

import "github.com/yousifnimah/NumToWordsGo/NumToWords/locales"

type EntryAr struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry *EntryAr) Translate(i int) string {
	//TODO implement me
	panic("implement me")
}

func (Entry EntryAr) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
