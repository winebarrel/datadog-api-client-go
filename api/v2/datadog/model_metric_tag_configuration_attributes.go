/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// MetricTagConfigurationAttributes Object containing the definition of a metric tag configuration attributes.
type MetricTagConfigurationAttributes struct {
	// Timestamp when the tag configuration was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Toggle to turn on/off percentile aggregations for distribution metrics. Only present when the `metric_type` is `distribution`.
	IncludePercentiles *bool                              `json:"include_percentiles,omitempty"`
	MetricType         *MetricTagConfigurationMetricTypes `json:"metric_type,omitempty"`
	// Timestamp when the tag configuration was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// List of tag keys on which to group.
	Tags *[]string `json:"tags,omitempty"`
}

// NewMetricTagConfigurationAttributes instantiates a new MetricTagConfigurationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTagConfigurationAttributes() *MetricTagConfigurationAttributes {
	this := MetricTagConfigurationAttributes{}
	var metricType MetricTagConfigurationMetricTypes = METRICTAGCONFIGURATIONMETRICTYPES_GAUGE
	this.MetricType = &metricType
	return &this
}

// NewMetricTagConfigurationAttributesWithDefaults instantiates a new MetricTagConfigurationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTagConfigurationAttributesWithDefaults() *MetricTagConfigurationAttributes {
	this := MetricTagConfigurationAttributes{}
	var metricType MetricTagConfigurationMetricTypes = METRICTAGCONFIGURATIONMETRICTYPES_GAUGE
	this.MetricType = &metricType
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MetricTagConfigurationAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MetricTagConfigurationAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MetricTagConfigurationAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetIncludePercentiles returns the IncludePercentiles field value if set, zero value otherwise.
func (o *MetricTagConfigurationAttributes) GetIncludePercentiles() bool {
	if o == nil || o.IncludePercentiles == nil {
		var ret bool
		return ret
	}
	return *o.IncludePercentiles
}

// GetIncludePercentilesOk returns a tuple with the IncludePercentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationAttributes) GetIncludePercentilesOk() (*bool, bool) {
	if o == nil || o.IncludePercentiles == nil {
		return nil, false
	}
	return o.IncludePercentiles, true
}

// HasIncludePercentiles returns a boolean if a field has been set.
func (o *MetricTagConfigurationAttributes) HasIncludePercentiles() bool {
	if o != nil && o.IncludePercentiles != nil {
		return true
	}

	return false
}

// SetIncludePercentiles gets a reference to the given bool and assigns it to the IncludePercentiles field.
func (o *MetricTagConfigurationAttributes) SetIncludePercentiles(v bool) {
	o.IncludePercentiles = &v
}

// GetMetricType returns the MetricType field value if set, zero value otherwise.
func (o *MetricTagConfigurationAttributes) GetMetricType() MetricTagConfigurationMetricTypes {
	if o == nil || o.MetricType == nil {
		var ret MetricTagConfigurationMetricTypes
		return ret
	}
	return *o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationAttributes) GetMetricTypeOk() (*MetricTagConfigurationMetricTypes, bool) {
	if o == nil || o.MetricType == nil {
		return nil, false
	}
	return o.MetricType, true
}

// HasMetricType returns a boolean if a field has been set.
func (o *MetricTagConfigurationAttributes) HasMetricType() bool {
	if o != nil && o.MetricType != nil {
		return true
	}

	return false
}

// SetMetricType gets a reference to the given MetricTagConfigurationMetricTypes and assigns it to the MetricType field.
func (o *MetricTagConfigurationAttributes) SetMetricType(v MetricTagConfigurationMetricTypes) {
	o.MetricType = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *MetricTagConfigurationAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *MetricTagConfigurationAttributes) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *MetricTagConfigurationAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetricTagConfigurationAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetricTagConfigurationAttributes) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetricTagConfigurationAttributes) SetTags(v []string) {
	o.Tags = &v
}

func (o MetricTagConfigurationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.IncludePercentiles != nil {
		toSerialize["include_percentiles"] = o.IncludePercentiles
	}
	if o.MetricType != nil {
		toSerialize["metric_type"] = o.MetricType
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableMetricTagConfigurationAttributes struct {
	value *MetricTagConfigurationAttributes
	isSet bool
}

func (v NullableMetricTagConfigurationAttributes) Get() *MetricTagConfigurationAttributes {
	return v.value
}

func (v *NullableMetricTagConfigurationAttributes) Set(val *MetricTagConfigurationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationAttributes(val *MetricTagConfigurationAttributes) *NullableMetricTagConfigurationAttributes {
	return &NullableMetricTagConfigurationAttributes{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
