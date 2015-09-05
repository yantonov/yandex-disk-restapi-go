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
	sort_mode    *SortMode
	limit        *uint32
	offset       *uint32
	fields       []string
	preview_size *PreviewSize
	preview_crop *bool
}

func (c *Client) NewResourceInfoRequest(path string, options ...ResourceRequestOptions) *ResourceInfoRequest {
	var parameters = make(map[string]interface{})
	parameters["path"] = path
	if len(options) > 0 {
		opt := options[0]
		if opt.sort_mode != nil {
			parameters["sort"] = opt.sort_mode.String()
		}
		if opt.limit != nil {
			parameters["limit"] = opt.limit
		}
		if opt.offset != nil {
			parameters["offset"] = opt.offset
		}
		if opt.preview_size != nil {
			parameters["preview_size"] = opt.preview_size.String()
		}
		if opt.preview_crop != nil {
			parameters["preview_crop"] = opt.preview_crop
		}
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
