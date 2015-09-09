package src

import "strings"

func createResourceInfoRequest(c *Client,
	apiPath string,
	path string,
	options ...ResourceInfoRequestOptions) *httpRequest {
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
		if opt.Fields != nil {
			parameters["fields"] = strings.Join(opt.Fields, ",")
		}
		if opt.Preview_size != nil {
			parameters["preview_size"] = opt.Preview_size.String()
		}
		if opt.Preview_crop != nil {
			parameters["preview_crop"] = opt.Preview_crop
		}
	}
	return createGetRequest(c, apiPath, parameters)
}
