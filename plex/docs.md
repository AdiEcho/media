# Overview

package `plex`

## Index

- [Types](#types)
  - [type Address](#type-address)
    - [func (a \*Address) Set(s string) error](#func-address-set)
    - [func (a \*Address) String() string](#func-address-string)
  - [type Anonymous](#type-anonymous)
    - [func (a \*Anonymous) Match(web \*Address) (\*DiscoverMatch, error)](#func-anonymous-match)
    - [func (a \*Anonymous) New() error](#func-anonymous-new)
    - [func (a \*Anonymous) Video(
  match \*DiscoverMatch, forward string,
) (\*OnDemand, error)](#func-anonymous-video)
  - [type DiscoverMatch](#type-discovermatch)
  - [type MediaPart](#type-mediapart)
    - [func (\*MediaPart) RequestHeader() (http.Header, error)](#func-mediapart-requestheader)
    - [func (m \*MediaPart) RequestUrl() (string, bool)](#func-mediapart-requesturl)
    - [func (\*MediaPart) UnwrapResponse(b []byte) ([]byte, error)](#func-mediapart-unwrapresponse)
    - [func (\*MediaPart) WrapRequest(b []byte) ([]byte, error)](#func-mediapart-wraprequest)
  - [type Namer](#type-namer)
    - [func (n Namer) Episode() int](#func-namer-episode)
    - [func (n Namer) Season() int](#func-namer-season)
    - [func (n Namer) Show() string](#func-namer-show)
    - [func (n Namer) Title() string](#func-namer-title)
    - [func (n Namer) Year() int](#func-namer-year)
  - [type OnDemand](#type-ondemand)
    - [func (o \*OnDemand) Dash() (\*MediaPart, bool)](#func-ondemand-dash)
  - [type Url](#type-url)
    - [func (u \*Url) UnmarshalText(text []byte) error](#func-url-unmarshaltext)
- [Source files](#source-files)

## Types

### type [Address](./plex.go#L42)

```go
type Address struct {
  Path string
}
```

### func (\*Address) [Set](./plex.go#L50)

```go
func (a *Address) Set(s string) error
```

### func (\*Address) [String](./plex.go#L46)

```go
func (a *Address) String() string
```

### type [Anonymous](./user.go#L10)

```go
type Anonymous struct {
  AuthToken string
}
```

### func (\*Anonymous) [Match](./user.go#L34)

```go
func (a *Anonymous) Match(web *Address) (*DiscoverMatch, error)
```

### func (\*Anonymous) [New](./user.go#L14)

```go
func (a *Anonymous) New() error
```

### func (\*Anonymous) [Video](./user.go#L66)

```go
func (a *Anonymous) Video(
  match *DiscoverMatch, forward string,
) (*OnDemand, error)
```

### type [DiscoverMatch](./plex.go#L9)

```go
type DiscoverMatch struct {
  GrandparentTitle string
  Index            int
  ParentIndex      int
  RatingKey        string
  Title            string
  Year             int
}
```

### type [MediaPart](./plex.go#L102)

```go
type MediaPart struct {
  Key     Url
  License *Url
}
```

### func (\*MediaPart) [RequestHeader](./plex.go#L72)

```go
func (*MediaPart) RequestHeader() (http.Header, error)
```

### func (\*MediaPart) [RequestUrl](./plex.go#L107)

```go
func (m *MediaPart) RequestUrl() (string, bool)
```

### func (\*MediaPart) [UnwrapResponse](./plex.go#L76)

```go
func (*MediaPart) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*MediaPart) [WrapRequest](./plex.go#L68)

```go
func (*MediaPart) WrapRequest(b []byte) ([]byte, error)
```

### type [Namer](./plex.go#L18)

```go
type Namer struct {
  Match *DiscoverMatch
}
```

### func (Namer) [Episode](./plex.go#L22)

```go
func (n Namer) Episode() int
```

### func (Namer) [Season](./plex.go#L26)

```go
func (n Namer) Season() int
```

### func (Namer) [Show](./plex.go#L30)

```go
func (n Namer) Show() string
```

### func (Namer) [Title](./plex.go#L34)

```go
func (n Namer) Title() string
```

### func (Namer) [Year](./plex.go#L38)

```go
func (n Namer) Year() int
```

### type [OnDemand](./plex.go#L80)

```go
type OnDemand struct {
  Media []struct {
    Part     []MediaPart
    Protocol string
  }
}
```

### func (\*OnDemand) [Dash](./plex.go#L57)

```go
func (o *OnDemand) Dash() (*MediaPart, bool)
```

### type [Url](./plex.go#L87)

```go
type Url struct {
  Url *url.URL
}
```

### func (\*Url) [UnmarshalText](./plex.go#L91)

```go
func (u *Url) UnmarshalText(text []byte) error
```

## Source files

[plex.go](./plex.go)
[user.go](./user.go)
