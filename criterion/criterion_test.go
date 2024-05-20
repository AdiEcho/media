package criterion

import (
	"154.pages.dev/encoding"
	"fmt"
	"os"
	"testing"
)

func TestVideo(t *testing.T) {
	var (
		token AuthToken
		err   error
	)
	token.Data, err = os.ReadFile("token.json")
	if err != nil {
		t.Fatal(err)
	}
	token.unmarshal()
	item, err := token.video(my_dinner)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", item)
	name, err := encoding.Name(item)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%q\n", name)
}

func TestToken(t *testing.T) {
	username := os.Getenv("criterion_username")
	if username == "" {
		t.Fatal("Getenv")
	}
	password := os.Getenv("criterion_password")
	var token AuthToken
	err := token.New(username, password)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("token.json", token.Data, 0666)
}

// criterionchannel.com/videos/my-dinner-with-andre
const my_dinner = "my-dinner-with-andre"
