package basexx_test

import (
	"fmt"
	"log"

	"github.com/bobg/basexx/v2"
)

func ExampleConvert() {
	const base10val = "12345"

	// The basexx package has no predefined Base20 type,
	// but any base 2 through 36 using alphanumeric digits
	// can be defined with basexx.Alnum.
	base20 := basexx.Alnum(20)

	result, err := basexx.Convert(base10val, basexx.Base10, base20)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s (base 10) = %s (base 20)", base10val, result)

	// Output: 12345 (base 10) = 1ah5 (base 20)
}
