package yandexdiskapi

import "encoding/json"

type ResourceInfoRequest struct {
	client      *Client
	httpRequest *httpRequest
}

func (r *ResourceInfoRequest) Request() *httpRequest {
	return r.httpRequest
}

func (c *Client) NewResourceInfoRequest(path string, options ...ResourceInfoRequestOptions) *ResourceInfoRequest {
	return &ResourceInfoRequest{
		client:      c,
		httpRequest: createResourceInfoRequest(c, "/resources", path, options...),
	}
}

func (req *ResourceInfoRequest) Exec() (*ResourceInfoResponse, error) {
	data, err := req.Request().run(req.client)
	if err != nil {
		return nil, err
	}

	var info ResourceInfoResponse
	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}
	if info.Custom_properties == nil {
		info.Custom_properties = make(map[string]interface{})
	}
	if info.Embedded != nil {
		if cap(info.Embedded.Items) == 0 {
			info.Embedded.Items = []ResourceInfoResponse{}
		}
	}
	return &info, nil
}
