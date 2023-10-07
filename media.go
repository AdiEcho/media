package media

import (
   "encoding/json"
   "os"
   "strconv"
   "time"
)

type Namer interface {
   Date() (time.Time, error)
   Episode() (int64, error)
   Season() (int64, error)
   Series() (string, bool)
   Title() string
}

func Name(n Namer) (string, error) {
   var b []byte
   title := Clean(n.Title())
   if series, ok := n.Series(); ok {
      b = append(b, series...)
      b = append(b, " - S"...)
      {
         s, err := n.Season()
         if err != nil {
            return "", err
         }
         b = strconv.AppendInt(b, s, 10)
      }
      b = append(b, " E"...)
      {
         e, err := n.Episode()
         if err != nil {
            return "", err
         }
         b = strconv.AppendInt(b, e, 10)
      }
      b = append(b, " - "...)
      b = append(b, title...)
   } else {
      b = append(b, title...)
      b = append(b, " - "...)
      {
         d, err := n.Date()
         if err != nil {
            return "", err
         }
         b = d.AppendFormat(b, "2006")
      }
   }
   return string(b), nil
}

func Clean(path string) string {
   m := map[byte]bool{
      '"': true,
      '*': true,
      '/': true,
      ':': true,
      '<': true,
      '>': true,
      '?': true,
      '\\': true,
      '|': true,
   }
   b := []byte(path)
   for k, v := range b {
      if m[v] {
         b[k] = '-'
      }
   }
   return string(b)
}

func User(name string) (map[string]string, error) {
   b, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   if err := json.Unmarshal(b, &m); err != nil {
      return nil, err
   }
   return m, nil
}
