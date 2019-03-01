package src

import (
	"github.com/logrusorgru/aurora"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Result struct {
	status	aurora.Value
	entry	Fingerprint
}

func checkSubdomain(subdomain string, settings Settings) Result {

	if isValidUrl(subdomain) == false{
		if settings.Https {
			subdomain = "https://" + subdomain
		} else {
			subdomain = "http://" + subdomain
		}
	}

	resp, err := http.Get(subdomain)
	if err != nil {
		return Result {aurora.Red("HTTP ERROR"), Fingerprint{}}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result {aurora.Red("RESPONSE ERROR"), Fingerprint{}}
	}

	return matchResponse(string(body))
}


func matchResponse(body string) Result {
	fingerprints := Fingerprints()

	for _, fingerprint := range fingerprints {
		if strings.Contains(body, fingerprint.fingerprint) {
			return Result {aurora.Green("VULNERABLE"), fingerprint}

		}
	}

	return Result {aurora.Red("NOT VULNERABLE"), Fingerprint{}}

}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}