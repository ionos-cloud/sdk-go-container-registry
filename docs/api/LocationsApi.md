# \LocationsApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**LocationsGet**](LocationsApi.md#LocationsGet) | **Get** /locations | Get container registry locations|



## LocationsGet

```go
var result LocationsResponse = LocationsGet(ctx)
                      .Execute()
```

Get container registry locations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-container-registry"
)

func main() {

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.LocationsApi.LocationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.LocationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `LocationsGet`: LocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.LocationsGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiLocationsGetRequest struct via the builder pattern


### Return type

[**LocationsResponse**](../models/LocationsResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"LocationsApiService.LocationsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "LocationsApiService.LocationsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "LocationsApiService.LocationsGet": {
    "port": "8443",
},
})
```

