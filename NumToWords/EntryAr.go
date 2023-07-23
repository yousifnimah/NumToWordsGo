package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryAr struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry *EntryAr) Translate(Input int) string {
	Result := ""

	switch {
	case Input < 10: //0-9
		{
			Result = Entry.handleUnits(Input)
		}
	case Input < 20: //10-19
		{
			Result = Entry.handleTeens(Input)
		}
	case Input < 100: //20-99
		{
			Result = Entry.handleTwenties(Input)
		}
	case Input < 1000: //100-999
		{
			Result = Entry.handleHundreds(Input)
		}
	case Input < 100000: //1000-9999
		{
			Result = Entry.handleThousands(Input)
		}
	}

	//if Input < 10 {
	//	result = Entry.handleUnits(Input)
	//} else if Input < 20 {
	//	result = Entry.LocalizedEntity.Teens[Input-10]
	//} else if Input < 100 {
	//	Seg01 := Entry.Translate(Input % 10)
	//	Seg02 := Entry.LocalizedEntity.Tens[Input/10]
	//	if Seg01 != "" && Seg02 != "" {
	//		And = Entry.LocalizedEntity.And
	//	}
	//	result = fmt.Sprintf("%s %s %s", Seg01, And, Seg02)
	//} else if Input < 1000 {
	//	result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/100), Entry.LocalizedEntity.Hundred, Entry.Translate(Input%100))
	//} else if Input < 1000000 {
	//	result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/1000), Entry.LocalizedEntity.Thousand, Entry.Translate(Input%1000))
	//} else if Input < 1000000000 {
	//	result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/1000000), Entry.LocalizedEntity.Million, Entry.Translate(Input%1000000))
	//} else if Input < 10000000000 {
	//	result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/1000000000), Entry.LocalizedEntity.Billion, Entry.Translate(Input%1000000000))
	//} else {
	//	result = fmt.Sprintf("%s %s %s", Entry.Translate(Input/1000000000000), Entry.LocalizedEntity.Trillion, Entry.Translate(Input%1000000000000))
	//}
	return strings.TrimSpace(Result)
}

func (EntAr EntryAr) handleUnits(Input int) string {
	return EntAr.LocalizedEntity.Units[Input]
}

func (EntAr EntryAr) handleTeens(Input int) string {
	return EntAr.LocalizedEntity.Teens[Input-10]
}

func (EntAr EntryAr) handleTwenties(Input int) string {
	And := ""
	Seg01 := EntAr.Translate(Input % 10)
	Seg02 := EntAr.LocalizedEntity.Tens[Input/10]
	if Seg01 != "" && Seg02 != "" {
		And = EntAr.LocalizedEntity.And
	}
	return fmt.Sprintf("%s %s %s", Seg01, And, Seg02)
}

func (EntAr EntryAr) handleHundreds(Input int) string {
	And := ""
	Seg01 := EntAr.LocalizedEntity.Hundreds[(Input / 100)]
	Seg02 := EntAr.Translate(Input % 100)
	if Seg01 != "" && Seg02 != "" {
		And = EntAr.LocalizedEntity.And
	}

	return fmt.Sprintf("%s %s %s", Seg01, And, Seg02)
}

func (EntAr EntryAr) handleThousands(Input int) string {
	And := ""
	Seg01 := EntAr.LocalizedEntity.Thousands[(Input / 1000)]
	Seg03 := EntAr.Translate(Input % 1000)
	if Seg01 != "" && Seg03 != "" {
		And = EntAr.LocalizedEntity.And
	}
	return fmt.Sprintf("%s %s %s", Seg01, And, Seg03)
}

func (Entry EntryAr) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
