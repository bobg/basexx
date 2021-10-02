package basexx

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestLength(t *testing.T) {
	cases := []struct {
		from, to int64
		n, want  int
	}{
		{from: 2, to: 10, n: 1, want: 1},
		{from: 2, to: 10, n: 2, want: 1},
		{from: 2, to: 10, n: 4, want: 2},
		{from: 2, to: 10, n: 100, want: 31},
		{from: 10, to: 2, n: 1, want: 4},
		{from: 10, to: 2, n: 100, want: 333},
		{from: 256, to: 50, n: 20, want: 29},
		{from: 256, to: 94, n: 20, want: 25},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case%02d", i+1), func(t *testing.T) {
			got := Length(c.from, c.to, c.n)
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	bases := []Base{
		Base10, Base2, Base8, Base12, Base16, Base36, Base62, Base94, Base50, Base30,
	}

	f, err := os.Open("testdata.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	do := func(val10 string, fromBase Base, fromVal string, toBase Base, want string) {
		t.Run(fmt.Sprintf("case_%s_base%d_to_base%d", val10, fromBase.N(), toBase.N()), func(t *testing.T) {
			src := NewBuffer([]byte(fromVal), fromBase)
			destBuf := make([]byte, Length(fromBase.N(), toBase.N(), len(fromVal)))
			dest := NewBuffer(destBuf[:], toBase)
			_, err := Convert(dest, src)
			if err != nil {
				t.Fatal(err)
			}
			got := dest.Written()
			if string(got) != want {
				t.Errorf("got %s, want %s", string(got), want)
			}
		})
	}

	for {
		vals, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if len(vals) != len(bases) {
			t.Fatalf("testdata row contains %d value(s), want %d", len(vals), len(bases))
		}

		for i := 0; i < len(bases)-1; i++ {
			for j := i + 1; j < len(bases); j++ {
				// Check converting from bases[i] to bases[j] and vice versa.
				do(vals[0], bases[i], vals[i], bases[j], vals[j])
				do(vals[0], bases[j], vals[j], bases[i], vals[i])
			}
		}

		// Check that converting between Base16 and Binary works the same as in encoding/hex.

		val16 := vals[4]
		if len(val16)%2 == 1 {
			val16 = "0" + val16
		}
		bin, err := hex.DecodeString(val16)
		if err != nil {
			t.Fatal(err)
		}

		do(vals[0], Base16, val16, Binary, string(bin))
		do(vals[0], Binary, string(bin), Base16, vals[4])
	}
}

func TestEncodeDecode(t *testing.T) {
	bases := []Base{Base30, Base50, Base62, Base94, Binary}
	for i := 2; i <= 36; i++ {
		bases = append(bases, Alnum(i))
	}
	for i, base := range bases {
		t.Run(fmt.Sprintf("base_%d_%d", base.N(), i+1), func(t *testing.T) {
			cases := []struct {
				n       int64
				s       string
				wantErr bool
			}{{
				n: 0,
				s: "0",
			}, {
				n: 1,
				s: "1",
			}, {
				n:       -1,
				wantErr: true,
			}, {
				n:       base.N(),
				wantErr: true,
			}}

			for j, tc := range cases {
				t.Run(fmt.Sprintf("case_%d", j+1), func(t *testing.T) {
					got, err := base.Encode(tc.n)
					if tc.wantErr && err != nil {
						return
					}
					if err != nil {
						t.Fatal(err)
					}
					if tc.wantErr {
						t.Fatal("got no error but wanted one")
					}

					// Handle base94 and base256 specially.
					want := tc.s
					switch base.N() {
					case 94:
						want = string([]byte{'!' + byte(tc.n)})
					case 256:
						want = string([]byte{byte(tc.n)})
					}

					if string(got) != want {
						t.Fatalf("got %s, want %s", string(got), want)
					}

					got2, err := base.Decode(got)
					if err != nil {
						t.Fatal(err)
					}
					if got2 != tc.n {
						t.Errorf("got base %d, want %d", got2, tc.n)
					}
				})
			}
		})
	}
}
