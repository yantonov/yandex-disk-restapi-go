package src

import "encoding/json"

type TrashResourceInfoRequest struct {
	client      *Client
	httpRequest *httpRequest
}

func (r *TrashResourceInfoRequest) Request() *httpRequest {
	return r.httpRequest
}

func (c *Client) NewTrashResourceInfoRequest(path string, options ...ResourceInfoRequestOptions) *TrashResourceInfoRequest {
	return &TrashResourceInfoRequest{
		client:      c,
		httpRequest: createResourceInfoRequest(c, "/trash/resources", path, options...),
	}
}

func (req *TrashResourceInfoRequest) Exec() (*ResourceInfoResponse, error) {
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
