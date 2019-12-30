package src

import "encoding/json"

type DownloadRequest struct {
	client      *Client
	httpRequest *httpRequest
}

func (r *DownloadRequest) request() *httpRequest {
	return r.httpRequest
}

func (c *Client) NewDownloadRequest(path string) *DownloadRequest {
	var parameters = make(map[string]interface{})
	parameters["path"] = path
	return &DownloadRequest{
		client:      c,
		httpRequest: createGetRequest(c, "/resources/download", parameters),
	}
}

func (req *DownloadRequest) Exec() (*LinkResponse, error) {
	data, err := req.request().run(req.client)
	if err != nil {
		return nil, err
	}

	var info LinkResponse
	err = json.Unmarshal(data, &info)

	if err != nil {
		return nil, err
	}

	return &info, nil
}
