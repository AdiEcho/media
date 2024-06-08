package rtbf

import (
	"154.pages.dev/widevine"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestEntitlement(t *testing.T) {
	text, err := os.ReadFile("account.json")
	if err != nil {
		t.Fatal(err)
	}
	var account AccountLogin
	err = account.Unmarshal(text)
	if err != nil {
		t.Fatal(err)
	}
	token, err := account.Token()
	if err != nil {
		t.Fatal(err)
	}
	gigya, err := token.Login()
	if err != nil {
		t.Fatal(err)
	}
	page, err := NewPage(media[0].path)
	if err != nil {
		t.Fatal(err)
	}
	title, err := gigya.Entitlement(page)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", title)
	fmt.Println(title.dash())
}

func TestAccountsLogin(t *testing.T) {
	username := os.Getenv("rtbf_username")
	if username == "" {
		t.Fatal("Getenv")
	}
	password := os.Getenv("rtbf_password")
	var login AccountLogin
	err := login.New(username, password)
	if err != nil {
		t.Fatal(err)
	}
	text, err := login.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("account.json", text, 0666)
}

func TestWidevine(t *testing.T) {
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
	medium := media[0]
	key_id, err := base64.StdEncoding.DecodeString(medium.key_id)
	if err != nil {
		t.Fatal(err)
	}
	var module widevine.CDM
	err = module.New(private_key, client_id, widevine.PSSH(key_id, nil))
	if err != nil {
		t.Fatal(err)
	}
	text, err := os.ReadFile("account.json")
	if err != nil {
		t.Fatal(err)
	}
	var account AccountLogin
	err = account.Unmarshal(text)
	if err != nil {
		t.Fatal(err)
	}
	token, err := account.Token()
	if err != nil {
		t.Fatal(err)
	}
	gigya, err := token.Login()
	if err != nil {
		t.Fatal(err)
	}
	page, err := NewPage(medium.path)
	if err != nil {
		t.Fatal(err)
	}
	title, err := gigya.Entitlement(page)
	if err != nil {
		t.Fatal(err)
	}
	key, err := module.Key(title, key_id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%x\n", key)
}

func TestWebToken(t *testing.T) {
	text, err := os.ReadFile("account.json")
	if err != nil {
		t.Fatal(err)
	}
	var account AccountLogin
	err = account.Unmarshal(text)
	if err != nil {
		t.Fatal(err)
	}
	token, err := account.Token()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", token)
}
