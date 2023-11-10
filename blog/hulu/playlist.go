package hulu

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a authenticate) playlist(d deep_link) (*http.Response, error) {
   var p playlist_request
   p.Content_EAB_ID = d.EAB_ID
   p.Deejay_Device_ID = 166
   p.Token = a.Data.User_Token
   p.Unencrypted = true
   p.Version = 5012541
   p.Playback.Audio.Codecs.Selection_Mode = "ONE"
   //p.playback.audio.codecs.values[0].type = "AAC";
   //p.playback.drm.selection_mode = "ONE";
   //p.playback.drm.values[0].security_level = "L3";
   //p.playback.drm.values[0].type = "WIDEVINE";
   //p.playback.drm.values[0].version = "MODULAR";
   //p.playback.manifest.type = "DASH";
   //p.playback.segments.selection_mode = "ONE";
   //p.playback.segments.values[0].encryption.mode = "CENC";
   //p.playback.segments.values[0].encryption.type = "CENC";
   //p.playback.segments.values[0].type = "FMP4";
   //p.playback.version = 2;
   //p.playback.video.codecs.selection_mode = "FIRST";
   //p.playback.video.codecs.values[0].level = "5.2";
   //p.playback.video.codecs.values[0].profile = "HIGH";
   //p.playback.video.codecs.values[0].type = "H264";
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Scheme = "https"
   req.URL.Host = "play.hulu.com"
   req.URL.Path = "/v6/playlist"
   req.Body = io.NopCloser(strings.NewReader(playlist_body))
   req.Header["Content-Type"] = []string{"application/json"}
   return new(http.Transport).RoundTrip(&req)
}
