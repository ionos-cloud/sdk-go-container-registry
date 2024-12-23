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

// ApiErrorMessage struct for ApiErrorMessage
type ApiErrorMessage struct {
	ErrorCode *string `json:"errorCode"`
	Message   *string `json:"message"`
}

// NewApiErrorMessage instantiates a new ApiErrorMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorMessage(errorCode string, message string) *ApiErrorMessage {
	this := ApiErrorMessage{}

	this.ErrorCode = &errorCode
	this.Message = &message

	return &this
}

// NewApiErrorMessageWithDefaults instantiates a new ApiErrorMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorMessageWithDefaults() *ApiErrorMessage {
	this := ApiErrorMessage{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiErrorMessage) GetErrorCode() *string {
	if o == nil {
		return nil
	}

	return o.ErrorCode

}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiErrorMessage) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ApiErrorMessage) SetErrorCode(v string) {

	o.ErrorCode = &v

}

// HasErrorCode returns a boolean if a field has been set.
func (o *ApiErrorMessage) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiErrorMessage) GetMessage() *string {
	if o == nil {
		return nil
	}

	return o.Message

}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiErrorMessage) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Message, true
}

// SetMessage sets field value
func (o *ApiErrorMessage) SetMessage(v string) {

	o.Message = &v

}

// HasMessage returns a boolean if a field has been set.
func (o *ApiErrorMessage) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

func (o ApiErrorMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}

	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	return json.Marshal(toSerialize)
}

type NullableApiErrorMessage struct {
	value *ApiErrorMessage
	isSet bool
}

func (v NullableApiErrorMessage) Get() *ApiErrorMessage {
	return v.value
}

func (v *NullableApiErrorMessage) Set(val *ApiErrorMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorMessage(val *ApiErrorMessage) *NullableApiErrorMessage {
	return &NullableApiErrorMessage{value: val, isSet: true}
}

func (v NullableApiErrorMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
