package main

import "fmt"

func handleShow(fn string) error {
	matches, err := findMatches(fn)
	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	}

	for _, s := range matches {
		fmt.Printf("%v\t%v\n", s.Kind, s.Txt)
	}
	return nil
}
