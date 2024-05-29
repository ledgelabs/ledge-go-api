# Leaderboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GameId** | **string** |  | 
**AltScoreTextAlias** | **NullableString** |  | 
**ScoreTextAlias** | **string** |  | 
**ImageUrl** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**Title** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 

## Methods

### NewLeaderboard

`func NewLeaderboard(gameId string, altScoreTextAlias NullableString, scoreTextAlias string, imageUrl NullableString, description NullableString, title string, updatedAt time.Time, createdAt time.Time, id string, ) *Leaderboard`

NewLeaderboard instantiates a new Leaderboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardWithDefaults

`func NewLeaderboardWithDefaults() *Leaderboard`

NewLeaderboardWithDefaults instantiates a new Leaderboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGameId

`func (o *Leaderboard) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *Leaderboard) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *Leaderboard) SetGameId(v string)`

SetGameId sets GameId field to given value.


### GetAltScoreTextAlias

`func (o *Leaderboard) GetAltScoreTextAlias() string`

GetAltScoreTextAlias returns the AltScoreTextAlias field if non-nil, zero value otherwise.

### GetAltScoreTextAliasOk

`func (o *Leaderboard) GetAltScoreTextAliasOk() (*string, bool)`

GetAltScoreTextAliasOk returns a tuple with the AltScoreTextAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltScoreTextAlias

`func (o *Leaderboard) SetAltScoreTextAlias(v string)`

SetAltScoreTextAlias sets AltScoreTextAlias field to given value.


### SetAltScoreTextAliasNil

`func (o *Leaderboard) SetAltScoreTextAliasNil(b bool)`

 SetAltScoreTextAliasNil sets the value for AltScoreTextAlias to be an explicit nil

### UnsetAltScoreTextAlias
`func (o *Leaderboard) UnsetAltScoreTextAlias()`

UnsetAltScoreTextAlias ensures that no value is present for AltScoreTextAlias, not even an explicit nil
### GetScoreTextAlias

`func (o *Leaderboard) GetScoreTextAlias() string`

GetScoreTextAlias returns the ScoreTextAlias field if non-nil, zero value otherwise.

### GetScoreTextAliasOk

`func (o *Leaderboard) GetScoreTextAliasOk() (*string, bool)`

GetScoreTextAliasOk returns a tuple with the ScoreTextAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreTextAlias

`func (o *Leaderboard) SetScoreTextAlias(v string)`

SetScoreTextAlias sets ScoreTextAlias field to given value.


### GetImageUrl

`func (o *Leaderboard) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Leaderboard) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Leaderboard) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### SetImageUrlNil

`func (o *Leaderboard) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *Leaderboard) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDescription

`func (o *Leaderboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Leaderboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Leaderboard) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Leaderboard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Leaderboard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTitle

`func (o *Leaderboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Leaderboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Leaderboard) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *Leaderboard) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Leaderboard) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Leaderboard) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Leaderboard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Leaderboard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Leaderboard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *Leaderboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Leaderboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Leaderboard) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


