from . import *
from . import vars

def get_proxy_delays(name):
    if not vars.api_url:
        return [{"delays": []}]
    response = requests.get(f"{vars.api_url}/proxies/{name}")
    if not 200 <= response.status_code <= 299:
        return [{"delays": []}]
    response = response.json()
    if "message" in response:
        return [{"delays": []}]
    
    proxy_delays = [history["delay"] for history in response["history"]]
    
    return [{"name": name, "delays": proxy_delays}]

def get_proxies_delays(name):
    if not vars.api_url:
        return []
    response = requests.get(f"{vars.api_url}/proxies/{name}")
    if not 200 <= response.status_code <= 299:
        return []
    response = response.json()
    if "message" in response:
        return []
    
    names = response["all"]
    for name in names:
        for exclude in vars.excludes:
            if exclude in name:
                try:
                    names.remove(name)
                except:
                    pass
                
    proxies_delays = []
    for name in names:
        proxy_delays = get_proxy_delays(name)
        for i, delay in enumerate(proxy_delays[0]["delays"]):
            if delay == 0:
                proxy_delays[0]["delays"][i] = 3000
        proxies_delays += proxy_delays
    
    return proxies_delays