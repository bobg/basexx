package basexx

// TableBase is a [Base] initialized from a string of digits.
type TableBase struct {
	digits string
	vals   [256]int64
}

// NewTableBase returns a new TableBase initialized from the given digits.
func NewTableBase(digits string) TableBase {
	b := TableBase{digits: digits}
	for i := 0; i < 256; i++ {
		b.vals[i] = -1
	}
	for i := 0; i < len(digits); i++ {
		b.vals[digits[i]] = int64(i)
	}
	return b
}

// N implements Base.N.
func (b TableBase) N() int64 { return int64(len(b.digits)) }

// Digit implements Base.Digit.
func (b TableBase) Digit(val int64) (byte, error) {
	if val < 0 || val >= int64(len(b.digits)) {
		return 0, ErrInvalid
	}
	return b.digits[val], nil
}

// Val implements Base.Val.
func (b TableBase) Val(digit byte) (int64, error) {
	val := b.vals[digit]
	if val < 0 {
		return 0, ErrInvalid
	}
	return val, nil
}

var (
	// Base30 uses digits 0-9, then lower-case bcdfghjkmnpqrstvwxyz.
	// It excludes vowels (to avoid inadvertently spelling naughty words) and the letter "l".
	// Note, this is not the same as basexx.Alnum(30),
	// which uses 0-9 and then abcdefghijklmnopqrst.
	Base30 = NewTableBase("0123456789bcdfghjkmnpqrstvwxyz")

	// Base50 uses digits 0-9, then lower-case bcdfghjkmnpqrstvwxyz, then upper-case BCDFGHJKMNPQRSTVWXYZ.
	// It excludes vowels (to avoid inadvertently spelling naughty words) plus lower- and upper-case L.
	Base50 = NewTableBase("0123456789bcdfghjkmnpqrstvwxyzBCDFGHJKMNPQRSTVWXYZ")
)
