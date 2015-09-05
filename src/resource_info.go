package src

import "encoding/json"

type ResourceInfoRequest struct {
	client      *Client
	httpRequest *httpRequest
}

func (r *ResourceInfoRequest) request() *httpRequest {
	return r.httpRequest
}

type SortMode struct {
	mode string
}

func (m *SortMode) ByName() *SortMode {
	return &SortMode{
		mode: "name",
	}
}

func (m *SortMode) String() string {
	if m.mode == "" {
		panic("undefined mode")
	}
	return m.mode
}

type ResourceRequestOptions struct {
	sortMode *SortMode
}

type ResourceInfoResponse struct {
	Public_key        string                 `json:"public_key"`
	Name              string                 `json:"name"`
	Created           string                 `json:"created"`
	Custom_properties map[string]interface{} `json:"custom_properties"`
	Public_url        string                 `json:"public_url"`
	Origin_path       string                 `json:"origin_path"`
	Modified          string                 `json:"modified"`
	Path              string                 `json:"path"`
	Md5               string                 `json:"md5"`
	Resource_type     string                 `json:"type"`
	Mime_type         string                 `json:"mime_type"`
	Size              uint64                 `json:"size"`
	// "_embedded": {объект ResourceList},
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
