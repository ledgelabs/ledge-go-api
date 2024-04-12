# CreateGoalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | **string** |  | 
**QuestId** | **string** |  | 
**BannerUrl** | Pointer to **string** |  | [optional] 
**Target** | **float64** |  | 
**Instruction** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewCreateGoalRequest

`func NewCreateGoalRequest(activity string, questId string, target float64, description string, title string, ) *CreateGoalRequest`

NewCreateGoalRequest instantiates a new CreateGoalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGoalRequestWithDefaults

`func NewCreateGoalRequestWithDefaults() *CreateGoalRequest`

NewCreateGoalRequestWithDefaults instantiates a new CreateGoalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *CreateGoalRequest) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *CreateGoalRequest) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *CreateGoalRequest) SetActivity(v string)`

SetActivity sets Activity field to given value.


### GetQuestId

`func (o *CreateGoalRequest) GetQuestId() string`

GetQuestId returns the QuestId field if non-nil, zero value otherwise.

### GetQuestIdOk

`func (o *CreateGoalRequest) GetQuestIdOk() (*string, bool)`

GetQuestIdOk returns a tuple with the QuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestId

`func (o *CreateGoalRequest) SetQuestId(v string)`

SetQuestId sets QuestId field to given value.


### GetBannerUrl

`func (o *CreateGoalRequest) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *CreateGoalRequest) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *CreateGoalRequest) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *CreateGoalRequest) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### GetTarget

`func (o *CreateGoalRequest) GetTarget() float64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateGoalRequest) GetTargetOk() (*float64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateGoalRequest) SetTarget(v float64)`

SetTarget sets Target field to given value.


### GetInstruction

`func (o *CreateGoalRequest) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *CreateGoalRequest) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *CreateGoalRequest) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *CreateGoalRequest) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetDescription

`func (o *CreateGoalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGoalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGoalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTitle

`func (o *CreateGoalRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGoalRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGoalRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


