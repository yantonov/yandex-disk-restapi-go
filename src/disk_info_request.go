package src

import "encoding/json"

type DiskInfoRequest struct {
	client      *Client
	httpRequest *httpRequest
}

func (r *DiskInfoRequest) request() *httpRequest {
	return r.httpRequest
}

type DiskInfoResponse struct {
	Trash_size     uint64            `json:"trash_size"`
	Total_space    uint64            `json:"total_space"`
	Used_space     uint64            `json:"used_space"`
	System_folders map[string]string `json:"system_folders"`
}

func (c *Client) NewDiskInfoRequest() *DiskInfoRequest {
	return &DiskInfoRequest{
		client:      c,
		httpRequest: createGetRequest(c, "/", nil),
	}
}

func (req *DiskInfoRequest) Exec() (*DiskInfoResponse, error) {
	data, err := req.request().run(req.client)
	if err != nil {
		return nil, err
	}

	var info DiskInfoResponse
	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}
	if info.System_folders == nil {
		info.System_folders = make(map[string]string)
	}

	return &info, nil
}
