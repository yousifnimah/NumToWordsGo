package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryDe struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntryDe) Translate(Input int) string {
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

func (Entry EntryDe) handleUnits(Input int) string {
	return Entry.LocalizedEntity.Units[Input]
}

func (Entry EntryDe) handleTeens(Input int) string {
	return Entry.LocalizedEntity.Teens[Input-10]
}

func (Entry EntryDe) handleTens(Input int) string {
	tensPart := Input / 10
	unitsPart := Input % 10
	tens := Entry.LocalizedEntity.Tens[tensPart]
	if unitsPart == 0 {
		return tens
	}
	units := Entry.Translate(unitsPart)
	if units == "eins" {
		units = "ein"
	}
	return fmt.Sprintf("%sund%s", units, tens)
}

func (Entry EntryDe) handleHundreds(Input int) string {
	hundredsPart := Input / 100
	remainder := Input % 100
	hundreds := Entry.LocalizedEntity.Hundreds[hundredsPart]
	if remainder == 0 {
		return hundreds
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s%s", hundreds, remStr)
}

func (Entry EntryDe) handleThousands(Input int) string {
	thousands := Input / 1000
	remainder := Input % 1000

	var thousandsPart string
	if thousands == 1 {
		thousandsPart = "eintausend"
	} else if thousands > 1 {
		thousandsStr := Entry.Translate(thousands)
		thousandsStr = strings.ReplaceAll(thousandsStr, "eins", "ein")
		thousandsPart = fmt.Sprintf("%stausend", thousandsStr)
	}

	if remainder == 0 {
		return thousandsPart
	}

	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s%s", thousandsPart, remStr)
}

func (Entry EntryDe) handleMillions(Input int) string {
	millions := Input / 1000000
	remainder := Input % 1000000

	var millionPart string
	if millions == 1 {
		millionPart = "eine Million"
	} else if millions > 1 {
		millionStr := Entry.Translate(millions)
		millionStr = strings.ReplaceAll(millionStr, "eins", "ein")
		plural := Entry.getMillionPlural(millions)
		millionPart = fmt.Sprintf("%s %s", millionStr, plural)
	}

	if remainder == 0 {
		return millionPart
	}

	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s", millionPart, remStr)
}

func (Entry EntryDe) getMillionPlural(n int) string {
	if n == 1 {
		return "Million"
	}
	return "Millionen"
}

func (Entry EntryDe) handleBillions(Input int) string {
	billions := Input / 1000000000
	remainder := Input % 1000000000

	var billionPart string
	if billions == 1 {
		billionPart = "eine Milliarde"
	} else if billions > 1 {
		billionStr := Entry.Translate(billions)
		billionStr = strings.ReplaceAll(billionStr, "eins", "ein")
		plural := Entry.getBillionPlural(billions)
		billionPart = fmt.Sprintf("%s %s", billionStr, plural)
	}

	if remainder == 0 {
		return billionPart
	}

	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s", billionPart, remStr)
}

func (Entry EntryDe) getBillionPlural(n int) string {
	if n == 1 {
		return "Milliarde"
	}
	return "Milliarden"
}

func (Entry EntryDe) handleTrillions(Input int) string {
	trillions := int(int64(Input) / 1000000000000)
	remainder := int(int64(Input) % 1000000000000)

	var trillionPart string
	if trillions == 1 {
		trillionPart = "eine Billion"
	} else if trillions > 1 {
		trillionStr := Entry.Translate(trillions)
		trillionStr = strings.ReplaceAll(trillionStr, "eins", "ein")
		plural := Entry.getTrillionPlural(trillions)
		trillionPart = fmt.Sprintf("%s %s", trillionStr, plural)
	}

	if remainder == 0 {
		return trillionPart
	}

	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s", trillionPart, remStr)
}

func (Entry EntryDe) getTrillionPlural(n int) string {
	if n == 1 {
		return "Billion"
	}
	return "Billionen"
}

func (Entry EntryDe) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
