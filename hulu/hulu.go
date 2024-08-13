package hulu

import (
   "bytes"
   "errors"
   "encoding/json"
   "net/http"
   "path"
   "time"
)

type codec_value struct {
   Height int `json:"height,omitempty"`
   Level   string `json:"level,omitempty"`
   Profile string `json:"profile,omitempty"`
   Tier string `json:"tier,omitempty"`
   Type    string `json:"type"`
   Width int `json:"width,omitempty"`
}

type Playlist struct {
   StreamUrl string `json:"stream_url"`
   WvServer string `json:"wv_server"`
}

func (Playlist) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playlist) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (Playlist) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type DeepLink struct {
   EabId string `json:"eab_id"`
}

type drm_value struct {
   SecurityLevel string `json:"security_level"`
   Type          string `json:"type"`
   Version       string `json:"version"`
}

type playlist_request struct {
   ContentEabId   string `json:"content_eab_id"`
   DeejayDeviceId int    `json:"deejay_device_id"`
   Unencrypted    bool   `json:"unencrypted"`
   Version        int    `json:"version"`
   Playback       struct {
      Audio struct {
         Codecs struct {
            SelectionMode string `json:"selection_mode"`
            Values []codec_value `json:"values"`
         } `json:"codecs"`
      } `json:"audio"`
      Video   struct {
         Codecs struct {
            SelectionMode string `json:"selection_mode"`
            Values []codec_value `json:"values"`
         } `json:"codecs"`
      } `json:"video"`
      Drm struct {
         SelectionMode string `json:"selection_mode"`
         Values []drm_value `json:"values"`
      } `json:"drm"`
      Manifest struct {
         Type string `json:"type"`
      } `json:"manifest"`
      Segments struct {
         SelectionMode string `json:"selection_mode"`
         Values []segment_value `json:"values"`
      } `json:"segments"`
      Version int `json:"version"`
   } `json:"playback"`
}

type segment_value struct {
   Encryption struct {
      Mode string `json:"mode"`
      Type string `json:"type"`
   } `json:"encryption"`
   Type string `json:"type"`
}

type Details struct {
   EpisodeName string `json:"episode_name"`
   EpisodeNumber int `json:"episode_number"`
   Headline string
   PremiereDate time.Time `json:"premiere_date"`
   SeasonNumber int `json:"season_number"`
   SeriesName string `json:"series_name"`
}

type EntityId struct {
   s string
}

func (e *EntityId) String() string {
   return e.s
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
func (e *EntityId) Set(s string) error {
   e.s = path.Base(s)
   return nil
}

func (d *Details) Show() string {
   return d.SeriesName
}

func (d *Details) Season() int {
   return d.SeasonNumber
}

func (d *Details) Episode() int {
   return d.EpisodeNumber
}

func (d *Details) Year() int {
   return d.PremiereDate.Year()
}

func (p *Playlist) RequestUrl() (string, bool) {
   return p.WvServer, true
}

func (d *Details) Title() string {
   if d.EpisodeName != "" {
      return d.EpisodeName
   }
   return d.Headline
}
