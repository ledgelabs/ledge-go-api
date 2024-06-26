/*
@ledge/external-api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgeapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the QuestSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestSchedule{}

// QuestSchedule Model QuestSchedule
type QuestSchedule struct {
	QuestId string `json:"questId"`
	Processed bool `json:"processed"`
	Recurring bool `json:"recurring"`
	EndTime NullableTime `json:"endTime"`
	StartTime time.Time `json:"startTime"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	Id string `json:"id"`
}

type _QuestSchedule QuestSchedule

// NewQuestSchedule instantiates a new QuestSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestSchedule(questId string, processed bool, recurring bool, endTime NullableTime, startTime time.Time, updatedAt time.Time, createdAt time.Time, id string) *QuestSchedule {
	this := QuestSchedule{}
	this.QuestId = questId
	this.Processed = processed
	this.Recurring = recurring
	this.EndTime = endTime
	this.StartTime = startTime
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	this.Id = id
	return &this
}

// NewQuestScheduleWithDefaults instantiates a new QuestSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestScheduleWithDefaults() *QuestSchedule {
	this := QuestSchedule{}
	return &this
}

// GetQuestId returns the QuestId field value
func (o *QuestSchedule) GetQuestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuestId
}

// GetQuestIdOk returns a tuple with the QuestId field value
// and a boolean to check if the value has been set.
func (o *QuestSchedule) GetQuestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuestId, true
}

// SetQuestId sets field value
func (o *QuestSchedule) SetQuestId(v string) {
	o.QuestId = v
}

// GetProcessed returns the Processed field value
func (o *QuestSchedule) GetProcessed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Processed
}

// GetProcessedOk returns a tuple with the Processed field value
// and a boolean to check if the value has been set.
func (o *QuestSchedule) GetProcessedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Processed, true
}

// SetProcessed sets field value
func (o *QuestSchedule) SetProcessed(v bool) {
	o.Processed = v
}

// GetRecurring returns the Recurring field value
func (o *QuestSchedule) GetRecurring() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value
// and a boolean to check if the value has been set.
func (o *QuestSchedule) GetRecurringOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurring, true
}

// SetRecurring sets field value
func (o *QuestSchedule) SetRecurring(v bool) {
	o.Recurring = v
}

// GetEndTime returns the EndTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *QuestSchedule) GetEndTime() time.Time {
	if o == nil || o.EndTime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuestSchedule) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// SetEndTime sets field value
func (o *QuestSchedule) SetEndTime(v time.Time) {
	o.EndTime.Set(&v)
}

// GetStartTime returns the StartTime field value
func (o *QuestSchedule) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *QuestSchedule) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *QuestSchedule) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *QuestSchedule) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *QuestSchedule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *QuestSchedule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *QuestSchedule) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *QuestSchedule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *QuestSchedule) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *QuestSchedule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QuestSchedule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QuestSchedule) SetId(v string) {
	o.Id = v
}

func (o QuestSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["questId"] = o.QuestId
	toSerialize["processed"] = o.Processed
	toSerialize["recurring"] = o.Recurring
	toSerialize["endTime"] = o.EndTime.Get()
	toSerialize["startTime"] = o.StartTime
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *QuestSchedule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"questId",
		"processed",
		"recurring",
		"endTime",
		"startTime",
		"updatedAt",
		"createdAt",
		"id",
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

	varQuestSchedule := _QuestSchedule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuestSchedule)

	if err != nil {
		return err
	}

	*o = QuestSchedule(varQuestSchedule)

	return err
}

type NullableQuestSchedule struct {
	value *QuestSchedule
	isSet bool
}

func (v NullableQuestSchedule) Get() *QuestSchedule {
	return v.value
}

func (v *NullableQuestSchedule) Set(val *QuestSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestSchedule(val *QuestSchedule) *NullableQuestSchedule {
	return &NullableQuestSchedule{value: val, isSet: true}
}

func (v NullableQuestSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


