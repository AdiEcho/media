package tubi

import (
	"154.pages.dev/widevine"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

const resolution_data = `
{
   "codec": "VIDEO_CODEC_H264",
   "resolution": "VIDEO_RESOLUTION_720P",
   "ssai_version": "",
   "titan_version": "",
   "type": "dash_widevine"
}
`

func TestResolution(t *testing.T) {
	var v struct {
		Resolution Resolution
	}
	err := json.Unmarshal([]byte(resolution_data), &v)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", v)
}

func TestLicense(t *testing.T) {
	test := tests["movie"]
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
