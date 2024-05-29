# QuestScheduleCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recurring** | Pointer to **bool** |  | [optional] 
**QuestId** | **string** |  | 
**EndTime** | **time.Time** |  | 
**StartTime** | **time.Time** |  | 

## Methods

### NewQuestScheduleCreateInput

`func NewQuestScheduleCreateInput(questId string, endTime time.Time, startTime time.Time, ) *QuestScheduleCreateInput`

NewQuestScheduleCreateInput instantiates a new QuestScheduleCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestScheduleCreateInputWithDefaults

`func NewQuestScheduleCreateInputWithDefaults() *QuestScheduleCreateInput`

NewQuestScheduleCreateInputWithDefaults instantiates a new QuestScheduleCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurring

`func (o *QuestScheduleCreateInput) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *QuestScheduleCreateInput) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *QuestScheduleCreateInput) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *QuestScheduleCreateInput) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetQuestId

`func (o *QuestScheduleCreateInput) GetQuestId() string`

GetQuestId returns the QuestId field if non-nil, zero value otherwise.

### GetQuestIdOk

`func (o *QuestScheduleCreateInput) GetQuestIdOk() (*string, bool)`

GetQuestIdOk returns a tuple with the QuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestId

`func (o *QuestScheduleCreateInput) SetQuestId(v string)`

SetQuestId sets QuestId field to given value.


### GetEndTime

`func (o *QuestScheduleCreateInput) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *QuestScheduleCreateInput) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *QuestScheduleCreateInput) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetStartTime

`func (o *QuestScheduleCreateInput) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *QuestScheduleCreateInput) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *QuestScheduleCreateInput) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


