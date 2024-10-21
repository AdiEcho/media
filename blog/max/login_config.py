
MAX_DEVICE_LOGIN_URL = 'https://default.prd.api.discomax.com/authentication/linkDevice/login'
MAX_CODE_PAIR = 'https://default.prd.api.discomax.com/authentication/linkDevice/initiate'
MAX_BOOTSTRAP_URL = 'https://default.prd.api.discomax.com/session-context/headwaiter/v1/bootstrap'
MAX_TOKENS_URL = 'https://default.prd.api.discomax.com/token'
MAX_TOKENS = "max_tokens.txt"


MAX_HEADERS = {
    'origin': 'https://webos.play.max.com',
    'tracestate': 'wbd=session:93df873e-b37e-428c-a4e5-c327388546fe',
    'x-disco-client': 'LGTV:4.9.0-05.00.03:beam:1.1.2.2',
    'accept-language': 'en-GB',
    'traceparent': '00-d1f9f363654e3b2592f297b6460b9eaa-43e9ba466b673f91-01',
    'x-disco-params': 'realm=bolt,bid=beam,siteLookupKey=beam_us,features=ar',
    'content-type': 'application/json',
    'accept': 'application/json, text/plain, */*',
    'x-device-info': 'beam/1.1.2.2 (LG/OLED55C9PVA; webOS/4.9.0-05.00.03; 04137aa2-1e1e-6f52-7a08-12249c864690/9e1b83a7-ddde-42c9-b335-a54232bd2a9f)',
    'user-agent': 'Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.34 Safari/537.36 WebAppManager',
    'referer': 'https://webos.play.max.com/',
}


def get_headers():
    return MAX_HEADERS.copy()


def get_tokens_using_code(init_token, x_wbd_session_state):
    data = {}
    headers = get_headers()
    url = MAX_DEVICE_LOGIN_URL
    headers['Authorization'] = f'Bearer {init_token}'
    headers['x-wbd-session-state'] = x_wbd_session_state

    return {
        'url': url,
        'json': data,
        'headers': headers,
    }


def get_init_token_data():
        headers = MAX_HEADERS.copy()

        params = {
            'realm': 'bolt',
        }

        url = MAX_TOKENS_URL
        return {
            'url': url,
            'params': params,
            'headers': headers,
        }


def get_code_pair(init_access_token, x_wbd_session_state):
        url = MAX_CODE_PAIR
        headers = MAX_HEADERS.copy()
        headers['Authorization'] = f'Bearer {init_access_token}'
        headers['x-wbd-session-state'] = x_wbd_session_state
        json_data = {}

        return {
            'url': url,
            'json': json_data,
            'headers': headers,
        }


def get_wbd_session_state_info(access_token):
        headers = MAX_HEADERS.copy()
        headers['authorization'] = f"Bearer {access_token}"
        json_data = {}

        url = MAX_BOOTSTRAP_URL
        return {
            'url': url,
            'json': json_data,
            'headers': headers,
        }
