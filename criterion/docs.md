# Overview

package `criterion`

## Index

- [Types](#types)
  - [type AuthToken](#type-authtoken)
    - [func (a \*AuthToken) Files(item \*EmbedItem) (VideoFiles, error)](#func-authtoken-files)
    - [func (AuthToken) Marshal(username, password string) ([]byte, error)](#func-authtoken-marshal)
    - [func (a \*AuthToken) Unmarshal(data []byte) error](#func-authtoken-unmarshal)
    - [func (a \*AuthToken) Video(slug string) (\*EmbedItem, error)](#func-authtoken-video)
  - [type EmbedItem](#type-embeditem)
    - [func (\*EmbedItem) Episode() int](#func-embeditem-episode)
    - [func (\*EmbedItem) Season() int](#func-embeditem-season)
    - [func (\*EmbedItem) Show() string](#func-embeditem-show)
    - [func (e \*EmbedItem) Title() string](#func-embeditem-title)
    - [func (e \*EmbedItem) Year() int](#func-embeditem-year)
  - [type VideoFile](#type-videofile)
    - [func (\*VideoFile) RequestHeader() (http.Header, error)](#func-videofile-requestheader)
    - [func (v \*VideoFile) RequestUrl() (string, bool)](#func-videofile-requesturl)
    - [func (\*VideoFile) UnwrapResponse(b []byte) ([]byte, error)](#func-videofile-unwrapresponse)
    - [func (\*VideoFile) WrapRequest(b []byte) ([]byte, error)](#func-videofile-wraprequest)
  - [type VideoFiles](#type-videofiles)
    - [func (v \*VideoFiles) Dash() (\*VideoFile, bool)](#func-videofiles-dash)
- [Source files](#source-files)

## Types

### type [AuthToken](./criterion.go#L38)

```go
type AuthToken struct {
  AccessToken string `json:"access_token"`
}
```

### func (\*AuthToken) [Files](./criterion.go#L14)

```go
func (a *AuthToken) Files(item *EmbedItem) (VideoFiles, error)
```

### func (AuthToken) [Marshal](./criterion.go#L46)

```go
func (AuthToken) Marshal(username, password string) ([]byte, error)
```

### func (\*AuthToken) [Unmarshal](./criterion.go#L42)

```go
func (a *AuthToken) Unmarshal(data []byte) error
```

### func (\*AuthToken) [Video](./criterion.go#L60)

```go
func (a *AuthToken) Video(slug string) (*EmbedItem, error)
```

### type [EmbedItem](./criterion.go#L108)

```go
type EmbedItem struct {
  Links struct {
    Files struct {
      Href string
    }
  } `json:"_links"`
  Metadata struct {
    YearReleased int `json:"year_released"`
  }
  Name string
}
```

### func (\*EmbedItem) [Episode](./criterion.go#L88)

```go
func (*EmbedItem) Episode() int
```

### func (\*EmbedItem) [Season](./criterion.go#L92)

```go
func (*EmbedItem) Season() int
```

### func (\*EmbedItem) [Show](./criterion.go#L96)

```go
func (*EmbedItem) Show() string
```

### func (\*EmbedItem) [Title](./criterion.go#L100)

```go
func (e *EmbedItem) Title() string
```

### func (\*EmbedItem) [Year](./criterion.go#L104)

```go
func (e *EmbedItem) Year() int
```

### type [VideoFile](./criterion.go#L126)

```go
type VideoFile struct {
  DrmAuthorizationToken string `json:"drm_authorization_token"`
  Links                 struct {
    Source struct {
      Href string
    }
  } `json:"_links"`
  Method string
}
```

### func (\*VideoFile) [RequestHeader](./criterion.go#L136)

```go
func (*VideoFile) RequestHeader() (http.Header, error)
```

### func (\*VideoFile) [RequestUrl](./criterion.go#L120)

```go
func (v *VideoFile) RequestUrl() (string, bool)
```

### func (\*VideoFile) [UnwrapResponse](./criterion.go#L140)

```go
func (*VideoFile) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*VideoFile) [WrapRequest](./criterion.go#L144)

```go
func (*VideoFile) WrapRequest(b []byte) ([]byte, error)
```

### type [VideoFiles](./criterion.go#L148)

```go
type VideoFiles []VideoFile
```

### func (\*VideoFiles) [Dash](./criterion.go#L150)

```go
func (v *VideoFiles) Dash() (*VideoFile, bool)
```

## Source files

[criterion.go](./criterion.go)
