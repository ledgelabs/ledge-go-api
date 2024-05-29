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

// checks if the Leaderboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Leaderboard{}

// Leaderboard Model Leaderboard
type Leaderboard struct {
	GameId string `json:"gameId"`
	AltScoreTextAlias NullableString `json:"altScoreTextAlias"`
	ScoreTextAlias string `json:"scoreTextAlias"`
	ImageUrl NullableString `json:"imageUrl"`
	Description NullableString `json:"description"`
	Title string `json:"title"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	Id string `json:"id"`
}

type _Leaderboard Leaderboard

// NewLeaderboard instantiates a new Leaderboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeaderboard(gameId string, altScoreTextAlias NullableString, scoreTextAlias string, imageUrl NullableString, description NullableString, title string, updatedAt time.Time, createdAt time.Time, id string) *Leaderboard {
	this := Leaderboard{}
	this.GameId = gameId
	this.AltScoreTextAlias = altScoreTextAlias
	this.ScoreTextAlias = scoreTextAlias
	this.ImageUrl = imageUrl
	this.Description = description
	this.Title = title
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	this.Id = id
	return &this
}

// NewLeaderboardWithDefaults instantiates a new Leaderboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeaderboardWithDefaults() *Leaderboard {
	this := Leaderboard{}
	return &this
}

// GetGameId returns the GameId field value
func (o *Leaderboard) GetGameId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value
// and a boolean to check if the value has been set.
func (o *Leaderboard) GetGameIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameId, true
}

// SetGameId sets field value
func (o *Leaderboard) SetGameId(v string) {
	o.GameId = v
}

// GetAltScoreTextAlias returns the AltScoreTextAlias field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Leaderboard) GetAltScoreTextAlias() string {
	if o == nil || o.AltScoreTextAlias.Get() == nil {
		var ret string
		return ret
	}

	return *o.AltScoreTextAlias.Get()
}

// GetAltScoreTextAliasOk returns a tuple with the AltScoreTextAlias field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Leaderboard) GetAltScoreTextAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AltScoreTextAlias.Get(), o.AltScoreTextAlias.IsSet()
}

// SetAltScoreTextAlias sets field value
func (o *Leaderboard) SetAltScoreTextAlias(v string) {
	o.AltScoreTextAlias.Set(&v)
}

// GetScoreTextAlias returns the ScoreTextAlias field value
func (o *Leaderboard) GetScoreTextAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScoreTextAlias
}

// GetScoreTextAliasOk returns a tuple with the ScoreTextAlias field value
// and a boolean to check if the value has been set.
func (o *Leaderboard) GetScoreTextAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScoreTextAlias, true
}

// SetScoreTextAlias sets field value
func (o *Leaderboard) SetScoreTextAlias(v string) {
	o.ScoreTextAlias = v
}

// GetImageUrl returns the ImageUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Leaderboard) GetImageUrl() string {
	if o == nil || o.ImageUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Leaderboard) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// SetImageUrl sets field value
func (o *Leaderboard) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Leaderboard) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Leaderboard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Leaderboard) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetTitle returns the Title field value
func (o *Leaderboard) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Leaderboard) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Leaderboard) SetTitle(v string) {
	o.Title = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Leaderboard) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Leaderboard) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Leaderboard) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Leaderboard) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Leaderboard) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Leaderboard) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *Leaderboard) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Leaderboard) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Leaderboard) SetId(v string) {
	o.Id = v
}

func (o Leaderboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Leaderboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gameId"] = o.GameId
	toSerialize["altScoreTextAlias"] = o.AltScoreTextAlias.Get()
	toSerialize["scoreTextAlias"] = o.ScoreTextAlias
	toSerialize["imageUrl"] = o.ImageUrl.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["title"] = o.Title
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *Leaderboard) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gameId",
		"altScoreTextAlias",
		"scoreTextAlias",
		"imageUrl",
		"description",
		"title",
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

	varLeaderboard := _Leaderboard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLeaderboard)

	if err != nil {
		return err
	}

	*o = Leaderboard(varLeaderboard)

	return err
}

type NullableLeaderboard struct {
	value *Leaderboard
	isSet bool
}

func (v NullableLeaderboard) Get() *Leaderboard {
	return v.value
}

func (v *NullableLeaderboard) Set(val *Leaderboard) {
	v.value = val
	v.isSet = true
}

func (v NullableLeaderboard) IsSet() bool {
	return v.isSet
}

func (v *NullableLeaderboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeaderboard(val *Leaderboard) *NullableLeaderboard {
	return &NullableLeaderboard{value: val, isSet: true}
}

func (v NullableLeaderboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeaderboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


