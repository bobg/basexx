package basexx

type base62 struct{}

func (b base62) N() int64 { return 62 }

func (b base62) Encode(val int64) (byte, error) {
	if val < 0 || val > 61 {
		return 0, ErrInvalid
	}
	if val < 10 {
		return byte(val) + '0', nil
	}
	if val < 36 {
		return byte(val) - 10 + 'a', nil
	}
	return byte(val) - 36 + 'A', nil
}

func (b base62) Decode(digit byte) (int64, error) {
	switch {
	case '0' <= digit && digit <= '9':
		return int64(digit - '0'), nil
	case 'a' <= digit && digit <= 'z':
		return int64(digit - 'a' + 10), nil
	case 'A' <= digit && digit <= 'Z':
		return int64(digit - 'A' + 36), nil
	default:
		return 0, ErrInvalid
	}
}

// Base62 uses digits 0..9, then a..z, then A..Z.
var Base62 base62
