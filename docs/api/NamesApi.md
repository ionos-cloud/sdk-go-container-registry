# \NamesApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**NamesCheckUsage**](NamesApi.md#NamesCheckUsage) | **Head** /names/{name} | Get container registry name availability|



## NamesCheckUsage

```go
var result  = NamesCheckUsage(ctx, name)
                      .Execute()
```

Get container registry name availability



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
    name := "name_example" // string | The desired registry name

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.NamesApi.NamesCheckUsage(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.NamesCheckUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**name** | **string** | The desired registry name | |

### Other Parameters

Other parameters are passed through a pointer to an apiNamesCheckUsageRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"NamesApiService.NamesCheckUsage"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "NamesApiService.NamesCheckUsage": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "NamesApiService.NamesCheckUsage": {
    "port": "8443",
},
})
```

