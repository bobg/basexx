// Package basexx provides functions for converting numbers to and from strings in arbitrary bases.
package basexx

import (
	"bufio"
	"io"
	"math/big"
	"strings"

	"github.com/bobg/go-generics/v2/slices"
	"github.com/pkg/errors"
)

// Base is the abstract type of a number base.
type Base interface {
	// N is the number of digits in the base.
	// It must be at least 2 and at most 256.
	N() int64

	// Val is the value of the given digit in this base.
	Val(byte) (int64, error)

	// Digit is the digit with the given value in this base.
	Digit(int64) (byte, error)
}

// ErrInvalid is returned when a digit or a value is out of range.
var ErrInvalid = errors.New("invalid")

var zero = new(big.Int)

// Decode decodes an integer expressed as a string in the given base.
func Decode(inp io.Reader, base Base) (*big.Int, error) {
	var rr io.ByteReader
	if r, ok := inp.(io.ByteReader); ok {
		rr = r
	} else {
		rr = bufio.NewReader(inp)
	}

	var (
		result = new(big.Int)
		n      = big.NewInt(base.N())
	)
	for {
		digit, err := rr.ReadByte()
		if errors.Is(err, io.EOF) {
			return result, nil
		}
		if err != nil {
			return nil, errors.Wrap(err, "reading input")
		}
		val, err := base.Val(digit)
		if err != nil {
			return nil, errors.Wrap(err, "invalid digit")
		}
		result.Mul(result, n)
		result.Add(result, big.NewInt(val))
	}
}

// DecodeString decodes a string in the given base.
func DecodeString(s string, base Base) (*big.Int, error) {
	return Decode(strings.NewReader(s), base)
}

// Encode encodes inp as a string in the given base.
// If inp is negative, it is silently made positive.
func Encode(out io.Writer, inp *big.Int, base Base) error {
	var (
		digits []byte
		n      = big.NewInt(base.N())
	)

	switch inp.Sign() {
	case -1:
		inp = new(big.Int).Neg(inp)
	case 0:
		result, err := base.Digit(0)
		if err != nil {
			return errors.Wrap(err, "invalid digit")
		}
		_, err = out.Write([]byte{result})
		return errors.Wrap(err, "writing output")
	}

	for inp.Cmp(zero) > 0 {
		_, m := inp.DivMod(inp, n, new(big.Int))
		digit, err := base.Digit(m.Int64())
		if err != nil {
			return errors.Wrap(err, "invalid digit")
		}
		digits = append(digits, digit)
	}

	slices.Reverse(digits)

	_, err := out.Write(digits)
	return errors.Wrap(err, "writing output")
}

// EncodeInt64 encodes an integer as a string in the given base.
// If inp is negative, it is silently made positive.
func EncodeInt64(out io.Writer, inp int64, base Base) error {
	return Encode(out, big.NewInt(inp), base)
}

// Convert converts a string from one base to another.
func Convert(inp string, from, to Base) (string, error) {
	n, err := DecodeString(inp, from)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	err = Encode(&sb, n, to)
	return sb.String(), errors.Wrap(err, "encoding")
}
