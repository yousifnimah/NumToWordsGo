# NumToWordsGo

NumToWordsGo is a lightweight and efficient Go library that provides a simple function to convert numeric values into their corresponding word representation. With this library, you can easily convert numbers to words, making it useful for applications that require human-readable number formatting or text generation.

## Features

- Convert numbers to their word representation
- Supports numbers up to billions
- Efficient and easy-to-use
- Supports two languages English and Arabic

## Installation

To use NumToWordsGo in your Go project, you need to install it using `go get` command:

```bash
go get github.com/yousifnimah/NumToWordsGo
```

## English Language Usage

Here's an example of how to use NumToWordsGo in your Go code:

```go
package main

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords"
)

func main() {
	input := 123456789
	LanguageCode := "en"
	words, err := NumToWords.Convert(input, LanguageCode)
	if err != nil {
		return
	}
	fmt.Println(words)
}
```

Output:
```go
one hundred twenty three million four hundred fifty six thousand seven hundred eighty nine
```

## Arabic Language Usage

Here's an example of how to use NumToWordsGo in your Go code:

```go
package main

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords"
)

func main() {
	input := 50010
	LanguageCode := "ar"
	words, err := NumToWords.Convert(input, LanguageCode)
	if err != nil {
		return
	}
	fmt.Println(words)
}
```

Output:
```go
خمسون ألف و عشرة
```

## Screenshot:
![img.png](https://i.imgur.com/g0kPxcB.png)
![img.png](https://i.imgur.com/IJyLyrn.gif)

***In the example above***, we import the NumToWordsGo package and use the NumToWords function to convert the number 50010 into its word representation in Arabic and English languages. The result is then printed to the console.


## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing
Contributions are welcome! If you find any issues or want to add new features or improvements, please create a pull request.

## Acknowledgments
This library is inspired by the need for number-to-word conversion in various applications.
Contact
For any questions or inquiries, please feel free to contact me.