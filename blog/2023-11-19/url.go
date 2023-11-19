package url

import "net/url"

func path(raw_URL string) (string, error) {
   u, err := url.Parse(raw_URL)
   if err != nil {
      return "", err
   }
   return u.Path, nil
}
