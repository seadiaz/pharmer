package linodego

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// LinodeConfigService is an interface generated for "github.com/appscode/linodego".LinodeConfigService.
type LinodeConfigInterface interface {
	Create(int, int, string, map[string]string) (*LinodeConfigResponse, error)
	Delete(int, int) (*LinodeConfigResponse, error)
	List(int, int) (*LinodeConfigListResponse, error)
	Update(int, int, int, map[string]string) (*LinodeConfigResponse, error)
}

// Linode Config Service
type LinodeConfigService struct {
	client *Client
}

var _ LinodeConfigInterface = &LinodeConfigService{}

// Response for linode.config.list API
type LinodeConfigListResponse struct {
	Response
	LinodeConfigs []LinodeConfig
}

// Response for general config APIs
type LinodeConfigResponse struct {
	Response
	LinodeConfigId LinodeConfigId
}

// Get Config List. If configId is greater than 0, limit results to given config.
func (t *LinodeConfigService) List(linodeId int, configId int) (*LinodeConfigListResponse, error) {
	u := &url.Values{}
	if configId > 0 {
		u.Add("ConfigID", strconv.Itoa(configId))
	}
	u.Add("LinodeID", strconv.Itoa(linodeId))
	v := LinodeConfigListResponse{}
	if err := t.client.do("linode.config.list", u, &v.Response); err != nil {
		return nil, err
	}

	v.LinodeConfigs = make([]LinodeConfig, 5)
	if err := json.Unmarshal(v.RawData, &v.LinodeConfigs); err != nil {
		return nil, err
	}
	return &v, nil
}

// Create Config
func (t *LinodeConfigService) Create(linodeId int, kernelId int, label string, args map[string]string) (*LinodeConfigResponse, error) {
	u := &url.Values{}
	u.Add("LinodeID", strconv.Itoa(linodeId))
	u.Add("KernelID", strconv.Itoa(kernelId))
	u.Add("Label", label)
	// add optional parameters
	processOptionalArgs(args, u)
	v := LinodeConfigResponse{}
	if err := t.client.do("linode.config.create", u, &v.Response); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(v.RawData, &v.LinodeConfigId); err != nil {
		return nil, err
	}
	return &v, nil
}

// Update Config. See https://www.linode.com/api/linode/linode.config.update for allowed arguments.
func (t *LinodeConfigService) Update(configId int, linodeId int, kernelId int, args map[string]string) (*LinodeConfigResponse, error) {
	u := &url.Values{}
	u.Add("ConfigID", strconv.Itoa(configId))
	if linodeId > 0 {
		u.Add("LinodeID", strconv.Itoa(linodeId))
	}
	if kernelId > 0 {
		u.Add("KernelID", strconv.Itoa(kernelId))
	}

	// add optional parameters
	processOptionalArgs(args, u)
	v := LinodeConfigResponse{}
	if err := t.client.do("linode.config.update", u, &v.Response); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(v.RawData, &v.LinodeConfigId); err != nil {
		return nil, err
	}
	return &v, nil
}

// Delete Config
func (t *LinodeConfigService) Delete(linodeId int, configId int) (*LinodeConfigResponse, error) {
	u := &url.Values{}
	u.Add("LinodeID", strconv.Itoa(linodeId))
	u.Add("ConfigID", strconv.Itoa(configId))
	v := LinodeConfigResponse{}
	if err := t.client.do("linode.config.delete", u, &v.Response); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(v.RawData, &v.LinodeConfigId); err != nil {
		return nil, err
	}
	return &v, nil
}

func processOptionalArgs(args map[string]string, parameters *url.Values) {
	for k, v := range args {
		if k == "helper_xen" {
			parameters.Add("helper_distro", v)
		}
		parameters.Add(k, v)
	}
}
