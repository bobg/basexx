package basexx

import (
	"errors"
	"io"
	"math"
	"math/big"
)

type Source interface {
	Read() (int64, error)
	Base() int64
}

type Dest interface {
	Prepend(int64) error
	Base() int64
}

type Base interface {
	N() int64
	Encode(int64) ([]byte, error)
	Decode([]byte) (int64, error)
}

var ErrInvalid = errors.New("invalid")

var zero = new(big.Int)

func Convert(dest Dest, src Source) (int, error) {
	var (
		accum    = new(big.Int)
		srcBase  = big.NewInt(src.Base())
		destBase = big.NewInt(dest.Base())
	)
	for {
		digit, err := src.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		accum.Mul(accum, srcBase)
		if digit != 0 {
			accum.Add(accum, big.NewInt(digit))
		}
	}
	var written int
	for accum.Cmp(zero) > 0 {
		r := new(big.Int)
		accum.QuoRem(accum, destBase, r)
		err := dest.Prepend(r.Int64())
		if err != nil {
			return written, err
		}
		written++
	}
	if written == 0 {
		err := dest.Prepend(0)
		if err != nil {
			return written, err
		}
		written++
	}
	return written, nil
}

// Length computes the maximum length of a digit string converted from n `from`-base digits to base `to`.
func Length(from, to int64, n int) int {
	ratio := math.Log(float64(from)) / math.Log(float64(to))
	result := float64(n) * ratio
	return int(math.Ceil(result))
}
