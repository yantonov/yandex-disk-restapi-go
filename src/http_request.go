package src

type httpRequest struct {
	method     string
	path       string
	parameters map[string]interface{}
	headers    map[string][]string
}

func createGetRequest(client *Client, path string, params map[string]interface{}) *httpRequest {
	return createRequest(client, "GET", path, params)
}

func createPostRequest(client *Client, path string, params map[string]interface{}) *httpRequest {
	return createRequest(client, "POST", path, params)
}

func createRequest(client *Client, method string, path string, parameters map[string]interface{}) *httpRequest {
	var headers = make(map[string][]string)
	headers["Authorization"] = []string{"OAuth " + client.token}
	return &httpRequest{
		method:     method,
		path:       path,
		parameters: parameters,
		headers:    headers,
	}
}
