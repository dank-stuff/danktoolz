package actions

import (
	"net"
	"net/http"
	"strconv"
	"strings"
)

type IpInfo struct {
	RawRemoteAddr string   `json:"raw_remote_addr,omitempty"`
	ClientIP      string   `json:"client_ip,omitempty"`
	ForwardedFor  []string `json:"forwarded_for,omitempty"`
	RealIP        string   `json:"real_ip,omitempty"`
	TrueClientIP  string   `json:"true_client_ip,omitempty"`
	Authorization string   `json:"authorization,omitempty"`
	Referer       string   `json:"referer,omitempty"`
	Origin        string   `json:"origin,omitempty"`

	UA              string `json:"user_agent,omitempty"`
	BrandList       string `json:"ua_brand_list,omitempty"`
	Platform        string `json:"ua_platform,omitempty"`
	IsMobile        bool   `json:"ua_is_mobile"`
	Model           string `json:"ua_model,omitempty"`
	Architecture    string `json:"ua_architecture,omitempty"`
	Bitness         string `json:"ua_bitness,omitempty"`
	FullVersionList string `json:"ua_full_versions,omitempty"`
	FormFactors     string `json:"ua_form_factors,omitempty"`
	DeviceMemoryGB  string `json:"device_memory_gb,omitempty"`
	ViewportWidth   int    `json:"viewport_width_px,omitempty"`

	AcceptLanguage string `json:"accept_language,omitempty"`
	TimezoneOffset int    `json:"timezone_offset,omitempty"`

	CountryCode   string  `json:"country_code,omitempty"`
	ContinentCode string  `json:"continent_code,omitempty"`
	City          string  `json:"city,omitempty"`
	Region        string  `json:"region,omitempty"`
	PostalCode    string  `json:"postal_code,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	CloudTimezone string  `json:"cloud_timezone,omitempty"`

	SaveData          bool    `json:"save_data"`
	EffectiveConnType string  `json:"effective_conn_type,omitempty"`
	RoundTripTimeMs   int     `json:"rtt_ms,omitempty"`
	DownlinkMbps      float64 `json:"downlink_mbps,omitempty"`

	GlobalPrivacyControl bool `json:"global_privacy_control"`
	DoNotTrack           bool `json:"do_not_track"`
}

func ExtractClientMetadata(r *http.Request) IpInfo {
	var meta IpInfo

	meta.RawRemoteAddr = r.RemoteAddr
	meta.RealIP = r.Header.Get("X-Real-IP")
	meta.TrueClientIP = r.Header.Get("True-Client-IP")
	if meta.TrueClientIP == "" {
		meta.TrueClientIP = r.Header.Get("CF-Connecting-IP")
	}

	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		for part := range strings.SplitSeq(xff, ",") {
			meta.ForwardedFor = append(meta.ForwardedFor, strings.TrimSpace(part))
		}
	}
	meta.ClientIP = determineBestClientIP(r, meta)

	meta.Authorization = r.Header.Get("Authorization")
	meta.Referer = r.Header.Get("Referer")
	meta.Origin = r.Header.Get("Origin")
	meta.AcceptLanguage = r.Header.Get("Accept-Language")

	meta.UA = r.Header.Get("User-Agent")
	meta.BrandList = r.Header.Get("Sec-CH-UA")
	meta.Platform = r.Header.Get("Sec-CH-UA-Platform")
	meta.IsMobile = r.Header.Get("Sec-CH-UA-Mobile") == "?1"
	meta.Model = r.Header.Get("Sec-CH-UA-Model")
	meta.Architecture = r.Header.Get("Sec-CH-UA-Arch")
	meta.Bitness = r.Header.Get("Sec-CH-UA-Bitness")
	meta.FullVersionList = r.Header.Get("Sec-CH-UA-Full-Version-List")
	meta.FormFactors = r.Header.Get("Sec-CH-UA-Form-Factors")

	meta.DeviceMemoryGB = r.Header.Get("Sec-CH-Device-Memory")
	meta.ViewportWidth, _ = strconv.Atoi(r.Header.Get("Sec-CH-Viewport-Width"))
	meta.TimezoneOffset, _ = strconv.Atoi(r.Header.Get("Sec-CH-UA-Timezone-Offset"))

	meta.CountryCode = r.Header.Get("CF-IPCountry")
	meta.ContinentCode = r.Header.Get("CF-IPContinent")

	if awsCountry := r.Header.Get("CloudFront-Viewer-Country"); awsCountry != "" {
		meta.CountryCode = awsCountry
	}
	meta.City = r.Header.Get("CloudFront-Viewer-City")
	meta.Region = r.Header.Get("CloudFront-Viewer-Country-Region")
	meta.PostalCode = r.Header.Get("CloudFront-Viewer-Postal-Code")
	meta.CloudTimezone = r.Header.Get("CloudFront-Viewer-Timezone")
	meta.Latitude, _ = strconv.ParseFloat(r.Header.Get("CloudFront-Viewer-Latitude"), 64)
	meta.Longitude, _ = strconv.ParseFloat(r.Header.Get("CloudFront-Viewer-Longitude"), 64)

	meta.SaveData = r.Header.Get("Save-Data") == "on"
	meta.EffectiveConnType = r.Header.Get("ECT")
	meta.RoundTripTimeMs, _ = strconv.Atoi(r.Header.Get("RTT"))
	meta.DownlinkMbps, _ = strconv.ParseFloat(r.Header.Get("Downlink"), 64)

	meta.GlobalPrivacyControl = r.Header.Get("Sec-GPC") == "1"
	meta.DoNotTrack = r.Header.Get("DNT") == "1"

	return meta
}

func determineBestClientIP(r *http.Request, meta IpInfo) string {
	if meta.TrueClientIP != "" {
		return meta.TrueClientIP
	}

	if len(meta.ForwardedFor) > 0 && meta.ForwardedFor[0] != "" {
		return meta.ForwardedFor[0]
	}

	if meta.RealIP != "" {
		return meta.RealIP
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
