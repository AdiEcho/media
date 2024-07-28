import hashlib
import hmac
import time

drm_proxy_secret = b'Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6'

now = int(time.time() * 1000)

print(now)

signature = hmac.new(
   drm_proxy_secret,
   msg = f'{now}widevine'.encode(),
   digestmod = hashlib.sha256,
).hexdigest()

print(signature)
