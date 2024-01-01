package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func Test_Player_Android(t *testing.T) {
   var req Request
   req.Android()
   req.Video_ID = android_IDs[0]
   for range [9]struct{}{} {
      p, err := req.Player(nil)
      if err != nil {
         t.Fatal(err)
      }
      if len(p.Streaming_Data.Adaptive_Formats) == 0 {
         t.Fatal("adaptiveFormats")
      }
      if p.Video_Details.View_Count == 0 {
         t.Fatal("viewCount")
      }
      time.Sleep(time.Second)
   }
}

const prompt = `1. Go to
%v

2. Enter this code
%v
`

func Test_Code(t *testing.T) {
   code, err := New_Device_Code()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(prompt, code.Verification_URL, code.User_Code)
   for range [9]bool{} {
      time.Sleep(9 * time.Second)
      tok, err := code.Token()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", tok)
      if tok.Access_Token != "" {
         break
      }
   }
}

const image_test = "UpNXI3_ctAc"

func Test_Image(t *testing.T) {
   req, err := http.NewRequest("HEAD", "", nil)
   if err != nil {
      t.Fatal(err)
   }
   for _, img := range Images {
      req.URL = img.URL(image_test)
      fmt.Println("HEAD", req.URL)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         t.Fatal(res.Status)
      }
      time.Sleep(99 * time.Millisecond)
   }
}

var id_tests = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

func Test_ID(t *testing.T) {
   var req Request
   for _, test := range id_tests {
      err := req.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(req.Video_ID)
   }
}
