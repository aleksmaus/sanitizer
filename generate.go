package main

import (
	"bufio"
	"crypto/rand"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"math/big"
	"net"
)

//go:embed data
var dataFS embed.FS

var words []string
var domains []string

func init() {
	var err error
	words, err = readFile(dataFS, "data/words")
	if err != nil {
		panic(err)
	}
	domains, err = readFile(dataFS, "data/domains")
	if err != nil {
		panic(err)
	}
}

func handleGenerate(fn string) error {
	gi, err := genInfo(fn)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(gi, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func genInfo(fn string) ([]GenInfo, error) {
	matches, err := findMatches(fn)
	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	}

	var gi []GenInfo

	for _, m := range matches {
		var genFn func(string) (string, error)
		switch m.Kind {
		case KindIPV4:
			genFn = genIPV4
		case KindIPV6:
			genFn = genIPV6
		case KindHost:
			genFn = genHost
		default:
			return nil, fmt.Errorf("unsupported generator for kind: %v", m.Kind)
		}

		replacement, err := genFn(m.Txt)
		if err != nil {
			return nil, err
		}
		gi = append(gi, GenInfo{
			Kind:        m.Kind,
			Txt:         m.Txt,
			Replacement: replacement,
		})
	}
	return gi, nil
}

type GenInfo struct {
	Kind        Kind   `json:"kind"`
	Txt         string `json:"text"`
	Replacement string `json:"replacement"`
}

func genIPV4(_ string) (string, error) {
	ip := make([]byte, 4)
	_, err := rand.Read(ip)
	if err != nil {
		return "", err
	}

	return net.IP(ip).String(), nil
}

func genIPV6(_ string) (string, error) {
	ip := make([]byte, 16)
	_, err := rand.Read(ip)
	if err != nil {
		return "", err
	}
	return net.IP(ip).String(), nil
}

func genHost(origin string) (string, error) {
	n, err := genRandomNoneZero(len(words))
	if err != nil {
		return "", err
	}

	host := words[n]

	n, err = genRandomNoneZero(len(domains))
	if err != nil {
		return "", err
	}

	domain := domains[n]

	n, err = genRandomNoneZero(24)
	if err != nil {
		return "", err
	}
	random, err := randomString(n)
	if err != nil {
		return "", err
	}
	return random + `.` + host + `.` + domain, nil
}

func readFile(rootFs fs.FS, fn string) ([]string, error) {
	file, err := rootFs.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res, nil
}

func genRandomNoneZero(max int) (int, error) {
	for {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			return 0, err
		}

		if num.Cmp(big.NewInt(0)) > 0 {
			return int(num.Int64()), nil
		}
	}
}

func randomString(sz int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, sz)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		b[i] = chars[num.Int64()]
	}
	return string(b), nil
}
