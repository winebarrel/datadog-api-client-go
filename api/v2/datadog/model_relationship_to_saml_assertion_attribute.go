/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// RelationshipToSAMLAssertionAttribute AuthN Mapping relationship to SAML Assertion Attribute.
type RelationshipToSAMLAssertionAttribute struct {
	Data RelationshipToSAMLAssertionAttributeData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewRelationshipToSAMLAssertionAttribute instantiates a new RelationshipToSAMLAssertionAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipToSAMLAssertionAttribute(data RelationshipToSAMLAssertionAttributeData) *RelationshipToSAMLAssertionAttribute {
	this := RelationshipToSAMLAssertionAttribute{}
	this.Data = data
	return &this
}

// NewRelationshipToSAMLAssertionAttributeWithDefaults instantiates a new RelationshipToSAMLAssertionAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipToSAMLAssertionAttributeWithDefaults() *RelationshipToSAMLAssertionAttribute {
	this := RelationshipToSAMLAssertionAttribute{}
	return &this
}

// GetData returns the Data field value
func (o *RelationshipToSAMLAssertionAttribute) GetData() RelationshipToSAMLAssertionAttributeData {
	if o == nil {
		var ret RelationshipToSAMLAssertionAttributeData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RelationshipToSAMLAssertionAttribute) GetDataOk() (*RelationshipToSAMLAssertionAttributeData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RelationshipToSAMLAssertionAttribute) SetData(v RelationshipToSAMLAssertionAttributeData) {
	o.Data = v
}

func (o RelationshipToSAMLAssertionAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *RelationshipToSAMLAssertionAttribute) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Data *RelationshipToSAMLAssertionAttributeData `json:"data"`
	}{}
	all := struct {
		Data RelationshipToSAMLAssertionAttributeData `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Data = all.Data
	return nil
}
