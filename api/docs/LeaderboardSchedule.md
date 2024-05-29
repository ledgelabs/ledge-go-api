# LeaderboardSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaderboardId** | **string** |  | 
**RewardsProcessed** | **bool** |  | 
**ScheduleProcessed** | **bool** |  | 
**Recurring** | **bool** |  | 
**EndTime** | **NullableTime** |  | 
**StartTime** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 

## Methods

### NewLeaderboardSchedule

`func NewLeaderboardSchedule(leaderboardId string, rewardsProcessed bool, scheduleProcessed bool, recurring bool, endTime NullableTime, startTime time.Time, updatedAt time.Time, createdAt time.Time, id string, ) *LeaderboardSchedule`

NewLeaderboardSchedule instantiates a new LeaderboardSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardScheduleWithDefaults

`func NewLeaderboardScheduleWithDefaults() *LeaderboardSchedule`

NewLeaderboardScheduleWithDefaults instantiates a new LeaderboardSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaderboardId

`func (o *LeaderboardSchedule) GetLeaderboardId() string`

GetLeaderboardId returns the LeaderboardId field if non-nil, zero value otherwise.

### GetLeaderboardIdOk

`func (o *LeaderboardSchedule) GetLeaderboardIdOk() (*string, bool)`

GetLeaderboardIdOk returns a tuple with the LeaderboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboardId

`func (o *LeaderboardSchedule) SetLeaderboardId(v string)`

SetLeaderboardId sets LeaderboardId field to given value.


### GetRewardsProcessed

`func (o *LeaderboardSchedule) GetRewardsProcessed() bool`

GetRewardsProcessed returns the RewardsProcessed field if non-nil, zero value otherwise.

### GetRewardsProcessedOk

`func (o *LeaderboardSchedule) GetRewardsProcessedOk() (*bool, bool)`

GetRewardsProcessedOk returns a tuple with the RewardsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsProcessed

`func (o *LeaderboardSchedule) SetRewardsProcessed(v bool)`

SetRewardsProcessed sets RewardsProcessed field to given value.


### GetScheduleProcessed

`func (o *LeaderboardSchedule) GetScheduleProcessed() bool`

GetScheduleProcessed returns the ScheduleProcessed field if non-nil, zero value otherwise.

### GetScheduleProcessedOk

`func (o *LeaderboardSchedule) GetScheduleProcessedOk() (*bool, bool)`

GetScheduleProcessedOk returns a tuple with the ScheduleProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleProcessed

`func (o *LeaderboardSchedule) SetScheduleProcessed(v bool)`

SetScheduleProcessed sets ScheduleProcessed field to given value.


### GetRecurring

`func (o *LeaderboardSchedule) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *LeaderboardSchedule) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *LeaderboardSchedule) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.


### GetEndTime

`func (o *LeaderboardSchedule) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LeaderboardSchedule) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LeaderboardSchedule) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *LeaderboardSchedule) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *LeaderboardSchedule) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStartTime

`func (o *LeaderboardSchedule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LeaderboardSchedule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LeaderboardSchedule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetUpdatedAt

`func (o *LeaderboardSchedule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LeaderboardSchedule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LeaderboardSchedule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *LeaderboardSchedule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LeaderboardSchedule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LeaderboardSchedule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *LeaderboardSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeaderboardSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeaderboardSchedule) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


