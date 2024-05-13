package member

import (
	"fmt"
	"os"
	"testing"
)

func TestAsset(t *testing.T) {
	article, err := american_hustle.article()
	if err != nil {
		t.Fatal(err)
	}
	var auth Authenticate
	auth.data, err = os.ReadFile("authenticate.json")
	if err != nil {
		t.Fatal(err)
	}
	auth.unmarshal()
	asset, ok := article.film()
	if !ok {
		t.Fatal("data_article.film")
	}
	play, err := auth.play(asset)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(play.dash())
}

func TestAuthenticate(t *testing.T) {
	username := os.Getenv("cineMember_username")
	if username == "" {
		t.Fatal("Getenv")
	}
	password := os.Getenv("cineMember_password")
	var auth Authenticate
	err := auth.New(username, password)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("authenticate.json", auth.data, 0666)
}
