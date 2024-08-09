from . import vars
import requests

def set_api_url(host="127.0.0.1", port="9090", https=False):
    if not https:
        vars.api_url = f"http://{host}:{port}"
    else:
        vars.api_url = f"https://{host}:{port}"
    
    try:
        response = requests.get(f"{vars.api_url}/version")
        if 200 <= response.status_code <= 299:
            return response.json()["version"]
    except:
        pass
        
    vars.api_url = ""
    
def set_excludes(excludes_=[]):
    vars.excludes = excludes_