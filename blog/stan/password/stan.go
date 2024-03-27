package stan

import (
)

func login() {
   payload = {
      'email': username,
      'password': password,
   }
   enable_h265 = settings.getBool('enable_h265', True)
   enable_4k = settings.getBool('enable_4k', True) if (req_wv_level(WV_L1) and req_hdcp_level(HDCP_2_2)) else False
   return {
      'audioCodecs': 'aac',
      'capabilities.audioCodec': 'aac',
      'capabilities.drm': 'widevine',
      'capabilities.screenSize': '3840x2160',
      'capabilities.videoCodec': 'h264,decode,h263,h265,hevc,mjpeg,mpeg2v,mp4,mpeg4,vc1,vp8,vp9',
      'colorSpace': 'hdr10',
      'drm': 'widevine',
      'features': 'hdr10,hevc',
      'hdcpVersion': '2.2' if enable_4k else '0', #0, 1, 2, 2.2
      'manufacturer': 'NVIDIA', #NVIDIA, Sony
      'model': 'SHIELD Android TV' if (enable_4k or enable_h265) else '', #SHIELD Android TV, BRAVIA 4K 2020
      'os': 'Android-9',
      'screenSize': '3840x2160',
      'sdk': '28',
      'stanName': STAN_NAME,
      'stanVersion': '4.9.1',
      'type': 'console', #console, tv
      'videoCodecs': 'h264,decode,h263,h265,hevc,mjpeg,mpeg2v,mp4,mpeg4,vc1,vp8,vp9',
   }
   data = self._session.post('/login/v1/sessions/app', data=payload).json()
}
