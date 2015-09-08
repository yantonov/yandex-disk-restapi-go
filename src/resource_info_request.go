package src

import "encoding/json"

type ResourceInfoRequest struct {
	client      *Client
	httpRequest *httpRequest
}

func (r *ResourceInfoRequest) Request() *httpRequest {
	return r.httpRequest
}

type ResourceInfoRequestOptions struct {
	Sort_mode    *SortMode
	Limit        *uint32
	Offset       *uint32
	Fields       []string
	Preview_size *PreviewSize
	Preview_crop *bool
}

func (c *Client) NewResourceInfoRequest(path string, options ...ResourceInfoRequestOptions) *ResourceInfoRequest {
	var parameters = make(map[string]interface{})
	parameters["path"] = path
	if len(options) > 0 {
		opt := options[0]
		if opt.Sort_mode != nil {
			parameters["sort"] = opt.Sort_mode.String()
		}
		if opt.Limit != nil {
			parameters["limit"] = opt.Limit
		}
		if opt.Offset != nil {
			parameters["offset"] = opt.Offset
		}
		if opt.Preview_size != nil {
			parameters["preview_size"] = opt.Preview_size.String()
		}
		if opt.Preview_crop != nil {
			parameters["preview_crop"] = opt.Preview_crop
		}
	}
	return &ResourceInfoRequest{
		client:      c,
		httpRequest: createGetRequest(c, "/resources", parameters),
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

	return &info, nil
}
