/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature
 *
 * API version: 1.1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// RegistryPagination struct for RegistryPagination
type RegistryPagination struct {
	// The maximum number of elements to return (used together with pagination.token for pagination)
	Limit *int32 `json:"limit"`
	// An opaque token used to iterate the set of results (used together with limit for pagination)
	Token *string `json:"token"`
}

// NewRegistryPagination instantiates a new RegistryPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryPagination(limit int32, token string) *RegistryPagination {
	this := RegistryPagination{}

	this.Limit = &limit
	this.Token = &token

	return &this
}

// NewRegistryPaginationWithDefaults instantiates a new RegistryPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryPaginationWithDefaults() *RegistryPagination {
	this := RegistryPagination{}
	return &this
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RegistryPagination) GetLimit() *int32 {
	if o == nil {
		return nil
	}

	return o.Limit

}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryPagination) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Limit, true
}

// SetLimit sets field value
func (o *RegistryPagination) SetLimit(v int32) {

	o.Limit = &v

}

// HasLimit returns a boolean if a field has been set.
func (o *RegistryPagination) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// GetToken returns the Token field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RegistryPagination) GetToken() *string {
	if o == nil {
		return nil
	}

	return o.Token

}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryPagination) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Token, true
}

// SetToken sets field value
func (o *RegistryPagination) SetToken(v string) {

	o.Token = &v

}

// HasToken returns a boolean if a field has been set.
func (o *RegistryPagination) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

func (o RegistryPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	return json.Marshal(toSerialize)
}

type NullableRegistryPagination struct {
	value *RegistryPagination
	isSet bool
}

func (v NullableRegistryPagination) Get() *RegistryPagination {
	return v.value
}

func (v *NullableRegistryPagination) Set(val *RegistryPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryPagination(val *RegistryPagination) *NullableRegistryPagination {
	return &NullableRegistryPagination{value: val, isSet: true}
}

func (v NullableRegistryPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
