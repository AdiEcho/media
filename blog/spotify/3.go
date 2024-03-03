package spotify

import (
	"154.pages.dev/protobuf"
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
)

func Three() {
	var req http.Request
	req.Header = make(http.Header)
	req.Method = "POST"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.URL = new(url.URL)
	req.URL.Host = "login5.spotify.com"
	req.URL.Path = "/v3/login"
	req.URL.Scheme = "https"
	req.Body = io.NopCloser(bytes.NewReader(body3.Encode()))
	res, err := new(http.Transport).RoundTrip(&req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	res.Write(os.Stdout)
}

var body3 = protobuf.Message{
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Message{
		protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9a8d2f0ce77a4e248bb71fefcb557637")},
	}},
	protobuf.Field{Number: 100, Type: 2, Value: protobuf.Message{
		protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("AQC16FhLEAnXz3AmrmM8VsWA85EqrwpOiS3HUFM2R-escWtteW_eiNTawRqy18tS9CqHdkIalCp31A0S_gy2sqOLhxGTcloaYX4wN8zNbZUuf6kyXQ5CIoHScC9sGykWauDtBSwy9v5gIG7GfofLjwjW5BQOH5xgm5-ywHiHS-G1DRBjR6Asud75ThdbDpeg8oiAdxjmmYAQ7oAuyyKcgg4pkQ")},
	}},
}
