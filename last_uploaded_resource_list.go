package yandexdiskapi

type LastUploadedResourceListResponse struct {
	Items []ResourceInfoResponse `json:"items"`
	Limit *uint64                `json:"limit"`
}
