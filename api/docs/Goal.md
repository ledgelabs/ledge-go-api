# Goal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestId** | **string** |  | 
**ObjectId** | **NullableString** |  | 
**BannerUrl** | **NullableString** |  | 
**Instructions** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**Activity** | [**ActivityType**](ActivityType.md) |  | 
**Target** | **float64** |  | 
**Title** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 

## Methods

### NewGoal

`func NewGoal(questId string, objectId NullableString, bannerUrl NullableString, instructions NullableString, description NullableString, activity ActivityType, target float64, title string, updatedAt time.Time, createdAt time.Time, id string, ) *Goal`

NewGoal instantiates a new Goal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalWithDefaults

`func NewGoalWithDefaults() *Goal`

NewGoalWithDefaults instantiates a new Goal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestId

`func (o *Goal) GetQuestId() string`

GetQuestId returns the QuestId field if non-nil, zero value otherwise.

### GetQuestIdOk

`func (o *Goal) GetQuestIdOk() (*string, bool)`

GetQuestIdOk returns a tuple with the QuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestId

`func (o *Goal) SetQuestId(v string)`

SetQuestId sets QuestId field to given value.


### GetObjectId

`func (o *Goal) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *Goal) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *Goal) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### SetObjectIdNil

`func (o *Goal) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *Goal) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetBannerUrl

`func (o *Goal) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *Goal) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *Goal) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.


### SetBannerUrlNil

`func (o *Goal) SetBannerUrlNil(b bool)`

 SetBannerUrlNil sets the value for BannerUrl to be an explicit nil

### UnsetBannerUrl
`func (o *Goal) UnsetBannerUrl()`

UnsetBannerUrl ensures that no value is present for BannerUrl, not even an explicit nil
### GetInstructions

`func (o *Goal) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Goal) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Goal) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### SetInstructionsNil

`func (o *Goal) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *Goal) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetDescription

`func (o *Goal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Goal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Goal) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Goal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Goal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetActivity

`func (o *Goal) GetActivity() ActivityType`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Goal) GetActivityOk() (*ActivityType, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Goal) SetActivity(v ActivityType)`

SetActivity sets Activity field to given value.


### GetTarget

`func (o *Goal) GetTarget() float64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Goal) GetTargetOk() (*float64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Goal) SetTarget(v float64)`

SetTarget sets Target field to given value.


### GetTitle

`func (o *Goal) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Goal) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Goal) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *Goal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Goal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Goal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Goal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Goal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Goal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *Goal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Goal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Goal) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


