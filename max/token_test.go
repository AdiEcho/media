package max

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	var login DefaultLogin
	login.Credentials.Username = os.Getenv("max_username")
	if login.Credentials.Username == "" {
		t.Fatal("Getenv")
	}
	login.Credentials.Password = os.Getenv("max_password")
	var key PublicKey
	err := key.New()
	if err != nil {
		t.Fatal(err)
	}
	var token DefaultToken
	err = token.New()
	if err != nil {
		t.Fatal(err)
	}
	err = token.Login(key, login)
	if err != nil {
		t.Fatal(err)
	}
	text, err := token.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("token.json", text, 0666)
}

func TestConfig(t *testing.T) {
	var token DefaultToken
	err := token.New()
	if err != nil {
		t.Fatal(err)
	}
	decision, err := token.decision()
	if err != nil {
		t.Fatal(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(decision)
}

func TestToken(t *testing.T) {
	var token DefaultToken
	err := token.New()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", token)
}
