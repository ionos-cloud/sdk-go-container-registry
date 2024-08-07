# Error

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | Pointer to **int32** | The HTTP status code of the operation. | [optional] |
|**Messages** | Pointer to [**[]ErrorMessages**](ErrorMessages.md) | A list of error messages.  | [optional] |

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatus

`func (o *Error) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *Error) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *Error) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *Error) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessages

`func (o *Error) GetMessages() []ErrorMessages`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Error) GetMessagesOk() (*[]ErrorMessages, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Error) SetMessages(v []ErrorMessages)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Error) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


