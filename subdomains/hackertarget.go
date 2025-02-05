package subdomains

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"oott/helper"
	"strings"
)

type Hackertarget struct {
	// any necessary fields specific
}

func (s *Hackertarget) ScanSubdomains(domain string) ([]SubDomainDetails, error) {
	helper.InfoPrintln("[+] Scanning subdomains on Hackertarget:", domain)

	// Make the API request
	url := fmt.Sprintf("https://api.hackertarget.com/hostsearch/?q=%s", domain)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if !strings.Contains(string(body), domain) {
		helper.ErrorPrintln(string(body))
		return nil, nil
	}

	var subdomains []SubDomainDetails

	subdomainsString := strings.Split(string(body), "\n")
	for _, subdomainString := range subdomainsString {
		domainIP := strings.Split(subdomainString, ",")
		if len(domainIP) > 1 {
			domain := strings.TrimSpace(domainIP[0])
			ip := strings.TrimSpace(domainIP[1])

			subdomain := SubDomainDetails{
				DomainName: domain,
				Address:    ip,
				Source:     "Hackertarget",
			}

			subdomains = append(subdomains, subdomain)
		}
	}

	return subdomains, nil
}
