# Overview

package `max`

## Index

- [Types](#types)
  - [type Address](#type-address)
    - [func (a \*Address) MarshalText() ([]byte, error)](#func-address-marshaltext)
    - [func (a \*Address) UnmarshalText(text []byte) error](#func-address-unmarshaltext)
  - [type BoltToken](#type-bolttoken)
    - [func (b \*BoltToken) Initiate() (\*LinkInitiate, error)](#func-bolttoken-initiate)
    - [func (b \*BoltToken) New() error](#func-bolttoken-new)
  - [type DefaultRoutes](#type-defaultroutes)
    - [func (d \*DefaultRoutes) Episode() int](#func-defaultroutes-episode)
    - [func (d \*DefaultRoutes) Season() int](#func-defaultroutes-season)
    - [func (d \*DefaultRoutes) Show() string](#func-defaultroutes-show)
    - [func (d \*DefaultRoutes) Title() string](#func-defaultroutes-title)
    - [func (d \*DefaultRoutes) Year() int](#func-defaultroutes-year)
  - [type LinkInitiate](#type-linkinitiate)
  - [type LinkLogin](#type-linklogin)
    - [func (LinkLogin) Marshal(token \*BoltToken) ([]byte, error)](#func-linklogin-marshal)
    - [func (v \*LinkLogin) Playback(web \*Address) (\*Playback, error)](#func-linklogin-playback)
    - [func (v \*LinkLogin) Routes(web \*Address) (\*DefaultRoutes, error)](#func-linklogin-routes)
    - [func (v \*LinkLogin) Unmarshal(data []byte) error](#func-linklogin-unmarshal)
  - [type Manifest](#type-manifest)
    - [func (m \*Manifest) UnmarshalText(text []byte) error](#func-manifest-unmarshaltext)
  - [type Playback](#type-playback)
    - [func (\*Playback) RequestHeader() (http.Header, error)](#func-playback-requestheader)
    - [func (p \*Playback) RequestUrl() (string, bool)](#func-playback-requesturl)
    - [func (\*Playback) UnwrapResponse(b []byte) ([]byte, error)](#func-playback-unwrapresponse)
    - [func (\*Playback) WrapRequest(b []byte) ([]byte, error)](#func-playback-wraprequest)
  - [type RouteInclude](#type-routeinclude)
- [Source files](#source-files)

## Types

### type [Address](./login.go#L236)

```go
type Address struct {
  EditId  string
  VideoId string
}
```

### func (\*Address) [MarshalText](./login.go#L223)

```go
func (a *Address) MarshalText() ([]byte, error)
```

### func (\*Address) [UnmarshalText](./login.go#L241)

```go
func (a *Address) UnmarshalText(text []byte) error
```

### type [BoltToken](./token.go#L10)

```go
type BoltToken struct {
  St string
}
```

### func (\*BoltToken) [Initiate](./initiate.go#L8)

```go
func (b *BoltToken) Initiate() (*LinkInitiate, error)
```

### func (\*BoltToken) [New](./token.go#L14)

```go
func (b *BoltToken) New() error
```

### type [DefaultRoutes](./login.go#L214)

```go
type DefaultRoutes struct {
  Data struct {
    Attributes struct {
      Url Address
    }
  }
  Included []RouteInclude
}
```

### func (\*DefaultRoutes) [Episode](./login.go#L180)

```go
func (d *DefaultRoutes) Episode() int
```

### func (\*DefaultRoutes) [Season](./login.go#L173)

```go
func (d *DefaultRoutes) Season() int
```

### func (\*DefaultRoutes) [Show](./login.go#L201)

```go
func (d *DefaultRoutes) Show() string
```

### func (\*DefaultRoutes) [Title](./login.go#L187)

```go
func (d *DefaultRoutes) Title() string
```

### func (\*DefaultRoutes) [Year](./login.go#L194)

```go
func (d *DefaultRoutes) Year() int
```

### type [LinkInitiate](./initiate.go#L32)

```go
type LinkInitiate struct {
  Data struct {
    Attributes struct {
      LinkingCode string
      TargetUrl   string
    }
  }
}
```

### type [LinkLogin](./login.go#L293)

```go
type LinkLogin struct {
  Data struct {
    Attributes struct {
      Token string
    }
  }
}
```

### func (LinkLogin) [Marshal](./login.go#L17)

```go
func (LinkLogin) Marshal(token *BoltToken) ([]byte, error)
```

you must
/authentication/linkDevice/initiate
first or this will always fail

### func (\*LinkLogin) [Playback](./login.go#L33)

```go
func (v *LinkLogin) Playback(web *Address) (*Playback, error)
```

### func (\*LinkLogin) [Routes](./login.go#L257)

```go
func (v *LinkLogin) Routes(web *Address) (*DefaultRoutes, error)
```

### func (\*LinkLogin) [Unmarshal](./login.go#L301)

```go
func (v *LinkLogin) Unmarshal(data []byte) error
```

### type [Manifest](./login.go#L78)

```go
type Manifest struct {
  Url string
}
```

### func (\*Manifest) [UnmarshalText](./login.go#L73)

```go
func (m *Manifest) UnmarshalText(text []byte) error
```

### type [Playback](./login.go#L82)

```go
type Playback struct {
  Drm struct {
    Schemes struct {
      Widevine struct {
        LicenseUrl string
      }
    }
  }
  Fallback struct {
    Manifest struct {
      Url Manifest
    }
  }
}
```

### func (\*Playback) [RequestHeader](./login.go#L105)

```go
func (*Playback) RequestHeader() (http.Header, error)
```

### func (\*Playback) [RequestUrl](./login.go#L109)

```go
func (p *Playback) RequestUrl() (string, bool)
```

### func (\*Playback) [UnwrapResponse](./login.go#L101)

```go
func (*Playback) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*Playback) [WrapRequest](./login.go#L97)

```go
func (*Playback) WrapRequest(b []byte) ([]byte, error)
```

### type [RouteInclude](./login.go#L113)

```go
type RouteInclude struct {
  Attributes struct {
    AirDate       time.Time
    Name          string
    EpisodeNumber int
    SeasonNumber  int
  }
  Id            string
  Relationships *struct {
    Show *struct {
      Data struct {
        Id string
      }
    }
  }
}
```

## Source files

[initiate.go](./initiate.go)
[login.go](./login.go)
[token.go](./token.go)
