package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bobg/basexx/v2"
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
	dest, err := basexx.Convert(inp, from, to)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dest)
}
