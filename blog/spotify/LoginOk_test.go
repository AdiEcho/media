package spotify

import (
	"154.pages.dev/protobuf"
	"fmt"
	"os"
	"testing"
)

func TestChallenge(t *testing.T) {
	username := os.Getenv("spotify_username")
	if username == "" {
		t.Fatal("spotify_username")
	}
	password := os.Getenv("spotify_password")
	var response login_response
	err := response.New(username, password)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := response.ok(username, password)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ok.access_token())
}

var _ = protobuf.Message{
	protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
		protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("tzjngbdihh9uk2wd4w5016f21")},
		protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("BQDB3nakoC8zFIBMOMHmMNe-QCK3jRl-0R_IT_kYPGI19FnRaTlMqO_MjTL3oBXpt0uqa_byAcKSLxSrjBxCCIFKHmVN5LpYPci6p3nzuclQjhwSHu76LoWK2PhB9i5wRbzFKTDD3T94q766Zsgqq6oMjtntf4qEjGzm-KU4ku9dovqsdJuCvYixFp8rLRnkajLywPUIJo-hXj6EYg-Z0eCbvfSiCDyLqIFVEqUY4wBtS_mZYeyFbRGIvlrTLyby5xG5JGX6NxpuqcR8lZHUvmTlyclCtaVEedsQQlgXGVh2jWLvdLybye_-WeLWDKXf0zKC9E5OyOqTB_seXioN")},
		protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("AQD2qEd61e2koLBcTZSNSZF8ejR9jO6ofT8BB4BgFa5S0pUt3hcF6QIalzJn4CnJtf-cXQjwjZ47rllCcIfnNrTC2-T_aaWOIRxp6X5_drG_UN1NpcJEzRBNqR1cjWoMJv90Ud7c3UEB7hAU-RYr8j-bFYvHV-j3SBbjx8F67sroo42hkew9HBiHhkRUPvqWNZnU8AdyVaZd6iUMtIVgDVEHOg")},
		protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(3600)},
	}},
}
