# encoding

## json.Unmarshal

header, body:

~~~go
type response struct {
   date value[time.Time]
   body value[struct {
      Slideshow struct {
         Date string
         Title string
      }
   }]
}
~~~

body:

~~~go
type response struct {
   body *struct {
      Slideshow struct {
         Date string
         Title string
      }
   }
   raw []byte
}
~~~

## json.Unmarshal, json.Marshal, json.Unmarshal

header, body:

~~~go
type response struct {
   date time.Time
   body struct {
      Slideshow struct {
         Date string
         Title string
      }
   }
}
~~~

body:

~~~go
type response struct {
   Slideshow struct {
      Date string
      Title string
   }
}
~~~
