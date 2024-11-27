# Overview

package `rtbf`

## Index

- [Types](#types)
  - [type Address](#type-address)
    - [func (a Address) Page() (\*AuvioPage, error)](#func-address-page)
    - [func (a \*Address) Set(s string) error](#func-address-set)
    - [func (a \*Address) String() string](#func-address-string)
  - [type AuvioAuth](#type-auvioauth)
    - [func (a \*AuvioAuth) Entitlement(asset_id string) (\*Entitlement, error)](#func-auvioauth-entitlement)
  - [type AuvioLogin](#type-auviologin)
    - [func (AuvioLogin) Marshal(id, password string) ([]byte, error)](#func-auviologin-marshal)
    - [func (a \*AuvioLogin) Token() (\*WebToken, error)](#func-auviologin-token)
    - [func (a \*AuvioLogin) Unmarshal(data []byte) error](#func-auviologin-unmarshal)
  - [type AuvioPage](#type-auviopage)
    - [func (a \*AuvioPage) GetAssetId() (string, bool)](#func-auviopage-getassetid)
  - [type Entitlement](#type-entitlement)
    - [func (e \*Entitlement) Dash() (string, bool)](#func-entitlement-dash)
    - [func (\*Entitlement) RequestHeader() (http.Header, error)](#func-entitlement-requestheader)
    - [func (e \*Entitlement) RequestUrl() (string, bool)](#func-entitlement-requesturl)
    - [func (\*Entitlement) UnwrapResponse(b []byte) ([]byte, error)](#func-entitlement-unwrapresponse)
    - [func (\*Entitlement) WrapRequest(b []byte) ([]byte, error)](#func-entitlement-wraprequest)
  - [type Namer](#type-namer)
    - [func (n Namer) Episode() int](#func-namer-episode)
    - [func (n Namer) Season() int](#func-namer-season)
    - [func (n Namer) Show() string](#func-namer-show)
    - [func (n Namer) Title() string](#func-namer-title)
    - [func (Namer) Year() int](#func-namer-year)
  - [type Subtitle](#type-subtitle)
    - [func (s \*Subtitle) UnmarshalText(text []byte) error](#func-subtitle-unmarshaltext)
  - [type Title](#type-title)
    - [func (t \*Title) UnmarshalText(text []byte) error](#func-title-unmarshaltext)
  - [type WebToken](#type-webtoken)
    - [func (w \*WebToken) Auth() (\*AuvioAuth, error)](#func-webtoken-auth)
- [Source files](#source-files)

## Types

### type [Address](./rtbf.go#L199)

```go
type Address struct {
  Path string
}
```

### func (Address) [Page](./rtbf.go#L29)

```go
func (a Address) Page() (*AuvioPage, error)
```

### func (\*Address) [Set](./rtbf.go#L52)

```go
func (a *Address) Set(s string) error
```

### func (\*Address) [String](./rtbf.go#L313)

```go
func (a *Address) String() string
```

### type [AuvioAuth](./rtbf.go#L58)

```go
type AuvioAuth struct {
  SessionToken string
}
```

### func (\*AuvioAuth) [Entitlement](./rtbf.go#L225)

```go
func (a *AuvioAuth) Entitlement(asset_id string) (*Entitlement, error)
```

### type [AuvioLogin](./rtbf.go#L281)

```go
type AuvioLogin struct {
  ErrorMessage string
  SessionInfo  struct {
    CookieValue string
  }
}
```

### func (AuvioLogin) [Marshal](./rtbf.go#L14)

```go
func (AuvioLogin) Marshal(id, password string) ([]byte, error)
```

### func (\*AuvioLogin) [Token](./rtbf.go#L259)

```go
func (a *AuvioLogin) Token() (*WebToken, error)
```

### func (\*AuvioLogin) [Unmarshal](./rtbf.go#L288)

```go
func (a *AuvioLogin) Unmarshal(data []byte) error
```

### type [AuvioPage](./rtbf.go#L206)

```go
type AuvioPage struct {
  AssetId string
  Media   *struct {
    AssetId string
  }
  Subtitle Subtitle
  Title    Title
}
```

### func (\*AuvioPage) [GetAssetId](./rtbf.go#L215)

```go
func (a *AuvioPage) GetAssetId() (string, bool)
```

### type [Entitlement](./rtbf.go#L83)

```go
type Entitlement struct {
  AssetId   string
  PlayToken string
  Formats   []struct {
    Format       string
    MediaLocator string
  }
}
```

### func (\*Entitlement) [Dash](./rtbf.go#L62)

```go
func (e *Entitlement) Dash() (string, bool)
```

### func (\*Entitlement) [RequestHeader](./rtbf.go#L307)

```go
func (*Entitlement) RequestHeader() (http.Header, error)
```

### func (\*Entitlement) [RequestUrl](./rtbf.go#L71)

```go
func (e *Entitlement) RequestUrl() (string, bool)
```

### func (\*Entitlement) [UnwrapResponse](./rtbf.go#L303)

```go
func (*Entitlement) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*Entitlement) [WrapRequest](./rtbf.go#L299)

```go
func (*Entitlement) WrapRequest(b []byte) ([]byte, error)
```

### type [Namer](./rtbf.go#L92)

```go
type Namer struct {
  Page *AuvioPage
}
```

### func (Namer) [Episode](./rtbf.go#L101)

```go
func (n Namer) Episode() int
```

### func (Namer) [Season](./rtbf.go#L105)

```go
func (n Namer) Season() int
```

### func (Namer) [Show](./rtbf.go#L109)

```go
func (n Namer) Show() string
```

### func (Namer) [Title](./rtbf.go#L116)

```go
func (n Namer) Title() string
```

### func (Namer) [Year](./rtbf.go#L97)

```go
func (Namer) Year() int
```

its just not available from what I can tell

### type [Subtitle](./rtbf.go#L123)

```go
type Subtitle struct {
  Episode  int
  Subtitle string
}
```

### func (\*Subtitle) [UnmarshalText](./rtbf.go#L130)

```go
func (s *Subtitle) UnmarshalText(text []byte) error
```

json.data.content.subtitle = "06 - Les ombres de la guerre";
json.data.content.subtitle = "Avec Rosamund Pike";

### type [Title](./rtbf.go#L141)

```go
type Title struct {
  Season int
  Title  string
}
```

### func (\*Title) [UnmarshalText](./rtbf.go#L148)

```go
func (t *Title) UnmarshalText(text []byte) error
```

json.data.content.title = "Grantchester S01";
json.data.content.title = "I care a lot";

### type [WebToken](./rtbf.go#L159)

```go
type WebToken struct {
  ErrorMessage string
  IdToken      string `json:"id_token"`
}
```

### func (\*WebToken) [Auth](./rtbf.go#L164)

```go
func (w *WebToken) Auth() (*AuvioAuth, error)
```

## Source files

[rtbf.go](./rtbf.go)
