package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryCs struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntryCs) Translate(Input int) string {
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

func (Entry EntryCs) handleUnits(Input int) string {
	return Entry.LocalizedEntity.Units[Input]
}

func (Entry EntryCs) handleTeens(Input int) string {
	return Entry.LocalizedEntity.Teens[Input-10]
}

func (Entry EntryCs) handleTens(Input int) string {
	tensPart := Input / 10
	unitsPart := Input % 10
	tens := Entry.LocalizedEntity.Tens[tensPart]
	if unitsPart == 0 {
		return tens
	}
	units := Entry.Translate(unitsPart)
	return fmt.Sprintf("%s %s", tens, units)
}

func (Entry EntryCs) handleHundreds(Input int) string {
	hundredsPart := Input / 100
	remainder := Input % 100
	hundreds := Entry.LocalizedEntity.Hundreds[hundredsPart]
	if remainder == 0 {
		return hundreds
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s", hundreds, remStr)
}

func (Entry EntryCs) handleThousands(Input int) string {
	thousands := Input / 1000
	remainder := Input % 1000
	thousandsStr := Entry.Translate(thousands)
	plural := getThousandPlural(thousands)
	if remainder == 0 {
		return fmt.Sprintf("%s %s", thousandsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", thousandsStr, plural, remStr)
}

func getThousandPlural(n int) string {
	if n == 1 {
		return "tisíc"
	}
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 14 {
		return "tisíc"
	}
	last := n % 10
	if last >= 2 && last <= 4 {
		return "tisíce"
	}
	return "tisíc"
}

func (Entry EntryCs) handleMillions(Input int) string {
	millions := Input / 1000000
	remainder := Input % 1000000
	millionsStr := Entry.Translate(millions)
	plural := Entry.getMillionPlural(millions)
	if remainder == 0 {
		return fmt.Sprintf("%s %s", millionsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", millionsStr, plural, remStr)
}

func (Entry EntryCs) getMillionPlural(n int) string {
	if n == 1 {
		return "milion"
	}
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 14 {
		return "milionů"
	}
	last := n % 10
	if last >= 2 && last <= 4 {
		return "miliony"
	}
	return "milionů"
}

func (Entry EntryCs) handleBillions(Input int) string {
	billions := Input / 1000000000
	remainder := Input % 1000000000
	billionsStr := Entry.Translate(billions)
	plural := Entry.getBillionPlural(billions)
	if remainder == 0 {
		return fmt.Sprintf("%s %s", billionsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", billionsStr, plural, remStr)
}

func (Entry EntryCs) getBillionPlural(n int) string {
	if n == 1 {
		return "miliarda"
	}
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 14 {
		return "miliard"
	}
	last := n % 10
	if last >= 2 && last <= 4 {
		return "miliardy"
	}
	return "miliard"
}

func (Entry EntryCs) handleTrillions(Input int) string {
	trillions := int(int64(Input) / 1000000000000)
	remainder := int(int64(Input) % 1000000000000)
	trillionsStr := Entry.Translate(trillions)
	plural := Entry.getTrillionPlural(trillions)
	if remainder == 0 {
		return fmt.Sprintf("%s %s", trillionsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", trillionsStr, plural, remStr)
}

func (Entry EntryCs) getTrillionPlural(n int) string {
	if n == 1 {
		return "bilion"
	}
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 14 {
		return "bilionů"
	}
	last := n % 10
	if last >= 2 && last <= 4 {
		return "biliony"
	}
	return "bilionů"
}

func (Entry EntryCs) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
