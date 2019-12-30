package src

import (
	"encoding/json"
	"strings"
)

type FlatFileListRequest struct {
	client      *Client
	httpRequest *httpRequest
}

type FlatFileListRequestOptions struct {
	Media_type   []MediaType
	Limit        *uint32
	Offset       *uint32
	Fields       []string
	Preview_size *PreviewSize
	Preview_crop *bool
}

func (r *FlatFileListRequest) Request() *httpRequest {
	return r.httpRequest
}

func (c *Client) NewFlatFileListRequest(options ...FlatFileListRequestOptions) *FlatFileListRequest {
	var parameters = make(map[string]interface{})
	if len(options) > 0 {
		opt := options[0]
		if opt.Limit != nil {
			parameters["limit"] = *opt.Limit
		}
		if opt.Offset != nil {
			parameters["offset"] = *opt.Offset
		}
		if opt.Fields != nil {
			parameters["fields"] = strings.Join(opt.Fields, ",")
		}
		if opt.Preview_size != nil {
			parameters["preview_size"] = opt.Preview_size.String()
		}
		if opt.Preview_crop != nil {
			parameters["preview_crop"] = *opt.Preview_crop
		}
		if opt.Media_type != nil {
			var str_media_types = make([]string, len(opt.Media_type))
			for i, t := range opt.Media_type {
				str_media_types[i] = t.String()
			}
			parameters["media_type"] = strings.Join(str_media_types, ",")
		}
	}
	return &FlatFileListRequest{
		client:      c,
		httpRequest: createGetRequest(c, "/resources/files", parameters),
	}
}

func (req *FlatFileListRequest) Exec() (*FilesResourceListResponse, error) {
	data, err := req.Request().run(req.client)
	if err != nil {
		return nil, err
	}
	var info FilesResourceListResponse
	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}
	if cap(info.Items) == 0 {
		info.Items = []ResourceInfoResponse{}
	}
	return &info, nil
}
