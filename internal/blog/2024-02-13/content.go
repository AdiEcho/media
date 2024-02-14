package roku

import (
   "net/url"
   "strings"
)

func content(id string) (*url.URL, error) {
   a, err := url.Parse("https://therokuchannel.roku.com/api/v2/homescreen/content")
   if err != nil {
      return nil, err
   }
   include := []string{
      "episodeNumber",
      "releaseDate",
      "seasonNumber",
      "series.title",
      "title",
      "viewOptions",
   }
   expand := url.URL{
      Scheme: "https",
      Host: "content.sr.roku.com",
      Path: "/content/v1/roku-trc/" + id,
      RawQuery: url.Values{
         "expand": {"series"},
         "include": {strings.Join(include, ",")},
      }.Encode(),
   }
   homescreen := url.PathEscape(expand.String())
   return a.JoinPath(homescreen), nil
}
