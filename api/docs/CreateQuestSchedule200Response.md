# CreateQuestSchedule200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestSchedule** | Pointer to [**QuestSchedule**](QuestSchedule.md) |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewCreateQuestSchedule200Response

`func NewCreateQuestSchedule200Response(message string, ) *CreateQuestSchedule200Response`

NewCreateQuestSchedule200Response instantiates a new CreateQuestSchedule200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuestSchedule200ResponseWithDefaults

`func NewCreateQuestSchedule200ResponseWithDefaults() *CreateQuestSchedule200Response`

NewCreateQuestSchedule200ResponseWithDefaults instantiates a new CreateQuestSchedule200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestSchedule

`func (o *CreateQuestSchedule200Response) GetQuestSchedule() QuestSchedule`

GetQuestSchedule returns the QuestSchedule field if non-nil, zero value otherwise.

### GetQuestScheduleOk

`func (o *CreateQuestSchedule200Response) GetQuestScheduleOk() (*QuestSchedule, bool)`

GetQuestScheduleOk returns a tuple with the QuestSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestSchedule

`func (o *CreateQuestSchedule200Response) SetQuestSchedule(v QuestSchedule)`

SetQuestSchedule sets QuestSchedule field to given value.

### HasQuestSchedule

`func (o *CreateQuestSchedule200Response) HasQuestSchedule() bool`

HasQuestSchedule returns a boolean if a field has been set.

### GetMessage

`func (o *CreateQuestSchedule200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateQuestSchedule200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateQuestSchedule200Response) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


