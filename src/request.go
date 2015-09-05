package src

type apiRequest struct {
	method     string
	path       string
	parameters map[string]interface{}
	headers    map[string][]string
}

func createGetRequest(client *Client, path string, params map[string]interface{}) *apiRequest {
	return createRequest(client, "GET", path, params)
}

func createPostRequest(client *Client, path string, params map[string]interface{}) *apiRequest {
	return createRequest(client, "POST", path, params)
}

func createRequest(client *Client, method string, path string, parameters map[string]interface{}) *apiRequest {
	var headers = make(map[string][]string)
	headers["Authorization"] = []string{"OAuth " + client.token}
	return &apiRequest{
		method:     method,
		path:       path,
		parameters: parameters,
		headers:    headers,
	}
}
