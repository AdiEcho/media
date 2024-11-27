# Overview

package `hulu`

## Index

- [Types](#types)
  - [type Authenticate](#type-authenticate)
    - [func (a \*Authenticate) DeepLink(id \*EntityId) (\*DeepLink, error)](#func-authenticate-deeplink)
    - [func (a \*Authenticate) Details(link \*DeepLink) (\*Details, error)](#func-authenticate-details)
    - [func (Authenticate) Marshal(email, password string) ([]byte, error)](#func-authenticate-marshal)
    - [func (a \*Authenticate) Playlist(link \*DeepLink) (\*Playlist, error)](#func-authenticate-playlist)
    - [func (a \*Authenticate) Unmarshal(data []byte) error](#func-authenticate-unmarshal)
  - [type DeepLink](#type-deeplink)
  - [type Details](#type-details)
    - [func (d \*Details) Episode() int](#func-details-episode)
    - [func (d \*Details) Season() int](#func-details-season)
    - [func (d \*Details) Show() string](#func-details-show)
    - [func (d \*Details) Title() string](#func-details-title)
    - [func (d \*Details) Year() int](#func-details-year)
  - [type EntityId](#type-entityid)
    - [func (e \*EntityId) Set(s string) error](#func-entityid-set)
    - [func (e \*EntityId) String() string](#func-entityid-string)
  - [type Playlist](#type-playlist)
    - [func (\*Playlist) RequestHeader() (http.Header, error)](#func-playlist-requestheader)
    - [func (p \*Playlist) RequestUrl() (string, bool)](#func-playlist-requesturl)
    - [func (\*Playlist) UnwrapResponse(b []byte) ([]byte, error)](#func-playlist-unwrapresponse)
    - [func (\*Playlist) WrapRequest(b []byte) ([]byte, error)](#func-playlist-wraprequest)
- [Source files](#source-files)

## Types

### type [Authenticate](./authenticate.go#L174)

```go
type Authenticate struct {
  Data struct {
    UserToken string `json:"user_token"`
  }
}
```

### func (\*Authenticate) [DeepLink](./authenticate.go#L145)

```go
func (a *Authenticate) DeepLink(id *EntityId) (*DeepLink, error)
```

### func (\*Authenticate) [Details](./authenticate.go#L92)

```go
func (a *Authenticate) Details(link *DeepLink) (*Details, error)
```

### func (Authenticate) [Marshal](./authenticate.go#L124)

```go
func (Authenticate) Marshal(email, password string) ([]byte, error)
```

### func (\*Authenticate) [Playlist](./authenticate.go#L13)

```go
func (a *Authenticate) Playlist(link *DeepLink) (*Playlist, error)
```

### func (\*Authenticate) [Unmarshal](./authenticate.go#L180)

```go
func (a *Authenticate) Unmarshal(data []byte) error
```

### type [DeepLink](./hulu.go#L30)

```go
type DeepLink struct {
  EabId string `json:"eab_id"`
}
```

### type [Details](./hulu.go#L34)

```go
type Details struct {
  EpisodeName   string `json:"episode_name"`
  EpisodeNumber int    `json:"episode_number"`
  Headline      string
  PremiereDate  time.Time `json:"premiere_date"`
  SeasonNumber  int       `json:"season_number"`
  SeriesName    string    `json:"series_name"`
}
```

### func (\*Details) [Episode](./hulu.go#L51)

```go
func (d *Details) Episode() int
```

### func (\*Details) [Season](./hulu.go#L47)

```go
func (d *Details) Season() int
```

### func (\*Details) [Show](./hulu.go#L43)

```go
func (d *Details) Show() string
```

### func (\*Details) [Title](./hulu.go#L59)

```go
func (d *Details) Title() string
```

### func (\*Details) [Year](./hulu.go#L55)

```go
func (d *Details) Year() int
```

### type [EntityId](./hulu.go#L66)

```go
type EntityId struct {
  Text string
}
```

### func (\*EntityId) [Set](./hulu.go#L75)

```go
func (e *EntityId) Set(s string) error
```

hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d

### func (\*EntityId) [String](./hulu.go#L70)

```go
func (e *EntityId) String() string
```

### type [Playlist](./hulu.go#L9)

```go
type Playlist struct {
  StreamUrl string `json:"stream_url"`
  WvServer  string `json:"wv_server"`
}
```

### func (\*Playlist) [RequestHeader](./hulu.go#L22)

```go
func (*Playlist) RequestHeader() (http.Header, error)
```

### func (\*Playlist) [RequestUrl](./hulu.go#L14)

```go
func (p *Playlist) RequestUrl() (string, bool)
```

### func (\*Playlist) [UnwrapResponse](./hulu.go#L26)

```go
func (*Playlist) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*Playlist) [WrapRequest](./hulu.go#L18)

```go
func (*Playlist) WrapRequest(b []byte) ([]byte, error)
```

## Source files

[authenticate.go](./authenticate.go)
[hulu.go](./hulu.go)
