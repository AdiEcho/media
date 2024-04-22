package main

import (
   "bufio"
   "fmt"
   "net/http"
   "slices"
   "strconv"
   "strings"
   "time"
)

var country = map[string]string{
   "Argentina": "country_ar.netset",
   "Australia": "country_au.netset",
   "Bolivia": "country_bo.netset",
   "Brazil": "country_br.netset",
   "Canada": "country_ca.netset",
   "Chile": "country_cl.netset",
   "Colombia": "country_co.netset",
   "Costa Rica": "country_cr.netset",
   "Denmark": "country_dk.netset",
   "Ecuador": "country_ec.netset",
   "Germany": "country_de.netset",
   "Guatemala": "country_gt.netset",
   "Ireland": "country_ie.netset",
   "Mexico": "country_mx.netset",
   "Norway": "country_no.netset",
   "Peru": "country_pe.netset",
   "Sweden": "country_se.netset",
   "United Kingdom": "country_gb.netset",
   "Venezuela": "country_ve.netset",
}

func main() {
   for name, netset := range country {
      address := func() string {
         var b strings.Builder
         b.WriteString("https://raw.githubusercontent.com/firehol/")
         b.WriteString("blocklist-ipsets/master/geolite2_country/")
         b.WriteString(netset)
         return b.String()
      }()
      func() {
         res, err := http.Get(address)
         if err != nil {
            panic(err)
         }
         defer res.Body.Close()
         var blocks []block
         scan := bufio.NewScanner(res.Body)
         for scan.Scan() {
            text := scan.Text()
            if !strings.HasPrefix(text, "#") {
               var (
                  b block
                  ok bool
                  raw_size string
               )
               if b.ip, raw_size, ok = strings.Cut(text, "/"); ok {
                  b.size, err = strconv.Atoi(raw_size)
                  if err != nil {
                     panic(err)
                  }
                  blocks = append(blocks, b)
               }
            }
         }
         slices.SortFunc(blocks, func(a, b block) int {
            return a.size - b.size
         })
         fmt.Printf("%v %+v\n\n", name, blocks[:9])
      }()
      time.Sleep(time.Second)
   }
}

type block struct {
   size int
   ip string
}
