package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "mt.ssai-oneapp.nbcuni.com"
   req.URL.Path = "/prod/nbc/gLU/RcQ/9000283422/1698569087378-MEWw4/cmaf/mpeg_cenc_2sec/master_cmaf.mpd"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   val := make(url.Values)
   val["mt.config"] = []string{"oneapp-atp-dash-vod-2s-generic"}
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

var req_body = strings.NewReader(`
{
 "reportingMode": "client",
 "availSuppression": {
  "mode": "BEHIND_LIVE_EDGE",
  "value": "00:00:00"
 },
 "playerParams": {
  "origin_domain": "vod-lf-oneapp2-prd.akamaized.net"
 },
 "adsParams": {
  "caid": "NBC_VOD_9000283422",
  "mode": "on-demand",
  "pvrn": "394767237",
  "vprn": "572200546",
  "vdur": "4067",
  "vip": "72.181.23.38",
  "sfid": "9244655",
  "csid": "oneapp_desktop_computer_web_ondemand",
  "crtp": "vast3ap",
  "nw": "169843",
  "prof": "169843%3Aoneapp_web_vod_mt_atp",
  "metr": "1031",
  "flag": "%2Bsltp%2Bemcr%2Bslcb%2Bsbid-fbad%2Baeti%2Bslif%2Bvicb%2Bamcb%2Bplay-uapl%2Bdtrd",
  "resp": "vmap1",
  "afid": "200265138",
  "vcid": "-1",
  "am_abrspec": "not_required",
  "am_appv": "1.224.3",
  "am_buildv": "1.224.3",
  "am_cdn": "akamai",
  "am_crmid": "-1",
  "am_crmid_type": "mParticleId",
  "am_pubid": "-1",
  "am_playerv": "v3.0.5-v58.hotfix",
  "am_sdkv": "2.12.1-peacock",
  "_fw_player_height": "1080",
  "_fw_player_width": "1920",
  "am_pub": "nbcu",
  "am_appname": "oneapp",
  "am_programtype": "television",
  "am_nielsen_genre": "",
  "nielsen_device_group": "",
  "nielsen_platform": "",
  "tms_id": "",
  "gc_id": "",
  "_fw_h_user_agent": "Mozilla%2F5.0%20(Windows%20NT%2010.0%3B%20Win64%3B%20x64%3B%20rv%3A101.0)%20Gecko%2F20100101%20Firefox%2F101.0",
  "_fw_app_bundle": "",
  "_fw_nielsen_app_id": "PAD3C6E72-ED61-417F-A865-3AB63FDB6197",
  "_fw_cdn_provider": "nbcu_akamai",
  "_fw_vcid2": "-1",
  "_fw_h_referer": "https%3A%2F%2Fwww.nbc.com%2Fsaturday-night-live%2Fvideo%2Foctober-21-bad-bunny%2F9000283422",
  "_fw_coppa": "0",
  "_fw_ae": "",
  "am_bc": "0",
  "mini": "false",
  "am_sst": "fullEpisodePlayer",
  "pl": "n%2Fa",
  "bl_enabled": "false",
  "am_brand": "nbc",
  "enable_pabi": "false",
  "_fw_content_language": "en",
  "yo.nl": "none"
 }
}
`)
