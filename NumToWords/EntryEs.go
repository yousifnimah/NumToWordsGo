package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryEs struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntryEs) Translate(Input int) string {
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
	return strings.TrimSpace(result)
}

func (Entry EntryEs) handleUnits(Input int) string {
	return Entry.LocalizedEntity.Units[Input]
}

func (Entry EntryEs) handleTeens(Input int) string {
	return Entry.LocalizedEntity.Teens[Input-10]
}

func (Entry EntryEs) handleTens(Input int) string {
	tensPart := Input / 10
	unitsPart := Input % 10

	if tensPart == 2 {
		if unitsPart == 0 {
			return "veinte"
		}
		return fmt.Sprintf("veinti%s", Entry.handleUnits(unitsPart))
	}

	tens := Entry.LocalizedEntity.Tens[tensPart]
	if unitsPart == 0 {
		return tens
	}
	units := Entry.handleUnits(unitsPart)
	
	if tensPart >= 3 {
		return fmt.Sprintf("%s y %s", tens, units)
	}
	return fmt.Sprintf("%s%s", tens, units)
}

func (Entry EntryEs) handleHundreds(Input int) string {
	hundredsPart := Input / 100
	remainder := Input % 100

	var hundreds string
	if hundredsPart == 1 {
		if remainder == 0 {
			hundreds = "cien"
		} else {
			hundreds = "ciento"
		}
	} else {
		hundreds = Entry.LocalizedEntity.Hundreds[hundredsPart]
	}

	if remainder == 0 {
		return hundreds
	}
	return fmt.Sprintf("%s %s", hundreds, Entry.Translate(remainder))
}

func (Entry EntryEs) handleThousands(Input int) string {
	thousands := Input / 1000
	remainder := Input % 1000

	var thousandsStr string
	if thousands == 1 {
		thousandsStr = "mil"
	} else {
		thousandsStr = fmt.Sprintf("%s mil", Entry.Translate(thousands))
	}

	if remainder == 0 {
		return thousandsStr
	}
	return fmt.Sprintf("%s %s", thousandsStr, Entry.Translate(remainder))
}

func (Entry EntryEs) handleMillions(Input int) string {
	millions := Input / 1000000
	remainder := Input % 1000000

	millionsStr := Entry.Translate(millions)
	plural := "millón"
	if millions != 1 {
		plural = "millones"
		if millionsStr == "uno" {
			millionsStr = "un"
		}
	} else {
		millionsStr = "un"
	}

	if remainder == 0 {
		return fmt.Sprintf("%s %s", millionsStr, plural)
	}
	return fmt.Sprintf("%s %s %s", millionsStr, plural, Entry.Translate(remainder))
}

func (Entry EntryEs) handleBillions(Input int) string {
	billions := Input / 1000000000
	remainder := Input % 1000000000

	billionsStr := Entry.Translate(billions)
	billionWord := "mil millones"

	if billions == 1 {
		return fmt.Sprintf("%s %s", billionWord, Entry.Translate(remainder))
	}

	result := fmt.Sprintf("%s %s", billionsStr, billionWord)
	if remainder == 0 {
		return result
	}
	return fmt.Sprintf("%s %s", result, Entry.Translate(remainder))
}

func (Entry EntryEs) handleTrillions(Input int) string {
	trillions := int(int64(Input) / 1000000000000)
	remainder := int(int64(Input) % 1000000000000)

	trillionsStr := Entry.Translate(trillions)
	plural := "billón"
	if trillions != 1 {
		plural = "billones"
		if trillionsStr == "uno" {
			trillionsStr = "un"
		}
	} else {
		trillionsStr = "un"
	}

	if remainder == 0 {
		return fmt.Sprintf("%s %s", trillionsStr, plural)
	}
	return fmt.Sprintf("%s %s %s", trillionsStr, plural, Entry.Translate(remainder))
}

func (Entry EntryEs) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
