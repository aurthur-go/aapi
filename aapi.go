package aapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	// GET :
	GET = "GET"
	// POST :
	POST = "POST"
	// PUT :
	PUT = "PUT"
	// DELETE :
	DELETE = "DELETE"
)

// Resource :
type Resource interface {
	Get(values url.Values) (int, interface{})
	Post(values url.Values) (int, interface{})
	Put(values url.Values) (int, interface{})
	Delete(values url.Values) (int, interface{})
}

type (
	// GetNotSupported :
	GetNotSupported struct{}
	// PostNotSupported :
	PostNotSupported struct{}
	// PutNotSupported :
	PutNotSupported struct{}
	// DeleteNotSupported :
	DeleteNotSupported struct{}
)

// Get :
func (GetNotSupported) Get(values url.Values) (int, interface{}) {
	return 405, map[string]string{"status": "405", "msg": "Method Not allowed"}
}

// Post :
func (PostNotSupported) Post(values url.Values) (int, interface{}) {
	return 405, map[string]string{"status": "405", "msg": "Method Not allowed"}
}

// Put :
func (PutNotSupported) Put(values url.Values) (int, interface{}) {
	return 405, map[string]string{"status": "405", "msg": "Method Not allowed"}
}

// Delete :
func (DeleteNotSupported) Delete(values url.Values) (int, interface{}) {
	return 405, map[string]string{"status": "405", "msg": "Method Not allowed"}
}

// API :
type API struct{}

// Abort :
func (api *API) Abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}

func (api *API) requestHandler(resource Resource) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		var data interface{}
		var code int

		request.ParseForm()
		method := request.Method
		values := request.Form

		switch method {
		case GET:
			code, data = resource.Get(values)
		case POST:
			code, data = resource.Post(values)
		case PUT:
			code, data = resource.Put(values)
		case DELETE:
			code, data = resource.Delete(values)
		default:
			api.Abort(rw, 405)
			return
		}

		content, err := json.Marshal(data)
		if err != nil {
			api.Abort(rw, 500)
		}
		rw.WriteHeader(code)
		rw.Write(content)
	}
}

// AddResource :
func (api *API) AddResource(resource Resource, path string) {
	http.HandleFunc(path, api.requestHandler(resource))
}

// Start :
func (api *API) Start(port string) {
	http.ListenAndServe(":"+port, nil)
}
