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

// SyntheticsSSLCertificateSubject Object describing the SSL certificate used for the test.
type SyntheticsSSLCertificateSubject struct {
	// Country Name associated with the certificate.
	C *string `json:"C,omitempty"`
	// Common Name that associated with the certificate.
	CN *string `json:"CN,omitempty"`
	// Locality associated with the certificate.
	L *string `json:"L,omitempty"`
	// Organization associated with the certificate.
	O *string `json:"O,omitempty"`
	// Organizational Unit associated with the certificate.
	OU *string `json:"OU,omitempty"`
	// State Or Province Name associated with the certificate.
	ST *string `json:"ST,omitempty"`
	// Subject Alternative Name associated with the certificate.
	AltName *string `json:"altName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSyntheticsSSLCertificateSubject instantiates a new SyntheticsSSLCertificateSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsSSLCertificateSubject() *SyntheticsSSLCertificateSubject {
	this := SyntheticsSSLCertificateSubject{}
	return &this
}

// NewSyntheticsSSLCertificateSubjectWithDefaults instantiates a new SyntheticsSSLCertificateSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsSSLCertificateSubjectWithDefaults() *SyntheticsSSLCertificateSubject {
	this := SyntheticsSSLCertificateSubject{}
	return &this
}

// GetC returns the C field value if set, zero value otherwise.
func (o *SyntheticsSSLCertificateSubject) GetC() string {
	if o == nil || o.C == nil {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSSLCertificateSubject) GetCOk() (*string, bool) {
	if o == nil || o.C == nil {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *SyntheticsSSLCertificateSubject) HasC() bool {
	if o != nil && o.C != nil {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *SyntheticsSSLCertificateSubject) SetC(v string) {
	o.C = &v
}

// GetCN returns the CN field value if set, zero value otherwise.
func (o *SyntheticsSSLCertificateSubject) GetCN() string {
	if o == nil || o.CN == nil {
		var ret string
		return ret
	}
	return *o.CN
}

// GetCNOk returns a tuple with the CN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSSLCertificateSubject) GetCNOk() (*string, bool) {
	if o == nil || o.CN == nil {
		return nil, false
	}
	return o.CN, true
}

// HasCN returns a boolean if a field has been set.
func (o *SyntheticsSSLCertificateSubject) HasCN() bool {
	if o != nil && o.CN != nil {
		return true
	}

	return false
}

// SetCN gets a reference to the given string and assigns it to the CN field.
func (o *SyntheticsSSLCertificateSubject) SetCN(v string) {
	o.CN = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *SyntheticsSSLCertificateSubject) GetL() string {
	if o == nil || o.L == nil {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSSLCertificateSubject) GetLOk() (*string, bool) {
	if o == nil || o.L == nil {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *SyntheticsSSLCertificateSubject) HasL() bool {
	if o != nil && o.L != nil {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *SyntheticsSSLCertificateSubject) SetL(v string) {
	o.L = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *SyntheticsSSLCertificateSubject) GetO() string {
	if o == nil || o.O == nil {
		var ret string
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSSLCertificateSubject) GetOOk() (*string, bool) {
	if o == nil || o.O == nil {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *SyntheticsSSLCertificateSubject) HasO() bool {
	if o != nil && o.O != nil {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *SyntheticsSSLCertificateSubject) SetO(v string) {
	o.O = &v
}

// GetOU returns the OU field value if set, zero value otherwise.
func (o *SyntheticsSSLCertificateSubject) GetOU() string {
	if o == nil || o.OU == nil {
		var ret string
		return ret
	}
	return *o.OU
}

// GetOUOk returns a tuple with the OU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSSLCertificateSubject) GetOUOk() (*string, bool) {
	if o == nil || o.OU == nil {
		return nil, false
	}
	return o.OU, true
}

// HasOU returns a boolean if a field has been set.
func (o *SyntheticsSSLCertificateSubject) HasOU() bool {
	if o != nil && o.OU != nil {
		return true
	}

	return false
}

// SetOU gets a reference to the given string and assigns it to the OU field.
func (o *SyntheticsSSLCertificateSubject) SetOU(v string) {
	o.OU = &v
}

// GetST returns the ST field value if set, zero value otherwise.
func (o *SyntheticsSSLCertificateSubject) GetST() string {
	if o == nil || o.ST == nil {
		var ret string
		return ret
	}
	return *o.ST
}

// GetSTOk returns a tuple with the ST field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSSLCertificateSubject) GetSTOk() (*string, bool) {
	if o == nil || o.ST == nil {
		return nil, false
	}
	return o.ST, true
}

// HasST returns a boolean if a field has been set.
func (o *SyntheticsSSLCertificateSubject) HasST() bool {
	if o != nil && o.ST != nil {
		return true
	}

	return false
}

// SetST gets a reference to the given string and assigns it to the ST field.
func (o *SyntheticsSSLCertificateSubject) SetST(v string) {
	o.ST = &v
}

// GetAltName returns the AltName field value if set, zero value otherwise.
func (o *SyntheticsSSLCertificateSubject) GetAltName() string {
	if o == nil || o.AltName == nil {
		var ret string
		return ret
	}
	return *o.AltName
}

// GetAltNameOk returns a tuple with the AltName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSSLCertificateSubject) GetAltNameOk() (*string, bool) {
	if o == nil || o.AltName == nil {
		return nil, false
	}
	return o.AltName, true
}

// HasAltName returns a boolean if a field has been set.
func (o *SyntheticsSSLCertificateSubject) HasAltName() bool {
	if o != nil && o.AltName != nil {
		return true
	}

	return false
}

// SetAltName gets a reference to the given string and assigns it to the AltName field.
func (o *SyntheticsSSLCertificateSubject) SetAltName(v string) {
	o.AltName = &v
}

func (o SyntheticsSSLCertificateSubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.C != nil {
		toSerialize["C"] = o.C
	}
	if o.CN != nil {
		toSerialize["CN"] = o.CN
	}
	if o.L != nil {
		toSerialize["L"] = o.L
	}
	if o.O != nil {
		toSerialize["O"] = o.O
	}
	if o.OU != nil {
		toSerialize["OU"] = o.OU
	}
	if o.ST != nil {
		toSerialize["ST"] = o.ST
	}
	if o.AltName != nil {
		toSerialize["altName"] = o.AltName
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsSSLCertificateSubject) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		C       *string `json:"C,omitempty"`
		CN      *string `json:"CN,omitempty"`
		L       *string `json:"L,omitempty"`
		O       *string `json:"O,omitempty"`
		OU      *string `json:"OU,omitempty"`
		ST      *string `json:"ST,omitempty"`
		AltName *string `json:"altName,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.C = all.C
	o.CN = all.CN
	o.L = all.L
	o.O = all.O
	o.OU = all.OU
	o.ST = all.ST
	o.AltName = all.AltName
	return nil
}
