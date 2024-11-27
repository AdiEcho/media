# Overview

package `pluto`

## Index

- [Variables](#variables)
- [Types](#types)
  - [type Address](#type-address)
    - [func (a \*Address) Set(text string) error](#func-address-set)
    - [func (a Address) String() string](#func-address-string)
    - [func (a Address) Video(forward string) (\*OnDemand, error)](#func-address-video)
  - [type EpisodeClip](#type-episodeclip)
    - [func (e \*EpisodeClip) Dash() (\*url.URL, bool)](#func-episodeclip-dash)
  - [type FileBase](#type-filebase)
  - [type Namer](#type-namer)
    - [func (Namer) Episode() int](#func-namer-episode)
    - [func (Namer) Season() int](#func-namer-season)
    - [func (n Namer) Show() string](#func-namer-show)
    - [func (n Namer) Title() string](#func-namer-title)
    - [func (Namer) Year() int](#func-namer-year)
  - [type OnDemand](#type-ondemand)
    - [func (o OnDemand) Clip() (\*EpisodeClip, error)](#func-ondemand-clip)
  - [type Poster](#type-poster)
    - [func (Poster) RequestHeader() (http.Header, error)](#func-poster-requestheader)
    - [func (Poster) RequestUrl() (string, bool)](#func-poster-requesturl)
    - [func (Poster) UnwrapResponse(b []byte) ([]byte, error)](#func-poster-unwrapresponse)
    - [func (Poster) WrapRequest(b []byte) ([]byte, error)](#func-poster-wraprequest)
  - [type Url](#type-url)
    - [func (u \*Url) UnmarshalText(text []byte) error](#func-url-unmarshaltext)
  - [type VideoSeason](#type-videoseason)
- [Source files](#source-files)

## Variables

```go
var Base = []FileBase{
  {"http", "silo-hybrik.pluto.tv.s3.amazonaws.com", "200 OK"},
  {"http", "siloh-fs.plutotv.net", "403 OK"},
  {"http", "siloh-ns1.plutotv.net", "403 OK"},
  {"https", "siloh-fs.plutotv.net", "403 OK"},
  {"https", "siloh-ns1.plutotv.net", "403 OK"},
}
```

## Types

### type [Address](./pluto.go#L78)

```go
type Address [2]string
```

### func (\*Address) [Set](./pluto.go#L80)

```go
func (a *Address) Set(text string) error
```

### func (Address) [String](./pluto.go#L183)

```go
func (a Address) String() string
```

### func (Address) [Video](./pluto.go#L134)

```go
func (a Address) Video(forward string) (*OnDemand, error)
```

### type [EpisodeClip](./pluto.go#L104)

```go
type EpisodeClip struct {
  Sources []struct {
    File Url
    Type string
  }
}
```

### func (\*EpisodeClip) [Dash](./pluto.go#L111)

```go
func (e *EpisodeClip) Dash() (*url.URL, bool)
```

### type [FileBase](./pluto.go#L128)

```go
type FileBase struct {
  Scheme string
  Host   string
  Status string
}
```

### type [Namer](./pluto.go#L42)

```go
type Namer struct {
  Video *OnDemand
}
```

### func (Namer) [Episode](./pluto.go#L50)

```go
func (Namer) Episode() int
```

### func (Namer) [Season](./pluto.go#L46)

```go
func (Namer) Season() int
```

### func (Namer) [Show](./pluto.go#L58)

```go
func (n Namer) Show() string
```

### func (Namer) [Title](./pluto.go#L65)

```go
func (n Namer) Title() string
```

### func (Namer) [Year](./pluto.go#L54)

```go
func (Namer) Year() int
```

### type [OnDemand](./pluto.go#L69)

```go
type OnDemand struct {
  Episode string `json:"_id"`
  Id      string
  Name    string
  Seasons []*VideoSeason
  Slug    string
  // contains filtered or unexported fields
}
```

### func (OnDemand) [Clip](./pluto.go#L199)

```go
func (o OnDemand) Clip() (*EpisodeClip, error)
```

### type [Poster](./pluto.go#L11)

```go
type Poster struct{}
```

### func (Poster) [RequestHeader](./pluto.go#L17)

```go
func (Poster) RequestHeader() (http.Header, error)
```

### func (Poster) [RequestUrl](./pluto.go#L13)

```go
func (Poster) RequestUrl() (string, bool)
```

### func (Poster) [UnwrapResponse](./pluto.go#L25)

```go
func (Poster) UnwrapResponse(b []byte) ([]byte, error)
```

### func (Poster) [WrapRequest](./pluto.go#L21)

```go
func (Poster) WrapRequest(b []byte) ([]byte, error)
```

### type [Url](./pluto.go#L29)

```go
type Url struct {
  Url url.URL
}
```

### func (\*Url) [UnmarshalText](./pluto.go#L33)

```go
func (u *Url) UnmarshalText(text []byte) error
```

### type [VideoSeason](./pluto.go#L37)

```go
type VideoSeason struct {
  Episodes []*OnDemand
  // contains filtered or unexported fields
}
```

## Source files

[pluto.go](./pluto.go)
