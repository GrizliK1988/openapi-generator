/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"encoding/json"
)

// ClassModel Model for testing model with \"_class\" property
type ClassModel struct {
	Class *string `json:"_class,omitempty"`
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *ClassModel) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ClassModel) GetClassOk() (string, bool) {
	if o == nil || o.Class == nil {
		var ret string
		return ret, false
	}
	return *o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *ClassModel) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *ClassModel) SetClass(v string) {
	o.Class = &v
}

type NullableClassModel struct {
	Value ClassModel
	ExplicitNull bool
}

func (v NullableClassModel) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableClassModel) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

