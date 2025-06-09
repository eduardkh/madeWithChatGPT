import requests

url = "https://run.hfd.co.il/RunCom.Server/Request.aspx?APPNAME=RUN&PRGNAME=ship_locate_blank&arguments=-AEP902050457,-AHfd003064592,-Ay,-Ajson"

payload = {}
headers = {
    'Host': 'run.hfd.co.il',
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:139.0) Gecko/20100101 Firefox/139.0',
    'Accept': 'application/json, text/javascript, */*; q=0.01',
    'Accept-Language': 'he,he-IL;q=0.8,en-US;q=0.5,en;q=0.3',
    'Accept-Encoding': 'gzip, deflate',
    'X-Requested-With': 'XMLHttpRequest',
    'Referer': 'https://run.hfd.co.il/EPOST_TRACK/',
    'Sec-Fetch-Dest': 'empty',
    'Sec-Fetch-Mode': 'cors',
    'Sec-Fetch-Site': 'same-origin',
    'Priority': 'u=0',
    'Te': 'trailers'
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
