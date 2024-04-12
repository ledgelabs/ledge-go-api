# CreateQuestScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recurring** | **bool** |  | 
**QuestId** | **string** |  | 
**EndTime** | **time.Time** |  | 
**StartTime** | **time.Time** |  | 

## Methods

### NewCreateQuestScheduleRequest

`func NewCreateQuestScheduleRequest(recurring bool, questId string, endTime time.Time, startTime time.Time, ) *CreateQuestScheduleRequest`

NewCreateQuestScheduleRequest instantiates a new CreateQuestScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuestScheduleRequestWithDefaults

`func NewCreateQuestScheduleRequestWithDefaults() *CreateQuestScheduleRequest`

NewCreateQuestScheduleRequestWithDefaults instantiates a new CreateQuestScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurring

`func (o *CreateQuestScheduleRequest) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *CreateQuestScheduleRequest) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *CreateQuestScheduleRequest) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.


### GetQuestId

`func (o *CreateQuestScheduleRequest) GetQuestId() string`

GetQuestId returns the QuestId field if non-nil, zero value otherwise.

### GetQuestIdOk

`func (o *CreateQuestScheduleRequest) GetQuestIdOk() (*string, bool)`

GetQuestIdOk returns a tuple with the QuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestId

`func (o *CreateQuestScheduleRequest) SetQuestId(v string)`

SetQuestId sets QuestId field to given value.


### GetEndTime

`func (o *CreateQuestScheduleRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateQuestScheduleRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateQuestScheduleRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetStartTime

`func (o *CreateQuestScheduleRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateQuestScheduleRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateQuestScheduleRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


