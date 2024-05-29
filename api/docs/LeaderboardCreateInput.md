# LeaderboardCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AltScoreTextAlias** | Pointer to **string** |  | [optional] 
**ScoreTextAlias** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewLeaderboardCreateInput

`func NewLeaderboardCreateInput(scoreTextAlias string, title string, ) *LeaderboardCreateInput`

NewLeaderboardCreateInput instantiates a new LeaderboardCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardCreateInputWithDefaults

`func NewLeaderboardCreateInputWithDefaults() *LeaderboardCreateInput`

NewLeaderboardCreateInputWithDefaults instantiates a new LeaderboardCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageUrl

`func (o *LeaderboardCreateInput) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *LeaderboardCreateInput) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *LeaderboardCreateInput) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *LeaderboardCreateInput) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetDescription

`func (o *LeaderboardCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LeaderboardCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LeaderboardCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LeaderboardCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAltScoreTextAlias

`func (o *LeaderboardCreateInput) GetAltScoreTextAlias() string`

GetAltScoreTextAlias returns the AltScoreTextAlias field if non-nil, zero value otherwise.

### GetAltScoreTextAliasOk

`func (o *LeaderboardCreateInput) GetAltScoreTextAliasOk() (*string, bool)`

GetAltScoreTextAliasOk returns a tuple with the AltScoreTextAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltScoreTextAlias

`func (o *LeaderboardCreateInput) SetAltScoreTextAlias(v string)`

SetAltScoreTextAlias sets AltScoreTextAlias field to given value.

### HasAltScoreTextAlias

`func (o *LeaderboardCreateInput) HasAltScoreTextAlias() bool`

HasAltScoreTextAlias returns a boolean if a field has been set.

### GetScoreTextAlias

`func (o *LeaderboardCreateInput) GetScoreTextAlias() string`

GetScoreTextAlias returns the ScoreTextAlias field if non-nil, zero value otherwise.

### GetScoreTextAliasOk

`func (o *LeaderboardCreateInput) GetScoreTextAliasOk() (*string, bool)`

GetScoreTextAliasOk returns a tuple with the ScoreTextAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreTextAlias

`func (o *LeaderboardCreateInput) SetScoreTextAlias(v string)`

SetScoreTextAlias sets ScoreTextAlias field to given value.


### GetTitle

`func (o *LeaderboardCreateInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LeaderboardCreateInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LeaderboardCreateInput) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


