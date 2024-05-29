# \ExternalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterUser**](ExternalAPI.md#RegisterUser) | **Post** /external/users | 
[**TrackActivity**](ExternalAPI.md#TrackActivity) | **Post** /external/activities | 



## RegisterUser

> RegisterUser200Response RegisterUser(ctx).ApiKey(apiKey).ExternalUser(externalUser).Execute()



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
	externalUser := *openapiclient.NewExternalUser("UserId_example", "Username_example") // ExternalUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.RegisterUser(context.Background()).ApiKey(apiKey).ExternalUser(externalUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.RegisterUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterUser`: RegisterUser200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.RegisterUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | 
 **externalUser** | [**ExternalUser**](ExternalUser.md) |  | 

### Return type

[**RegisterUser200Response**](RegisterUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackActivity

> TrackActivity200Response TrackActivity(ctx).ApiKey(apiKey).TrackActivityInput(trackActivityInput).Execute()



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
	trackActivityInput := *openapiclient.NewTrackActivityInput("Occurrence_example", "ActivityId_example", "UserId_example") // TrackActivityInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAPI.TrackActivity(context.Background()).ApiKey(apiKey).TrackActivityInput(trackActivityInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.TrackActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackActivity`: TrackActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.TrackActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | 
 **trackActivityInput** | [**TrackActivityInput**](TrackActivityInput.md) |  | 

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

