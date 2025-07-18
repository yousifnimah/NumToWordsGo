package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryFr struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntryFr) Translate(Input int) string {
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

func (Entry EntryFr) handleUnits(Input int) string {
	return Entry.LocalizedEntity.Units[Input]
}

func (Entry EntryFr) handleTeens(Input int) string {
	return Entry.LocalizedEntity.Teens[Input-10]
}

func (Entry EntryFr) handleTens(Input int) string {
	if Input >= 70 && Input < 80 {
		n := Input - 60
		if n == 10 {
			return "soixante-dix"
		} else if n == 11 {
			return "soixante et onze"
		} else {
			return fmt.Sprintf("soixante-%s", Entry.LocalizedEntity.Teens[n-10])
		}
	} else if Input >= 80 && Input < 90 {
		if Input == 80 {
			return "quatre-vingts"
		}
		unit := Input - 80
		if unit == 1 {
			return "quatre-vingt-un"
		}
		return fmt.Sprintf("quatre-vingt-%s", Entry.handleUnits(unit))
	} else if Input >= 90 && Input < 100 {
		n := Input - 80
		if n == 10 {
			return "quatre-vingt-dix"
		}
		teen := n - 10
		return fmt.Sprintf("quatre-vingt-%s", Entry.LocalizedEntity.Teens[teen])
	} else {
		tensPart := Input / 10
		unitsPart := Input % 10
		tens := Entry.LocalizedEntity.Tens[tensPart]
		if unitsPart == 0 {
			return tens
		}
		units := Entry.handleUnits(unitsPart)
		if unitsPart == 1 && tensPart != 8 && tensPart != 9 {
			return fmt.Sprintf("%s et %s", tens, units)
		}
		return fmt.Sprintf("%s-%s", tens, units)
	}
}

func (Entry EntryFr) handleHundreds(Input int) string {
	hundredsPart := Input / 100
	remainder := Input % 100
	hundreds := Entry.LocalizedEntity.Hundreds[hundredsPart]
	if remainder == 0 {
		return hundreds
	}
	if hundredsPart > 1 && strings.HasSuffix(hundreds, "s") {
		hundreds = hundreds[:len(hundreds)-1]
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s", hundreds, remStr)
}

func (Entry EntryFr) handleThousands(Input int) string {
	thousands := Input / 1000
	remainder := Input % 1000

	var thousandsStr string
	if thousands == 1 {
		thousandsStr = "mille"
	} else {
		thousandsStr = Entry.Translate(thousands) + " mille"
	}

	if remainder == 0 {
		return thousandsStr
	}
	return fmt.Sprintf("%s %s", thousandsStr, Entry.Translate(remainder))
}

func (Entry EntryFr) handleMillions(Input int) string {
	millions := Input / 1000000
	remainder := Input % 1000000

	plural := "million"
	if millions > 1 {
		plural += "s"
	}
	millionsStr := fmt.Sprintf("%s %s", Entry.Translate(millions), plural)

	if remainder == 0 {
		return millionsStr
	}
	return fmt.Sprintf("%s %s", millionsStr, Entry.Translate(remainder))
}

func (Entry EntryFr) handleBillions(Input int) string {
	billions := Input / 1000000000
	remainder := Input % 1000000000

	plural := "milliard"
	if billions > 1 {
		plural += "s"
	}
	billionsStr := fmt.Sprintf("%s %s", Entry.Translate(billions), plural)

	if remainder == 0 {
		return billionsStr
	}
	return fmt.Sprintf("%s %s", billionsStr, Entry.Translate(remainder))
}

func (Entry EntryFr) handleTrillions(Input int) string {
	trillions := int(int64(Input) / 1000000000000)
	remainder := int(int64(Input) % 1000000000000)

	plural := "billion"
	if trillions > 1 {
		plural += "s"
	}
	trillionsStr := fmt.Sprintf("%s %s", Entry.Translate(trillions), plural)

	if remainder == 0 {
		return trillionsStr
	}
	return fmt.Sprintf("%s %s", trillionsStr, Entry.Translate(remainder))
}

func (Entry EntryFr) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
