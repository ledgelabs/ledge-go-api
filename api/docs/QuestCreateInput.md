# QuestCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**QuestType**](QuestType.md) |  | 
**Description** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewQuestCreateInput

`func NewQuestCreateInput(type_ QuestType, description string, title string, ) *QuestCreateInput`

NewQuestCreateInput instantiates a new QuestCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestCreateInputWithDefaults

`func NewQuestCreateInputWithDefaults() *QuestCreateInput`

NewQuestCreateInputWithDefaults instantiates a new QuestCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QuestCreateInput) GetType() QuestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuestCreateInput) GetTypeOk() (*QuestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuestCreateInput) SetType(v QuestType)`

SetType sets Type field to given value.


### GetDescription

`func (o *QuestCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuestCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuestCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTitle

`func (o *QuestCreateInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QuestCreateInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QuestCreateInput) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


