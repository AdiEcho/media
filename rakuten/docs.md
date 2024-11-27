# Overview

package `rakuten`

## Index

- [Types](#types)
  - [type Address](#type-address)
    - [func (a \*Address) Fhd() \*OnDemand](#func-address-fhd)
    - [func (a \*Address) Hd() \*OnDemand](#func-address-hd)
    - [func (a \*Address) Movie() (\*GizmoMovie, error)](#func-address-movie)
    - [func (a \*Address) Set(s string) error](#func-address-set)
    - [func (a \*Address) String() string](#func-address-string)
  - [type GizmoMovie](#type-gizmomovie)
    - [func (\*GizmoMovie) Episode() int](#func-gizmomovie-episode)
    - [func (\*GizmoMovie) Season() int](#func-gizmomovie-season)
    - [func (\*GizmoMovie) Show() string](#func-gizmomovie-show)
    - [func (g \*GizmoMovie) Title() string](#func-gizmomovie-title)
    - [func (g \*GizmoMovie) Year() int](#func-gizmomovie-year)
  - [type OnDemand](#type-ondemand)
    - [func (o \*OnDemand) Info() (\*StreamInfo, error)](#func-ondemand-info)
  - [type StreamInfo](#type-streaminfo)
    - [func (\*StreamInfo) RequestHeader() (http.Header, error)](#func-streaminfo-requestheader)
    - [func (s \*StreamInfo) RequestUrl() (string, bool)](#func-streaminfo-requesturl)
    - [func (\*StreamInfo) UnwrapResponse(b []byte) ([]byte, error)](#func-streaminfo-unwrapresponse)
    - [func (\*StreamInfo) WrapRequest(b []byte) ([]byte, error)](#func-streaminfo-wraprequest)
- [Source files](#source-files)

## Types

### type [Address](./address.go#L79)

```go
type Address struct {
  ClassificationId int
  ContentId        string
  MarketCode       string
}
```

### func (\*Address) [Fhd](./address.go#L75)

```go
func (a *Address) Fhd() *OnDemand
```

### func (\*Address) [Hd](./address.go#L71)

```go
func (a *Address) Hd() *OnDemand
```

### func (\*Address) [Movie](./address.go#L85)

```go
func (a *Address) Movie() (*GizmoMovie, error)
```

### func (\*Address) [Set](./address.go#L26)

```go
func (a *Address) Set(s string) error
```

### func (\*Address) [String](./address.go#L43)

```go
func (a *Address) String() string
```

### type [GizmoMovie](./rakuten.go#L10)

```go
type GizmoMovie struct {
  Data struct {
    Title string
    Year  int
  }
}
```

### func (\*GizmoMovie) [Episode](./rakuten.go#L88)

```go
func (*GizmoMovie) Episode() int
```

### func (\*GizmoMovie) [Season](./rakuten.go#L84)

```go
func (*GizmoMovie) Season() int
```

### func (\*GizmoMovie) [Show](./rakuten.go#L80)

```go
func (*GizmoMovie) Show() string
```

### func (\*GizmoMovie) [Title](./rakuten.go#L17)

```go
func (g *GizmoMovie) Title() string
```

### func (\*GizmoMovie) [Year](./rakuten.go#L21)

```go
func (g *GizmoMovie) Year() int
```

### type [OnDemand](./rakuten.go#L56)

```go
type OnDemand struct {
  AudioLanguage            string `json:"audio_language"`
  AudioQuality             string `json:"audio_quality"`
  ClassificationId         int    `json:"classification_id"`
  ContentId                string `json:"content_id"`
  ContentType              string `json:"content_type"`
  DeviceIdentifier         string `json:"device_identifier"`
  DeviceSerial             string `json:"device_serial"`
  DeviceStreamVideoQuality string `json:"device_stream_video_quality"`
  Player                   string `json:"player"`
  SubtitleLanguage         string `json:"subtitle_language"`
  VideoType                string `json:"video_type"`
}
```

### func (\*OnDemand) [Info](./rakuten.go#L26)

```go
func (o *OnDemand) Info() (*StreamInfo, error)
```

geo block

### type [StreamInfo](./rakuten.go#L70)

```go
type StreamInfo struct {
  LicenseUrl   string `json:"license_url"`
  Url          string
  VideoQuality string `json:"video_quality"`
}
```

### func (\*StreamInfo) [RequestHeader](./rakuten.go#L97)

```go
func (*StreamInfo) RequestHeader() (http.Header, error)
```

github.com/mitmproxy/mitmproxy/blob/main/mitmproxy/contentviews/protobuf.py

### func (\*StreamInfo) [RequestUrl](./rakuten.go#L76)

```go
func (s *StreamInfo) RequestUrl() (string, bool)
```

### func (\*StreamInfo) [UnwrapResponse](./rakuten.go#L103)

```go
func (*StreamInfo) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*StreamInfo) [WrapRequest](./rakuten.go#L92)

```go
func (*StreamInfo) WrapRequest(b []byte) ([]byte, error)
```

## Source files

[address.go](./address.go)
[rakuten.go](./rakuten.go)
