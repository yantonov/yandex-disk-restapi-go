package src

import (
	"encoding/json"
	"strings"
)

type LastUploadedResourceListRequest struct {
	client      *Client
	httpRequest *httpRequest
}

type LastUploadedResourceListRequestOptions struct {
	Media_type   []MediaType
	Limit        *uint32
	Fields       []string
	Preview_size *PreviewSize
	Preview_crop *bool
}

func (r *LastUploadedResourceListRequest) Request() *httpRequest {
	return r.httpRequest
}

func (c *Client) NewLastUploadedResourceListRequest(options ...LastUploadedResourceListRequestOptions) *LastUploadedResourceListRequest {
	var parameters = make(map[string]interface{})
	if len(options) > 0 {
		opt := options[0]
		if opt.Limit != nil {
			parameters["limit"] = opt.Limit
		}
		if opt.Fields != nil {
			parameters["fields"] = strings.Join(opt.Fields, ",")
		}
		if opt.Preview_size != nil {
			parameters["preview_size"] = opt.Preview_size.String()
		}
		if opt.Preview_crop != nil {
			parameters["preview_crop"] = opt.Preview_crop
		}
		if opt.Media_type != nil {
			var str_media_types = make([]string, len(opt.Media_type))
			for i, t := range opt.Media_type {
				str_media_types[i] = t.String()
			}
			parameters["media_type"] = strings.Join(str_media_types, ",")
		}
	}
	return &LastUploadedResourceListRequest{
		client:      c,
		httpRequest: createGetRequest(c, "/resources/last-uploaded", parameters),
	}
}

func (req *LastUploadedResourceListRequest) Exec() (*LastUploadedResourceListResponse, error) {
	data, err := req.Request().run(req.client)
	if err != nil {
		return nil, err
	}
	var info LastUploadedResourceListResponse
	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}
	if cap(info.Items) == 0 {
		info.Items = []ResourceInfoResponse{}
	}
	return &info, nil
}
