package main

import (
	"bufio"
	"os"
)

func findMatches(fn string) ([]Match, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	collector := NewCollector()

	for scanner.Scan() {
		matches, err := getMatches(scanner.Text())
		if err != nil {
			return nil, err
		}
		collector.Push(matches...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return collector.Result(), nil
}

func getMatches(s string) ([]Match, error) {
	var res []Match
	for _, rxi := range rxInfos {
		matches := rxi.Rx.FindAllString(s, -1)
		for _, s := range matches {
			res = append(res, Match{Txt: s, Kind: rxi.Kind})
		}
	}
	return res, nil
}
