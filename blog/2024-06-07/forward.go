package main

import (
   "fmt"
   "io"
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
   {"Argentina", "ipip_country_ar.netset"},
   {"Australia", "ipip_country_au.netset"},
   {"Bolivia", "ipip_country_bo.netset"},
   {"Brazil", "ipip_country_br.netset"},
   {"Canada", "ipip_country_ca.netset"},
   {"Chile", "ipip_country_cl.netset"},
   {"Colombia", "ipip_country_co.netset"},
   {"Costa Rica", "ipip_country_cr.netset"},
   {"Denmark", "ipip_country_dk.netset"},
   {"Ecuador", "ipip_country_ec.netset"},
   {"Egypt", "ipip_country_eg.netset"},
   {"Germany", "ipip_country_de.netset"},
   {"Guatemala", "ipip_country_gt.netset"},
   {"India", "ipip_country_in.netset"},
   {"Indonesia", "ipip_country_id.netset"},
   {"Ireland", "ipip_country_ie.netset"},
   {"Italy", "ipip_country_it.netset"},
   {"Latvia", "ipip_country_lv.netset"},
   {"Malaysia", "ipip_country_my.netset"},
   {"Mexico", "ipip_country_mx.netset"},
   {"Netherlands", "ipip_country_nl.netset"},
   {"New Zealand", "ipip_country_nz.netset"},
   {"Norway", "ipip_country_no.netset"},
   {"Peru", "ipip_country_pe.netset"},
   {"Russia", "ipip_country_ru.netset"},
   {"South Africa", "ipip_country_za.netset"},
   {"South Korea", "ipip_country_kr.netset"},
   {"Spain", "ipip_country_es.netset"},
   {"Sweden", "ipip_country_se.netset"},
   {"Taiwan", "ipip_country_tw.netset"},
   {"United Kingdom", "ipip_country_gb.netset"},
   {"Venezuela", "ipip_country_ve.netset"},
}

func main() {
   for _, country := range countries {
      address := func() string {
         var b strings.Builder
         b.WriteString("https://raw.githubusercontent.com/firehol/")
         b.WriteString("blocklist-ipsets/master/ipip_country/")
         b.WriteString(country.name)
         return b.String()
      }()
      blocks, err := get(address)
      if err != nil {
         panic(err)
      }
      slices.SortFunc(blocks, func(a, b blocklist) int {
         if v := a.size - b.size; v != 0 {
            return v
         }
         return len(a.ip) - len(b.ip)
      })
      fmt.Printf("{%q, %q},\n", country.country, blocks[0].ip)
      time.Sleep(99 * time.Millisecond)
   }
}

func get(address string) ([]blocklist, error) {
   resp, err := http.Get(address)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   text, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var blocks []blocklist
   for _, line := range strings.Split(string(text), "\n") {
      if !strings.HasPrefix(line, "#") {
         var (
            b blocklist
            ok bool
            raw_size string
         )
         if b.ip, raw_size, ok = strings.Cut(line, "/"); ok {
            b.size, err = strconv.Atoi(raw_size)
            if err != nil {
               return nil, err
            }
            blocks = append(blocks, b)
         }
      }
   }
   return blocks, nil
}

type blocklist struct {
   size int
   ip string
}

