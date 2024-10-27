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
var asset_types = []string{
   "",
   "DASH_CENC",
   "DASH_CENC_HDR10",
   "DASH_CENC_PRECON",
   "DASH_CENC_PS4",
   "DASH_LIVE",
   "DASH_TA",
}

var united_states = []string{
   "Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
   "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   "rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
}

const france = "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ"

// formats=MPEG-DASH
// passes response status, Period and ContentProtection tests
func TestFrance(t *testing.T) {
   for _, asset := range asset_types {
      err := get(france, asset)
      fmt.Printf("%v %q\n", err, asset)
      if err == nil {
         break
      }
      time.Sleep(time.Second)
   }
}

// formats=MPEG-DASH&assetTypes=DASH_CENC
// passes response status, Period and ContentProtection tests
func TestUnitedStates(t *testing.T) {
   for _, asset := range asset_types {
      ok := func() bool {
         for _, id := range united_states {
            err := get(id, asset)
            fmt.Printf("%v %v %q\n", err, id, asset)
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
