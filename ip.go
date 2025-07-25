package utils

import (
	"encoding/json"
	"net"
	"os"
)

func GetLocalIP() string {
	var localIP string

	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			localIP = ipv4.String()
			break
		}
	}

	return localIP
}

//{
//	"traits": {
//		"ip_address": "41.203.87.90",
//		"network": "41.203.64.0/18"
//	},
//	"continent": {
//		"names": {
//			"de": "Afrika",
//			"en": "Africa",
//			"es": "África",
//			"fr": "Afrique",
//			"ja": "アフリカ",
//			"pt-BR": "África",
//			"ru": "Африка",
//			"zh-CN": "非洲"
//		},
//		"code": "AF",
//		"geoname_id": 6255146
//	},
//	"country": {
//		"names": {
//			"de": "Nigeria",
//			"en": "Nigeria",
//			"es": "Nigeria",
//			"fr": "Nigeria",
//			"ja": "ナイジェリア連邦共和国",
//			"pt-BR": "Nigéria",
//			"ru": "Нигерия",
//			"zh-CN": "尼日利亚"
//		},
//		"iso_code": "NG",
//		"geoname_id": 2328926
//	},
//	"registered_country": {
//		"names": {
//			"de": "Nigeria",
//			"en": "Nigeria",
//			"es": "Nigeria",
//			"fr": "Nigeria",
//			"ja": "ナイジェリア連邦共和国",
//			"pt-BR": "Nigéria",
//			"ru": "Нигерия",
//			"zh-CN": "尼日利亚"
//		},
//		"iso_code": "NG",
//		"geoname_id": 2328926
//	}
//}

type IPAddressResponse struct {
	Traits struct {
		IpAddress string `json:"ip_address"`
		Network   string `json:"network"`
	} `json:"traits"`
	Continent struct {
		Names struct {
			De   string `json:"de"`
			En   string `json:"en"`
			Es   string `json:"es"`
			Fr   string `json:"fr"`
			Ja   string `json:"ja"`
			PtBR string `json:"pt-BR"`
			Ru   string `json:"ru"`
			ZhCN string `json:"zh-CN"`
		} `json:"names"`
		Code      string `json:"code"`
		GeonameId int    `json:"geoname_id"`
	} `json:"continent"`
	Country struct {
		Names struct {
			De   string `json:"de"`
			En   string `json:"en"`
			Es   string `json:"es"`
			Fr   string `json:"fr"`
			Ja   string `json:"ja"`
			PtBR string `json:"pt-BR"`
			Ru   string `json:"ru"`
			ZhCN string `json:"zh-CN"`
		} `json:"names"`
		IsoCode   string `json:"iso_code"`
		GeonameId int    `json:"geoname_id"`
	} `json:"country"`
	RegisteredCountry struct {
		Names struct {
			De   string `json:"de"`
			En   string `json:"en"`
			Es   string `json:"es"`
			Fr   string `json:"fr"`
			Ja   string `json:"ja"`
			PtBR string `json:"pt-BR"`
			Ru   string `json:"ru"`
			ZhCN string `json:"zh-CN"`
		} `json:"names"`
		IsoCode   string `json:"iso_code"`
		GeonameId int    `json:"geoname_id"`
	} `json:"registered_country"`
}

type GetIPAddressResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		GeoData string `json:"geo_data"`
		IPAddressResponse
	}
}

func GetIPAddress(ip string) (GetIPAddressResponse, error) {
	uri := "http://" + "geoip:8080" + "/api/geo/geo-country"

	param := make(map[string]string)
	param["ip"] = ip

	_, response, err := GetRequest(RequestMethodGet, uri, param, nil)
	if err != nil {
		return GetIPAddressResponse{}, err
	}

	var ipAddressResponse GetIPAddressResponse
	json.Unmarshal([]byte(response), &ipAddressResponse)
	if ipAddressResponse.Code == 200 {
		json.Unmarshal([]byte(ipAddressResponse.Data.GeoData), &ipAddressResponse.Data.IPAddressResponse)
	}

	return ipAddressResponse, nil
}
