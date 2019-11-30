package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bobg/basexx"
)

func main() {
	bases := map[string]basexx.Base{
		"30":  basexx.Base30,
		"50":  basexx.Base50,
		"62":  basexx.Base62,
		"94":  basexx.Base94,
		"bin": basexx.Binary,
	}

	for i := 2; i <= 36; i++ {
		bases["a"+strconv.Itoa(i)] = basexx.Alnum(i)
	}

	var (
		fromstr = flag.String("from", "", "convert from base")
		tostr   = flag.String("to", "", "convert to base")
	)

	flag.Parse()

	from, ok := bases[*fromstr]
	if !ok {
		log.Fatalf(`unknown "from" base "%s"`, *fromstr)
	}

	to, ok := bases[*tostr]
	if !ok {
		log.Fatalf(`unknown "to" base "%s"`, *tostr)
	}

	if flag.NArg() == 0 {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			inp := s.Text()
			do(inp, from, to)
		}
	} else {
		for _, inp := range flag.Args() {
			do(inp, from, to)
		}
	}
}

func do(inp string, from, to basexx.Base) {
	var (
		src     = basexx.NewBuffer([]byte(inp), from)
		destBuf = make([]byte, basexx.Length(from.N(), to.N(), len(inp)))
		dest    = basexx.NewBuffer(destBuf, to)
	)
	_, err := basexx.Convert(dest, src)
	if err != nil {
		log.Fatal(err)
	}
	result := dest.Written()
	fmt.Println(string(result))
}
