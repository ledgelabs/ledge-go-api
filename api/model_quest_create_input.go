/*
@ledge/external-api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgeapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the QuestCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestCreateInput{}

// QuestCreateInput struct for QuestCreateInput
type QuestCreateInput struct {
	Type QuestType `json:"type"`
	Description string `json:"description"`
	Title string `json:"title"`
}

type _QuestCreateInput QuestCreateInput

// NewQuestCreateInput instantiates a new QuestCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestCreateInput(type_ QuestType, description string, title string) *QuestCreateInput {
	this := QuestCreateInput{}
	this.Type = type_
	this.Description = description
	this.Title = title
	return &this
}

// NewQuestCreateInputWithDefaults instantiates a new QuestCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestCreateInputWithDefaults() *QuestCreateInput {
	this := QuestCreateInput{}
	return &this
}

// GetType returns the Type field value
func (o *QuestCreateInput) GetType() QuestType {
	if o == nil {
		var ret QuestType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *QuestCreateInput) GetTypeOk() (*QuestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *QuestCreateInput) SetType(v QuestType) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *QuestCreateInput) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *QuestCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *QuestCreateInput) SetDescription(v string) {
	o.Description = v
}

// GetTitle returns the Title field value
func (o *QuestCreateInput) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *QuestCreateInput) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *QuestCreateInput) SetTitle(v string) {
	o.Title = v
}

func (o QuestCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["description"] = o.Description
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *QuestCreateInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"description",
		"title",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQuestCreateInput := _QuestCreateInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuestCreateInput)

	if err != nil {
		return err
	}

	*o = QuestCreateInput(varQuestCreateInput)

	return err
}

type NullableQuestCreateInput struct {
	value *QuestCreateInput
	isSet bool
}

func (v NullableQuestCreateInput) Get() *QuestCreateInput {
	return v.value
}

func (v *NullableQuestCreateInput) Set(val *QuestCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestCreateInput(val *QuestCreateInput) *NullableQuestCreateInput {
	return &NullableQuestCreateInput{value: val, isSet: true}
}

func (v NullableQuestCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

