package basexx

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
)

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
			got, err := Convert(fromVal, fromBase, toBase)
			if err != nil {
				t.Fatal(err)
			}
			if got != want {
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

// func TestEncodeDecode(t *testing.T) {
// 	bases := []Base{Base30, Base50, Base62, Base94, Binary}
// 	for i := 2; i <= 36; i++ {
// 		bases = append(bases, Alnum(i))
// 	}
// 	for i, base := range bases {
// 		t.Run(fmt.Sprintf("base_%d_%d", base.N(), i+1), func(t *testing.T) {
// 			cases := []struct {
// 				n       int64
// 				s       string
// 				wantErr bool
// 			}{{
// 				n: 0,
// 				s: "0",
// 			}, {
// 				n: 1,
// 				s: "1",
// 			}, {
// 				n:       -1,
// 				wantErr: true,
// 			}, {
// 				n:       base.N(),
// 				wantErr: true,
// 			}}

// 			for j, tc := range cases {
// 				t.Run(fmt.Sprintf("case_%d", j+1), func(t *testing.T) {
// 					got, err := base.Encode(tc.n)
// 					if tc.wantErr && err != nil {
// 						return
// 					}
// 					if err != nil {
// 						t.Fatal(err)
// 					}
// 					if tc.wantErr {
// 						t.Fatal("got no error but wanted one")
// 					}

// 					// Handle base94 and base256 specially.
// 					want := tc.s
// 					switch base.N() {
// 					case 94:
// 						want = string([]byte{'!' + byte(tc.n)})
// 					case 256:
// 						want = string([]byte{byte(tc.n)})
// 					}

// 					if string(got) != want {
// 						t.Fatalf("got %s, want %s", string(got), want)
// 					}

// 					got2, err := base.Decode(got)
// 					if err != nil {
// 						t.Fatal(err)
// 					}
// 					if got2 != tc.n {
// 						t.Errorf("got base %d, want %d", got2, tc.n)
// 					}
// 				})
// 			}
// 		})
// 	}
// }

// func TestDigits(t *testing.T) {
// 	cases := []struct {
// 		val  int64
// 		base Base
// 		want string
// 	}{{
// 		val: 0, base: Base2, want: "0",
// 	}, {
// 		val: 0, base: Base36, want: "0",
// 	}, {
// 		val: 1, base: Base2, want: "1",
// 	}, {
// 		val: 1, base: Base36, want: "1",
// 	}, {
// 		val: 10, base: Base2, want: "1010",
// 	}, {
// 		val: 10, base: Base36, want: "a",
// 	}, {
// 		val: 42, base: Base2, want: "101010",
// 	}, {
// 		val: 42, base: Base36, want: "16",
// 	}, {
// 		val: 42, base: Base30, want: "1d",
// 	}, {
// 		val: 42, base: Base50, want: "R",
// 	}}

// 	for i, tc := range cases {
// 		t.Run(fmt.Sprintf("case_%d", i+1), func(t *testing.T) {
// 			got, err := Digits(tc.val, tc.base)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			if got != tc.want {
// 				t.Fatalf("got %s, want %s", got, tc.want)
// 			}
// 			got2, err := Value(got, tc.base)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			if got2 != tc.val {
// 				t.Errorf("got back %d, want %d", got2, tc.val)
// 			}
// 		})
// 	}
// }
