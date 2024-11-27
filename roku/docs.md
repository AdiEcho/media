# Overview

package `roku`

## Index

- [Types](#types)
  - [type AccountAuth](#type-accountauth)
    - [func (\*AccountAuth) Marshal(token \*AccountToken) ([]byte, error)](#func-accountauth-marshal)
    - [func (a \*AccountAuth) Playback(roku_id string) (\*Playback, error)](#func-accountauth-playback)
    - [func (a \*AccountAuth) Unmarshal(data []byte) error](#func-accountauth-unmarshal)
  - [type AccountCode](#type-accountcode)
    - [func (\*AccountCode) Marshal(auth \*AccountAuth) ([]byte, error)](#func-accountcode-marshal)
    - [func (a \*AccountCode) String() string](#func-accountcode-string)
    - [func (a \*AccountCode) Unmarshal(data []byte) error](#func-accountcode-unmarshal)
  - [type AccountToken](#type-accounttoken)
    - [func (AccountToken) Marshal(
  auth \*AccountAuth, code \*AccountCode,
) ([]byte, error)](#func-accounttoken-marshal)
    - [func (a \*AccountToken) Unmarshal(data []byte) error](#func-accounttoken-unmarshal)
  - [type HomeScreen](#type-homescreen)
    - [func (h \*HomeScreen) New(id string) error](#func-homescreen-new)
  - [type Namer](#type-namer)
    - [func (n \*Namer) Episode() int](#func-namer-episode)
    - [func (n \*Namer) Season() int](#func-namer-season)
    - [func (n \*Namer) Show() string](#func-namer-show)
    - [func (n \*Namer) Title() string](#func-namer-title)
    - [func (n \*Namer) Year() int](#func-namer-year)
  - [type Playback](#type-playback)
    - [func (\*Playback) RequestHeader() (http.Header, error)](#func-playback-requestheader)
    - [func (p \*Playback) RequestUrl() (string, bool)](#func-playback-requesturl)
    - [func (\*Playback) UnwrapResponse(b []byte) ([]byte, error)](#func-playback-unwrapresponse)
    - [func (\*Playback) WrapRequest(b []byte) ([]byte, error)](#func-playback-wraprequest)
- [Source files](#source-files)

## Types

### type [AccountAuth](./auth.go#L70)

```go
type AccountAuth struct {
  AuthToken string
}
```

### func (\*AccountAuth) [Marshal](./auth.go#L13)

```go
func (*AccountAuth) Marshal(token *AccountToken) ([]byte, error)
```

token can be nil

### func (\*AccountAuth) [Playback](./auth.go#L31)

```go
func (a *AccountAuth) Playback(roku_id string) (*Playback, error)
```

### func (\*AccountAuth) [Unmarshal](./auth.go#L74)

```go
func (a *AccountAuth) Unmarshal(data []byte) error
```

### type [AccountCode](./code.go#L47)

```go
type AccountCode struct {
  Code string
}
```

### func (\*AccountCode) [Marshal](./code.go#L11)

```go
func (*AccountCode) Marshal(auth *AccountAuth) ([]byte, error)
```

### func (\*AccountCode) [String](./code.go#L36)

```go
func (a *AccountCode) String() string
```

### func (\*AccountCode) [Unmarshal](./code.go#L51)

```go
func (a *AccountCode) Unmarshal(data []byte) error
```

### type [AccountToken](./token.go#L106)

```go
type AccountToken struct {
  Token string
}
```

### func (AccountToken) [Marshal](./token.go#L13)

```go
func (AccountToken) Marshal(
  auth *AccountAuth, code *AccountCode,
) ([]byte, error)
```

### func (\*AccountToken) [Unmarshal](./token.go#L110)

```go
func (a *AccountToken) Unmarshal(data []byte) error
```

### type [HomeScreen](./token.go#L35)

```go
type HomeScreen struct {
  EpisodeNumber int       `json:",string"`
  ReleaseDate   time.Time // 2007-01-01T000000Z
  SeasonNumber  int       `json:",string"`
  Series        *struct {
    Title string
  }
  Title string
}
```

### func (\*HomeScreen) [New](./token.go#L45)

```go
func (h *HomeScreen) New(id string) error
```

### type [Namer](./token.go#L89)

```go
type Namer struct {
  Home HomeScreen
}
```

### func (\*Namer) [Episode](./token.go#L122)

```go
func (n *Namer) Episode() int
```

### func (\*Namer) [Season](./token.go#L118)

```go
func (n *Namer) Season() int
```

### func (\*Namer) [Show](./token.go#L130)

```go
func (n *Namer) Show() string
```

### func (\*Namer) [Title](./token.go#L126)

```go
func (n *Namer) Title() string
```

### func (\*Namer) [Year](./token.go#L114)

```go
func (n *Namer) Year() int
```

### type [Playback](./token.go#L93)

```go
type Playback struct {
  Drm struct {
    Widevine struct {
      LicenseServer string
    }
  }
  Url string
}
```

### func (\*Playback) [RequestHeader](./token.go#L137)

```go
func (*Playback) RequestHeader() (http.Header, error)
```

### func (\*Playback) [RequestUrl](./token.go#L102)

```go
func (p *Playback) RequestUrl() (string, bool)
```

### func (\*Playback) [UnwrapResponse](./token.go#L145)

```go
func (*Playback) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*Playback) [WrapRequest](./token.go#L141)

```go
func (*Playback) WrapRequest(b []byte) ([]byte, error)
```

## Source files

[auth.go](./auth.go)
[code.go](./code.go)
[token.go](./token.go)
