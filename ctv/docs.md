# Overview

package `ctv`

## Index

- [Types](#types)
  - [type Address](#type-address)
    - [func (a Address) Resolve() (\*ResolvePath, error)](#func-address-resolve)
    - [func (a \*Address) Set(s string) error](#func-address-set)
    - [func (a \*Address) String() string](#func-address-string)
  - [type AxisContent](#type-axiscontent)
    - [func (a \*AxisContent) Manifest(media \*MediaContent) (string, error)](#func-axiscontent-manifest)
  - [type Date](#type-date)
    - [func (d \*Date) MarshalText() ([]byte, error)](#func-date-marshaltext)
    - [func (d \*Date) UnmarshalText(text []byte) error](#func-date-unmarshaltext)
  - [type MediaContent](#type-mediacontent)
    - [func (\*MediaContent) Marshal(axis \*AxisContent) ([]byte, error)](#func-mediacontent-marshal)
    - [func (m \*MediaContent) Unmarshal(data []byte) error](#func-mediacontent-unmarshal)
  - [type Namer](#type-namer)
    - [func (n \*Namer) Episode() int](#func-namer-episode)
    - [func (n \*Namer) Season() int](#func-namer-season)
    - [func (n \*Namer) Show() string](#func-namer-show)
    - [func (n \*Namer) Title() string](#func-namer-title)
    - [func (n \*Namer) Year() int](#func-namer-year)
  - [type Poster](#type-poster)
    - [func (Poster) RequestHeader() (http.Header, error)](#func-poster-requestheader)
    - [func (Poster) RequestUrl() (string, bool)](#func-poster-requesturl)
    - [func (Poster) UnwrapResponse(b []byte) ([]byte, error)](#func-poster-unwrapresponse)
    - [func (Poster) WrapRequest(b []byte) ([]byte, error)](#func-poster-wraprequest)
  - [type ResolvePath](#type-resolvepath)
    - [func (r \*ResolvePath) Axis() (\*AxisContent, error)](#func-resolvepath-axis)
- [Source files](#source-files)

## Types

### type [Address](./ctv.go#L53)

```go
type Address struct {
  Path string
}
```

### func (Address) [Resolve](./ctv.go#L69)

```go
func (a Address) Resolve() (*ResolvePath, error)
```

### func (\*Address) [Set](./ctv.go#L58)

```go
func (a *Address) Set(s string) error
```

https://www.ctv.ca/shows/friends/the-one-with-the-bullies-s2e21

### func (\*Address) [String](./ctv.go#L65)

```go
func (a *Address) String() string
```

### type [AxisContent](./ctv.go#L153)

```go
type AxisContent struct {
  AxisId                int64
  AxisPlaybackLanguages []struct {
    DestinationCode string
  }
}
```

### func (\*AxisContent) [Manifest](./ctv.go#L120)

```go
func (a *AxisContent) Manifest(media *MediaContent) (string, error)
```

hard geo block

### type [Date](./ctv.go#L160)

```go
type Date struct {
  Time time.Time
}
```

### func (\*Date) [MarshalText](./ctv.go#L173)

```go
func (d *Date) MarshalText() ([]byte, error)
```

### func (\*Date) [UnmarshalText](./ctv.go#L164)

```go
func (d *Date) UnmarshalText(text []byte) error
```

### type [MediaContent](./ctv.go#L177)

```go
type MediaContent struct {
  BroadcastDate   Date
  ContentPackages []struct {
    Id int64
  }
  Episode int
  Media   struct {
    Name string
    Type string
  }
  Name   string
  Season struct {
    Number int
  }
}
```

### func (\*MediaContent) [Marshal](./ctv.go#L197)

```go
func (*MediaContent) Marshal(axis *AxisContent) ([]byte, error)
```

### func (\*MediaContent) [Unmarshal](./ctv.go#L193)

```go
func (m *MediaContent) Unmarshal(data []byte) error
```

### type [Namer](./ctv.go#L218)

```go
type Namer struct {
  Media MediaContent
}
```

### func (\*Namer) [Episode](./ctv.go#L237)

```go
func (n *Namer) Episode() int
```

### func (\*Namer) [Season](./ctv.go#L222)

```go
func (n *Namer) Season() int
```

### func (\*Namer) [Show](./ctv.go#L226)

```go
func (n *Namer) Show() string
```

### func (\*Namer) [Title](./ctv.go#L241)

```go
func (n *Namer) Title() string
```

### func (\*Namer) [Year](./ctv.go#L233)

```go
func (n *Namer) Year() int
```

### type [Poster](./ctv.go#L248)

```go
type Poster struct{}
```

### func (Poster) [RequestHeader](./ctv.go#L250)

```go
func (Poster) RequestHeader() (http.Header, error)
```

### func (Poster) [RequestUrl](./ctv.go#L254)

```go
func (Poster) RequestUrl() (string, bool)
```

### func (Poster) [UnwrapResponse](./ctv.go#L262)

```go
func (Poster) UnwrapResponse(b []byte) ([]byte, error)
```

### func (Poster) [WrapRequest](./ctv.go#L258)

```go
func (Poster) WrapRequest(b []byte) ([]byte, error)
```

### type [ResolvePath](./ctv.go#L273)

```go
type ResolvePath struct {
  Id                   string
  FirstPlayableContent *struct {
    Id string
  }
}
```

### func (\*ResolvePath) [Axis](./ctv.go#L280)

```go
func (r *ResolvePath) Axis() (*AxisContent, error)
```

## Source files

[ctv.go](./ctv.go)
