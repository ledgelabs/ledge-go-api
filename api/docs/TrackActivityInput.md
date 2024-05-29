# TrackActivityInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Occurrence** | **string** |  | 
**Count** | Pointer to **float64** |  | [optional] 
**ActivityId** | **string** |  | 
**UserId** | **string** |  | 

## Methods

### NewTrackActivityInput

`func NewTrackActivityInput(occurrence string, activityId string, userId string, ) *TrackActivityInput`

NewTrackActivityInput instantiates a new TrackActivityInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackActivityInputWithDefaults

`func NewTrackActivityInputWithDefaults() *TrackActivityInput`

NewTrackActivityInputWithDefaults instantiates a new TrackActivityInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurrence

`func (o *TrackActivityInput) GetOccurrence() string`

GetOccurrence returns the Occurrence field if non-nil, zero value otherwise.

### GetOccurrenceOk

`func (o *TrackActivityInput) GetOccurrenceOk() (*string, bool)`

GetOccurrenceOk returns a tuple with the Occurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrence

`func (o *TrackActivityInput) SetOccurrence(v string)`

SetOccurrence sets Occurrence field to given value.


### GetCount

`func (o *TrackActivityInput) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TrackActivityInput) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TrackActivityInput) SetCount(v float64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TrackActivityInput) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetActivityId

`func (o *TrackActivityInput) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *TrackActivityInput) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *TrackActivityInput) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetUserId

`func (o *TrackActivityInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TrackActivityInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TrackActivityInput) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


