package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Origin"] = []string{"https://app.10ft.itv.com"}
   req.Header["Referer"] = []string{"https://app.10ft.itv.com/3.416.0/androidtv/programmes/10_4008/10a4008a0001"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36"}
   req.Header["X-Requested-With"] = []string{"air.ITVMobilePlayer"}
   req.Method = "GET"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "content-inventory.prd.oasvc.itv.com"
   req.URL.Path = "/discovery"
   req.URL.RawPath = ""
   value := url.Values{}
   value["query"] = []string{"query { brands( filter: { legacyId: \"10/4008\" features: [OUTBAND_WEBVTT MPEG_DASH WIDEVINE] available: \"NOW\" broadcaster: ITV platform: DOTCOM tiers: [\"FREE\" \"PAID\"] } ){ ccid title imageUrl(imageType:ITVX) numberOfAvailableSeries synopses{ ninety epg } legacyId categories tier contentOwner genres { id name } visuallySigned partnership earliestAvailableSeries { earliestAvailableEpisode { broadcastDateTime seriesNumber episodeNumber } } latestAvailableSeries { longRunning fullSeries earliestAvailableEpisode { title broadcastDateTime seriesNumber episodeNumber } latestAvailableEpisode { title broadcastDateTime seriesNumber episodeNumber } } series (sortBy: SEQUENCE_ASC){ ccid seriesNumber legacyId fullSeries seriesType longRunning numberOfAvailableEpisodes latestAvailableEpisode { title broadcastDateTime seriesNumber episodeNumber } titles (sortBy: SEQUENCE_ASC) { ... on Episode { episodeNumber series { longRunning seriesNumber seriesType } contentOwner partnership versions { scheduleEvent { originalBroadcastDateTime } } } ...on Special { categories contentOwner episodeNumber partnership productionYear genres { id name } versions { scheduleEvent { originalBroadcastDateTime } } } ...on Film { categories contentOwner partnership productionYear genres { id name hubCategory } } ccid titleType legacyId brandLegacyId title channel {name} contentOwner partnership regionalisation broadcastDateTime imageUrl(imageType:ITVX) tier visuallySigned nextAvailableTitle { legacyId } series { fullSeries tier seriesNumber longRunning } brand { numberOfAvailableSeries } synopses { ninety epg } latestAvailableVersion { legacyId duration linearContent playlistUrl visuallySigned tier availability { start end } subtitled audioDescribed compliance{ displayableGuidance } variants { features protection } bsl { playlistUrl } } } } } titles (filter: { titleTypes: [SPECIAL FILM] brandLegacyId: \"10/4008\" available: \"NOW\" broadcaster: ITV platform: DOTCOM tiers: [\"FREE\" \"PAID\"] }){ ... on Episode { episodeNumber series { longRunning seriesNumber seriesType } contentOwner partnership versions { scheduleEvent { originalBroadcastDateTime } } } ...on Special { categories contentOwner episodeNumber partnership productionYear genres { id name } versions { scheduleEvent { originalBroadcastDateTime } } } ...on Film { categories contentOwner partnership productionYear genres { id name hubCategory } } ccid titleType legacyId brandLegacyId title channel {name} contentOwner partnership regionalisation broadcastDateTime imageUrl(imageType:ITVX) tier visuallySigned nextAvailableTitle { legacyId } series { fullSeries tier seriesNumber longRunning } brand { numberOfAvailableSeries } synopses { ninety epg } latestAvailableVersion { legacyId duration linearContent playlistUrl visuallySigned tier availability { start end } subtitled audioDescribed compliance{ displayableGuidance } variants { features protection } bsl { playlistUrl } } } titlesToRedirect: titles(filter: { legacyId: \"10/4008\" titleTypes: [SPECIAL FILM] available: \"NOW\" tiers: [\"PAID\" \"FREE\"] } ) { legacyId brandLegacyId title titleType brand { title legacyId } } }"}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   req.Body = nil
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader("")
