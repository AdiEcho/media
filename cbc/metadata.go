package cbc

import (
   "strconv"
   "strings"
)

type Catalog_Gem struct {
   Content []struct {
      Lineups []struct {
         Items []Lineup_Item
      }
   }
   Selected_URL string `json:"selectedUrl"`
   Structured_Metadata Metadata `json:"structuredMetadata"`
}

type Metadata struct {
   Part_Of_Series *struct {
      Name string // The Fall
   } `json:"partofSeries"`
   Part_Of_Season *struct {
      Season_Number int `json:"seasonNumber"`
   } `json:"partofSeason"`
   Episode_Number int `json:"episodeNumber"`
   Name string
   Date_Created string `json:"dateCreated"` // 2014-01-01T00:00:00
}

func (Metadata) Owner() (string, bool) {
   return "", false
}

func (m Metadata) Title() (string, bool) {
   return m.Name, true
}

func (m Metadata) Episode() (string, bool) {
   if m.Episode_Number >= 1 {
      return strconv.Itoa(m.Episode_Number), true
   }
   return "", false
}

func (m Metadata) Season() (string, bool) {
   if p := m.Part_Of_Season; p != nil {
      return strconv.Itoa(p.Season_Number), true
   }
   return "", false
}

func (m Metadata) Show() (string, bool) {
   if m.Part_Of_Series != nil {
      return m.Part_Of_Series.Name, true
   }
   return "", false
}

func (m Metadata) Year() (string, bool) {
   if m.Part_Of_Series != nil {
      return "", false
   }
   year, _, _ := strings.Cut(m.Date_Created, "-")
   return year, true
}
