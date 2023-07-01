package NumToWords

import (
	"errors"
	"fmt"
	"strings"
)

type NumToWord struct{}

func Convert(Input int, Language string) (string, error) {
	if Language == "" {
		err := errors.New("language is not provided")
		panic(err)
		return "", err
	}
	Entry := new(Entry)
	Entry.Language = Language
	Entry.BootstrapLanguage()
	return Entry.Calculate(Input), nil
}

func (Entry Entry) Calculate(Input int) string {
	result := ""
	if Input == 0 {
		result = Entry.LocalizedEntity.Zero
	} else if Input < 10 {
		result = Entry.LocalizedEntity.Units[Input]
	} else if Input < 20 {
		result = Entry.LocalizedEntity.Teens[Input-10]
	} else if Input < 100 {
		result = Entry.LocalizedEntity.Tens[Input/10] + " " + Entry.Calculate(Input%10)
	} else if Input < 1000 {
		result = fmt.Sprintf("%s %s %s", Entry.Calculate(Input/100), Entry.LocalizedEntity.Hundred, Entry.Calculate(Input%100))
	} else if Input < 1000000 {
		result = fmt.Sprintf("%s %s %s", Entry.Calculate(Input/1000), Entry.LocalizedEntity.Thousand, Entry.Calculate(Input%1000))
	} else if Input < 1000000000 {
		result = fmt.Sprintf("%s %s %s", Entry.Calculate(Input/1000000), Entry.LocalizedEntity.Million, Entry.Calculate(Input%1000000))
	} else {
		result = fmt.Sprintf("%s %s %s", Entry.Calculate(Input/1000000000), Entry.LocalizedEntity.Billion, Entry.Calculate(Input%1000000000))
	}
	return strings.TrimSpace(result)
}
