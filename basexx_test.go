package basexx

import (
	"encoding/csv"
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
		Base10, Base2, Base8, Base12, Base16, Base36, Base62, Base94,
	}

	f, err := os.Open("testdata.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
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

		for i := 0; i < len(bases)-1; i++ {
			for j := i + 1; j < len(bases); j++ {
				do(vals[0], bases[i], vals[i], bases[j], vals[j])
				do(vals[0], bases[j], vals[j], bases[i], vals[i])
			}
		}
	}
}
