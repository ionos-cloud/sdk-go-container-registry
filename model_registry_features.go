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

// RegistryFeatures Optional registry features.  __Note__: some may incur additional charges - see individual feature descriptions for details
type RegistryFeatures struct {
	VulnerabilityScanning *FeatureVulnerabilityScanning `json:"vulnerabilityScanning,omitempty"`
}

// NewRegistryFeatures instantiates a new RegistryFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryFeatures() *RegistryFeatures {
	this := RegistryFeatures{}

	return &this
}

// NewRegistryFeaturesWithDefaults instantiates a new RegistryFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryFeaturesWithDefaults() *RegistryFeatures {
	this := RegistryFeatures{}
	return &this
}

// GetVulnerabilityScanning returns the VulnerabilityScanning field value
// If the value is explicit nil, the zero value for FeatureVulnerabilityScanning will be returned
func (o *RegistryFeatures) GetVulnerabilityScanning() *FeatureVulnerabilityScanning {
	if o == nil {
		return nil
	}

	return o.VulnerabilityScanning

}

// GetVulnerabilityScanningOk returns a tuple with the VulnerabilityScanning field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistryFeatures) GetVulnerabilityScanningOk() (*FeatureVulnerabilityScanning, bool) {
	if o == nil {
		return nil, false
	}

	return o.VulnerabilityScanning, true
}

// SetVulnerabilityScanning sets field value
func (o *RegistryFeatures) SetVulnerabilityScanning(v FeatureVulnerabilityScanning) {

	o.VulnerabilityScanning = &v

}

// HasVulnerabilityScanning returns a boolean if a field has been set.
func (o *RegistryFeatures) HasVulnerabilityScanning() bool {
	if o != nil && o.VulnerabilityScanning != nil {
		return true
	}

	return false
}

func (o RegistryFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VulnerabilityScanning != nil {
		toSerialize["vulnerabilityScanning"] = o.VulnerabilityScanning
	}

	return json.Marshal(toSerialize)
}

type NullableRegistryFeatures struct {
	value *RegistryFeatures
	isSet bool
}

func (v NullableRegistryFeatures) Get() *RegistryFeatures {
	return v.value
}

func (v *NullableRegistryFeatures) Set(val *RegistryFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryFeatures(val *RegistryFeatures) *NullableRegistryFeatures {
	return &NullableRegistryFeatures{value: val, isSet: true}
}

func (v NullableRegistryFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
