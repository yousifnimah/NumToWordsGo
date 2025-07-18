package NumToWords

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords/locales"
	"strings"
)

type EntryUk struct {
	Language        string
	LocalizedEntity locales.Lang
}

func (Entry EntryUk) Translate(Input int) string {
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

func (Entry EntryUk) handleUnits(Input int) string {
	return Entry.LocalizedEntity.Units[Input]
}

func (Entry EntryUk) handleTeens(Input int) string {
	return Entry.LocalizedEntity.Teens[Input-10]
}

func (Entry EntryUk) handleTens(Input int) string {
	tensPart := Input / 10
	unitsPart := Input % 10
	tens := Entry.LocalizedEntity.Tens[tensPart]
	if unitsPart == 0 {
		return tens
	}
	units := Entry.Translate(unitsPart)
	return fmt.Sprintf("%s %s", tens, units)
}

func (Entry EntryUk) handleHundreds(Input int) string {
	hundredsPart := Input / 100
	remainder := Input % 100
	hundreds := Entry.LocalizedEntity.Hundreds[hundredsPart]
	if remainder == 0 {
		return hundreds
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s", hundreds, remStr)
}

func (Entry EntryUk) handleThousands(Input int) string {
	thousands := Input / 1000
	remainder := Input % 1000

	thousandsStr := Entry.getFemaleForm(Entry.Translate(thousands))
	plural := getThousandPlural(thousands)

	if thousands == 1 && remainder == 0 {
		return Entry.LocalizedEntity.Thousand
	}

	if remainder == 0 {
		return fmt.Sprintf("%s %s", thousandsStr, plural)
	}
	remStr := Entry.Translate(remainder)
	return fmt.Sprintf("%s %s %s", thousandsStr, plural, remStr)
}

func (Entry EntryUk) getThousandPlural(n int) string {
	mod100 := n % 100
	mod10 := n % 10

	if mod10 == 1 && mod100 != 11 {
		return "тисяча"
	}
	if mod10 >= 2 && mod10 <= 4 && (mod100 < 10 || mod100 >= 20) {
		return "тисячі"
	}
	return "тисяч"
}

func (Entry EntryUk) handleMillions(Input int) string {
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

func (Entry EntryUk) getMillionPlural(n int) string {
	mod100 := n % 100
	mod10 := n % 10

	if mod10 == 1 && mod100 != 11 {
		return "мільйон"
	}
	if mod10 >= 2 && mod10 <= 4 && (mod100 < 10 || mod100 >= 20) {
		return "мільйони"
	}
	return "мільйонів"
}

func (Entry EntryUk) handleBillions(Input int) string {
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

func (Entry EntryUk) getBillionPlural(n int) string {
	mod100 := n % 100
	mod10 := n % 10

	if mod10 == 1 && mod100 != 11 {
		return "мільярд"
	}
	if mod10 >= 2 && mod10 <= 4 && (mod100 < 10 || mod100 >= 20) {
		return "мільярди"
	}
	return "мільярдів"
}

func (Entry EntryUk) handleTrillions(Input int) string {
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

func (Entry EntryUk) getTrillionPlural(n int) string {
	mod100 := n % 100
	mod10 := n % 10

	if mod10 == 1 && mod100 != 11 {
		return "трильйон"
	}
	if mod10 >= 2 && mod10 <= 4 && (mod100 < 10 || mod100 >= 20) {
		return "трильйони"
	}
	return "трильйонів"
}

func (Entry EntryUk) getFemaleForm(s string) string {
	s = strings.ReplaceAll(s, "один", "одна")
	s = strings.ReplaceAll(s, "два", "дві")
	return s
}

func (Entry EntryUk) ZeroResponse() string {
	return Entry.LocalizedEntity.Zero
}
