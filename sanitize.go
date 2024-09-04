package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func handleSanitize(fn, cfgfn string) error {
	var (
		gi  []GenInfo
		err error
	)

	if cfgfn == "" {
		gi, err = genInfo(fn)
	} else {
		var b []byte
		b, err = os.ReadFile(cfgfn)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &gi)
	}

	if err != nil {
		return err
	}

	lookup := genReplacementLookup(gi)

	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		for _, rxi := range rxInfos {
			matches := rxi.Rx.FindAllStringIndex(s, -1)
			if len(matches) == 0 {
				continue
			}

			var sbuf strings.Builder

			pos := 0
			for _, m := range matches {
				matchStart := m[0]
				matchEnd := m[1]

				_, _ = sbuf.WriteString(s[pos:matchStart])
				replacement, ok := lookup[Match{Kind: rxi.Kind, Txt: s[matchStart:matchEnd]}]
				if ok {
					_, _ = sbuf.WriteString(replacement)
				} else {
					_, _ = sbuf.WriteString(s[matchStart:matchEnd])
				}
				pos = matchEnd
			}
			_, _ = sbuf.WriteString(s[pos:])
			s = sbuf.String()
		}
		fmt.Println(s)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func genReplacementLookup(gi []GenInfo) map[Match]string {
	res := make(map[Match]string, len(gi))
	for _, e := range gi {
		res[Match{Kind: e.Kind, Txt: e.Txt}] = e.Replacement
	}

	return res
}
