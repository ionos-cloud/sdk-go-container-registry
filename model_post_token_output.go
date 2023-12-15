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

// PostTokenOutput struct for PostTokenOutput
type PostTokenOutput struct {
	Href       *string              `json:"href,omitempty"`
	Id         *string              `json:"id,omitempty"`
	Metadata   *ApiResourceMetadata `json:"metadata"`
	Properties *TokenProperties     `json:"properties"`
	Type       *string              `json:"type,omitempty"`
}

// NewPostTokenOutput instantiates a new PostTokenOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTokenOutput(metadata NullableApiResourceMetadata, properties NullableTokenProperties) *PostTokenOutput {
	this := PostTokenOutput{}

	this.Metadata = metadata.value
	this.Properties = properties.value

	return &this
}

// NewPostTokenOutputWithDefaults instantiates a new PostTokenOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTokenOutputWithDefaults() *PostTokenOutput {
	this := PostTokenOutput{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostTokenOutput) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTokenOutput) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *PostTokenOutput) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *PostTokenOutput) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostTokenOutput) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTokenOutput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *PostTokenOutput) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *PostTokenOutput) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for ApiResourceMetadata will be returned
func (o *PostTokenOutput) GetMetadata() *ApiResourceMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTokenOutput) GetMetadataOk() (*ApiResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *PostTokenOutput) SetMetadata(v ApiResourceMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *PostTokenOutput) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for TokenProperties will be returned
func (o *PostTokenOutput) GetProperties() *TokenProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTokenOutput) GetPropertiesOk() (*TokenProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *PostTokenOutput) SetProperties(v TokenProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *PostTokenOutput) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostTokenOutput) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTokenOutput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *PostTokenOutput) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *PostTokenOutput) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o PostTokenOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	toSerialize["metadata"] = o.Metadata

	toSerialize["properties"] = o.Properties

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullablePostTokenOutput struct {
	value *PostTokenOutput
	isSet bool
}

func (v NullablePostTokenOutput) Get() *PostTokenOutput {
	return v.value
}

func (v *NullablePostTokenOutput) Set(val *PostTokenOutput) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTokenOutput) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTokenOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTokenOutput(val *PostTokenOutput) *NullablePostTokenOutput {
	return &NullablePostTokenOutput{value: val, isSet: true}
}

func (v NullablePostTokenOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTokenOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
