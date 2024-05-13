package member

import (
	"fmt"
	"os"
	"testing"
)

func TestAsset(t *testing.T) {
	article, err := american_hustle.Article()
	if err != nil {
		t.Fatal(err)
	}
	var auth Authenticate
	auth.Data, err = os.ReadFile("authenticate.json")
	if err != nil {
		t.Fatal(err)
	}
	auth.Unmarshal()
	asset, ok := article.Film()
	if !ok {
		t.Fatal("data_article.film")
	}
	play, err := auth.Play(asset)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(play.DASH())
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
	os.WriteFile("authenticate.json", auth.Data, 0666)
}
