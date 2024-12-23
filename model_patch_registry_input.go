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

// PatchRegistryInput struct for PatchRegistryInput
type PatchRegistryInput struct {
	GarbageCollectionSchedule *WeeklySchedule   `json:"garbageCollectionSchedule,omitempty"`
	Features                  *RegistryFeatures `json:"features,omitempty"`
	// Subnets and IPs that are allowed to access the registry API, supports IPv4 and IPv6. Maximum of 25 items may be specified. If no CIDR is given /32 and /128 are assumed for IPv4 and IPv6 respectively. 0.0.0.0/0 can be used to deny all traffic. __Note__: If this list is empty or not set, there are no restrictions.
	ApiSubnetAllowList *[]string `json:"apiSubnetAllowList,omitempty"`
}

// NewPatchRegistryInput instantiates a new PatchRegistryInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchRegistryInput() *PatchRegistryInput {
	this := PatchRegistryInput{}

	return &this
}

// NewPatchRegistryInputWithDefaults instantiates a new PatchRegistryInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchRegistryInputWithDefaults() *PatchRegistryInput {
	this := PatchRegistryInput{}
	return &this
}

// GetGarbageCollectionSchedule returns the GarbageCollectionSchedule field value
// If the value is explicit nil, the zero value for WeeklySchedule will be returned
func (o *PatchRegistryInput) GetGarbageCollectionSchedule() *WeeklySchedule {
	if o == nil {
		return nil
	}

	return o.GarbageCollectionSchedule

}

// GetGarbageCollectionScheduleOk returns a tuple with the GarbageCollectionSchedule field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchRegistryInput) GetGarbageCollectionScheduleOk() (*WeeklySchedule, bool) {
	if o == nil {
		return nil, false
	}

	return o.GarbageCollectionSchedule, true
}

// SetGarbageCollectionSchedule sets field value
func (o *PatchRegistryInput) SetGarbageCollectionSchedule(v WeeklySchedule) {

	o.GarbageCollectionSchedule = &v

}

// HasGarbageCollectionSchedule returns a boolean if a field has been set.
func (o *PatchRegistryInput) HasGarbageCollectionSchedule() bool {
	if o != nil && o.GarbageCollectionSchedule != nil {
		return true
	}

	return false
}

// GetFeatures returns the Features field value
// If the value is explicit nil, the zero value for RegistryFeatures will be returned
func (o *PatchRegistryInput) GetFeatures() *RegistryFeatures {
	if o == nil {
		return nil
	}

	return o.Features

}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchRegistryInput) GetFeaturesOk() (*RegistryFeatures, bool) {
	if o == nil {
		return nil, false
	}

	return o.Features, true
}

// SetFeatures sets field value
func (o *PatchRegistryInput) SetFeatures(v RegistryFeatures) {

	o.Features = &v

}

// HasFeatures returns a boolean if a field has been set.
func (o *PatchRegistryInput) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// GetApiSubnetAllowList returns the ApiSubnetAllowList field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *PatchRegistryInput) GetApiSubnetAllowList() *[]string {
	if o == nil {
		return nil
	}

	return o.ApiSubnetAllowList

}

// GetApiSubnetAllowListOk returns a tuple with the ApiSubnetAllowList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchRegistryInput) GetApiSubnetAllowListOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ApiSubnetAllowList, true
}

// SetApiSubnetAllowList sets field value
func (o *PatchRegistryInput) SetApiSubnetAllowList(v []string) {

	o.ApiSubnetAllowList = &v

}

// HasApiSubnetAllowList returns a boolean if a field has been set.
func (o *PatchRegistryInput) HasApiSubnetAllowList() bool {
	if o != nil && o.ApiSubnetAllowList != nil {
		return true
	}

	return false
}

func (o PatchRegistryInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GarbageCollectionSchedule != nil {
		toSerialize["garbageCollectionSchedule"] = o.GarbageCollectionSchedule
	}

	if o.Features != nil {
		toSerialize["features"] = o.Features
	}

	if o.ApiSubnetAllowList != nil {
		toSerialize["apiSubnetAllowList"] = o.ApiSubnetAllowList
	}

	return json.Marshal(toSerialize)
}

type NullablePatchRegistryInput struct {
	value *PatchRegistryInput
	isSet bool
}

func (v NullablePatchRegistryInput) Get() *PatchRegistryInput {
	return v.value
}

func (v *NullablePatchRegistryInput) Set(val *PatchRegistryInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchRegistryInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchRegistryInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchRegistryInput(val *PatchRegistryInput) *NullablePatchRegistryInput {
	return &NullablePatchRegistryInput{value: val, isSet: true}
}

func (v NullablePatchRegistryInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchRegistryInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
