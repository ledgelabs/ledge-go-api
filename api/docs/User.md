# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogin** | **NullableTime** |  | 
**Verified** | **bool** |  | 
**MergedWith** | **NullableString** |  | 
**GameId** | **NullableString** |  | 
**ExternalId** | **NullableString** |  | 
**ReferredById** | **NullableString** |  | 
**EnableNotifications** | **bool** |  | 
**HasAcceptedLegal** | **bool** |  | 
**HasOnboarded** | **bool** |  | 
**Gender** | [**Gender**](Gender.md) |  | 
**BirthYear** | **float64** |  | 
**Avatar** | **string** |  | 
**RemainingReferrals** | **float64** |  | 
**Code** | **string** |  | 
**Usertag** | **float64** |  | 
**Username** | **string** |  | 
**Name** | **string** |  | 
**PhoneNumber** | **NullableString** |  | 
**Email** | **NullableString** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 

## Methods

### NewUser

`func NewUser(lastLogin NullableTime, verified bool, mergedWith NullableString, gameId NullableString, externalId NullableString, referredById NullableString, enableNotifications bool, hasAcceptedLegal bool, hasOnboarded bool, gender Gender, birthYear float64, avatar string, remainingReferrals float64, code string, usertag float64, username string, name string, phoneNumber NullableString, email NullableString, updatedAt time.Time, createdAt time.Time, id string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogin

`func (o *User) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.


### SetLastLoginNil

`func (o *User) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *User) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetVerified

`func (o *User) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *User) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *User) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetMergedWith

`func (o *User) GetMergedWith() string`

GetMergedWith returns the MergedWith field if non-nil, zero value otherwise.

### GetMergedWithOk

`func (o *User) GetMergedWithOk() (*string, bool)`

GetMergedWithOk returns a tuple with the MergedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedWith

`func (o *User) SetMergedWith(v string)`

SetMergedWith sets MergedWith field to given value.


### SetMergedWithNil

`func (o *User) SetMergedWithNil(b bool)`

 SetMergedWithNil sets the value for MergedWith to be an explicit nil

### UnsetMergedWith
`func (o *User) UnsetMergedWith()`

UnsetMergedWith ensures that no value is present for MergedWith, not even an explicit nil
### GetGameId

`func (o *User) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *User) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *User) SetGameId(v string)`

SetGameId sets GameId field to given value.


### SetGameIdNil

`func (o *User) SetGameIdNil(b bool)`

 SetGameIdNil sets the value for GameId to be an explicit nil

### UnsetGameId
`func (o *User) UnsetGameId()`

UnsetGameId ensures that no value is present for GameId, not even an explicit nil
### GetExternalId

`func (o *User) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *User) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *User) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *User) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *User) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetReferredById

`func (o *User) GetReferredById() string`

GetReferredById returns the ReferredById field if non-nil, zero value otherwise.

### GetReferredByIdOk

`func (o *User) GetReferredByIdOk() (*string, bool)`

GetReferredByIdOk returns a tuple with the ReferredById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredById

`func (o *User) SetReferredById(v string)`

SetReferredById sets ReferredById field to given value.


### SetReferredByIdNil

`func (o *User) SetReferredByIdNil(b bool)`

 SetReferredByIdNil sets the value for ReferredById to be an explicit nil

### UnsetReferredById
`func (o *User) UnsetReferredById()`

UnsetReferredById ensures that no value is present for ReferredById, not even an explicit nil
### GetEnableNotifications

`func (o *User) GetEnableNotifications() bool`

GetEnableNotifications returns the EnableNotifications field if non-nil, zero value otherwise.

### GetEnableNotificationsOk

`func (o *User) GetEnableNotificationsOk() (*bool, bool)`

GetEnableNotificationsOk returns a tuple with the EnableNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifications

`func (o *User) SetEnableNotifications(v bool)`

SetEnableNotifications sets EnableNotifications field to given value.


### GetHasAcceptedLegal

`func (o *User) GetHasAcceptedLegal() bool`

GetHasAcceptedLegal returns the HasAcceptedLegal field if non-nil, zero value otherwise.

### GetHasAcceptedLegalOk

`func (o *User) GetHasAcceptedLegalOk() (*bool, bool)`

GetHasAcceptedLegalOk returns a tuple with the HasAcceptedLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAcceptedLegal

`func (o *User) SetHasAcceptedLegal(v bool)`

SetHasAcceptedLegal sets HasAcceptedLegal field to given value.


### GetHasOnboarded

`func (o *User) GetHasOnboarded() bool`

GetHasOnboarded returns the HasOnboarded field if non-nil, zero value otherwise.

### GetHasOnboardedOk

`func (o *User) GetHasOnboardedOk() (*bool, bool)`

GetHasOnboardedOk returns a tuple with the HasOnboarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOnboarded

`func (o *User) SetHasOnboarded(v bool)`

SetHasOnboarded sets HasOnboarded field to given value.


### GetGender

`func (o *User) GetGender() Gender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *User) GetGenderOk() (*Gender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *User) SetGender(v Gender)`

SetGender sets Gender field to given value.


### GetBirthYear

`func (o *User) GetBirthYear() float64`

GetBirthYear returns the BirthYear field if non-nil, zero value otherwise.

### GetBirthYearOk

`func (o *User) GetBirthYearOk() (*float64, bool)`

GetBirthYearOk returns a tuple with the BirthYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthYear

`func (o *User) SetBirthYear(v float64)`

SetBirthYear sets BirthYear field to given value.


### GetAvatar

`func (o *User) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *User) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *User) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetRemainingReferrals

`func (o *User) GetRemainingReferrals() float64`

GetRemainingReferrals returns the RemainingReferrals field if non-nil, zero value otherwise.

### GetRemainingReferralsOk

`func (o *User) GetRemainingReferralsOk() (*float64, bool)`

GetRemainingReferralsOk returns a tuple with the RemainingReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingReferrals

`func (o *User) SetRemainingReferrals(v float64)`

SetRemainingReferrals sets RemainingReferrals field to given value.


### GetCode

`func (o *User) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *User) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *User) SetCode(v string)`

SetCode sets Code field to given value.


### GetUsertag

`func (o *User) GetUsertag() float64`

GetUsertag returns the Usertag field if non-nil, zero value otherwise.

### GetUsertagOk

`func (o *User) GetUsertagOk() (*float64, bool)`

GetUsertagOk returns a tuple with the Usertag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsertag

`func (o *User) SetUsertag(v float64)`

SetUsertag sets Usertag field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.


### GetPhoneNumber

`func (o *User) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *User) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *User) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### SetPhoneNumberNil

`func (o *User) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *User) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *User) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *User) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


