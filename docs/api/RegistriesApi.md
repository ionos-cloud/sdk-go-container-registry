# \RegistriesApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RegistriesDelete**](RegistriesApi.md#RegistriesDelete) | **Delete** /registries/{registryId} | Delete registry|
|[**RegistriesFindById**](RegistriesApi.md#RegistriesFindById) | **Get** /registries/{registryId} | Get a registry|
|[**RegistriesGet**](RegistriesApi.md#RegistriesGet) | **Get** /registries | List all container registries|
|[**RegistriesPatch**](RegistriesApi.md#RegistriesPatch) | **Patch** /registries/{registryId} | Update the properties of a registry|
|[**RegistriesPost**](RegistriesApi.md#RegistriesPost) | **Post** /registries | Create container registry|
|[**RegistriesPut**](RegistriesApi.md#RegistriesPut) | **Put** /registries/{registryId} | Create or replace a container registry|



## RegistriesDelete

```go
var result  = RegistriesDelete(ctx, registryId)
                      .Execute()
```

Delete registry



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
    registryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the registry

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.RegistriesApi.RegistriesDelete(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RegistriesApiService.RegistriesDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RegistriesApiService.RegistriesDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RegistriesApiService.RegistriesDelete": {
    "port": "8443",
},
})
```


## RegistriesFindById

```go
var result RegistryResponse = RegistriesFindById(ctx, registryId)
                      .Execute()
```

Get a registry



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
    registryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the registry

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesFindById(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesFindById`: RegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RegistryResponse**](../models/RegistryResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RegistriesApiService.RegistriesFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RegistriesApiService.RegistriesFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RegistriesApiService.RegistriesFindById": {
    "port": "8443",
},
})
```


## RegistriesGet

```go
var result RegistriesResponse = RegistriesGet(ctx)
                      .FilterName(filterName)
                      .Limit(limit)
                      .PaginationToken(paginationToken)
                      .Execute()
```

List all container registries



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
    filterName := "my-registry" // string | The registry name to search for (optional)
    limit := "limit_example" // string | The maximum number of elements to return (used together with pagination.token for pagination) (optional) (default to "100")
    paginationToken := "eyJ2IjoibWV0YS5rOHMuaW8vdjEiLCJydiI6MTYzMjQ0OTk2ODAsInN0YXJ0IjoiM2RmYTc3YjctZGIwNS00MjMwLThmMjAtOGU3NjJlOTUxOTUzXHUwMDAwIn0" // string | An opaque token used to iterate the set of results (used together with limit for pagination) (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesGet(context.Background()).FilterName(filterName).Limit(limit).PaginationToken(paginationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesGet`: RegistriesResponse
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **filterName** | **string** | The registry name to search for | |
| **limit** | **string** | The maximum number of elements to return (used together with pagination.token for pagination) | [default to &quot;100&quot;]|
| **paginationToken** | **string** | An opaque token used to iterate the set of results (used together with limit for pagination) | |

### Return type

[**RegistriesResponse**](../models/RegistriesResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RegistriesApiService.RegistriesGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RegistriesApiService.RegistriesGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RegistriesApiService.RegistriesGet": {
    "port": "8443",
},
})
```


## RegistriesPatch

```go
var result RegistryResponse = RegistriesPatch(ctx, registryId)
                      .PatchRegistryInput(patchRegistryInput)
                      .Execute()
```

Update the properties of a registry



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
    registryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the registry
    patchRegistryInput := *openapiclient.NewPatchRegistryInput() // PatchRegistryInput | 

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesPatch(context.Background(), registryId).PatchRegistryInput(patchRegistryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesPatch`: RegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **patchRegistryInput** | [**PatchRegistryInput**](../models/PatchRegistryInput.md) |  | |

### Return type

[**RegistryResponse**](../models/RegistryResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RegistriesApiService.RegistriesPatch"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RegistriesApiService.RegistriesPatch": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RegistriesApiService.RegistriesPatch": {
    "port": "8443",
},
})
```


## RegistriesPost

```go
var result PostRegistryOutput = RegistriesPost(ctx)
                      .PostRegistryInput(postRegistryInput)
                      .Execute()
```

Create container registry



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
    postRegistryInput := *openapiclient.NewPostRegistryInput(*openapiclient.NewPostRegistryProperties("de/txl", "my-registry")) // PostRegistryInput | 

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesPost(context.Background()).PostRegistryInput(postRegistryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesPost`: PostRegistryOutput
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **postRegistryInput** | [**PostRegistryInput**](../models/PostRegistryInput.md) |  | |

### Return type

[**PostRegistryOutput**](../models/PostRegistryOutput.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RegistriesApiService.RegistriesPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RegistriesApiService.RegistriesPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RegistriesApiService.RegistriesPost": {
    "port": "8443",
},
})
```


## RegistriesPut

```go
var result PutRegistryOutput = RegistriesPut(ctx, registryId)
                      .PutRegistryInput(putRegistryInput)
                      .Execute()
```

Create or replace a container registry



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
    registryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the registry
    putRegistryInput := *openapiclient.NewPutRegistryInput(*openapiclient.NewPostRegistryProperties("de/txl", "my-registry")) // PutRegistryInput | 

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesPut(context.Background(), registryId).PutRegistryInput(putRegistryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesPut`: PutRegistryOutput
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **putRegistryInput** | [**PutRegistryInput**](../models/PutRegistryInput.md) |  | |

### Return type

[**PutRegistryOutput**](../models/PutRegistryOutput.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RegistriesApiService.RegistriesPut"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RegistriesApiService.RegistriesPut": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RegistriesApiService.RegistriesPut": {
    "port": "8443",
},
})
```

