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

func TestSlyGuy(t *testing.T) {
   for _, a := range sly_guy {
      for _, b := range sly_guy {
         fmt.Print(a, "|", b, "\n")
      }
   }
}

// github.com/matthuisman/slyguy.addons/blob/master/slyguy.paramount.plus/resources/lib/api.py
var sly_guy = []string{
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
   {true, ""},
   {true, "&assetTypes=DASH_CENC_PRECON"},
   {true, "&assetTypes=DASH_CENC|DASH_CENC_PRECON"},
   {true, "&assetTypes=DASH_CENC_PRECON|DASH_CENC"},
   {true, "&assetTypes=DASH_CENC_HDR10|DASH_CENC_PRECON"},
   {true, "&assetTypes=DASH_CENC_PRECON|DASH_CENC_HDR10"},
   {false, "&assetTypes=DASH_CENC"},
   {false, "&assetTypes=DASH_CENC_HDR10"},
   {false, "&assetTypes=DASH_CENC|DASH_CENC_HDR10"},
   {false, "&assetTypes=DASH_CENC_HDR10|DASH_CENC"},
}

const france = "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ"

func TestFrance(t *testing.T) {
   for _, asset := range asset_types {
      err := get(france, asset.value)
      fmt.Printf("%v %q\n", err, asset.value)
      time.Sleep(time.Second)
   }
}

var united_states = []string{
   "Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
   "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   "rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
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

func TestUnitedStates(t *testing.T) {
   for _, asset := range asset_types {
      if asset.france {
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
}
