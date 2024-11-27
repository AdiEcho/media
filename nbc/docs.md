# Overview

package `nbc`

## Index

- [Types](#types)
  - [type CoreVideo](#type-corevideo)
    - [func (c \*CoreVideo) New()](#func-corevideo-new)
    - [func (\*CoreVideo) RequestHeader() (http.Header, error)](#func-corevideo-requestheader)
    - [func (c \*CoreVideo) RequestUrl() (string, bool)](#func-corevideo-requesturl)
    - [func (\*CoreVideo) UnwrapResponse(b []byte) ([]byte, error)](#func-corevideo-unwrapresponse)
    - [func (\*CoreVideo) WrapRequest(b []byte) ([]byte, error)](#func-corevideo-wraprequest)
  - [type Metadata](#type-metadata)
    - [func (m \*Metadata) Episode() int](#func-metadata-episode)
    - [func (m \*Metadata) New(guid int) error](#func-metadata-new)
    - [func (m \*Metadata) OnDemand() (\*OnDemand, error)](#func-metadata-ondemand)
    - [func (m \*Metadata) Season() int](#func-metadata-season)
    - [func (m \*Metadata) Show() string](#func-metadata-show)
    - [func (m \*Metadata) Title() string](#func-metadata-title)
    - [func (m \*Metadata) Year() int](#func-metadata-year)
  - [type OnDemand](#type-ondemand)
- [Source files](#source-files)

## Types

### type [CoreVideo](./nbc.go#L53)

```go
type CoreVideo struct {
  DrmProxyUrl    string
  DrmProxySecret string
}
```

### func (\*CoreVideo) [New](./nbc.go#L70)

```go
func (c *CoreVideo) New()
```

### func (\*CoreVideo) [RequestHeader](./nbc.go#L96)

```go
func (*CoreVideo) RequestHeader() (http.Header, error)
```

### func (\*CoreVideo) [RequestUrl](./nbc.go#L58)

```go
func (c *CoreVideo) RequestUrl() (string, bool)
```

### func (\*CoreVideo) [UnwrapResponse](./nbc.go#L106)

```go
func (*CoreVideo) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*CoreVideo) [WrapRequest](./nbc.go#L102)

```go
func (*CoreVideo) WrapRequest(b []byte) ([]byte, error)
```

### type [Metadata](./metadata.go#L13)

```go
type Metadata struct {
  AirDate          time.Time
  EpisodeNumber    int `json:",string"`
  MovieShortTitle  string
  MpxAccountId     int64 `json:",string"`
  MpxGuid          int64 `json:",string"`
  ProgrammingType  string
  SeasonNumber     int `json:",string"`
  SecondaryTitle   string
  SeriesShortTitle string
}
```

### func (\*Metadata) [Episode](./metadata.go#L33)

```go
func (m *Metadata) Episode() int
```

### func (\*Metadata) [New](./metadata.go#L80)

```go
func (m *Metadata) New(guid int) error
```

### func (\*Metadata) [OnDemand](./metadata.go#L48)

```go
func (m *Metadata) OnDemand() (*OnDemand, error)
```

### func (\*Metadata) [Season](./metadata.go#L29)

```go
func (m *Metadata) Season() int
```

### func (\*Metadata) [Show](./metadata.go#L25)

```go
func (m *Metadata) Show() string
```

### func (\*Metadata) [Title](./metadata.go#L41)

```go
func (m *Metadata) Title() string
```

### func (\*Metadata) [Year](./metadata.go#L37)

```go
func (m *Metadata) Year() int
```

### type [OnDemand](./nbc.go#L80)

```go
type OnDemand struct {
  PlaybackUrl string
}
```

## Source files

[metadata.go](./metadata.go)
[nbc.go](./nbc.go)
