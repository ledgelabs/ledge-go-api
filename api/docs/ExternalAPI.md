# \ExternalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGoal**](ExternalAPI.md#CreateGoal) | **Post** /external/{key}/goal | 
[**CreateQuest**](ExternalAPI.md#CreateQuest) | **Post** /external/quests | 
[**CreateQuestSchedule**](ExternalAPI.md#CreateQuestSchedule) | **Post** /external/{key}/questSchedule | 
[**GetExternalUser**](ExternalAPI.md#GetExternalUser) | **Get** /external/{key}/external-user | 
[**GetMyGoals**](ExternalAPI.md#GetMyGoals) | **Get** /external/my-goals | 
[**GetMyQuests**](ExternalAPI.md#GetMyQuests) | **Get** /external/my-quests | 
[**RegisterUser**](ExternalAPI.md#RegisterUser) | **Post** /external/{key}/user | 
[**TrackActivity**](ExternalAPI.md#TrackActivity) | **Post** /external/{key}/activity | 



## CreateGoal

> TrackActivity200Response CreateGoal(ctx, key).CreateGoalRequest(createGoalRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	key := "key_example" // string | 
	createGoalRequest := *openapiclient.NewCreateGoalRequest("Activity_example", "QuestId_example", float64(123), "Description_example", "Title_example") // CreateGoalRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.CreateGoal(context.Background(), key).CreateGoalRequest(createGoalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.CreateGoal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGoal`: TrackActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.CreateGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGoalRequest** | [**CreateGoalRequest**](CreateGoalRequest.md) |  | 

### Return type

[**TrackActivity200Response**](TrackActivity200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuest

> TrackActivity200Response CreateQuest(ctx).ApiKey(apiKey).CreateQuestRequest(createQuestRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiKey := "apiKey_example" // string | 
	createQuestRequest := *openapiclient.NewCreateQuestRequest("Description_example", "Title_example") // CreateQuestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.CreateQuest(context.Background()).ApiKey(apiKey).CreateQuestRequest(createQuestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.CreateQuest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuest`: TrackActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.CreateQuest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | 
 **createQuestRequest** | [**CreateQuestRequest**](CreateQuestRequest.md) |  | 

### Return type

[**TrackActivity200Response**](TrackActivity200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuestSchedule

> CreateQuestSchedule200Response CreateQuestSchedule(ctx, key).CreateQuestScheduleRequest(createQuestScheduleRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	key := "key_example" // string | 
	createQuestScheduleRequest := *openapiclient.NewCreateQuestScheduleRequest(false, "QuestId_example", time.Now(), time.Now()) // CreateQuestScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.CreateQuestSchedule(context.Background(), key).CreateQuestScheduleRequest(createQuestScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.CreateQuestSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuestSchedule`: CreateQuestSchedule200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.CreateQuestSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createQuestScheduleRequest** | [**CreateQuestScheduleRequest**](CreateQuestScheduleRequest.md) |  | 

### Return type

[**CreateQuestSchedule200Response**](CreateQuestSchedule200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalUser

> User GetExternalUser(ctx, key).ExternalId(externalId).GameId(gameId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	key := "key_example" // string | 
	externalId := "externalId_example" // string | 
	gameId := "gameId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.GetExternalUser(context.Background(), key).ExternalId(externalId).GameId(gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.GetExternalUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalUser`: User
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.GetExternalUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalId** | **string** |  | 
 **gameId** | **string** |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyGoals

> []Goal GetMyGoals(ctx).ApiKey(apiKey).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiKey := "apiKey_example" // string | given by the Ledge admins.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.GetMyGoals(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.GetMyGoals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyGoals`: []Goal
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.GetMyGoals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyGoalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | given by the Ledge admins. | 

### Return type

[**[]Goal**](Goal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyQuests

> []Quest GetMyQuests(ctx).ApiKey(apiKey).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiKey := "apiKey_example" // string | given by the Ledge admins.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.GetMyQuests(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.GetMyQuests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyQuests`: []Quest
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.GetMyQuests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyQuestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | given by the Ledge admins. | 

### Return type

[**[]Quest**](Quest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterUser

> User RegisterUser(ctx, key).ExternalUser(externalUser).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	key := "key_example" // string | 
	externalUser := *openapiclient.NewExternalUser("UserId_example", "Username_example") // ExternalUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.RegisterUser(context.Background(), key).ExternalUser(externalUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.RegisterUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterUser`: User
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.RegisterUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalUser** | [**ExternalUser**](ExternalUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackActivity

> TrackActivity200Response TrackActivity(ctx, key).ExternalActivity(externalActivity).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	key := "key_example" // string | 
	externalActivity := *openapiclient.NewExternalActivity("Activity_example", "UserId_example") // ExternalActivity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.TrackActivity(context.Background(), key).ExternalActivity(externalActivity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.TrackActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackActivity`: TrackActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.TrackActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalActivity** | [**ExternalActivity**](ExternalActivity.md) |  | 

### Return type

[**TrackActivity200Response**](TrackActivity200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

