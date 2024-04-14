package tubi

import (
	"154.pages.dev/widevine"
	"encoding/hex"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestResolution(t *testing.T) {
	for _, test := range tests {
		cms := new(Content)
		err := cms.New(test.content_id)
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)
		if cms.Episode() {
			err := cms.New(cms.Series_ID)
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Second)
			var ok bool
			cms, ok = cms.Get(test.content_id)
			if !ok {
				t.Fatal("get")
			}
		}
		fmt.Println(cms.Video().Resolution, test.url)
	}
}

func TestLicense(t *testing.T) {
	test := tests["the-mask"]
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
	if err := module.New(private_key, client_id, key_id); err != nil {
		t.Fatal(err)
	}
	var cms Content
	if err := cms.New(test.content_id); err != nil {
		t.Fatal(err)
	}
	license, err := module.License(cms.Video())
	if err != nil {
		t.Fatal(err)
	}
	key, ok := module.Key(license)
	fmt.Printf("%x %v\n", key, ok)
}
