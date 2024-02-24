package main

import (
   "io"
   "net/http"
   "os"
   "strings"
)

// SkyOTT client="NBCU-WEB-v8",signature="i9ZXGWBOY0IQM7Eehx47Kv9vfqw=",timestamp="1708737167",version="1.0"
func calculate_signature(
   method, path string, headers http.Header, payload []byte,
) string {
   timestamp = time.Now().Unix()
   var text_headers []string
   for key := range headers {
      if strings.HasPrefix(key, "x-skyott") {
         text_headers = append(text_headers, key, ": ", headers.Get(key), "\n")
      }
   }
   slices.Sort(keys)
   headers_md5 = hashlib.md5(text_headers.encode()).hexdigest()
   if sys.version_info[0] > 2 and isinstance(payload, str) {
      payload = payload.encode('utf-8')
   }
   payload_md5 = hashlib.md5(payload).hexdigest()
   to_hash = (
      '{method}\n{path}\n{response_code}\n{app_id}\n{version}\n{headers_md5}\n{timestamp}\n{payload_md5}\n'
   ).format(
      method=method, path=path, response_code='', app_id=self.app_id,
      version=self.sig_version, headers_md5=headers_md5, timestamp=timestamp,
      payload_md5=payload_md5,
   )
   hashed = hmac.new(self.signature_key, to_hash.encode('utf8'), hashlib.sha1).digest()
   signature = base64.b64encode(hashed).decode('utf8')
   return {
      'x-sky-signature': 'SkyOTT client="{}",signature="{}",timestamp="{}",version="{}"'.format(
         self.app_id, signature, timestamp, self.sig_version
      )
   }
}
