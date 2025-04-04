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
	case Input < 1000000: //1000-9999
		{
			Result = Entry.handleThousands(Input)
		}
	case Input < 1000000000: //1000000-9999999
		{
			Result = Entry.handleMillions(Input)
		}
	case int64(Input) < 1000000000000: //1000000000-9999999999
		{
			Result = Entry.handleBillions(Input)
		}
	}
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
	var Seg01 string
	var Seg03 string
	var result string
	And := ""
	if Input/1000 > 2 {
		Seg01 = EntAr.Translate(Input / 1000)
		Seg03 = EntAr.Translate(Input % 1000)
		if Seg01 != "" && Seg03 != "" {
			And = EntAr.LocalizedEntity.And
		}
		result = fmt.Sprintf("%s %s %s %s", Seg01, EntAr.LocalizedEntity.Thousand, And, Seg03)
	} else {
		Seg01 = EntAr.LocalizedEntity.Thousands[(Input / 1000)]
		Seg03 = EntAr.Translate(Input % 1000)
		if Seg01 != "" && Seg03 != "" {
			And = EntAr.LocalizedEntity.And
		}
		result = fmt.Sprintf("%s %s %s", Seg01, And, Seg03)
	}
	return result
}

func (EntAr EntryAr) handleMillions(Input int) string {
	var Seg01 string
	var Seg03 string
	var result string
	And := ""
	if Input/1000000 > 2 {
		Seg01 = EntAr.Translate(Input / 1000000)
		Seg03 = EntAr.Translate(Input % 1000000)
		if Seg01 != "" && Seg03 != "" {
			And = EntAr.LocalizedEntity.And
		}
		result = fmt.Sprintf("%s %s %s %s", Seg01, EntAr.LocalizedEntity.Million, And, Seg03)
	} else {
		Seg01 = EntAr.LocalizedEntity.Millions[(Input / 1000000)]
		Seg03 = EntAr.Translate(Input % 1000000)
		if Seg01 != "" && Seg03 != "" {
			And = EntAr.LocalizedEntity.And
		}
		result = fmt.Sprintf("%s %s %s", Seg01, And, Seg03)
	}
	return result
}

func (EntAr EntryAr) handleBillions(Input int) string {
	var Seg01 string
	var Seg03 string
	var result string
	And := ""
	if Input/1000000000 > 2 {
		Seg01 = EntAr.Translate(Input / 1000000000)
		Seg03 = EntAr.Translate(Input % 1000000000)
		if Seg01 != "" && Seg03 != "" {
			And = EntAr.LocalizedEntity.And
		}
		result = fmt.Sprintf("%s %s %s %s", Seg01, EntAr.LocalizedEntity.Billion, And, Seg03)
	} else {
		Seg01 = EntAr.LocalizedEntity.Billions[(Input / 1000000000)]
		Seg03 = EntAr.Translate(Input % 1000000000)
		if Seg01 != "" && Seg03 != "" {
			And = EntAr.LocalizedEntity.And
		}
		result = fmt.Sprintf("%s %s %s", Seg01, And, Seg03)
	}
	return result
}

func (Entry EntryAr) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
