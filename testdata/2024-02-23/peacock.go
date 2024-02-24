package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

  def calculate_signature(self, method, url, headers, payload='', timestamp=None):
    if not timestamp:
      timestamp = int(time.time())


    if url.startswith('http'):
      parsed_url = urlparse(url)
      path = parsed_url.path
    else:
      path = url


    #print('path: {}'.format(path))


    text_headers = ''
    for key in sorted(headers.keys()):
      if key.lower().startswith('x-skyott'):
        text_headers += key + ': ' + headers[key] + '\n'
    #print(text_headers)
    headers_md5 = hashlib.md5(text_headers.encode()).hexdigest()
    #print(headers_md5)


    if sys.version_info[0] > 2 and isinstance(payload, str):
      payload = payload.encode('utf-8')
    payload_md5 = hashlib.md5(payload).hexdigest()


    to_hash = ('{method}\n{path}\n{response_code}\n{app_id}\n{version}\n{headers_md5}\n'
              '{timestamp}\n{payload_md5}\n').format(method=method, path=path,
                response_code='', app_id=self.app_id, version=self.sig_version,
                headers_md5=headers_md5, timestamp=timestamp, payload_md5=payload_md5)
    #print(to_hash)


    hashed = hmac.new(self.signature_key, to_hash.encode('utf8'), hashlib.sha1).digest()
    signature = base64.b64encode(hashed).decode('utf8')


    return {'x-sky-signature': 'SkyOTT client="{}",signature="{}",timestamp="{}",version="{}"'.format(
        self.app_id, signature, timestamp, self.sig_version)}
