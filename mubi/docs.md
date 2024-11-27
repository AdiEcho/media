# Overview

package `mubi`

## Index

- [Variables](#variables)
- [Types](#types)
  - [type Address](#type-address)
    - [func (a \*Address) Film() (\*FilmResponse, error)](#func-address-film)
    - [func (a \*Address) Set(text string) error](#func-address-set)
    - [func (a \*Address) String() string](#func-address-string)
  - [type Authenticate](#type-authenticate)
    - [func (Authenticate) Marshal(code \*LinkCode) ([]byte, error)](#func-authenticate-marshal)
    - [func (a \*Authenticate) RequestHeader() (http.Header, error)](#func-authenticate-requestheader)
    - [func (\*Authenticate) RequestUrl() (string, bool)](#func-authenticate-requesturl)
    - [func (a \*Authenticate) Unmarshal(data []byte) error](#func-authenticate-unmarshal)
    - [func (\*Authenticate) UnwrapResponse(b []byte) ([]byte, error)](#func-authenticate-unwrapresponse)
    - [func (a \*Authenticate) Viewing(film \*FilmResponse) error](#func-authenticate-viewing)
    - [func (\*Authenticate) WrapRequest(b []byte) ([]byte, error)](#func-authenticate-wraprequest)
  - [type FilmResponse](#type-filmresponse)
  - [type LinkCode](#type-linkcode)
    - [func (\*LinkCode) Marshal() ([]byte, error)](#func-linkcode-marshal)
    - [func (c \*LinkCode) String() string](#func-linkcode-string)
    - [func (c \*LinkCode) Unmarshal(data []byte) error](#func-linkcode-unmarshal)
  - [type Namer](#type-namer)
    - [func (Namer) Episode() int](#func-namer-episode)
    - [func (Namer) Season() int](#func-namer-season)
    - [func (Namer) Show() string](#func-namer-show)
    - [func (n Namer) Title() string](#func-namer-title)
    - [func (n Namer) Year() int](#func-namer-year)
  - [type SecureUrl](#type-secureurl)
    - [func (SecureUrl) Marshal(
  auth \*Authenticate, film \*FilmResponse,
) ([]byte, error)](#func-secureurl-marshal)
    - [func (s \*SecureUrl) Unmarshal(data []byte) error](#func-secureurl-unmarshal)
  - [type TextTrack](#type-texttrack)
    - [func (t \*TextTrack) String() string](#func-texttrack-string)
- [Source files](#source-files)

## Variables

```go
var ClientCountry = "US"
```

## Types

### type [Address](./code.go#L40)

```go
type Address struct {
  Text string
}
```

### func (\*Address) [Film](./code.go#L84)

```go
func (a *Address) Film() (*FilmResponse, error)
```

### func (\*Address) [Set](./code.go#L44)

```go
func (a *Address) Set(text string) error
```

### func (\*Address) [String](./code.go#L64)

```go
func (a *Address) String() string
```

### type [Authenticate](./auth.go#L111)

```go
type Authenticate struct {
  Token string
  User  struct {
    Id int
  }
}
```

### func (Authenticate) [Marshal](./auth.go#L14)

```go
func (Authenticate) Marshal(code *LinkCode) ([]byte, error)
```

### func (\*Authenticate) [RequestHeader](./auth.go#L43)

```go
func (a *Authenticate) RequestHeader() (http.Header, error)
```

### func (\*Authenticate) [RequestUrl](./auth.go#L96)

```go
func (*Authenticate) RequestUrl() (string, bool)
```

final slash is needed

### func (\*Authenticate) [Unmarshal](./auth.go#L118)

```go
func (a *Authenticate) Unmarshal(data []byte) error
```

### func (\*Authenticate) [UnwrapResponse](./auth.go#L100)

```go
func (*Authenticate) UnwrapResponse(b []byte) ([]byte, error)
```

### func (\*Authenticate) [Viewing](./auth.go#L62)

```go
func (a *Authenticate) Viewing(film *FilmResponse) error
```

Mubi do this sneaky thing. you cannot download a video unless you have told
the API that you are watching it. so you have to call
`/v3/films/%v/viewing`, otherwise it wont let you get the MPD. if you have
already viewed the video on the website that counts, but if you only use the
tool it will error

### func (\*Authenticate) [WrapRequest](./auth.go#L91)

```go
func (*Authenticate) WrapRequest(b []byte) ([]byte, error)
```

### type [FilmResponse](./code.go#L58)

```go
type FilmResponse struct {
  Id    int64
  Title string
  Year  int
}
```

### type [LinkCode](./code.go#L129)

```go
type LinkCode struct {
  AuthToken string `json:"auth_token"`
  LinkCode  string `json:"link_code"`
}
```

### func (\*LinkCode) [Marshal](./code.go#L11)

```go
func (*LinkCode) Marshal() ([]byte, error)
```

### func (\*LinkCode) [String](./code.go#L107)

```go
func (c *LinkCode) String() string
```

### func (\*LinkCode) [Unmarshal](./code.go#L134)

```go
func (c *LinkCode) Unmarshal(data []byte) error
```

### type [Namer](./code.go#L68)

```go
type Namer struct {
  Film *FilmResponse
}
```

### func (Namer) [Episode](./code.go#L117)

```go
func (Namer) Episode() int
```

### func (Namer) [Season](./code.go#L121)

```go
func (Namer) Season() int
```

### func (Namer) [Show](./code.go#L125)

```go
func (Namer) Show() string
```

### func (Namer) [Title](./code.go#L72)

```go
func (n Namer) Title() string
```

### func (Namer) [Year](./code.go#L76)

```go
func (n Namer) Year() int
```

### type [SecureUrl](./secure.go#L43)

```go
type SecureUrl struct {
  TextTrackUrls []TextTrack `json:"text_track_urls"`
  Url           string
}
```

### func (SecureUrl) [Marshal](./secure.go#L12)

```go
func (SecureUrl) Marshal(
  auth *Authenticate, film *FilmResponse,
) ([]byte, error)
```

### func (\*SecureUrl) [Unmarshal](./secure.go#L48)

```go
func (s *SecureUrl) Unmarshal(data []byte) error
```

### type [TextTrack](./code.go#L53)

```go
type TextTrack struct {
  Id  string
  Url string
}
```

### func (\*TextTrack) [String](./code.go#L80)

```go
func (t *TextTrack) String() string
```

## Source files

[auth.go](./auth.go)
[code.go](./code.go)
[secure.go](./secure.go)
