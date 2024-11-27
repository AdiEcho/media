# Overview

package `paramount`

## Index

- [Types](#types)
  - [type AppToken](#type-apptoken)
    - [func (a \*AppToken) ComCbsApp() error](#func-apptoken-comcbsapp)
    - [func (a \*AppToken) ComCbsCa() error](#func-apptoken-comcbsca)
    - [func (a \*AppToken) New(app_secret string) error](#func-apptoken-new)
    - [func (a AppToken) Session(content_id string) (\*SessionToken, error)](#func-apptoken-session)
  - [type Number](#type-number)
    - [func (n Number) MarshalText() ([]byte, error)](#func-number-marshaltext)
    - [func (n \*Number) UnmarshalText(text []byte) error](#func-number-unmarshaltext)
  - [type SessionToken](#type-sessiontoken)
    - [func (s \*SessionToken) RequestHeader() (http.Header, error)](#func-sessiontoken-requestheader)
    - [func (s \*SessionToken) RequestUrl() (string, bool)](#func-sessiontoken-requesturl)
    - [func (\*SessionToken) UnwrapResponse(b []byte) ([]byte, error)](#func-sessiontoken-unwrapresponse)
    - [func (\*SessionToken) WrapRequest(b []byte) ([]byte, error)](#func-sessiontoken-wraprequest)
  - [type VideoItem](#type-videoitem)
    - [func (v \*VideoItem) Episode() int](#func-videoitem-episode)
    - [func (\*VideoItem) Marshal(token AppToken, cid string) ([]byte, error)](#func-videoitem-marshal)
    - [func (v \*VideoItem) Mpd() string](#func-videoitem-mpd)
    - [func (v \*VideoItem) Season() int](#func-videoitem-season)
    - [func (v \*VideoItem) Show() string](#func-videoitem-show)
    - [func (v \*VideoItem) Title() string](#func-videoitem-title)
    - [func (v \*VideoItem) Unmarshal(data []byte) error](#func-videoitem-unmarshal)
    - [func (v \*VideoItem) Year() int](#func-videoitem-year)
- [Source files](#source-files)

## Types

### type [AppToken](./paramount.go#L26)

```go
type AppToken struct {
  Values url.Values
}
```

### func (\*AppToken) [ComCbsApp](./paramount.go#L88)

```go
func (a *AppToken) ComCbsApp() error
```

15.0.28

### func (\*AppToken) [ComCbsCa](./paramount.go#L93)

```go
func (a *AppToken) ComCbsCa() error
```

15.0.28

### func (\*AppToken) [New](./paramount.go#L62)

```go
func (a *AppToken) New(app_secret string) error
```

### func (AppToken) [Session](./paramount.go#L31)

```go
func (a AppToken) Session(content_id string) (*SessionToken, error)
```

must use app token and IP address for US

### type [Number](./paramount.go#L112)

```go
type Number int64
```

### func (Number) [MarshalText](./paramount.go#L97)

```go
func (n Number) MarshalText() ([]byte, error)
```

### func (\*Number) [UnmarshalText](./paramount.go#L101)

```go
func (n *Number) UnmarshalText(text []byte) error
```

### type [SessionToken](./paramount.go#L126)

```go
type SessionToken struct {
  LsSession string `json:"ls_session"`
  Url       string
}
```

### func (\*SessionToken) [RequestHeader](./paramount.go#L131)

```go
func (s *SessionToken) RequestHeader() (http.Header, error)
```

### func (\*SessionToken) [RequestUrl](./paramount.go#L122)

```go
func (s *SessionToken) RequestUrl() (string, bool)
```

### func (\*SessionToken) [UnwrapResponse](./paramount.go#L114)

```go
func (*SessionToken) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*SessionToken) [WrapRequest](./paramount.go#L118)

```go
func (*SessionToken) WrapRequest(b []byte) ([]byte, error)
```

### type [VideoItem](./item.go#L42)

```go
type VideoItem struct {
  AirDateIso   time.Time `json:"_airDateISO"`
  CmsAccountId string
  ContentId    string
  EpisodeNum   Number
  Label        string
  MediaType    string
  SeasonNum    Number
  SeriesTitle  string
}
```

### func (\*VideoItem) [Episode](./item.go#L23)

```go
func (v *VideoItem) Episode() int
```

### func (\*VideoItem) [Marshal](./item.go#L73)

```go
func (*VideoItem) Marshal(token AppToken, cid string) ([]byte, error)
```

must use app token and IP address for correct location

### func (\*VideoItem) [Mpd](./mpd.go#L23)

```go
func (v *VideoItem) Mpd() string
```

hard geo block

### func (\*VideoItem) [Season](./item.go#L19)

```go
func (v *VideoItem) Season() int
```

### func (\*VideoItem) [Show](./item.go#L35)

```go
func (v *VideoItem) Show() string
```

### func (\*VideoItem) [Title](./item.go#L27)

```go
func (v *VideoItem) Title() string
```

### func (\*VideoItem) [Unmarshal](./item.go#L53)

```go
func (v *VideoItem) Unmarshal(data []byte) error
```

### func (\*VideoItem) [Year](./item.go#L31)

```go
func (v *VideoItem) Year() int
```

## Source files

[item.go](./item.go)
[mpd.go](./mpd.go)
[paramount.go](./paramount.go)
