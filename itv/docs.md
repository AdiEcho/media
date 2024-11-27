# Overview

package `itv`

## Index

- [Types](#types)
  - [type DiscoveryTitle](#type-discoverytitle)
    - [func (d \*DiscoveryTitle) Playlist() (\*Playlist, error)](#func-discoverytitle-playlist)
  - [type LegacyId](#type-legacyid)
    - [func (i LegacyId) Discovery() (\*DiscoveryTitle, error)](#func-legacyid-discovery)
    - [func (i \*LegacyId) Set(text string) error](#func-legacyid-set)
    - [func (i LegacyId) String() string](#func-legacyid-string)
  - [type Namer](#type-namer)
    - [func (n Namer) Episode() int](#func-namer-episode)
    - [func (n Namer) Season() int](#func-namer-season)
    - [func (n Namer) Show() string](#func-namer-show)
    - [func (n Namer) Title() string](#func-namer-title)
    - [func (n Namer) Year() int](#func-namer-year)
  - [type Playlist](#type-playlist)
    - [func (p \*Playlist) Resolution720() (string, bool)](#func-playlist-resolution720)
  - [type Poster](#type-poster)
    - [func (Poster) RequestHeader() (http.Header, error)](#func-poster-requestheader)
    - [func (Poster) RequestUrl() (string, bool)](#func-poster-requesturl)
    - [func (Poster) UnwrapResponse(b []byte) ([]byte, error)](#func-poster-unwrapresponse)
    - [func (Poster) WrapRequest(b []byte) ([]byte, error)](#func-poster-wraprequest)
- [Source files](#source-files)

## Types

### type [DiscoveryTitle](./itv.go#L134)

```go
type DiscoveryTitle struct {
  LatestAvailableVersion struct {
    PlaylistUrl string
  }
  Brand *struct {
    Title string
  }
  EpisodeNumber  int
  ProductionYear int
  SeriesNumber   int
  Title          string
}
```

### func (\*DiscoveryTitle) [Playlist](./itv.go#L83)

```go
func (d *DiscoveryTitle) Playlist() (*Playlist, error)
```

hard geo block

### type [LegacyId](./itv.go#L80)

```go
type LegacyId [3]string
```

### func (LegacyId) [Discovery](./itv.go#L205)

```go
func (i LegacyId) Discovery() (*DiscoveryTitle, error)
```

### func (\*LegacyId) [Set](./itv.go#L67)

```go
func (i *LegacyId) Set(text string) error
```

### func (LegacyId) [String](./itv.go#L54)

```go
func (i LegacyId) String() string
```

### type [Namer](./itv.go#L147)

```go
type Namer struct {
  Discovery *DiscoveryTitle
}
```

### func (Namer) [Episode](./itv.go#L162)

```go
func (n Namer) Episode() int
```

### func (Namer) [Season](./itv.go#L158)

```go
func (n Namer) Season() int
```

### func (Namer) [Show](./itv.go#L151)

```go
func (n Namer) Show() string
```

### func (Namer) [Title](./itv.go#L166)

```go
func (n Namer) Title() string
```

### func (Namer) [Year](./itv.go#L170)

```go
func (n Namer) Year() int
```

### type [Playlist](./itv.go#L174)

```go
type Playlist struct {
  Playlist struct {
    Video struct {
      MediaFiles []struct {
        Href       string
        Resolution string
      }
    }
  }
}
```

### func (\*Playlist) [Resolution720](./itv.go#L45)

```go
func (p *Playlist) Resolution720() (string, bool)
```

### type [Poster](./itv.go#L185)

```go
type Poster struct{}
```

### func (Poster) [RequestHeader](./itv.go#L195)

```go
func (Poster) RequestHeader() (http.Header, error)
```

### func (Poster) [RequestUrl](./itv.go#L36)

```go
func (Poster) RequestUrl() (string, bool)
```

### func (Poster) [UnwrapResponse](./itv.go#L191)

```go
func (Poster) UnwrapResponse(b []byte) ([]byte, error)
```

### func (Poster) [WrapRequest](./itv.go#L187)

```go
func (Poster) WrapRequest(b []byte) ([]byte, error)
```

## Source files

[itv.go](./itv.go)
