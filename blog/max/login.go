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
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "default.any-amer.prd.api.max.com"
   req.URL.Path = "/login"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   
   req.Header["Accept"] = []string{"application/json, text/plain, */*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{"wbd-return-url=https%3A%2F%2Fplay.max.com%2Fmovie%2F127b00c5-0131-4bac-b2d1-40762deefe09", "transientID=bf61708a-cf38-4abe-8dc4-638f63766fc8", "session={\"uuid\":\"b7804758-3377-4190-b429-ea7dee273880\",\"created\":1717904271712,\"expires\":1717906086026}", "pageLanguage=en-US", "dotcomtrack=", "usprivacy=1---", "AMCV_BC501253513148ED0A490D45%40AdobeOrg=-1124106680%7CMCIDTS%7C19884%7CMCMID%7C27251383584571565567340020280998448675%7CMCOPTOUT-1717911476s%7CNONE%7CvVersion%7C5.2.0", "AMCVS_BC501253513148ED0A490D45%40AdobeOrg=1", "GI_WEB_SDK_SONIC_DEVICE_ID=d242a38d-0693-4e3c-9607-ae9e549d0040", "st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi1kNDJkMjdkYS01NzMwLTQ3MjYtYWUyMy05YjhiNjk4YWZjMzUiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6ZmM2ZTVmNWMtZjdiYi00ZDZmLTg3MTktN2U3ODQyZTM3NzQ2IiwiaWF0IjoxNzE3OTA0Mjc0LCJleHAiOjIwMzMyNjQyNzQsInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6dHJ1ZSwiZGV2aWNlSWQiOiJkMjQyYTM4ZC0wNjkzLTRlM2MtOTYwNy1hZTllNTQ5ZDAwNDAifQ.jgXYw-3x77f5xJ0Xzgkcgtfq7qs2MzJwj1ZJ85_aYzBjrc2xLpEe2v3wqMxNcipyoNTvvL78wXcDajmKtSuxOMxIRx5L2sGNZDdj5kyEZtdAjJuxOxrbfzG5US-S59wYmyk29lHGwPyflxdlc2N3YTlrL1Y7v7g5D9wmLOxa2mzvouUfFa_4817sMozoDCuiuZuI0I_bHhXUe1DwSfAJFHGo_Nbd3kdnK0vCfGHhvXYiHmRezPJFLp1fg6yXV1aE4Ygwef5nevbL7aYARytpntFkYO3nIZTWSXIEC4sfj3uwsRhJ9KjHUR3osOAoA8K5dalcn2aYXzlSj5dOSk4bfw", "gi_ls=0", "anonymousConsentElections=MjAyNC0wNi0wOVQwMzozNzo1Nlp8VVMtVFh8U2V8"}
   req.Header["Origin"] = []string{"https://auth.max.com"}
   req.Header["Referer"] = []string{"https://auth.max.com/"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-site"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["Traceparent"] = []string{"00-56448d7486444c663b66cc73bb567e80-b196433e8105b6b1-01"}
   req.Header["Tracestate"] = []string{"wbd=session:b7804758-3377-4190-b429-ea7dee273880"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Header["X-Device-Info"] = []string{"beam/4.1.0 (desktop/desktop; Windows/10; d242a38d-0693-4e3c-9607-ae9e549d0040/da0cdd94-5a39-42ef-aa68-54cbc1b852c3)"}
   req.Header["X-Disco-Arkose-Sitekey"] = []string{"B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"}
   req.Header["X-Disco-Arkose-Token"] = []string{"71117d738d82cf455.4147107801|r=us-east-1|meta=3|meta_width=300|metabgclr=transparent|metaiconclr=%23555555|guitextcolor=%23000000|lang=en|pk=B0217B00-2CA4-41CC-925D-1EEB57BFFC2F|at=40|sup=1|rid=39|ag=101|cdn_url=https%3A%2F%2Fwbd-api.arkoselabs.com%2Fcdn%2Ffc|lurl=https%3A%2F%2Faudio-us-east-1.arkoselabs.com|surl=https%3A%2F%2Fwbd-api.arkoselabs.com|smurl=https%3A%2F%2Fwbd-api.arkoselabs.com%2Fcdn%2Ffc%2Fassets%2Fstyle-manager"}
   req.Header["X-Disco-Client"] = []string{"WEB:10:beam:4.1.0"}
   req.Header["X-Disco-Client-Id"] = []string{"web1_prd:1717904286:2eb1d51660e2b286209aff73650cc08891f15cbcf75ff5b6dae34e1c039465d6"}
   req.Header["X-Disco-Params"] = []string{"realm=bolt,bid=beam,siteLookupKey=beam_us,features=ar"}
   req.Header["X-Gisdk"] = []string{"clientId=9f964812-8935-4293-a135-81be80f14c77"}
   req.Header["X-Wbd-Ace"] = []string{"MjAyNC0wNi0wOVQwMzozNzo1Nlp8VVMtVFh8U2V8"}
   req.Header["X-Wbd-Preferred-Language"] = []string{"en-US,en"}
   req.Header["X-Wbd-Session-State"] = []string{"localization:eyJhbGciOiJIUzI1NiJ9.eyJtZXRhZGF0YSI6eyJpZCI6ImxvY2FsaXphdGlvbiIsInZlcnNpb24iOiJ2MiIsInN0YXR1cyI6IlNUQVRVU19BQ1RJVkUifSwic2VsZWN0ZWRMYW5ndWFnZXMiOlsiZW4tVVMiXSwiZm9ybWF0IjoiZW4tVVMifQ.FraAyXAvLZCA0Ry1N-ayRBTJzFuDeWxUp_eMWWGYMy0;overrides:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..QqHs8KzfzfV19RgN.knBEmyON1klghkQzxOjw-t3NGoyyqZhr5HxcQsuEMhMtPE6mCO8yK-ttv5PQIcd_LXqnlxgTy4UeAy_0LwijVsRB0aPpa63ca7nGvY-j5akepq_8GnYEDmwSua8J_nTMgM_qJcdZYhf1YDVJhUh8bQu4pn1eOdDaKxlUgHF26kKjN6vfh0p0MziVJFHRLxR6IjLOwyjdht95dXh6VVcz1pIr-tzHRLnA70Md6hzTPxwddDGEpffeBYVOtQJAipUqv6cT6c3tRXj5Piq7E446chGFkvsbR-AowgT4I-DuzaxoxJDHyp8eI5AB8uIEk_1rvUG8O64lb-yMfCxeI9_wmGl73lGpLf7ENLhIwpcRgifitWPMWWk8eYvk9PC1YBIHp1ACZYNhwq5S1XMdEgz7E8t4z0l0Ps06AnWjpgh6BFE7jEjnUw.vZOtoGPOj3cTQBMNQPj9tg;experience:eyJhbGciOiJIUzI1NiJ9.eyJtZXRhZGF0YSI6eyJpZCI6ImV4cGVyaWVuY2UiLCJ2ZXJzaW9uIjoidjEiLCJzdGF0dXMiOiJTVEFUVVNfQUNUSVZFIn0sImFwcEV4cGVyaWVuY2UiOiJtYXgifQ.CjvqZF6TuvLhWuoYEYr_r02i1sW8yi1EwzQZ46A6jlg;user:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..NPBm1OvQ9D8TkfJC.AcCTZQk8aUlrlCymTml3Vz2GAC-Iam_SVmmFpcnIbYVu-Xf992hbyFX6Zl0FZMhggI4cvw3DWRNERMhE_0QhBx_o7czoSIFMSPT9enGaq46R3wNR4afqGFfqvrM-yvHA6PWgj4mdy_m6CPe9sn-UEmJW0NKc0OvtXXxD8tPStquaX7Zx61TFjAKz07uL2hUVRvaGqWqAUyArW1xJIm21XxiKhw.IF0jI53AW7TtAVVZTXk3ng;device:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..ZQBy7VLWVZjgP5Bh.zGozNh46cIsytAILG8Q_z2bJSF8-0W7CaL4nowBNwagnHu2f6VEjRubyTv-RqJtwmRUPIArFxSlYJACYmuKIRgnxL6vbtP_327ZPoSJkqCfD4Z8WwHjEkMOCtkWz6UVQv7PdnJnvsMCus5kmC4sXQjGznHtrB1GkVW3ty3nCAKOI8MTnSTgZ8WpgU-xvtT62VdbpNDCn6N4CZCgZljy5w2BVFHVXFjDeeVS4kzshkzDPT1z1XNBY74PUaSIjCHE7fUSrl3eosYbO2iHCpRFMc2xa9TwjVjPfKpTHyFguME8tAECMnMUT9QC4KhxNtu89A3kNYLVo6F3616o5aKIqrP164ZHYl8TvfKaccX4X9cm0NW2SjC-uA2r_KiKOTH0lkQllcja4pH2363h_GOP8NtNckPVX84r-S-zC0J1VDTtx6jyz0lpcudNefbFpMOGfGY3Er-SvuAP2OpQMFyT9jHqTPq7mGWXmgtWPuvaGGc-2MEByJ0u5Gd40Rz1QdACz1_z4YcspKw.W2aHWtx_yKqrUv7qSCJVjA;token:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..a52P1bKdZ7p8x4Fl.wCD4irm_9ejG5RpRoERDX-fj-O0HQif3yVS1kuaKNKacYmyYfimdESLzC9wHzES6BAbQ5nox6sMiQGwxJEZeofBTT3Q2DQ_6ZSJj2oTnDEWrbVIZV5cyFpDP5XbjtLdYdVT5QHAnvcS88GQUqexMG2u_lsB5aPPiLjZx--mi-ZaiBh8kbbTrIQLCTROPOYW3m79RXZfS8KPsVqXPIhgPosqPt5PzhS3wopb_X3oIaUCdxLJdPFJuMw3fcGQDEcAw.EvUv2r1XJ-u_gw4IRVDBtQ;geolocation:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0.._Ejv9BZkz3eJ5CGz.6xnRZioSr9xYT_Wr9zzesQlL39JVCYu71XL2bDc6uYGa64rN7JS5Sz2U_6iV6J5SomkZRd-Tp2Ne5riHoV_m5nHiYpICZHNzMWC4Ri9uRruxUK6BE6xgLyiZMZZoHoBT97pzDdudKws9IFBWJcjWYt3ZjYZrgZ1xYbAcdDKcz05gdhPIqW88x-uAXATvHSNVMWeKz-qN4YJzUOGY7N33NSjEEfCIZ99zsXx7c3NoAO8304xXHnvfWmVz3PrbSggjg8FCQiwYD8e-6xlHeB1gFjIuKRv_4hnMNC1IAyll64GApI3UaPOH9mct19tTU1XBzYLq9-44KNs7v_6XkGZ3cdhCKfiMg_2tkPvpMC7YEWBUZVpffTcLPgKagmXQ9JjR6YN9NLAU4iUraRFZTQUQZzd3hui9PHN7O_xTbzDgf3bPTY416yJm0mNNz3SPODYP_nWkFrJGB3jhKb-Rc9_-LkpyYG0qbuEczo_TJtA2X_aYGfn0n-WwpXg.R2KPPJN-odS7sobW9cRE5Q;customer_segments:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..fV-jNnIVrs5pxV0e.z5BTGxOTrp9-vz2BpzX5nutQ8OtanVpRH1lWxaPzhN--Nac46Pqa9eTtwmqAJZbU9pJRLRjsP_x3ZknuZSTUpcps43TnjccXtOrp4wUfQy_3jFuYDG_6gh9lKBIlc4esZh_9eqbi93Blscg_HcC9sAj9J0wi-bbymqbcwoPAk4aw83xzYkaZbNj5-Eb0-TnQhDqGIxzDo7494rtM8rHWC388ojmk1HUon1cX.7S5itEeRtyJlDENWlU03-Q;capabilities:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..-zwy5SjMZgc41VqW.4ECNuwP_8LjUwZbLZsMfxX-ZA3cgSMBFYJBuZPvJotD70T3PETRRMVtf806d5gDqFsdHn2y0Atlvc0BDupX7Q_z_o9IQJrL3v-IGWk8tjGLr1xKG3cjm382B-gg7qp78uqoiO7CIzwn234ZmQMDyxSa3UvV12MckOSyEJAZo6_khsHfg-iEeA60J9_1PzJFHHwrULFaYeT3rnOCFLshtY668ayUYjYUHMk4KsOO3AKY5xx7HAiTJC6SMwROaDYG9ilARJD0eYu48GWC2NN-X6aqAG80hlBUidi0IbSaHAmgkUr4nuephXPr6Tj4eUiZcnw._3g8_ta-cL6l5gqX5s8cEg;mockruleset:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..e_dRU9pSlrok5xN4.hwc.RxiG3wyRTKRBmQvS4u_0jA;profile:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..vXHJd2brbHwd4FsX.hNbbA3uwxy8pyFdv_C3Rn0ZCQqYkEKG4xBrw0bKqOk7cvucZ1YgT-XHSB78Uo-TYuQxAKoqOr27PIvWjDmirC6YLC8KCQtIsVIXnUtkE2aELtM9WcFU_i_6-lO38AcNKIi1eBxU0cQzkIVX-qFcdy-pAk3djw96l4jpYcItH6DB_F4WaNzbL2IyrPGbsjBpRCMjL2F9qff0QssHJTNZc7O9jZJkM1J5lp3o-em0EELd-cxOHfVFTw9GlLixz1Dx7q9_JW55OUTqnN60uKq3ZVu8pNIRwqY7DP5p_tpyAxw.qM8FNtKl67HHCV2nuQUSRw;privacy:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..00OlnJWrnCkBAIzo.-Y8OByeyhHr4PFrcnprhYJUlrUN6P5I51UkBsbjVyZPE05KJei4NWxVdNzunMJBXeF-ar8zz1abn7Isq5xg97jryrynyLb9-K-JPcemPLElwDAzPGk-em4eFfxnnezET17Nb0lBwf_z_Yh_bwgaqiFMJqkRgvmii-XMDOahFtGbthxXJMVxydgxtz4YJUfgleZCBYt23hkyIFp-znnuIoHRbRzLrB6AcNp4-0X61aTtg9NOD7I0.y3eAwh5BdhtNX2h-_Zvp_w;admin:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..eVx-jYKjRu5eq4_U.TUm51xDn2om45AL-2OosP52WFDB3p95IvSFLJYYNmlGVVoz0ZgrYT9eW9yBOvY1mpHUTrQfEIBUtPuPLttJniEzygQ.qbd8ZFZ53mX8A3pIA6kyKg;legacy_request_context:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..2N30yshJr0ynREcF.y4OqPNKAa-hOqbz7fWyTpVz70oaPazHxC4ufbsPfEApT-owSPzi6ME7SMGEcSyeIiG4_9runu7qCZyosm-rKC2tiIzstE4SVYPmpAihIDSTlCKKi08U5JEMZ5ynsqa0s2NayebkybDz9eNylB6jvLEFy6dvAD6RMWd4RxvqJK31iDprzgTzThk2Hz-dopc9zjj5flT2ox5PispNEe0R-48kT6snCrVbLzLF_lKQPLoSL4WVr1svl7rPs-PFsjh3Jorb12Z9ghKbY43CMdxGnG-zi6W9TB7Hn_asoefN7bogHeEK89_aTJJSvqo7_AJPO4Ld54W9-XaDKVXRjWoTUKE01754YOryNknpM-sQQ7pNOf4XPiEN_tTn86qqde-XUxsTZ.xIGcbK30fB1qn4BbdpjqOg;content:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..eDLT6-Yg11QIn9eE.Tnh20XB5yz9dnbt3MUxXm-4kmxxC_H4IDXlb7yRTBOkDoHiE1-3H9ULHWA69lXOEOmQq8DF8XQGoBzL1clUwZm_GBkC8trjIONwOzVWmL2zTld_BGRcCCITzJSB6NS0MHTdDMXtbqRTew9-9MNkL3Lg5OQzZFpsnGyz87Ke0PEsWe4WWddletoGzuQUcMUqOlzFMA3Jjm8tD7ZcvF9E1Ncnr-4UsKSeAdsFJhm-CLVPI6RiPuAfb1ZEFGpk5o66S9_AAKNzbI0Jo242z.-A18Aym8YWOQWAf-RFmO8A;gauthodc:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..DScUfVH24CQQyaE9.gdLyy_XRutLTBci2fJ9Q13z92pufjwuKN_iUhZWkf5JPqLWXCiqeG5UIkj35lx__2z7CBxi352TaP-AwJvfJtrNmpyrOiw.81QcwnnB7MyprWYD8Fhq-g;consents:eyJhbGciOiJIUzI1NiJ9.eyJtZXRhZGF0YSI6eyJpZCI6ImNvbnNlbnRzIiwidmVyc2lvbiI6InYxIiwic3RhdHVzIjoiU1RBVFVTX0FDVElWRSIsImlhdCI6IjIwMjQtMDYtMDlUMDM6Mzc6NTZaIn0sInN0YXR1cyI6IkNPTlNFTlRfU1RBVFVTX0VMRUNUSU9OIiwiZ2VvIjoiVVMtVFgiLCJ0cHYiOnRydWUsImVtYyI6ZmFsc2V9.lsc33o96qq-4CGUBq01_YeBI-JQ4bjl9EygEv1QoiRw;diana:eyJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiZGlyIn0..RmNXRmmJaiOZA70k.k0iT23EirwGJKzDU1wtCPZRONabR0Kg2uXRHS9cNwR447N_yawraZ4nwIoUXtroTLSCyJn-Sc_fhU_c4ghs7v7RPY69cgEbC9n-f5OdtntuYBBnXZkmEA2vbmW-GJA-1-t9WUIb5aytkTsG1aPCHbSn3z6o.KIKOYQUchkhaBmS-J2HrpQ;"}
   
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(`
{
 "credentials": {
  "username": "EMAIL",
  "password": "PASSWORD"
 }
}
`)
