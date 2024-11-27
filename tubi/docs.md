# Overview

package `tubi`

## Index

- [Types](#types)
  - [type Namer](#type-namer)
    - [func (n Namer) Episode() int](#func-namer-episode)
    - [func (n Namer) Season() int](#func-namer-season)
    - [func (n Namer) Show() string](#func-namer-show)
    - [func (n Namer) Title() string](#func-namer-title)
    - [func (n Namer) Year() int](#func-namer-year)
  - [type Resolution](#type-resolution)
    - [func (r Resolution) MarshalText() ([]byte, error)](#func-resolution-marshaltext)
    - [func (r \*Resolution) UnmarshalText(text []byte) error](#func-resolution-unmarshaltext)
  - [type VideoContent](#type-videocontent)
    - [func (v \*VideoContent) Episode() bool](#func-videocontent-episode)
    - [func (v \*VideoContent) Get(id int) (\*VideoContent, bool)](#func-videocontent-get)
    - [func (\*VideoContent) Marshal(id int) ([]byte, error)](#func-videocontent-marshal)
    - [func (v \*VideoContent) Series() bool](#func-videocontent-series)
    - [func (v \*VideoContent) Unmarshal(data []byte) error](#func-videocontent-unmarshal)
    - [func (v \*VideoContent) Video() (\*VideoResource, bool)](#func-videocontent-video)
  - [type VideoResource](#type-videoresource)
    - [func (\*VideoResource) RequestHeader() (http.Header, error)](#func-videoresource-requestheader)
    - [func (v \*VideoResource) RequestUrl() (string, bool)](#func-videoresource-requesturl)
    - [func (\*VideoResource) UnwrapResponse(b []byte) ([]byte, error)](#func-videoresource-unwrapresponse)
    - [func (\*VideoResource) WrapRequest(b []byte) ([]byte, error)](#func-videoresource-wraprequest)
- [Source files](#source-files)

## Types

### type [Namer](./tubi.go#L9)

```go
type Namer struct {
  Content *VideoContent
}
```

### func (Namer) [Episode](./tubi.go#L13)

```go
func (n Namer) Episode() int
```

### func (Namer) [Season](./tubi.go#L17)

```go
func (n Namer) Season() int
```

### func (Namer) [Show](./tubi.go#L24)

```go
func (n Namer) Show() string
```

### func (Namer) [Title](./tubi.go#L32)

```go
func (n Namer) Title() string
```

S01:E03 - Hell Hath No Fury

### func (Namer) [Year](./tubi.go#L39)

```go
func (n Namer) Year() int
```

### type [Resolution](./tubi.go#L43)

```go
type Resolution struct {
  Int64 int64
}
```

### func (Resolution) [MarshalText](./tubi.go#L77)

```go
func (r Resolution) MarshalText() ([]byte, error)
```

### func (\*Resolution) [UnmarshalText](./tubi.go#L47)

```go
func (r *Resolution) UnmarshalText(text []byte) error
```

### type [VideoContent](./content.go#L51)

```go
type VideoContent struct {
  Children       []*VideoContent
  DetailedType   string `json:"detailed_type"`
  EpisodeNumber  int    `json:"episode_number,string"`
  Id             int    `json:",string"`
  SeriesId       int    `json:"series_id,string"`
  Title          string
  VideoResources []VideoResource `json:"video_resources"`
  Year           int
  // contains filtered or unexported fields
}
```

### func (\*VideoContent) [Episode](./content.go#L15)

```go
func (v *VideoContent) Episode() bool
```

### func (\*VideoContent) [Get](./content.go#L32)

```go
func (v *VideoContent) Get(id int) (*VideoContent, bool)
```

### func (\*VideoContent) [Marshal](./content.go#L72)

```go
func (*VideoContent) Marshal(id int) ([]byte, error)
```

### func (\*VideoContent) [Series](./content.go#L11)

```go
func (v *VideoContent) Series() bool
```

### func (\*VideoContent) [Unmarshal](./content.go#L63)

```go
func (v *VideoContent) Unmarshal(data []byte) error
```

### func (\*VideoContent) [Video](./content.go#L19)

```go
func (v *VideoContent) Video() (*VideoResource, bool)
```

### type [VideoResource](./tubi.go#L59)

```go
type VideoResource struct {
  LicenseServer *struct {
    Url string
  } `json:"license_server"`
  Manifest struct {
    Url string
  }
  Resolution Resolution
  Type       string
}
```

### func (\*VideoResource) [RequestHeader](./tubi.go#L83)

```go
func (*VideoResource) RequestHeader() (http.Header, error)
```

### func (\*VideoResource) [RequestUrl](./tubi.go#L70)

```go
func (v *VideoResource) RequestUrl() (string, bool)
```

### func (\*VideoResource) [UnwrapResponse](./tubi.go#L87)

```go
func (*VideoResource) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*VideoResource) [WrapRequest](./tubi.go#L91)

```go
func (*VideoResource) WrapRequest(b []byte) ([]byte, error)
```

## Source files

[content.go](./content.go)
[tubi.go](./tubi.go)
