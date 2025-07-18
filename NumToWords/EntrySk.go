package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntrySk struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntrySk) Translate(Input int) string {
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

func (Entry EntrySk) handleUnits(Input int) string {
	return Entry.LocalizedEntity.Units[Input]
}

func (Entry EntrySk) handleTeens(Input int) string {
	return Entry.LocalizedEntity.Teens[Input-10]
}

func (Entry EntrySk) handleTens(Input int) string {
	tensPart := Input / 10
	unitsPart := Input % 10
	tens := Entry.LocalizedEntity.Tens[tensPart]
	if unitsPart == 0 {
		return tens
	}
	units := Entry.Translate(unitsPart)
	return fmt.Sprintf("%s %s", tens, units)
}

func (Entry EntrySk) handleHundreds(Input int) string {
	hundredsPart := Input / 100
	remainder := Input % 100
	hundreds := Entry.LocalizedEntity.Hundreds[hundredsPart]
	if remainder == 0 {
		return hundreds
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s", hundreds, remStr)
}

func (Entry EntrySk) handleThousands(Input int) string {
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

func (Entry EntrySk) getThousandPlural(n int) string {
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

func (Entry EntrySk) handleMillions(Input int) string {
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

func (Entry EntrySk) getMillionPlural(n int) string {
	if n == 1 {
		return "milión"
	}
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 14 {
		return "miliónov"
	}
	last := n % 10
	if last >= 2 && last <= 4 {
		return "milióny"
	}
	return "miliónov"
}

func (Entry EntrySk) handleBillions(Input int) string {
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

func (Entry EntrySk) getBillionPlural(n int) string {
	if n == 1 {
		return "miliarda"
	}
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 14 {
		return "miliárd"
	}
	last := n % 10
	if last >= 2 && last <= 4 {
		return "miliardy"
	}
	return "miliárd"
}

func (Entry EntrySk) handleTrillions(Input int) string {
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

func (Entry EntrySk) getTrillionPlural(n int) string {
	if n == 1 {
		return "bilión"
	}
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 14 {
		return "biliónov"
	}
	last := n % 10
	if last >= 2 && last <= 4 {
		return "bilióny"
	}
	return "biliónov"
}

func (Entry EntrySk) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
