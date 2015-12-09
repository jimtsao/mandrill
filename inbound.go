package mandrill

import (
	"encoding/json"
)

type Inbound struct {
	m *Mandrill
}

type inboundDomainResponse struct {
	Domain    string `json:"domain"`
	CreatedAt string `json:"created_at"`
	ValidMX   bool   `json:"valid_mx"`
}

func (i *Inbound) Domains() ([]inboundDomainResponse, error) {
	var ret []inboundDomainResponse
	data := simpleRequest{i.m.APIKey}
	body, err := i.m.execute("/inbound/domains.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *Inbound) AddDomain(domain string) (inboundDomainResponse, error) {
	var ret inboundDomainResponse
	data := struct {
		APIKey string `json:"key"`
		Domain string `json:"domain"`
	}{i.m.APIKey, domain}
	body, err := i.m.execute("/inbound/add-domain.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *Inbound) CheckDomain(domain string) (inboundDomainResponse, error) {
	var ret inboundDomainResponse
	data := struct {
		APIKey string `json:"key"`
		Domain string `json:"domain"`
	}{i.m.APIKey, domain}
	body, err := i.m.execute("/inbound/check-domain.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *Inbound) DeleteDomain(domain string) (inboundDomainResponse, error) {
	var ret inboundDomainResponse
	data := struct {
		APIKey string `json:"key"`
		Domain string `json:"domain"`
	}{i.m.APIKey, domain}
	body, err := i.m.execute("/inbound/delete-domain.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type inboundRouteResponse struct {
	Id      string `json:"id"`
	Pattern string `json:"pattern"`
	URL     string `json:"url"`
}

func (i *Inbound) Routes(domain string) ([]inboundRouteResponse, error) {
	var ret []inboundRouteResponse
	data := struct {
		APIKey string `json:"key"`
		Domain string `json:"domain"`
	}{i.m.APIKey, domain}
	body, err := i.m.execute("/inbound/routes.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *Inbound) AddRoute(domain string, pattern string, url string) (inboundRouteResponse, error) {
	var ret inboundRouteResponse
	data := struct {
		APIKey  string `json:"key"`
		Domain  string `json:"domain"`
		Pattern string `json:"pattern"`
		URL     string `json:"url"`
	}{i.m.APIKey, domain, pattern, url}
	body, err := i.m.execute("/inbound/add-route.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *Inbound) UpdateRoute(id string, pattern string, url string) (inboundRouteResponse, error) {
	var ret inboundRouteResponse
	data := struct {
		APIKey  string `json:"key"`
		Id      string `json:"id"`
		Pattern string `json:"pattern,omitempty"`
		URL     string `json:"url,omitempty"`
	}{i.m.APIKey, id, pattern, url}
	body, err := i.m.execute("/inbound/update-route.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *Inbound) DeleteRoute(id string) (inboundRouteResponse, error) {
	var ret inboundRouteResponse
	data := struct {
		APIKey string `json:"key"`
		Id     string `json:"id"`
	}{i.m.APIKey, id}
	body, err := i.m.execute("/inbound/delete-route.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// Take a raw MIME document destined for a domain with inbound domains set up
// and send it to the inbound hook exactly as if it had been sent over SMTP
func (i *Inbound) SendRaw(rawMessage string, to []string, from string, helo string, clientAddr string) ([]inboundSendRawResponse, error) {
	var ret []inboundSendRawResponse
	data := struct {
		APIKey        string   `json:"key"`
		RawMessage    string   `json:"raw_message"`
		To            []string `json:"to,omitempty"`
		MailFrom      string   `json:"mail_from,omitempty"`
		Helo          string   `json:"helo,omitempty"`
		ClientAddress string   `json:"client_address,omitempty"`
	}{i.m.APIKey, rawMessage, to, from, helo, clientAddr}
	body, err := i.m.execute("/inbound/send-raw.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type inboundSendRawResponse struct {
	Email   string `json:"email"`
	Pattern string `json:"pattern"`
	URL     string `json:"url"`
}
