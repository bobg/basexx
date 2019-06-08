package basexx

type binary struct{}

func (b binary) N() int64 { return 256 }

func (b binary) Encode(val int64) ([]byte, error) {
	if val < 0 || val > 255 {
		return nil, ErrInvalid
	}
	return []byte{byte(val)}, nil
}

func (b binary) Decode(inp []byte) (int64, error) {
	if len(inp) != 1 {
		return 0, ErrInvalid
	}
	return int64(inp[0]), nil
}

var Binary binary
