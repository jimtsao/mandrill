package mandrill

import (
	"encoding/json"
)

type IPs struct {
	m *Mandrill
}

type ipsResponse struct {
	IP        string `json:"ip"`
	CreatedAt string `json:"created_at"`
	Pool      string `json:"pool"`
	Domain    string `json:"domain"`
	CustomDNS struct {
		Enabled bool   `json:"enabled"`
		Valid   bool   `json:"valid"`
		Error   string `json:"error"`
	} `json:"custom_dns"`
	Warmup struct {
		WarmingUp bool   `json:"warming_up"`
		StartAt   string `json:"start_at"`
		EndAt     string `json:"end_at"`
	} `json:"warmup"`
}

// List dedicated IPs
func (i *IPs) List() ([]ipsResponse, error) {
	var ret []ipsResponse
	data := simpleRequest{i.m.APIKey}
	body, err := i.m.execute("/ips/list.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *IPs) Info(ip string) (ipsResponse, error) {
	var ret ipsResponse
	data := struct {
		APIKey string `json:"key"`
		IP     string `json:"ip"`
	}{i.m.APIKey, ip}
	body, err := i.m.execute("/ips/info.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// Provision requests an additional dedicated IP for your account. Accounts may
// have one outstanding request at any time, and provisioning requests are
// processed within 24 hours.
func (i *IPs) Provision(warmup bool, pool string) (ipsProvisionResponse, error) {
	var ret ipsProvisionResponse
	data := struct {
		APIKey string `json:"key"`
		Warmup bool   `json:"warmup,omitempty"`
		Pool   string `json:"pool,omitempty"`
	}{i.m.APIKey, warmup, pool}
	body, err := i.m.execute("/ips/provision.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type ipsProvisionResponse struct {
	RequestedAt string `json:"requested_at"`
}

// StartWarmup begins warmup process for a dedicated IP. During the warmup process, Mandrill
// will gradually increase the percentage of your mail that is sent over the warming-up IP,
// over a period of roughly 30 days. The rest of your mail will be sent over shared IPs or other
// dedicated IPs in the same pool.
func (i *IPs) StartWarmup(ip string) (ipsResponse, error) {
	var ret ipsResponse
	data := struct {
		APIKey string `json:"key"`
		IP     string `json:"ip"`
	}{i.m.APIKey, ip}
	body, err := i.m.execute("/ips/start-warmup.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *IPs) CancelWarmup(ip string) (ipsResponse, error) {
	var ret ipsResponse
	data := struct {
		APIKey string `json:"key"`
		IP     string `json:"ip"`
	}{i.m.APIKey, ip}
	body, err := i.m.execute("/ips/cancel-warmup.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// SetPool moves a dedicated IP to a different pool
func (i *IPs) SetPool(ip string, pool string, createPool bool) (ipsResponse, error) {
	var ret ipsResponse
	data := struct {
		APIKey     string `json:"key"`
		IP         string `json:"ip"`
		Pool       string `json:"pool"`
		CreatePool bool   `json:"create_pool,omitempty"`
	}{i.m.APIKey, ip, pool, createPool}
	body, err := i.m.execute("/ips/set-pool.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// Delete a dedicated IP. This is permanent and cannot be undone.
func (i *IPs) Delete(ip string) (ipsDeleteResponse, error) {
	var ret ipsDeleteResponse
	data := struct {
		APIKey string `json:"key"`
		IP     string `json:"ip"`
	}{i.m.APIKey, ip}
	body, err := i.m.execute("/ips/delete.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type ipsDeleteResponse struct {
	IP      string `json:"ip"`
	Deleted bool   `json:"deleted"`
}

// ListPools returns list of dedicated ip pools
func (i *IPs) ListPools() ([]ipsPoolsResponse, error) {
	var ret []ipsPoolsResponse
	data := simpleRequest{i.m.APIKey}
	body, err := i.m.execute("/ips/list-pools.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type ipsPoolsResponse struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Ips       []struct {
		CreatedAt string `json:"created_at"`
		CustomDNS struct {
			Enabled bool   `json:"enabled"`
			Error   string `json:"error"`
			Valid   bool   `json:"valid"`
		} `json:"custom_dns"`
		Domain string `json:"domain"`
		IP     string `json:"ip"`
		Pool   string `json:"pool"`
		Warmup struct {
			EndAt     string `json:"end_at"`
			StartAt   string `json:"start_at"`
			WarmingUp bool   `json:"warming_up"`
		} `json:"warmup"`
	} `json:"ips"`
}

func (i *IPs) PoolInfo(pool string) (ipsPoolsResponse, error) {
	var ret ipsPoolsResponse
	data := struct {
		APIKey string `json:"key"`
		Pool   string `json:"pool"`
	}{i.m.APIKey, pool}
	body, err := i.m.execute("/ips/pool-info.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *IPs) CreatePool(pool string) (ipsPoolsResponse, error) {
	var ret ipsPoolsResponse
	data := struct {
		APIKey string `json:"key"`
		Pool   string `json:"pool"`
	}{i.m.APIKey, pool}
	body, err := i.m.execute("/ips/create-pool.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (i *IPs) DeletePool(pool string) (ipsDeletePoolResponse, error) {
	var ret ipsDeletePoolResponse
	data := struct {
		APIKey string `json:"key"`
		Pool   string `json:"pool"`
	}{i.m.APIKey, pool}
	body, err := i.m.execute("/ips/delete-pool.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type ipsDeletePoolResponse struct {
	Pool    string `json:"pool"`
	Deleted bool   `json:"deleted"`
}

func (i *IPs) CheckCustomDNS(ip string, domain string) (ipsCheckDNSResponse, error) {
	var ret ipsCheckDNSResponse
	data := struct {
		APIKey string `json:"key"`
		IP     string `json:"ip"`
		Domain string `json:"domain"`
	}{i.m.APIKey, ip, domain}
	body, err := i.m.execute("/ips/check-custom-dns.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

type ipsCheckDNSResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error"`
}

func (i *IPs) SetCustomDNS(ip string, domain string) (ipsResponse, error) {
	var ret ipsResponse
	data := struct {
		APIKey string `json:"key"`
		IP     string `json:"ip"`
		Domain string `json:"domain"`
	}{i.m.APIKey, ip, domain}
	body, err := i.m.execute("/ips/check-custom-dns.json", data)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
