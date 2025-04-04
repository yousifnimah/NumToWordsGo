package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryEn struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntryEn) Translate(Input int) string {
	result := ""
	if Input < 10 {
		result = Entry.LocalizedEntity.Units[Input]
	} else if Input < 20 {
		result = Entry.LocalizedEntity.Teens[Input-10]
	} else if Input < 100 {
		result = Entry.LocalizedEntity.Tens[Input/10] + " " + Entry.Translate(Input%10)
	} else if Input < 1000 {
		result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/100), Entry.LocalizedEntity.Hundred, Entry.Translate(Input%100))
	} else if Input < 1000000 {
		result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/1000), Entry.LocalizedEntity.Thousand, Entry.Translate(Input%1000))
	} else if Input < 1000000000 {
		result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/1000000), Entry.LocalizedEntity.Million, Entry.Translate(Input%1000000))
	} else if int64(Input) < 10000000000 {
		result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/1000000000), Entry.LocalizedEntity.Billion, Entry.Translate(Input%1000000000))
	} else {
		result = fmt.Sprintf("%s %s %s", Entry.Translate(int(int64(Input)/1000000000000)), Entry.LocalizedEntity.Trillion, Entry.Translate(int(int64(Input)%1000000000000)))
	}
	return strings.TrimSpace(result)
}

func (Entry EntryEn) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
