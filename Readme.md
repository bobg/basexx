# Basexx - Convert between digit strings and various number bases

[![Go Reference](https://pkg.go.dev/badge/github.com/bobg/basexx.svg)](https://pkg.go.dev/github.com/bobg/basexx)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobg/basexx)](https://goreportcard.com/report/github.com/bobg/basexx)
[![Tests](https://github.com/bobg/basexx/actions/workflows/go.yml/badge.svg)](https://github.com/bobg/basexx/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/bobg/basexx/badge.svg?branch=master)](https://coveralls.io/github/bobg/basexx?branch=master)

This is basexx,
a package for converting numbers to digit strings in various bases
and vice versa.

## Usage

To get the Base30 encoding of the number 412:

```go
result, err := basexx.Digits(412, basexx.Base30)
```

To decode the Base30 digit string `"fr"`:

```go
result, err := basexx.Value("fr", basexx.Base30)
```

To convert a digit string `x` in base `from` to a new digit string in base `to`:

```go
var (
  src     = basexx.NewBuffer(x, from)
  destbuf = make([]byte, basexx.Length(from, to, len(x)))
  dest    = basexx.NewBuffer(destbuf, to)
)
_, err := Convert(dest, src)
if err != nil { ... }
result := dest.Written()
```

To define your own new number base:

```go
// ReverseBase10 uses digits '9' through '0' just to mess with you.
type ReverseBase10 struct{}

func (ReverseBase10) N() int64 { return 10 }

func (ReverseBase10) Encode(val int64) (byte, error) {
  if val < 0 || val > 9 {
    return 0, errors.New("digit value out of range")
  }
  return byte('9' - val), nil
}

func (ReverseBase10) Decode(digit byte) (int64, error) {
  if digit < '0’ || digit > '9' {
    return 0, errors.New("invalid encoded digit")
  }
  return int64(9 - (digit - '0’))
}
```