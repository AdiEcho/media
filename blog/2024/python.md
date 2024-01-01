# Python

If you want to MITM a Python program, you might need to set one or more of
these:

~~~ps1
$env:HTTPS_PROXY = 'http://127.0.0.1:8080'
$env:REQUESTS_CA_BUNDLE = 'C:\Users\Hello\.mitmproxy\mitmproxy-ca.pem'
$env:SSL_CERT_FILE = 'C:\Users\Hello\.mitmproxy\mitmproxy-ca.pem'
~~~
