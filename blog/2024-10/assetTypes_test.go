package paramount

import (
   "errors"
   "fmt"
   "io"
   "net/http"
   "strings"
   "testing"
   "time"
)

type asset_type [2]string

func (a asset_type) String() string {
   var data []byte
   if a[0] != "" {
      data = append(data, a[0]...)
   }
   if a[1] != "" {
      if a[1] != a[0] {
         if data != nil {
            data = append(data, '|')
         }
         data = append(data, a[1]...)
      }
   }
   if data != nil {
      data = append([]byte("&assetTypes="), data...)
   }
   return string(data)
}

const france = "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ"

func TestSlyGuy(t *testing.T) {
   for _, a := range sly_guy {
      for _, b := range sly_guy {
         fmt.Printf("%q\n", asset_type{a, b})
      }
   }
}

func get(id, asset string) error {
   req, err := http.NewRequest("", "https://link.theplatform.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/s/dJ5BDC/media/guid/2198311517/" + id
   req.URL.RawQuery = "formats=MPEG-DASH" + asset
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   if count := strings.Count(string(data), "</Period>"); count != 1 {
      return fmt.Errorf("%v Period", count)
   }
   return nil
}

// github.com/matthuisman/slyguy.addons/blob/master/slyguy.paramount.plus/resources/lib/api.py
var sly_guy = []string{
   "",
   "DASH_CENC",
   "DASH_CENC_HDR10",
   "DASH_CENC_PRECON",
   "DASH_CENC_PS4",
   "DASH_LIVE",
   "DASH_TA",
}

var asset_types = []struct{
   france bool
   value string
}{
   //{false, "&assetTypes=DASH_CENC_HDR10|DASH_CENC"},
   //{false, "&assetTypes=DASH_CENC_HDR10|DASH_CENC_PS4"},
   //{false, "&assetTypes=DASH_CENC_HDR10|DASH_LIVE"},
   //{false, "&assetTypes=DASH_CENC_HDR10|DASH_TA"},
   //{false, "&assetTypes=DASH_CENC_PS4|DASH_CENC"},
   //{false, "&assetTypes=DASH_CENC_PS4|DASH_CENC_HDR10"},
   //{false, "&assetTypes=DASH_CENC_PS4|DASH_LIVE"},
   //{false, "&assetTypes=DASH_CENC_PS4|DASH_TA"},
   //{false, "&assetTypes=DASH_CENC|DASH_CENC_HDR10"},
   //{false, "&assetTypes=DASH_CENC|DASH_CENC_PS4"},
   //{false, "&assetTypes=DASH_CENC|DASH_LIVE"},
   //{false, "&assetTypes=DASH_CENC|DASH_TA"},
   //{false, "&assetTypes=DASH_LIVE|DASH_CENC"},
   //{false, "&assetTypes=DASH_LIVE|DASH_CENC_HDR10"},
   //{false, "&assetTypes=DASH_LIVE|DASH_CENC_PS4"},
   //{false, "&assetTypes=DASH_LIVE|DASH_TA"},
   //{false, "&assetTypes=DASH_TA|DASH_CENC"},
   //{false, "&assetTypes=DASH_TA|DASH_CENC_HDR10"},
   //{false, "&assetTypes=DASH_TA|DASH_CENC_PS4"},
   //{false, "&assetTypes=DASH_TA|DASH_LIVE"},
   //{true, "&assetTypes=DASH_CENC_HDR10|DASH_CENC_PRECON"},
   //{true, "&assetTypes=DASH_CENC_PRECON|DASH_CENC"},
   //{true, "&assetTypes=DASH_CENC_PRECON|DASH_CENC_HDR10"},
   //{true, "&assetTypes=DASH_CENC_PRECON|DASH_CENC_PS4"},
   //{true, "&assetTypes=DASH_CENC_PRECON|DASH_LIVE"},
   //{true, "&assetTypes=DASH_CENC_PRECON|DASH_TA"},
   //{true, "&assetTypes=DASH_CENC_PS4|DASH_CENC_PRECON"},
   //{true, "&assetTypes=DASH_CENC|DASH_CENC_PRECON"},
   //{true, "&assetTypes=DASH_LIVE|DASH_CENC_PRECON"},
   //{true, "&assetTypes=DASH_TA|DASH_CENC_PRECON"},
   {false, ""},
   {false, "&assetTypes=DASH_TA"},
   {false, "&assetTypes=DASH_CENC"},
   {false, "&assetTypes=DASH_LIVE"},
   {false, "&assetTypes=DASH_CENC_PS4"},
   {false, "&assetTypes=DASH_CENC_HDR10"},
   {false, "&assetTypes=DASH_CENC_PRECON"},
}

// ""
func TestFrance(t *testing.T) {
   for _, asset := range asset_types {
      err := get(france, asset.value)
      fmt.Printf("%v %q\n", err, asset.value)
      time.Sleep(time.Second)
   }
}

// DASH_CENC
func TestUnitedStates(t *testing.T) {
   for _, asset := range asset_types {
      ok := func() bool {
         for _, id := range united_states {
            err := get(id, asset.value)
            fmt.Printf("%v %v %q\n", err, id, asset.value)
            if err != nil {
               return false
            }
            time.Sleep(time.Second)
         }
         return true
      }()
      if ok {
         break
      }
      fmt.Println()
   }
}

var united_states = []string{
   "Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
   "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   "rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
}
