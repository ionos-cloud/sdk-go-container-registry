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

// PostRegistryProperties struct for PostRegistryProperties
type PostRegistryProperties struct {
	GarbageCollectionSchedule *WeeklySchedule   `json:"garbageCollectionSchedule,omitempty"`
	Location                  *string           `json:"location"`
	Name                      *string           `json:"name"`
	Features                  *RegistryFeatures `json:"features,omitempty"`
	// Subnets and IPs that are allowed to access the registry API, supports IPv4 and IPv6. Maximum of 25 items may be specified. If no CIDR is given /32 and /128 are assumed for IPv4 and IPv6 respectively. 0.0.0.0/0 can be used to deny all traffic. __Note__: If this list is empty or not set, there are no restrictions.
	ApiSubnetAllowList *[]string `json:"apiSubnetAllowList,omitempty"`
}

// NewPostRegistryProperties instantiates a new PostRegistryProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRegistryProperties(location string, name string) *PostRegistryProperties {
	this := PostRegistryProperties{}

	this.Location = &location
	this.Name = &name

	return &this
}

// NewPostRegistryPropertiesWithDefaults instantiates a new PostRegistryProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRegistryPropertiesWithDefaults() *PostRegistryProperties {
	this := PostRegistryProperties{}
	return &this
}

// GetGarbageCollectionSchedule returns the GarbageCollectionSchedule field value
// If the value is explicit nil, the zero value for WeeklySchedule will be returned
func (o *PostRegistryProperties) GetGarbageCollectionSchedule() *WeeklySchedule {
	if o == nil {
		return nil
	}

	return o.GarbageCollectionSchedule

}

// GetGarbageCollectionScheduleOk returns a tuple with the GarbageCollectionSchedule field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostRegistryProperties) GetGarbageCollectionScheduleOk() (*WeeklySchedule, bool) {
	if o == nil {
		return nil, false
	}

	return o.GarbageCollectionSchedule, true
}

// SetGarbageCollectionSchedule sets field value
func (o *PostRegistryProperties) SetGarbageCollectionSchedule(v WeeklySchedule) {

	o.GarbageCollectionSchedule = &v

}

// HasGarbageCollectionSchedule returns a boolean if a field has been set.
func (o *PostRegistryProperties) HasGarbageCollectionSchedule() bool {
	if o != nil && o.GarbageCollectionSchedule != nil {
		return true
	}

	return false
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostRegistryProperties) GetLocation() *string {
	if o == nil {
		return nil
	}

	return o.Location

}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostRegistryProperties) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Location, true
}

// SetLocation sets field value
func (o *PostRegistryProperties) SetLocation(v string) {

	o.Location = &v

}

// HasLocation returns a boolean if a field has been set.
func (o *PostRegistryProperties) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostRegistryProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostRegistryProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *PostRegistryProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *PostRegistryProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetFeatures returns the Features field value
// If the value is explicit nil, the zero value for RegistryFeatures will be returned
func (o *PostRegistryProperties) GetFeatures() *RegistryFeatures {
	if o == nil {
		return nil
	}

	return o.Features

}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostRegistryProperties) GetFeaturesOk() (*RegistryFeatures, bool) {
	if o == nil {
		return nil, false
	}

	return o.Features, true
}

// SetFeatures sets field value
func (o *PostRegistryProperties) SetFeatures(v RegistryFeatures) {

	o.Features = &v

}

// HasFeatures returns a boolean if a field has been set.
func (o *PostRegistryProperties) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// GetApiSubnetAllowList returns the ApiSubnetAllowList field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *PostRegistryProperties) GetApiSubnetAllowList() *[]string {
	if o == nil {
		return nil
	}

	return o.ApiSubnetAllowList

}

// GetApiSubnetAllowListOk returns a tuple with the ApiSubnetAllowList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostRegistryProperties) GetApiSubnetAllowListOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ApiSubnetAllowList, true
}

// SetApiSubnetAllowList sets field value
func (o *PostRegistryProperties) SetApiSubnetAllowList(v []string) {

	o.ApiSubnetAllowList = &v

}

// HasApiSubnetAllowList returns a boolean if a field has been set.
func (o *PostRegistryProperties) HasApiSubnetAllowList() bool {
	if o != nil && o.ApiSubnetAllowList != nil {
		return true
	}

	return false
}

func (o PostRegistryProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GarbageCollectionSchedule != nil {
		toSerialize["garbageCollectionSchedule"] = o.GarbageCollectionSchedule
	}

	if o.Location != nil {
		toSerialize["location"] = o.Location
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Features != nil {
		toSerialize["features"] = o.Features
	}

	if o.ApiSubnetAllowList != nil {
		toSerialize["apiSubnetAllowList"] = o.ApiSubnetAllowList
	}

	return json.Marshal(toSerialize)
}

type NullablePostRegistryProperties struct {
	value *PostRegistryProperties
	isSet bool
}

func (v NullablePostRegistryProperties) Get() *PostRegistryProperties {
	return v.value
}

func (v *NullablePostRegistryProperties) Set(val *PostRegistryProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRegistryProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRegistryProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRegistryProperties(val *PostRegistryProperties) *NullablePostRegistryProperties {
	return &NullablePostRegistryProperties{value: val, isSet: true}
}

func (v NullablePostRegistryProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRegistryProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
