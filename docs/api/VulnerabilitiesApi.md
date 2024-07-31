# \VulnerabilitiesApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**VulnerabilitiesFindByID**](VulnerabilitiesApi.md#VulnerabilitiesFindByID) | **Get** /vulnerabilities/{vulnerabilityId} | Retrieve Vulnerability|



## VulnerabilitiesFindByID

```go
var result VulnerabilityRead = VulnerabilitiesFindByID(ctx, vulnerabilityId)
                      .Execute()
```

Retrieve Vulnerability



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
    vulnerabilityId := "CVE-2019-1234" // string | The ID of the Vulnerability.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.VulnerabilitiesApi.VulnerabilitiesFindByID(context.Background(), vulnerabilityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesFindByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `VulnerabilitiesFindByID`: VulnerabilityRead
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesFindByID`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**vulnerabilityId** | **string** | The ID of the Vulnerability. | |

### Other Parameters

Other parameters are passed through a pointer to an apiVulnerabilitiesFindByIDRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**VulnerabilityRead**](../models/VulnerabilityRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"VulnerabilitiesApiService.VulnerabilitiesFindByID"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "VulnerabilitiesApiService.VulnerabilitiesFindByID": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "VulnerabilitiesApiService.VulnerabilitiesFindByID": {
    "port": "8443",
},
})
```

