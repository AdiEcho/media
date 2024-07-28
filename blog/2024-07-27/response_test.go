package hello

import (
   "fmt"
   "os"
   "testing"
)

const text = `
{
  "slideshow": {
    "date": "date of publication",
    "title": "Sample Slide Show"
  }
}
`

func TestWrite(t *testing.T) {
   var bin http_bin
   err := bin.New()
   if err != nil {
      t.Fatal(err)
   }
   text, err := bin.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(text)
}

func TestRead(t *testing.T) {
   var bin http_bin
   err := bin.unmarshal([]byte(text))
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", bin)
}
