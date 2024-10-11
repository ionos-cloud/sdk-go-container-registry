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

// ErrorMessages struct for ErrorMessages
type ErrorMessages struct {
	// Application internal error code
	ErrorCode *string `json:"errorCode,omitempty"`
	// A human readable explanation specific to this occurrence of the problem.
	Message *string `json:"message,omitempty"`
}

// NewErrorMessages instantiates a new ErrorMessages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessages() *ErrorMessages {
	this := ErrorMessages{}

	return &this
}

// NewErrorMessagesWithDefaults instantiates a new ErrorMessages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessagesWithDefaults() *ErrorMessages {
	this := ErrorMessages{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ErrorMessages) GetErrorCode() *string {
	if o == nil {
		return nil
	}

	return o.ErrorCode

}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorMessages) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ErrorMessages) SetErrorCode(v string) {

	o.ErrorCode = &v

}

// HasErrorCode returns a boolean if a field has been set.
func (o *ErrorMessages) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ErrorMessages) GetMessage() *string {
	if o == nil {
		return nil
	}

	return o.Message

}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorMessages) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Message, true
}

// SetMessage sets field value
func (o *ErrorMessages) SetMessage(v string) {

	o.Message = &v

}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorMessages) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

func (o ErrorMessages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}

	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	return json.Marshal(toSerialize)
}

type NullableErrorMessages struct {
	value *ErrorMessages
	isSet bool
}

func (v NullableErrorMessages) Get() *ErrorMessages {
	return v.value
}

func (v *NullableErrorMessages) Set(val *ErrorMessages) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessages) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessages(val *ErrorMessages) *NullableErrorMessages {
	return &NullableErrorMessages{value: val, isSet: true}
}

func (v NullableErrorMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
