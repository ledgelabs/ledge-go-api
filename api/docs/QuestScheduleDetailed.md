# QuestScheduleDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestId** | **string** |  | 
**Processed** | **bool** |  | 
**Recurring** | **bool** |  | 
**EndTime** | **NullableTime** |  | 
**StartTime** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Quest** | [**Quest**](Quest.md) |  | 

## Methods

### NewQuestScheduleDetailed

`func NewQuestScheduleDetailed(questId string, processed bool, recurring bool, endTime NullableTime, startTime time.Time, updatedAt time.Time, createdAt time.Time, id string, quest Quest, ) *QuestScheduleDetailed`

NewQuestScheduleDetailed instantiates a new QuestScheduleDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestScheduleDetailedWithDefaults

`func NewQuestScheduleDetailedWithDefaults() *QuestScheduleDetailed`

NewQuestScheduleDetailedWithDefaults instantiates a new QuestScheduleDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestId

`func (o *QuestScheduleDetailed) GetQuestId() string`

GetQuestId returns the QuestId field if non-nil, zero value otherwise.

### GetQuestIdOk

`func (o *QuestScheduleDetailed) GetQuestIdOk() (*string, bool)`

GetQuestIdOk returns a tuple with the QuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestId

`func (o *QuestScheduleDetailed) SetQuestId(v string)`

SetQuestId sets QuestId field to given value.


### GetProcessed

`func (o *QuestScheduleDetailed) GetProcessed() bool`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *QuestScheduleDetailed) GetProcessedOk() (*bool, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *QuestScheduleDetailed) SetProcessed(v bool)`

SetProcessed sets Processed field to given value.


### GetRecurring

`func (o *QuestScheduleDetailed) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *QuestScheduleDetailed) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *QuestScheduleDetailed) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.


### GetEndTime

`func (o *QuestScheduleDetailed) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *QuestScheduleDetailed) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *QuestScheduleDetailed) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *QuestScheduleDetailed) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *QuestScheduleDetailed) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStartTime

`func (o *QuestScheduleDetailed) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *QuestScheduleDetailed) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *QuestScheduleDetailed) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetUpdatedAt

`func (o *QuestScheduleDetailed) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QuestScheduleDetailed) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QuestScheduleDetailed) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *QuestScheduleDetailed) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuestScheduleDetailed) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuestScheduleDetailed) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *QuestScheduleDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestScheduleDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestScheduleDetailed) SetId(v string)`

SetId sets Id field to given value.


### GetQuest

`func (o *QuestScheduleDetailed) GetQuest() Quest`

GetQuest returns the Quest field if non-nil, zero value otherwise.

### GetQuestOk

`func (o *QuestScheduleDetailed) GetQuestOk() (*Quest, bool)`

GetQuestOk returns a tuple with the Quest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuest

`func (o *QuestScheduleDetailed) SetQuest(v Quest)`

SetQuest sets Quest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


