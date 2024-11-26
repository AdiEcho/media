# Overview

package `member`

## Index

- [Types](#types)
  - [type Address](#type-address)
    - [func (a *Address) Set(s string) error](#func-address-set)
    - [func (a *Address) String() string](#func-address-string)
  - [type ArticleAsset](#type-articleasset)
  - [type Entitlement](#type-entitlement)
    - [func (*Entitlement) RequestHeader() (http.Header, error)](#func-entitlement-requestheader)
    - [func (e *Entitlement) RequestUrl() (string, bool)](#func-entitlement-requesturl)
    - [func (*Entitlement) UnwrapResponse(b []byte) ([]byte, error)](#func-entitlement-unwrapresponse)
    - [func (*Entitlement) WrapRequest(b []byte) ([]byte, error)](#func-entitlement-wraprequest)
  - [type OperationArticle](#type-operationarticle)
    - [func (*OperationArticle) Episode() int](#func-operationarticle-episode)
    - [func (o *OperationArticle) Film() (*ArticleAsset, bool)](#func-operationarticle-film)
    - [func (*OperationArticle) Marshal(web *Address) ([]byte, error)](#func-operationarticle-marshal)
    - [func (*OperationArticle) Season() int](#func-operationarticle-season)
    - [func (*OperationArticle) Show() string](#func-operationarticle-show)
    - [func (o *OperationArticle) Title() string](#func-operationarticle-title)
    - [func (o *OperationArticle) Unmarshal(data []byte) error](#func-operationarticle-unmarshal)
    - [func (o *OperationArticle) Year() int](#func-operationarticle-year)
  - [type OperationPlay](#type-operationplay)
    - [func (o *OperationPlay) Dash() (*Entitlement, bool)](#func-operationplay-dash)
    - [func (OperationPlay) Marshal(
  user *OperationUser, asset *ArticleAsset,
) ([]byte, error)](#func-operationplay-marshal)
    - [func (o *OperationPlay) Unmarshal(data []byte) error](#func-operationplay-unmarshal)
  - [type OperationUser](#type-operationuser)
    - [func (OperationUser) Marshal(email, password string) ([]byte, error)](#func-operationuser-marshal)
    - [func (o *OperationUser) Unmarshal(data []byte) error](#func-operationuser-unmarshal)
- [Source files](#source-files)

## Types

### type [Address](./article.go#L45)

```go
type Address struct {
  Path string
}
```

### func (*Address) [Set](./article.go#L36)

```go
func (a *Address) Set(s string) error
```

### func (*Address) [String](./article.go#L49)

```go
func (a *Address) String() string
```

### type [ArticleAsset](./article.go#L53)

```go
type ArticleAsset struct {
  Id         int
  LinkedType string `json:"linked_type"`
  // contains filtered or unexported fields
}
```

### type [Entitlement](./play.go#L95)

```go
type Entitlement struct {
  KeyDeliveryUrl string `json:"key_delivery_url"`
  Manifest       string
  Protocol       string
}
```

### func (*Entitlement) [RequestHeader](./play.go#L83)

```go
func (*Entitlement) RequestHeader() (http.Header, error)
```

### func (*Entitlement) [RequestUrl](./play.go#L79)

```go
func (e *Entitlement) RequestUrl() (string, bool)
```

### func (*Entitlement) [UnwrapResponse](./play.go#L91)

```go
func (*Entitlement) UnwrapResponse(b []byte) ([]byte, error)
```

### func (*Entitlement) [WrapRequest](./play.go#L87)

```go
func (*Entitlement) WrapRequest(b []byte) ([]byte, error)
```

### type [OperationArticle](./article.go#L119)

```go
type OperationArticle struct {
  Assets         []*ArticleAsset
  CanonicalTitle string `json:"canonical_title"`
  Id             int
  Metas          []struct {
    Key   string
    Value string
  }
}
```

### func (*OperationArticle) [Episode](./article.go#L107)

```go
func (*OperationArticle) Episode() int
```

### func (*OperationArticle) [Film](./article.go#L83)

```go
func (o *OperationArticle) Film() (*ArticleAsset, bool)
```

### func (*OperationArticle) [Marshal](./article.go#L59)

```go
func (*OperationArticle) Marshal(web *Address) ([]byte, error)
```

### func (*OperationArticle) [Season](./article.go#L111)

```go
func (*OperationArticle) Season() int
```

### func (*OperationArticle) [Show](./article.go#L115)

```go
func (*OperationArticle) Show() string
```

### func (*OperationArticle) [Title](./article.go#L92)

```go
func (o *OperationArticle) Title() string
```

### func (*OperationArticle) [Unmarshal](./article.go#L129)

```go
func (o *OperationArticle) Unmarshal(data []byte) error
```

### func (*OperationArticle) [Year](./article.go#L96)

```go
func (o *OperationArticle) Year() int
```

### type [OperationPlay](./play.go#L11)

```go
type OperationPlay struct {
  Data struct {
    ArticleAssetPlay struct {
      Entitlements []Entitlement
    }
  }
  Errors []struct {
    Message string
  }
}
```

### func (*OperationPlay) [Dash](./play.go#L70)

```go
func (o *OperationPlay) Dash() (*Entitlement, bool)
```

### func (OperationPlay) [Marshal](./play.go#L23)

```go
func (OperationPlay) Marshal(
  user *OperationUser, asset *ArticleAsset,
) ([]byte, error)
```

hard geo block

### func (*OperationPlay) [Unmarshal](./play.go#L59)

```go
func (o *OperationPlay) Unmarshal(data []byte) error
```

### type [OperationUser](./user.go#L44)

```go
type OperationUser struct {
  Data struct {
    UserAuthenticate struct {
      AccessToken string `json:"access_token"`
    }
  }
}
```

### func (OperationUser) [Marshal](./user.go#L18)

```go
func (OperationUser) Marshal(email, password string) ([]byte, error)
```

### func (*OperationUser) [Unmarshal](./user.go#L52)

```go
func (o *OperationUser) Unmarshal(data []byte) error
```

## Source files

[article.go](./article.go)
[play.go](./play.go)
[user.go](./user.go)
