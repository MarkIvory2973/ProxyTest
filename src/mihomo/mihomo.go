package mihomo

import (
	"ProxyTest/httpx"
	"fmt"
)

func toStrSlice(interfaces []interface{}) []string {
	var slice []string
	for _, value := range interfaces {
		if str, ok := value.(string); ok {
			slice = append(slice, str)
		} else {
			return []string{}
		}
	}

	return slice
}

type MiHoMo struct {
	TLS  bool
	Host string
	Port int
}

func (mihomo MiHoMo) GetProxiesName(groupName string) ([]string, error) {
	var url string
	if mihomo.TLS {
		url = fmt.Sprintf("https://%s:%d/proxies", mihomo.Host, mihomo.Port)
	} else {
		url = fmt.Sprintf("http://%s:%d/proxies", mihomo.Host, mihomo.Port)
	}

	response, err := httpx.HTTPGet(url)
	if err != nil {
		return []string{}, err
	}
	jsonContent, err := response.ToJson()
	if err != nil {
		return []string{}, err
	}

	if all, ok := jsonContent["proxies"].(map[string]interface{}); ok {
		if group, ok := all[groupName].(map[string]interface{}); ok {
			if proxiesName, ok := group["all"].([]interface{}); ok {
				return toStrSlice(proxiesName), nil
			}
		}
	}

	return []string{}, fmt.Errorf("GetProxiesName: can't get proxies(%s)", groupName)
}

func (mihomo MiHoMo) GetProxyDelays(proxyName string) ([]float64, error) {
	var url string
	if mihomo.TLS {
		url = fmt.Sprintf("https://%s:%d/proxies/%s", mihomo.Host, mihomo.Port, proxyName)
	} else {
		url = fmt.Sprintf("http://%s:%d/proxies/%s", mihomo.Host, mihomo.Port, proxyName)
	}

	response, err := httpx.HTTPGet(url)
	if err != nil {
		return []float64{}, err
	}
	jsonContent, err := response.ToJson()
	if err != nil {
		return []float64{}, err
	}

	var proxyDelays []float64
	if proxyHistories, ok := jsonContent["history"].([]interface{}); ok {
		for _, proxyHistory := range proxyHistories {
			if proxyHistoryMap, ok := proxyHistory.(map[string]interface{}); ok {
				if proxyDelay, ok := proxyHistoryMap["delay"].(float64); ok {
					if proxyDelay == 0 {
						proxyDelay = 3000
					}

					proxyDelays = append(proxyDelays, proxyDelay)
				}
			}
		}

		return proxyDelays, nil
	}

	return []float64{}, fmt.Errorf("GetProxyDelays: Can't get delay for proxy(%s)", proxyName)
}
