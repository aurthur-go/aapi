## aapi

使用教程

```go
package main

import "github.com/aurthur-go/aapi"

// IndexResource : Only support get
type IndexResource struct {
    aapi.DeleteNotSupported
    aapi.PostNotSupported
    aapi.PutNotSupported
}

// StatusResponse : status and msg
type StatusResponse struct {
    Status int    `json:"status"`
    Msg    string `json:"msg"`
}

var statusResponse = StatusResponse{}

const (
    isOK     = 200
    isError  = 201
)

func main() {

    indexResource := new(IndexResource)

    var aapi = new(aapi.API)

    aapi.AddResource(indexResource, "/")

    aapi.Start("8080")
}

// Get : /
func (IndexResource) Get(values url.Values) (int, interface{}) {
    statusResponse.Msg = "welcome!"
    statusResponse.Status = isOK
    return isOK, statusResponse
}


```