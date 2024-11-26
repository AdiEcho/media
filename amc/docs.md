# Overview

package `amc`

## Index

- [Types](#types)
  - [type Address](#type-address)
    - [func (a \*Address) Set(s string) error](#func-address-set)
    - [func (a \*Address) String() string](#func-address-string)
  - [type Authorization](#type-authorization)
    - [func (a \*Authorization) Content(path string) (\*ContentCompiler, error)](#func-authorization-content)
    - [func (a \*Authorization) Login(email, password string) ([]byte, error)](#func-authorization-login)
    - [func (a \*Authorization) Playback(nid string) (\*Playback, error)](#func-authorization-playback)
    - [func (a \*Authorization) Refresh() ([]byte, error)](#func-authorization-refresh)
    - [func (a \*Authorization) Unauth() error](#func-authorization-unauth)
    - [func (a \*Authorization) Unmarshal(data []byte) error](#func-authorization-unmarshal)
  - [type ContentCompiler](#type-contentcompiler)
    - [func (c \*ContentCompiler) Video() (\*CurrentVideo, bool)](#func-contentcompiler-video)
  - [type CurrentVideo](#type-currentvideo)
    - [func (c \*CurrentVideo) Episode() int](#func-currentvideo-episode)
    - [func (c \*CurrentVideo) Season() int](#func-currentvideo-season)
    - [func (c \*CurrentVideo) Show() string](#func-currentvideo-show)
    - [func (c \*CurrentVideo) Title() string](#func-currentvideo-title)
    - [func (c \*CurrentVideo) Year() int](#func-currentvideo-year)
  - [type DataSource](#type-datasource)
  - [type Playback](#type-playback)
    - [func (p \*Playback) Dash() (\*DataSource, bool)](#func-playback-dash)
    - [func (p \*Playback) RequestHeader() (http.Header, error)](#func-playback-requestheader)
    - [func (p \*Playback) RequestUrl() (string, bool)](#func-playback-requesturl)
    - [func (\*Playback) UnwrapResponse(b []byte) ([]byte, error)](#func-playback-unwrapresponse)
    - [func (\*Playback) WrapRequest(b []byte) ([]byte, error)](#func-playback-wraprequest)
- [Source files](#source-files)

## Types

### type [Address](./amc.go#L27)

```go
type Address struct {
  Nid  string
  Path string
}
```

### func (\*Address) [Set](./amc.go#L15)

```go
func (a *Address) Set(s string) error
```

### func (\*Address) [String](./amc.go#L32)

```go
func (a *Address) String() string
```

### type [Authorization](./auth.go#L137)

```go
type Authorization struct {
  Data struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
  }
}
```

### func (\*Authorization) [Content](./auth.go#L12)

```go
func (a *Authorization) Content(path string) (*ContentCompiler, error)
```

### func (\*Authorization) [Login](./auth.go#L148)

```go
func (a *Authorization) Login(email, password string) ([]byte, error)
```

### func (\*Authorization) [Playback](./auth.go#L59)

```go
func (a *Authorization) Playback(nid string) (*Playback, error)
```

### func (\*Authorization) [Refresh](./auth.go#L187)

```go
func (a *Authorization) Refresh() ([]byte, error)
```

### func (\*Authorization) [Unauth](./auth.go#L113)

```go
func (a *Authorization) Unauth() error
```

### func (\*Authorization) [Unmarshal](./auth.go#L144)

```go
func (a *Authorization) Unmarshal(data []byte) error
```

### type [ContentCompiler](./amc.go#L36)

```go
type ContentCompiler struct {
  Children []struct {
    Properties struct {
      CurrentVideo CurrentVideo
    }
    Type string
  }
}
```

### func (\*ContentCompiler) [Video](./amc.go#L45)

```go
func (c *ContentCompiler) Video() (*CurrentVideo, bool)
```

### type [CurrentVideo](./amc.go#L54)

```go
type CurrentVideo struct {
  Meta struct {
    Airdate       time.Time // 1996-01-01T00:00:00.000Z
    EpisodeNumber int
    Season        int `json:",string"`
    ShowTitle     string
  }
  Text struct {
    Title string
  }
}
```

### func (\*CurrentVideo) [Episode](./amc.go#L74)

```go
func (c *CurrentVideo) Episode() int
```

### func (\*CurrentVideo) [Season](./amc.go#L82)

```go
func (c *CurrentVideo) Season() int
```

### func (\*CurrentVideo) [Show](./amc.go#L78)

```go
func (c *CurrentVideo) Show() string
```

### func (\*CurrentVideo) [Title](./amc.go#L66)

```go
func (c *CurrentVideo) Title() string
```

### func (\*CurrentVideo) [Year](./amc.go#L70)

```go
func (c *CurrentVideo) Year() int
```

### type [DataSource](./amc.go#L86)

```go
type DataSource struct {
  KeySystems *struct {
    Widevine struct {
      LicenseUrl string `json:"license_url"`
    } `json:"com.widevine.alpha"`
  } `json:"key_systems"`
  Src  string
  Type string
}
```

### type [Playback](./amc.go#L118)

```go
type Playback struct {
  AmcnBcJwt string `json:"-"`
  Data      struct {
    PlaybackJsonData struct {
      Sources []DataSource
    }
  }
}
```

### func (\*Playback) [Dash](./amc.go#L96)

```go
func (p *Playback) Dash() (*DataSource, bool)
```

### func (\*Playback) [RequestHeader](./amc.go#L112)

```go
func (p *Playback) RequestHeader() (http.Header, error)
```

### func (\*Playback) [RequestUrl](./amc.go#L105)

```go
func (p *Playback) RequestUrl() (string, bool)
```

### func (\*Playback) [UnwrapResponse](./amc.go#L131)

```go
func (*Playback) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*Playback) [WrapRequest](./amc.go#L127)

```go
func (*Playback) WrapRequest(b []byte) ([]byte, error)
```

## Source files

[amc.go](./amc.go)
[auth.go](./auth.go)
