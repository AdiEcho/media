from mitmproxy import ctx, http

def response(f: http.HTTPFlow) -> None:
   if f.request.path.startswith('/generetic/generated/chunks/12.ff734ba67f44a707e609.js'):
      f.response.text = open('2.js', 'r').read()
