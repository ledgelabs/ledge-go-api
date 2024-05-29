# GoalCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannerUrl** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActivityId** | **string** |  | 
**QuestId** | **string** |  | 
**Target** | **float64** |  | 
**Title** | **string** |  | 

## Methods

### NewGoalCreateInput

`func NewGoalCreateInput(activityId string, questId string, target float64, title string, ) *GoalCreateInput`

NewGoalCreateInput instantiates a new GoalCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalCreateInputWithDefaults

`func NewGoalCreateInputWithDefaults() *GoalCreateInput`

NewGoalCreateInputWithDefaults instantiates a new GoalCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannerUrl

`func (o *GoalCreateInput) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *GoalCreateInput) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *GoalCreateInput) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *GoalCreateInput) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### GetInstructions

`func (o *GoalCreateInput) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *GoalCreateInput) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *GoalCreateInput) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *GoalCreateInput) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetDescription

`func (o *GoalCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GoalCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GoalCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GoalCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActivityId

`func (o *GoalCreateInput) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *GoalCreateInput) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *GoalCreateInput) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetQuestId

`func (o *GoalCreateInput) GetQuestId() string`

GetQuestId returns the QuestId field if non-nil, zero value otherwise.

### GetQuestIdOk

`func (o *GoalCreateInput) GetQuestIdOk() (*string, bool)`

GetQuestIdOk returns a tuple with the QuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestId

`func (o *GoalCreateInput) SetQuestId(v string)`

SetQuestId sets QuestId field to given value.


### GetTarget

`func (o *GoalCreateInput) GetTarget() float64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GoalCreateInput) GetTargetOk() (*float64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GoalCreateInput) SetTarget(v float64)`

SetTarget sets Target field to given value.


### GetTitle

`func (o *GoalCreateInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GoalCreateInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GoalCreateInput) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


