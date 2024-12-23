/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature ### 1.2.0  - Added registry `apiSubnetAllowList` ### 1.2.1  - Amended `apiSubnetAllowList` Regex
 *
 * API version: 1.2.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ApiErrorResponse struct for ApiErrorResponse
type ApiErrorResponse struct {
	HttpStatus *int32             `json:"httpStatus"`
	Messages   *[]ApiErrorMessage `json:"messages"`
}

// NewApiErrorResponse instantiates a new ApiErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorResponse(httpStatus int32, messages []ApiErrorMessage) *ApiErrorResponse {
	this := ApiErrorResponse{}

	this.HttpStatus = &httpStatus
	this.Messages = &messages

	return &this
}

// NewApiErrorResponseWithDefaults instantiates a new ApiErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorResponseWithDefaults() *ApiErrorResponse {
	this := ApiErrorResponse{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ApiErrorResponse) GetHttpStatus() *int32 {
	if o == nil {
		return nil
	}

	return o.HttpStatus

}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiErrorResponse) GetHttpStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ApiErrorResponse) SetHttpStatus(v int32) {

	o.HttpStatus = &v

}

// HasHttpStatus returns a boolean if a field has been set.
func (o *ApiErrorResponse) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// GetMessages returns the Messages field value
// If the value is explicit nil, the zero value for []ApiErrorMessage will be returned
func (o *ApiErrorResponse) GetMessages() *[]ApiErrorMessage {
	if o == nil {
		return nil
	}

	return o.Messages

}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiErrorResponse) GetMessagesOk() (*[]ApiErrorMessage, bool) {
	if o == nil {
		return nil, false
	}

	return o.Messages, true
}

// SetMessages sets field value
func (o *ApiErrorResponse) SetMessages(v []ApiErrorMessage) {

	o.Messages = &v

}

// HasMessages returns a boolean if a field has been set.
func (o *ApiErrorResponse) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

func (o ApiErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpStatus != nil {
		toSerialize["httpStatus"] = o.HttpStatus
	}

	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}

	return json.Marshal(toSerialize)
}

type NullableApiErrorResponse struct {
	value *ApiErrorResponse
	isSet bool
}

func (v NullableApiErrorResponse) Get() *ApiErrorResponse {
	return v.value
}

func (v *NullableApiErrorResponse) Set(val *ApiErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorResponse(val *ApiErrorResponse) *NullableApiErrorResponse {
	return &NullableApiErrorResponse{value: val, isSet: true}
}

func (v NullableApiErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
