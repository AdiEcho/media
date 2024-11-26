# Overview

package `draken`

## Index

- [Types](#types)
  - [type AuthLogin](#type-authlogin)
    - [func (a \*AuthLogin) Entitlement(movie \*FullMovie) (\*Entitlement, error)](#func-authlogin-entitlement)
    - [func (AuthLogin) Marshal(identity, key string) ([]byte, error)](#func-authlogin-marshal)
    - [func (a \*AuthLogin) Playback(
  movie \*FullMovie, title \*Entitlement,
) (\*Playback, error)](#func-authlogin-playback)
    - [func (a \*AuthLogin) Unmarshal(data []byte) error](#func-authlogin-unmarshal)
  - [type Entitlement](#type-entitlement)
  - [type FullMovie](#type-fullmovie)
    - [func (f \*FullMovie) New(custom_id string) error](#func-fullmovie-new)
  - [type Namer](#type-namer)
    - [func (\*Namer) Episode() int](#func-namer-episode)
    - [func (\*Namer) Season() int](#func-namer-season)
    - [func (\*Namer) Show() string](#func-namer-show)
    - [func (n \*Namer) Title() string](#func-namer-title)
    - [func (n \*Namer) Year() int](#func-namer-year)
  - [type Playback](#type-playback)
  - [type Poster](#type-poster)
    - [func (p \*Poster) RequestHeader() (http.Header, error)](#func-poster-requestheader)
    - [func (\*Poster) RequestUrl() (string, bool)](#func-poster-requesturl)
    - [func (\*Poster) UnwrapResponse(b []byte) ([]byte, error)](#func-poster-unwrapresponse)
    - [func (\*Poster) WrapRequest(b []byte) ([]byte, error)](#func-poster-wraprequest)
- [Source files](#source-files)

## Types

### type [AuthLogin](./draken.go#L89)

```go
type AuthLogin struct {
  Token string
}
```

### func (\*AuthLogin) [Entitlement](./draken.go#L97)

```go
func (a *AuthLogin) Entitlement(movie *FullMovie) (*Entitlement, error)
```

### func (AuthLogin) [Marshal](./draken.go#L70)

```go
func (AuthLogin) Marshal(identity, key string) ([]byte, error)
```

### func (\*AuthLogin) [Playback](./draken.go#L34)

```go
func (a *AuthLogin) Playback(
  movie *FullMovie, title *Entitlement,
) (*Playback, error)
```

### func (\*AuthLogin) [Unmarshal](./draken.go#L93)

```go
func (a *AuthLogin) Unmarshal(data []byte) error
```

### type [Entitlement](./draken.go#L123)

```go
type Entitlement struct {
  Token string
}
```

### type [FullMovie](./draken.go#L172)

```go
type FullMovie struct {
  DefaultPlayable struct {
    Id string
  }
  ProductionYear int `json:",string"`
  Title          string
}
```

### func (\*FullMovie) [New](./draken.go#L127)

```go
func (f *FullMovie) New(custom_id string) error
```

### type [Namer](./draken.go#L200)

```go
type Namer struct {
  Movie FullMovie
}
```

### func (\*Namer) [Episode](./draken.go#L188)

```go
func (*Namer) Episode() int
```

### func (\*Namer) [Season](./draken.go#L192)

```go
func (*Namer) Season() int
```

### func (\*Namer) [Show](./draken.go#L196)

```go
func (*Namer) Show() string
```

### func (\*Namer) [Title](./draken.go#L180)

```go
func (n *Namer) Title() string
```

### func (\*Namer) [Year](./draken.go#L184)

```go
func (n *Namer) Year() int
```

### type [Playback](./draken.go#L204)

```go
type Playback struct {
  Headers  map[string]string
  Playlist string
}
```

### type [Poster](./draken.go#L209)

```go
type Poster struct {
  Login AuthLogin
  Play  *Playback
}
```

### func (\*Poster) [RequestHeader](./draken.go#L214)

```go
func (p *Poster) RequestHeader() (http.Header, error)
```

### func (\*Poster) [RequestUrl](./draken.go#L224)

```go
func (*Poster) RequestUrl() (string, bool)
```

### func (\*Poster) [UnwrapResponse](./draken.go#L228)

```go
func (*Poster) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*Poster) [WrapRequest](./draken.go#L232)

```go
func (*Poster) WrapRequest(b []byte) ([]byte, error)
```

## Source files

[draken.go](./draken.go)
