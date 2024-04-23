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

var countries = []struct{
   country string
   name string
}{
   {"Argentina", "country_ar.netset"},
   {"Australia", "country_au.netset"},
   {"Bolivia", "country_bo.netset"},
   {"Brazil", "country_br.netset"},
   {"Canada", "country_ca.netset"},
   {"Chile", "country_cl.netset"},
   {"Colombia", "country_co.netset"},
   {"Costa Rica", "country_cr.netset"},
   {"Denmark", "country_dk.netset"},
   {"Ecuador", "country_ec.netset"},
   {"Germany", "country_de.netset"},
   {"Guatemala", "country_gt.netset"},
   {"Ireland", "country_ie.netset"},
   {"Italy", "country_it.netset"},
   {"Mexico", "country_mx.netset"},
   {"Norway", "country_no.netset"},
   {"Peru", "country_pe.netset"},
   {"South Africa", "country_za.netset"},
   {"Spain", "country_es.netset"},
   {"Sweden", "country_se.netset"},
   {"United Kingdom", "country_gb.netset"},
   {"Venezuela", "country_ve.netset"},
}

func main() {
   for _, country := range countries {
      address := func() string {
         var b strings.Builder
         b.WriteString("https://raw.githubusercontent.com/firehol/")
         b.WriteString("blocklist-ipsets/master/geolite2_country/")
         b.WriteString(country.name)
         return b.String()
      }()
      var blocks []block
      func() {
         res, err := http.Get(address)
         if err != nil {
            panic(err)
         }
         defer res.Body.Close()
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
      }()
      min := slices.MinFunc(blocks, func(a, b block) int {
         if v := a.size - b.size; v != 0 {
            return v
         }
         return len(a.ip) - len(b.ip)
      })
      fmt.Printf("{%q, %q},\n", country.country, min.ip)
      time.Sleep(time.Second)
   }
}

type block struct {
   size int
   ip string
}
