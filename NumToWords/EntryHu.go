package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryHu struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntryHu) Translate(Input int) string {
	if Input == 0 {
		return Entry.ZeroResponse()
	}
	var result string
	switch {
	case Input < 10:
		result = Entry.handleUnits(Input)
	case Input < 20:
		result = Entry.handleTeens(Input)
	case Input < 100:
		result = Entry.handleTens(Input)
	case Input < 1000:
		result = Entry.handleHundreds(Input)
	case Input < 1000000:
		result = Entry.handleThousands(Input)
	case Input < 1000000000:
		result = Entry.handleMillions(Input)
	case int64(Input) < 1000000000000:
		result = Entry.handleBillions(Input)
	default:
		result = Entry.handleTrillions(Input)
	}
	return strings.ReplaceAll(strings.TrimSpace(result), " ", "-")
}

func (Entry EntryHu) handleUnits(Input int) string {
	return Entry.LocalizedEntity.Units[Input]
}

func (Entry EntryHu) handleTeens(Input int) string {
	return Entry.LocalizedEntity.Teens[Input-10]
}

func (Entry EntryHu) handleTens(Input int) string {
	tensPart := Input / 10
	unitsPart := Input % 10
	tens := Entry.LocalizedEntity.Tens[tensPart]
	
	if unitsPart == 0 {
		return tens
	}
	units := Entry.Translate(unitsPart)
	return fmt.Sprintf("%s%s", tens, units)
}

func (Entry EntryHu) handleHundreds(Input int) string {
	hundredsPart := Input / 100
	remainder := Input % 100
	hundreds := Entry.LocalizedEntity.Hundreds[hundredsPart]
	
	if remainder == 0 {
		return hundreds
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s%s", hundreds, remStr)
}

func (Entry EntryHu) handleThousands(Input int) string {
	thousands := Input / 1000
	remainder := Input % 1000

	thousandsStr := Entry.getMultiplierForm(Entry.Translate(thousands))
	
	if thousands == 1 {
		if remainder == 0 {
			return Entry.LocalizedEntity.Thousand
		}
		return fmt.Sprintf("%s-%s", Entry.LocalizedEntity.Thousand, Entry.Translate(remainder))
	}

	if remainder == 0 {
		return fmt.Sprintf("%s%s", thousandsStr, Entry.LocalizedEntity.Thousand)
	}
	return fmt.Sprintf("%s%s-%s", thousandsStr, Entry.LocalizedEntity.Thousand, Entry.Translate(remainder))
}

func (Entry EntryHu) handleMillions(Input int) string {
	millions := Input / 1000000
	remainder := Input % 1000000

	millionsStr := Entry.getMultiplierForm(Entry.Translate(millions))
	
	if millions == 1 {
		if remainder == 0 {
			return Entry.LocalizedEntity.Million
		}
		return fmt.Sprintf("%s-%s", Entry.LocalizedEntity.Million, Entry.Translate(remainder))
	}

	if remainder == 0 {
		return fmt.Sprintf("%s%s", millionsStr, Entry.LocalizedEntity.Million)
	}
	return fmt.Sprintf("%s%s-%s", millionsStr, Entry.LocalizedEntity.Million, Entry.Translate(remainder))
}

func (Entry EntryHu) handleBillions(Input int) string {
	billions := Input / 1000000000
	remainder := Input % 1000000000

	billionsStr := Entry.getMultiplierForm(Entry.Translate(billions))
	
	if billions == 1 {
		if remainder == 0 {
			return Entry.LocalizedEntity.Billion
		}
		return fmt.Sprintf("%s-%s", Entry.LocalizedEntity.Billion, Entry.Translate(remainder))
	}

	if remainder == 0 {
		return fmt.Sprintf("%s%s", billionsStr, Entry.LocalizedEntity.Billion)
	}
	return fmt.Sprintf("%s%s-%s", billionsStr, Entry.LocalizedEntity.Billion, Entry.Translate(remainder))
}

func (Entry EntryHu) handleTrillions(Input int) string {
	trillions := int(int64(Input) / 1000000000000)
	remainder := int(int64(Input) % 1000000000000)

	trillionsStr := Entry.getMultiplierForm(Entry.Translate(trillions))
	
	if trillions == 1 {
		if remainder == 0 {
			return Entry.LocalizedEntity.Trillion
		}
		return fmt.Sprintf("%s-%s", Entry.LocalizedEntity.Trillion, Entry.Translate(remainder))
	}

	if remainder == 0 {
		return fmt.Sprintf("%s%s", trillionsStr, Entry.LocalizedEntity.Trillion)
	}
	return fmt.Sprintf("%s%s-%s", trillionsStr, Entry.LocalizedEntity.Trillion, Entry.Translate(remainder))
}

func (Entry EntryHu) getMultiplierForm(s string) string {
	return strings.ReplaceAll(s, "kettő", "két")
}

func (Entry EntryHu) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
