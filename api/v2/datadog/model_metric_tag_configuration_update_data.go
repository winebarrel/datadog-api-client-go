/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// MetricTagConfigurationUpdateData Object for a single tag configuration to be edited.
type MetricTagConfigurationUpdateData struct {
	Attributes *MetricTagConfigurationUpdateAttributes `json:"attributes,omitempty"`
	// The metric name for this resource.
	Id   string                     `json:"id"`
	Type MetricTagConfigurationType `json:"type"`
}

// NewMetricTagConfigurationUpdateData instantiates a new MetricTagConfigurationUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTagConfigurationUpdateData(id string, type_ MetricTagConfigurationType) *MetricTagConfigurationUpdateData {
	this := MetricTagConfigurationUpdateData{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewMetricTagConfigurationUpdateDataWithDefaults instantiates a new MetricTagConfigurationUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTagConfigurationUpdateDataWithDefaults() *MetricTagConfigurationUpdateData {
	this := MetricTagConfigurationUpdateData{}
	var type_ MetricTagConfigurationType = "manage_tags"
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MetricTagConfigurationUpdateData) GetAttributes() MetricTagConfigurationUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret MetricTagConfigurationUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationUpdateData) GetAttributesOk() (*MetricTagConfigurationUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MetricTagConfigurationUpdateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MetricTagConfigurationUpdateAttributes and assigns it to the Attributes field.
func (o *MetricTagConfigurationUpdateData) SetAttributes(v MetricTagConfigurationUpdateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value
func (o *MetricTagConfigurationUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MetricTagConfigurationUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *MetricTagConfigurationUpdateData) GetType() MetricTagConfigurationType {
	if o == nil {
		var ret MetricTagConfigurationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationUpdateData) GetTypeOk() (*MetricTagConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetricTagConfigurationUpdateData) SetType(v MetricTagConfigurationType) {
	o.Type = v
}

func (o MetricTagConfigurationUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMetricTagConfigurationUpdateData struct {
	value *MetricTagConfigurationUpdateData
	isSet bool
}

func (v NullableMetricTagConfigurationUpdateData) Get() *MetricTagConfigurationUpdateData {
	return v.value
}

func (v *NullableMetricTagConfigurationUpdateData) Set(val *MetricTagConfigurationUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationUpdateData(val *MetricTagConfigurationUpdateData) *NullableMetricTagConfigurationUpdateData {
	return &NullableMetricTagConfigurationUpdateData{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
