package cbc

import "time"

type Metadata struct {
   Part_Of_Series *struct {
      Name string // The Fall
   } `json:"partofSeries"`
   Part_Of_Season struct {
      Season_Number int `json:"seasonNumber"`
   } `json:"partofSeason"`
   Episode_Number int `json:"episodeNumber"`
   Name string
   Date_Created string `json:"dateCreated"` // 2014-01-01T00:00:00
}

func (Metadata) Owner() (string, bool) {
   return "", false
}

type Catalog_Gem struct {
   Content []struct {
      Lineups []struct {
         Items []Lineup_Item
      }
   }
   Selected_URL string `json:"selectedUrl"`
   Structured_Metadata Metadata `json:"structuredMetadata"`
}

/////////////////////

func (m Metadata) Series() string {
   return m.Part_Of_Series.Name
}

func (m Metadata) Season() (int64, error) {
   return m.Part_Of_Season.Season_Number, nil
}

func (m Metadata) Episode() (int64, error) {
   return m.Episode_Number, nil
}

func (m Metadata) Title() string {
   return m.Name
}

func (m Metadata) Date() (time.Time, error) {
   return time.Parse("2006-01-02T15:04:05", m.Date_Created)
}
