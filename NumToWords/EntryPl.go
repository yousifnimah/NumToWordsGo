package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryPl struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntryPl) Translate(Input int) string {
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

func (Entry EntryPl) handleUnits(Input int) string {
	return Entry.LocalizedEntity.Units[Input]
}

func (Entry EntryPl) handleTeens(Input int) string {
	return Entry.LocalizedEntity.Teens[Input-10]
}

func (Entry EntryPl) handleTens(Input int) string {
	tensPart := Input / 10
	unitsPart := Input % 10
	tens := Entry.LocalizedEntity.Tens[tensPart]
	if unitsPart == 0 {
		return tens
	}
	units := Entry.Translate(unitsPart)
	return fmt.Sprintf("%s %s", tens, units)
}

func (Entry EntryPl) handleHundreds(Input int) string {
	hundredsPart := Input / 100
	remainder := Input % 100
	hundreds := Entry.LocalizedEntity.Hundreds[hundredsPart]
	if remainder == 0 {
		return hundreds
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s", hundreds, remStr)
}

func (Entry EntryPl) handleThousands(Input int) string {
	thousands := Input / 1000
	remainder := Input % 1000

	if thousands == 1 {
		if remainder == 0 {
			return Entry.LocalizedEntity.Thousand
		}
		return fmt.Sprintf("%s %s", Entry.LocalizedEntity.Thousand, Entry.Translate(remainder))
	}

	thousandsStr := Entry.Translate(thousands)
	plural := getThousandPlural(thousands)
	if remainder == 0 {
		return fmt.Sprintf("%s %s", thousandsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", thousandsStr, plural, remStr)
}

func (Entry EntryPl) getThousandPlural(n int) string {
	mod100 := n % 100
	if mod100 >= 10 && mod100 <= 20 {
		return "tysięcy"
	}
	mod10 := n % 10
	if mod10 >= 2 && mod10 <= 4 {
		return "tysiące"
	}
	return "tysięcy"
}

func (Entry EntryPl) handleMillions(Input int) string {
	millions := Input / 1000000
	remainder := Input % 1000000

	if millions == 1 {
		if remainder == 0 {
			return Entry.LocalizedEntity.Million
		}
		return fmt.Sprintf("%s %s", Entry.LocalizedEntity.Million, Entry.Translate(remainder))
	}

	millionsStr := Entry.Translate(millions)
	plural := Entry.getMillionPlural(millions)
	if remainder == 0 {
		return fmt.Sprintf("%s %s", millionsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", millionsStr, plural, remStr)
}

func (Entry EntryPl) getMillionPlural(n int) string {
	mod100 := n % 100
	if mod100 >= 10 && mod100 <= 20 {
		return "milionów"
	}
	mod10 := n % 10
	if mod10 >= 2 && mod10 <= 4 {
		return "miliony"
	}
	return "milionów"
}

func (Entry EntryPl) handleBillions(Input int) string {
	billions := Input / 1000000000
	remainder := Input % 1000000000

	if billions == 1 {
		if remainder == 0 {
			return Entry.LocalizedEntity.Billion
		}
		return fmt.Sprintf("%s %s", Entry.LocalizedEntity.Billion, Entry.Translate(remainder))
	}

	billionsStr := Entry.Translate(billions)
	plural := Entry.getBillionPlural(billions)
	if remainder == 0 {
		return fmt.Sprintf("%s %s", billionsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", billionsStr, plural, remStr)
}

func (Entry EntryPl) getBillionPlural(n int) string {
	mod100 := n % 100
	if mod100 >= 10 && mod100 <= 20 {
		return "miliardów"
	}
	mod10 := n % 10
	if mod10 >= 2 && mod10 <= 4 {
		return "miliardy"
	}
	return "miliardów"
}

func (Entry EntryPl) handleTrillions(Input int) string {
	trillions := int(int64(Input) / 1000000000000)
	remainder := int(int64(Input) % 1000000000000)

	if trillions == 1 {
		if remainder == 0 {
			return Entry.LocalizedEntity.Trillion
		}
		return fmt.Sprintf("%s %s", Entry.LocalizedEntity.Trillion, Entry.Translate(remainder))
	}

	trillionsStr := Entry.Translate(trillions)
	plural := Entry.getTrillionPlural(trillions)
	if remainder == 0 {
		return fmt.Sprintf("%s %s", trillionsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", trillionsStr, plural, remStr)
}

func (Entry EntryPl) getTrillionPlural(n int) string {
	mod100 := n % 100
	if mod100 >= 10 && mod100 <= 20 {
		return "bilionów"
	}
	mod10 := n % 10
	if mod10 >= 2 && mod10 <= 4 {
		return "biliony"
	}
	return "bilionów"
}

func (Entry EntryPl) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
