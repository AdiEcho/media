package roku

import (
	"154.pages.dev/widevine"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestPlayback(t *testing.T) {
	var token AccountToken
	err := token.New()
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		play, err := token.playback(path.Base(test.url))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%+v\n", play)
		time.Sleep(time.Second)
	}
}

func TestLicense(t *testing.T) {
	test := tests["episode"]
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
	if err != nil {
		t.Fatal(err)
	}
	key_id, err := hex.DecodeString(test.key_id)
	if err != nil {
		t.Fatal(err)
	}
	var module widevine.CDM
	err = module.New(private_key, client_id, widevine.PSSH(key_id))
	if err != nil {
		t.Fatal(err)
	}
	var token AccountToken
	err = token.New()
	if err != nil {
		t.Fatal(err)
	}
	play, err := token.playback(path.Base(test.url))
	if err != nil {
		t.Fatal(err)
	}
	key, err := module.Key(play, key_id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%x\n", key)
}
