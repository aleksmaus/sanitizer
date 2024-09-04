package main

import (
	"encoding/json"
	"regexp"
)

var (
	rxIPv4 = regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
	rxIPv6 = regexp.MustCompile(`\b((([0-9a-fA-F]{1,4}:){7}([0-9a-fA-F]{1,4}|:))|(([0-9a-fA-F]{1,4}:){1,7}:)|(([0-9a-fA-F]{1,4}:){1,6}(:[0-9a-fA-F]{1,4}){1,1})|(([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2})|(([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3})|(([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4})|(([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5})|([0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6}))|(:((:[0-9a-fA-F]{1,4}){1,7}|:)))(%.+)?\b`)
	rxHost = regexp.MustCompile(`\b(?:(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,63})\b`)

	rxInfos = []RxInfo{
		{KindIPV4, rxIPv4},
		{KindIPV6, rxIPv6},
		{KindHost, rxHost},
	}
)

type RxInfo struct {
	Kind Kind
	Rx   *regexp.Regexp
}

type Kind int

const (
	KindUndefined = iota
	KindIPV4
	KindIPV6
	KindHost
)

func (k Kind) String() string {
	switch k {
	case KindUndefined:
		return "undefined"
	case KindIPV4:
		return "ipv4"
	case KindIPV6:
		return "ipv6"
	case KindHost:
		return "host"
	}
	return ""
}

func (k Kind) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}

func (k *Kind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "ipv4":
		*k = KindIPV4
	case "ipv6":
		*k = KindIPV6
	case "host":
		*k = KindHost
	}
	return nil
}

type Match struct {
	Kind Kind
	Txt  string
}
