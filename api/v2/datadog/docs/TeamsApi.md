# \TeamsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeam**](TeamsApi.md#CreateTeam) | **Post** /api/v1/teams | Create a new team
[**DeleteTeam**](TeamsApi.md#DeleteTeam) | **Delete** /api/v1/teams/{team_id} | Delete an existing team
[**GetTeam**](TeamsApi.md#GetTeam) | **Get** /api/v1/teams/{team_id} | Get details of a team
[**GetTeams**](TeamsApi.md#GetTeams) | **Get** /api/v1/teams | Get a list of all teams
[**PatchTeam**](TeamsApi.md#PatchTeam) | **Patch** /api/v1/teams/{team_id} | Update an existing team



## CreateTeam

> TeamPayload CreateTeam(ctx).Body(body).Execute()

Create a new team



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TeamPayload**](TeamPayload.md) | Teams Payload. | 

### Return type

[**TeamPayload**](TeamPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DeleteTeam(ctx, teamId).Execute()

Delete an existing team



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> TeamPayload GetTeam(ctx, teamId).Execute()

Get details of a team



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamPayload**](TeamPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeams

> TeamListPayload GetTeams(ctx).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get a list of all teams



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withUsers** | **bool** | Include user data in api response. | [default to false]
 **pageLimit** | **int64** | maximum number of incidents to return. | [default to 1000]
 **pageOffset** | **int64** | index of the first element for the query. | [default to 0]

### Return type

[**TeamListPayload**](TeamListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTeam

> TeamPayload PatchTeam(ctx, teamId).Body(body).Execute()

Update an existing team



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TeamPayload**](TeamPayload.md) | Teams Payload. | 

### Return type

[**TeamPayload**](TeamPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

