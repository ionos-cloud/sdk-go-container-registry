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

// PaginationLinks struct for PaginationLinks
type PaginationLinks struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Self     *string `json:"self"`
}

// NewPaginationLinks instantiates a new PaginationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationLinks(next string, previous string, self string) *PaginationLinks {
	this := PaginationLinks{}

	this.Next = &next
	this.Previous = &previous
	this.Self = &self

	return &this
}

// NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationLinksWithDefaults() *PaginationLinks {
	this := PaginationLinks{}
	return &this
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationLinks) GetNext() *string {
	if o == nil {
		return nil
	}

	return o.Next

}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Next, true
}

// SetNext sets field value
func (o *PaginationLinks) SetNext(v string) {

	o.Next = &v

}

// HasNext returns a boolean if a field has been set.
func (o *PaginationLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// GetPrevious returns the Previous field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationLinks) GetPrevious() *string {
	if o == nil {
		return nil
	}

	return o.Previous

}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Previous, true
}

// SetPrevious sets field value
func (o *PaginationLinks) SetPrevious(v string) {

	o.Previous = &v

}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginationLinks) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// GetSelf returns the Self field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationLinks) GetSelf() *string {
	if o == nil {
		return nil
	}

	return o.Self

}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Self, true
}

// SetSelf sets field value
func (o *PaginationLinks) SetSelf(v string) {

	o.Self = &v

}

// HasSelf returns a boolean if a field has been set.
func (o *PaginationLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

func (o PaginationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}

	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}

	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	return json.Marshal(toSerialize)
}

type NullablePaginationLinks struct {
	value *PaginationLinks
	isSet bool
}

func (v NullablePaginationLinks) Get() *PaginationLinks {
	return v.value
}

func (v *NullablePaginationLinks) Set(val *PaginationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationLinks(val *PaginationLinks) *NullablePaginationLinks {
	return &NullablePaginationLinks{value: val, isSet: true}
}

func (v NullablePaginationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
