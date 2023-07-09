package basexx

type base94 struct{}

func (b base94) N() int64 { return 94 }

func (b base94) Digit(val int64) (byte, error) {
	if val < 0 || val > 93 {
		return 0, ErrInvalid
	}
	return byte(val + 33), nil
}

func (b base94) Val(digit byte) (int64, error) {
	if digit < 33 || digit > 126 {
		return 0, ErrInvalid
	}
	return int64(digit - 33), nil
}

// Base94 uses all printable ASCII characters (33 through 126) as digits.
var Base94 base94
