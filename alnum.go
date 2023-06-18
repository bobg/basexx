package basexx

// Alnum is a type for bases from 2 through 36,
// where the digits for the first 10 digit values are '0' through '9'
// and the remaining digits are 'a' through 'z'.
// For decoding, upper-case 'A' through 'Z' are the same as lower-case.
type Alnum int

// N implements Base.N.
func (a Alnum) N() int64 { return int64(a) }

// Digit implements Base.Digit.
func (a Alnum) Digit(val int64) (byte, error) {
	if val < 0 || val >= int64(a) {
		return 0, ErrInvalid
	}
	if val < 10 {
		return byte(val) + '0', nil
	}
	return byte(val) - 10 + 'a', nil
}

// Val implements Base.Val.
func (a Alnum) Val(digit byte) (int64, error) {
	switch {
	case '0' <= digit && digit <= '9':
		return int64(digit - '0'), nil
	case 'a' <= digit && digit <= 'z':
		return int64(digit - 'a' + 10), nil
	case 'A' <= digit && digit <= 'Z':
		return int64(digit - 'A' + 10), nil
	default:
		return 0, ErrInvalid
	}
}

const (
	// Base2 uses digits "0" and "1"
	Base2 = Alnum(2)

	// Base8 uses digits "0" through "7"
	Base8 = Alnum(8)

	// Base10 uses digits "0" through "9"
	Base10 = Alnum(10)

	// Base12 uses digits "0" through "9" plus "a" and "b"
	Base12 = Alnum(12)

	// Base16 uses digits "0" through "9" plus "a" through "f"
	Base16 = Alnum(16)

	// Base32 uses digits "0" through "9" plus "a" through "v"
	Base32 = Alnum(32)

	// Base36 uses digits "0" through "9" plus "a" through "z"
	Base36 = Alnum(36)
)
