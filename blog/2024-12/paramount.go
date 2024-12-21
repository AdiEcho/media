package main

import (
)

func get(content_id, asset_type string) error {
   req, err := http.NewRequest("", "http://link.theplatform.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/s/dJ5BDC/media/guid/2198311517/" + content_id
   req.URL.RawQuery = url.Values{
      "assetTypes": {asset_type},
      "formats": {"MPEG-DASH"},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return fmt.Errorf("Status = %v", resp.Status)
   }
   var data string
   {
      b, err := io.ReadAll(resp.Body)
      if err != nil {
         return err
      }
      data = string(b)
   }
   count := strings.Count(data, "<Period ")
   if count != 1 {
      return fmt.Errorf("Period = %v", count)
   }
   return nil
}

var asset_types = []string{
   "DASH_LIVE",
   "DASH_CENC_HDR10",
   "DASH_TA",
   "DASH_CENC",
   "DASH_CENC_PRECON",
   "DASH_CENC_PS4",
}

func main() {
   content_id := flag.String("c", "", "content ID")
   flag.Parse()
   if *content_id != "" {
      err := get(*content_id)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
