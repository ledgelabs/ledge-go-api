# QuestUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**QuestType**](QuestType.md) |  | 
**Description** | **string** |  | 
**Title** | **string** |  | 
**QuestId** | **string** |  | 

## Methods

### NewQuestUpdateInput

`func NewQuestUpdateInput(type_ QuestType, description string, title string, questId string, ) *QuestUpdateInput`

NewQuestUpdateInput instantiates a new QuestUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestUpdateInputWithDefaults

`func NewQuestUpdateInputWithDefaults() *QuestUpdateInput`

NewQuestUpdateInputWithDefaults instantiates a new QuestUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QuestUpdateInput) GetType() QuestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuestUpdateInput) GetTypeOk() (*QuestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuestUpdateInput) SetType(v QuestType)`

SetType sets Type field to given value.


### GetDescription

`func (o *QuestUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuestUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuestUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTitle

`func (o *QuestUpdateInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QuestUpdateInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QuestUpdateInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetQuestId

`func (o *QuestUpdateInput) GetQuestId() string`

GetQuestId returns the QuestId field if non-nil, zero value otherwise.

### GetQuestIdOk

`func (o *QuestUpdateInput) GetQuestIdOk() (*string, bool)`

GetQuestIdOk returns a tuple with the QuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestId

`func (o *QuestUpdateInput) SetQuestId(v string)`

SetQuestId sets QuestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


