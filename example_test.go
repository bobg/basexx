package basexx_test

import (
	"fmt"
	"log"

	"github.com/bobg/basexx"
)

func ExampleConvert() {
	const base10val = "12345"

	// The basexx package has no predefined Base20 type,
	// but any base 2 through 36 using alphanumeric digits
	// can be defined with basexx.Alnum.
	base20 := basexx.Alnum(20)

	// A Buffer can serve as a Source for Convert.
	src := basexx.NewBuffer([]byte(base10val), basexx.Base10)

	// Allocate enough space (according to basexx.Length) for holding the result.
	destBuf := make([]byte, basexx.Length(10, 20, len(base10val)))

	// A Buffer can also serve as a Dest for Convert.
	dest := basexx.NewBuffer(destBuf[:], base20)

	_, err := basexx.Convert(dest, src)
	if err != nil {
		log.Fatal(err)
	}

	// Use Written to get the written-to portion of the allocated byte slice (destBuf).
	result := dest.Written()

	fmt.Printf("%s (base 10) = %s (base 20)", base10val, string(result))

	// Output: 12345 (base 10) = 1ah5 (base 20)
}
