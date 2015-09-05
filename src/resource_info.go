package src

import "encoding/json"

type ResourceInfoRequest struct {
	client      *Client
	httpRequest *httpRequest
}

func (r *ResourceInfoRequest) request() *httpRequest {
	return r.httpRequest
}

type ResourceRequestOptions struct {
	sortMode *SortMode
}

func (c *Client) NewResourceInfoRequest(path string, options ...ResourceRequestOptions) *ResourceInfoRequest {
	var parameters = make(map[string]interface{})
	parameters["path"] = path
	if len(options) > 0 && options[0].sortMode != nil {
		parameters["sort"] = options[0].sortMode.String()
	}
	return &ResourceInfoRequest{
		client:      c,
		httpRequest: createGetRequest(c, "/resources", parameters),
	}
}

func (req *ResourceInfoRequest) Exec() (*ResourceInfoResponse, error) {
	data, err := req.request().run(req.client)
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

	return &info, nil
}
