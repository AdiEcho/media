# Hulu

this is it:

~~~
GET /hulu/v1/dash/196861183/manifest.mpd?ad_cdns=da%2Cfa&boundary_signaling=hulu_unified&cdns=fa%2Cda&cluster=green&is_noah=0&pb_config=CAISIgogChcIgg4QuAgiBEgyNjQqBEhJR0gyAzUuMhIFRklSU1QaDgoMCgUKA0FBQxIDT05FIh4KFwoIV0lERVZJTkUSB01PRFVMQVIaAkwzEgNPTkUqFAoEREFTSBABGAEgATgBQAFIAVgDMh0KFgoERk1QNBIMCgRDRU5DEgRDRU5DIAESA09ORUABUKYB&qos=Cg0KBEgyNjQQhgMYoJwBCg0KBEgyNjUQjgIYoJwB&user_id=252683275&auth=1699500109_cd89c7e39b49c1d1f74e1578ed07daee HTTP/1.1
Host: vodmanifest.hulustream.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/4.7.2
content-length: 0
~~~

from:

~~~
POST https://play.hulu.com/v6/playlist HTTP/2.0
x-hulu-user-agent: androidv4/5.3.0+12541-google/b3d7ca343f99384;OS_23,MODEL_Android SDK built for x86
user-agent: Hulu/5.3.0+12541-google (Android 6.0; en_US; Android SDK built for x86; Build/MASTER;)
content-type: application/json

{
  "version": 5012541,
  "cp_session_id": "781d378c-7887-4b7c-a167-eb3f1ea8bb29",
  "deejay_device_id": 166,
  "device_identifier": "8cd3d3f6-8ea2-387f-ac37-989991c17789",
  "content_eab_id": "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326",
  "format": "json",
  "device_ad_id": "c4f00ca6-5323-4813-a6c2-069d4968bdd9",
  "include_t2_rev_beacon": true,
  "ignore_kids_block": false,
  "is_tablet": false,
  "unencrypted": true,
  "limit_ad_tracking": false,
  "network_mode": "LTE",
  "playback": {
    "audio": {
      "codecs": {
        "selection_mode": "ONE",
        "values": [
          {
            "type": "AAC"
          }
        ]
      }
    },
    "drm": {
      "disable_representation_kids": false,
      "hdcp": false,
      "multi_key": false,
      "selection_mode": "ONE",
      "values": [
        {
          "security_level": "L3",
          "type": "WIDEVINE",
          "version": "MODULAR"
        }
      ]
    },
    "manifest": {
      "live_dai": true,
      "multiple_cdns": true,
      "patch_updates": true,
      "secondary_audio": true,
      "type": "DASH",
      "unified_asset_signaling": true
    },
    "segments": {
      "selection_mode": "ONE",
      "values": [
        {
          "encryption": {
            "mode": "CENC",
            "type": "CENC"
          },
          "https": true,
          "muxed": false,
          "type": "FMP4"
        }
      ]
    },
    "version": 2,
    "video": {
      "codecs": {
        "selection_mode": "FIRST",
        "values": [
          {
            "level": "5.2",
            "profile": "HIGH",
            "type": "H264",
            "height": 1080,
            "width": 1794
          }
        ]
      }
    }
  },
  "interface_version": "1.11.11",
  "privacy_opt_out": "NO",
  "support_brightline": false,
  "token": "nk77TZQgj1xc245GV2MmvXxCSFY-02iNFieIYbDEORlBOt32yw--pGK/_drQkpQ3f..."
}
~~~
