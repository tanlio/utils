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

func GetIPAddress(ip string) GetIPAddressResponse {
	uri := "http://" + "geoip:8080" + "/api/geo/geo-country"

	param := make(map[string]string)
	param["ip"] = ip

	_, response, err := GetRequest(RequestMethodGet, uri, param, nil)
	if err != nil {
		return GetIPAddressResponse{}
	}

	var ipAddressResponse GetIPAddressResponse
	json.Unmarshal([]byte(response), &ipAddressResponse)
	if ipAddressResponse.Code == 200 {
		json.Unmarshal([]byte(ipAddressResponse.Data.GeoData), &ipAddressResponse.Data.IPAddressResponse)
	}

	return ipAddressResponse
}

//{
//	"traits": {
//		"ip_address": "113.248.171.111",
//		"network": "113.248.160.0/19"
//	},
//	"continent": {
//		"names": {
//			"de": "Asien",
//			"en": "Asia",
//			"es": "Asia",
//			"fr": "Asie",
//			"ja": "アジア",
//			"pt-BR": "Ásia",
//			"ru": "Азия",
//			"zh-CN": "亚洲"
//		},
//		"code": "AS",
//		"geoname_id": 6255147
//	},
//	"city": {
//		"names": {
//			"de": "Chongqing",
//			"en": "Chongqing",
//			"es": "Chongqing",
//			"fr": "Chongqing",
//			"ja": "重慶市",
//			"pt-BR": "Chongqing",
//			"ru": "Чунцин",
//			"zh-CN": "重庆市"
//		},
//		"geoname_id": 1814906
//	},
//	"subdivisions": [{
//		"names": {
//			"en": "Chongqing",
//			"fr": "Municipalité de Chongqing",
//			"zh-CN": "重庆"
//		},
//		"iso_code": "CQ",
//		"geoname_id": 1814905
//	}],
//	"country": {
//		"names": {
//			"de": "China",
//			"en": "China",
//			"es": "China",
//			"fr": "Chine",
//			"ja": "中国",
//			"pt-BR": "China",
//			"ru": "Китай",
//			"zh-CN": "中国"
//		},
//		"iso_code": "CN",
//		"geoname_id": 1814991
//	},
//	"registered_country": {
//		"names": {
//			"de": "China",
//			"en": "China",
//			"es": "China",
//			"fr": "Chine",
//			"ja": "中国",
//			"pt-BR": "China",
//			"ru": "Китай",
//			"zh-CN": "中国"
//		},
//		"iso_code": "CN",
//		"geoname_id": 1814991
//	},
//	"location": {
//		"latitude": 29.5689,
//		"longitude": 106.5577,
//		"time_zone": "Asia/Shanghai",
//		"accuracy_radius": 1000
//	}
//}

type IPCityResponse struct {
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
	City struct {
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
		GeonameId int `json:"geoname_id"`
	} `json:"city"`
	Subdivisions []struct {
		Names struct {
			En   string `json:"en"`
			Fr   string `json:"fr"`
			ZhCN string `json:"zh-CN"`
		} `json:"names"`
		IsoCode   string `json:"iso_code"`
		GeonameId int    `json:"geoname_id"`
	} `json:"subdivisions"`
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
	Location struct {
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		TimeZone       string  `json:"time_zone"`
		AccuracyRadius int     `json:"accuracy_radius"`
	} `json:"location"`
}

type GetIPCityResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		GeoData string `json:"geo_data"`
		IPCityResponse
	}
}

func GetIPCity(ip string) GetIPCityResponse {
	uri := "http://" + "geoip:8080" + "/api/geo/geo-city"

	param := make(map[string]string)
	param["ip"] = ip

	_, response, err := GetRequest(RequestMethodGet, uri, param, nil)
	if err != nil {
		return GetIPCityResponse{}
	}

	var ipAddressResponse GetIPCityResponse
	json.Unmarshal([]byte(response), &ipAddressResponse)
	if ipAddressResponse.Code == 200 {
		json.Unmarshal([]byte(ipAddressResponse.Data.GeoData), &ipAddressResponse.Data.IPCityResponse)
	}

	return ipAddressResponse
}
