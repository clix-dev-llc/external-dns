package providers

import (
	"fmt"
)

const (
	//FIXME - set the root domain name based on the env vars
	RootDomainName = "rancher-test.com."
)

type Provider interface {
	AddRecord(record DnsRecord) error
	RemoveRecord(record DnsRecord) error
	UpdateRecord(record DnsRecord) error
	GetRecords() ([]DnsRecord, error)
	GetName() string
}

type DnsRecord struct {
	DomainName string
	Records    []string
	Type       string
	TTL        int
}

var (
	providers map[string]Provider
)

func GetProvider(name string) Provider {
	if provider, ok := providers[name]; ok {
		return provider
	}
	return providers["route53"]
}

func RegisterProvider(name string, provider Provider) error {
	if providers == nil {
		providers = make(map[string]Provider)
	}
	if _, exists := providers[name]; exists {
		return fmt.Errorf("provider already registered")
	}
	providers[name] = provider
	return nil
}