package crtsh

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const crtUrl = "https://crt.sh/"

var UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.1"

type Entry struct {
	IssuerCaId     int64  `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	Id             int64  `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
	Serial         string `json:"serial_number"`
}
type Response []Entry

func SearchJSON(domain string, expired bool) ([]byte, error) {
	flags := url.Values{}
	flags.Set("q", domain)
	flags.Set("output", "json")
	if !expired {
		flags.Set("exclude", "expired")
	}

	req, err := http.NewRequest("GET", crtUrl, nil)
	if err != nil {
		return []byte{}, err
	}

	req.URL.RawQuery = flags.Encode()
	req.Header = http.Header{}
	req.Header.Add("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	var rawJson []byte
	if resp.StatusCode == 200 {
		rawJson, err = io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}
		resp.Body.Close()
	} else {
		return []byte{}, fmt.Errorf("received %d status code from '%s'", resp.StatusCode, crtUrl)
	}

	return rawJson, nil
}

func Search(domain string, expired bool) (Response, error) {
	js, err := SearchJSON(domain, expired)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(js, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
