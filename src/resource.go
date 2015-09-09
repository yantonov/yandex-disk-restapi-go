package src

type ResourceInfoResponse struct {
	Public_key        string                 `json:"public_key"`
	Name              string                 `json:"name"`
	Created           string                 `json:"created"`
	Custom_properties map[string]interface{} `json:"custom_properties"`
	Preview           string                 `json:"preview"`
	Public_url        string                 `json:"public_url"`
	Origin_path       string                 `json:"origin_path"`
	Modified          string                 `json:"modified"`
	Path              string                 `json:"path"`
	Md5               string                 `json:"md5"`
	Resource_type     string                 `json:"type"`
	Mime_type         string                 `json:"mime_type"`
	Size              uint64                 `json:"size"`
	Embedded          *ResourceListResponse  `json:"_embedded"`
}
