import base64
import sys
import json, time
from loguru import logger
import requests
import pathlib
import login_config as client_config
session = requests.Session()

def get_device_code(init_access_token, wbd_session_state=None):
    code_pair_request = session.post(**client_config.get_code_pair(init_access_token, wbd_session_state))
    if code_pair_request:
        code = code_pair_request.json()['data']['attributes']['linkingCode']
        return code
    else:
        logger.error("Coudlnt retrieve code")
        sys.exit(1)

def get_wbd_session_state( access_token):
    token_response = session.post(**client_config.get_wbd_session_state_info(access_token))
    if not token_response:
        logger.error('error getting payload entitlemnt')
        logger.error(token_response.text)
        sys.exit(1)
    get_wbd_session_state_info = token_response.headers['x-wbd-session-state']
    return get_wbd_session_state_info

def get_init_access_token():
    token_response = session.get(**client_config.get_init_token_data())
    if not token_response:
        logger.error('error getting req init token ')
        logger.error(token_response.text)
        logger.info("If you are certain email and password are correct, change proxy !!!!")
        sys.exit(1)
    init_access_token = token_response.json()['data']['attributes']['token']
    return init_access_token

def save_token(json_response, access_token, expired):
        json_response['expires_in'] = expired
        json_response['access_token'] = access_token
        MAX_TOKENS_FILE = pathlib.Path(client_config.MAX_TOKENS)
        MAX_TOKENS_FILE.write_text(json.dumps(json_response))

        logger.info('HBO Tokens was saved to cookies_Folder')

def max_login():
    logger.info("Going to get tokens using code_pair scheme")
    token_response = code_pair_login_method_LG()

    if not token_response:
        logger.error('error getting access and refresh token')
        logger.error(token_response.text)
        sys.exit(1)

    json_response = token_response.json()
    access_token = json_response['data']['attributes']['token']

    needed = access_token.split(".")[1]
    try:
        somestring = base64.b64decode(needed).decode()
    except:
        somestring = base64.b64decode(needed + "==").decode()
    somedict = json.loads(somestring)
    expiry = somedict['exp']
    save_token(json_response, access_token, expiry)
    return access_token, expiry

def code_pair_login_method_LG():
   # First Init Access Token From Device
   init_access_token = get_init_access_token()
   # Get get_wbd_session_state values client Config
   x_wbd_session_state = get_wbd_session_state(init_access_token)
   # Getting Code
   code = get_device_code(init_access_token, x_wbd_session_state)
   # Following the flow and doing one more request
   session.post(**client_config.get_tokens_using_code(init_access_token, x_wbd_session_state))
   logger.info(f"Go to www.max.com/signin and input the {code}")
   data = ''
   print('Checking link ', end='')
   while data == '':
        for index, char in enumerate("." * 5):
            sys.stdout.write(char)
            sys.stdout.flush()
            # exchange access code for oauth token
            time.sleep(0.2)
        token_response = session.post(
            **client_config.get_tokens_using_code(init_access_token, x_wbd_session_state))
        logger.info(token_response.text)
        data = token_response.text
        index += 1  # lists are zero indexed, we need to increase by one for the accurate count
        # backtrack the written characters, overwrite them with space, backtrack again:
        sys.stdout.write("\b" * index + " " * index + "\b" * index)
        sys.stdout.flush()
   if json.loads(data)['data']['type'] == "token":
     print('\nSuccessfully linked!')
   return token_response

if __name__ == "__main__":
    max_login()
