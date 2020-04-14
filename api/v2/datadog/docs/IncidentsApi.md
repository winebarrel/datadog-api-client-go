# \IncidentsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServiceToIncident**](IncidentsApi.md#AddServiceToIncident) | **Post** /api/v1/incidents/{incident_id}/relationships/services | Add a new service to an incident
[**AddTeamToIncident**](IncidentsApi.md#AddTeamToIncident) | **Post** /api/v1/incidents/{incident_id}/relationships/teams | Add a new team to an incident
[**CreateIncident**](IncidentsApi.md#CreateIncident) | **Post** /api/v1/incidents | Create a new incident
[**CreateIncidentPostmortem**](IncidentsApi.md#CreateIncidentPostmortem) | **Post** /api/v1/incidents/{incident_id}/relationships/postmortems | Create a new postmortem
[**CreateIncidentSelection**](IncidentsApi.md#CreateIncidentSelection) | **Post** /api/v1/incidents/{incident_id}/relationships/selections | Create a new selection
[**CreateIncidentTodo**](IncidentsApi.md#CreateIncidentTodo) | **Post** /api/v1/incidents/{incident_id}/relationships/todos | Create a new todo
[**DeleteIncident**](IncidentsApi.md#DeleteIncident) | **Delete** /api/v1/incidents/{incident_id} | Delete an existing incident
[**DeleteIncidentPostmortem**](IncidentsApi.md#DeleteIncidentPostmortem) | **Delete** /api/v1/incidents/{incident_id}/relationships/postmortems/{postmortem_id} | Update an existing postmortem
[**DeleteIncidentSelection**](IncidentsApi.md#DeleteIncidentSelection) | **Delete** /api/v1/incidents/{incident_id}/relationships/selections/{selection_id} | Deletes an existing selection
[**DeleteIncidentTodo**](IncidentsApi.md#DeleteIncidentTodo) | **Delete** /api/v1/incidents/{incident_id}/relationships/todos/{todo_id} | Deletes an existing todo for the given incident
[**GetIncident**](IncidentsApi.md#GetIncident) | **Get** /api/v1/incidents/{incident_id} | Get the details of an incident
[**GetIncidentPostmortem**](IncidentsApi.md#GetIncidentPostmortem) | **Get** /api/v1/incidents/{incident_id}/relationships/postmortems/{postmortem_id} | Get details of a postmortem
[**GetIncidentPostmortems**](IncidentsApi.md#GetIncidentPostmortems) | **Get** /api/v1/incidents/{incident_id}/relationships/postmortems | Get a list of all Postmortems
[**GetIncidentSelections**](IncidentsApi.md#GetIncidentSelections) | **Get** /api/v1/incidents/{incident_id}/relationships/selections | Get all selections for the given incident
[**GetIncidentTodo**](IncidentsApi.md#GetIncidentTodo) | **Get** /api/v1/incidents/{incident_id}/relationships/todos/{todo_id} | Get the details of an incident todo
[**GetIncidentTodos**](IncidentsApi.md#GetIncidentTodos) | **Get** /api/v1/incidents/{incident_id}/relationships/todos | Get all todos for the given incident
[**GetIncidents**](IncidentsApi.md#GetIncidents) | **Get** /api/v1/incidents | Get a list of incidents
[**GetServicesForIncident**](IncidentsApi.md#GetServicesForIncident) | **Get** /api/v1/incidents/{incident_id}/relationships/services | Get all services for the given incident
[**GetTeamsForIncident**](IncidentsApi.md#GetTeamsForIncident) | **Get** /api/v1/incidents/{incident_id}/relationships/teams | Get all teams for the given incident
[**IncidentsSummary**](IncidentsApi.md#IncidentsSummary) | **Get** /api/v1/incidents/summary | Incidents Summary
[**PatchIncident**](IncidentsApi.md#PatchIncident) | **Patch** /api/v1/incidents/{incident_id} | Update an existing incident
[**PatchIncidentPostmortem**](IncidentsApi.md#PatchIncidentPostmortem) | **Patch** /api/v1/incidents/{incident_id}/relationships/postmortems/{postmortem_id} | Update an existing postmortem
[**PatchIncidentSelection**](IncidentsApi.md#PatchIncidentSelection) | **Patch** /api/v1/incidents/{incident_id}/relationships/selections/{selection_id} | Update an existing selection
[**PatchIncidentTodo**](IncidentsApi.md#PatchIncidentTodo) | **Patch** /api/v1/incidents/{incident_id}/relationships/todos/{todo_id} | Update an existing todo
[**RemoveServiceFromIncident**](IncidentsApi.md#RemoveServiceFromIncident) | **Delete** /api/v1/incidents/{incident_id}/relationships/services/{service_id} | Remove a service from an incident
[**RemoveTeamFromIncident**](IncidentsApi.md#RemoveTeamFromIncident) | **Delete** /api/v1/incidents/{incident_id}/relationships/teams/{team_id} | Remove a team from an incident
[**SearchIncidents**](IncidentsApi.md#SearchIncidents) | **Get** /api/v1/incidents/search | Search incidents



## AddServiceToIncident

> ServiceListPayload AddServiceToIncident(ctx, incidentId).Body(body).Execute()

Add a new service to an incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServiceToIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ServicePayload**](ServicePayload.md) | Services Payload. | 

### Return type

[**ServiceListPayload**](ServiceListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTeamToIncident

> TeamListPayload AddTeamToIncident(ctx, incidentId).Body(body).Execute()

Add a new team to an incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTeamToIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TeamPayload**](TeamPayload.md) | Teams Payload. | 

### Return type

[**TeamListPayload**](TeamListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncident

> IncidentPayload CreateIncident(ctx).Body(body).Execute()

Create a new incident



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IncidentCreateWithInitialDataPayload**](IncidentCreateWithInitialDataPayload.md) | Incident Payload. | 

### Return type

[**IncidentPayload**](IncidentPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncidentPostmortem

> IncidentPostmortemPayload CreateIncidentPostmortem(ctx, incidentId).Body(body).Execute()

Create a new postmortem



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentPostmortemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentPostmortemPayload**](IncidentPostmortemPayload.md) | Incident Postmortem Payload. | 

### Return type

[**IncidentPostmortemPayload**](IncidentPostmortemPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncidentSelection

> IncidentSelectionPayload CreateIncidentSelection(ctx, incidentId).Body(body).Execute()

Create a new selection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentSelectionPayload**](IncidentSelectionPayload.md) | Incident Selection Payload. | 

### Return type

[**IncidentSelectionPayload**](IncidentSelectionPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncidentTodo

> IncidentTodoPayload CreateIncidentTodo(ctx, incidentId).Body(body).Execute()

Create a new todo



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentTodoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentTodoPayload**](IncidentTodoPayload.md) | Incident Todo Payload. | 

### Return type

[**IncidentTodoPayload**](IncidentTodoPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncident

> DeleteIncident(ctx, incidentId).Execute()

Delete an existing incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentRequest struct via the builder pattern


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


## DeleteIncidentPostmortem

> DeleteIncidentPostmortem(ctx, incidentId, postmortemId).Execute()

Update an existing postmortem



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**postmortemId** | **string** | The ID of the postmortem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentPostmortemRequest struct via the builder pattern


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


## DeleteIncidentSelection

> DeleteIncidentSelection(ctx, incidentId, selectionId).Execute()

Deletes an existing selection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**selectionId** | **string** | The ID of the incident field selection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentSelectionRequest struct via the builder pattern


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


## DeleteIncidentTodo

> DeleteIncidentTodo(ctx, incidentId, todoId).Execute()

Deletes an existing todo for the given incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**todoId** | **string** | The ID of the incident todo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentTodoRequest struct via the builder pattern


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


## GetIncident

> IncidentPayload GetIncident(ctx, incidentId).IsPublicId(isPublicId).Execute()

Get the details of an incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isPublicId** | **bool** | Set &#x60;true&#x60; if provided &#x60;incident_id&#x60; is a &#x60;public_id&#x60;. | 

### Return type

[**IncidentPayload**](IncidentPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentPostmortem

> IncidentPostmortemPayload GetIncidentPostmortem(ctx, incidentId, postmortemId).Execute()

Get details of a postmortem



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**postmortemId** | **string** | The ID of the postmortem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentPostmortemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IncidentPostmortemPayload**](IncidentPostmortemPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentPostmortems

> IncidentPostmortemListPayload GetIncidentPostmortems(ctx, incidentId).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get a list of all Postmortems



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentPostmortemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withUsers** | **bool** | Include user data in api response. | [default to false]
 **pageLimit** | **int64** | maximum number of incidents to return. | [default to 1000]
 **pageOffset** | **int64** | index of the first element for the query. | [default to 0]

### Return type

[**IncidentPostmortemListPayload**](IncidentPostmortemListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentSelections

> IncidentSelectionListPayload GetIncidentSelections(ctx, incidentId).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get all selections for the given incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentSelectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withUsers** | **bool** | Include user data in api response. | [default to false]
 **pageLimit** | **int64** | maximum number of incidents to return. | [default to 1000]
 **pageOffset** | **int64** | index of the first element for the query. | [default to 0]

### Return type

[**IncidentSelectionListPayload**](IncidentSelectionListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentTodo

> IncidentTodoPayload GetIncidentTodo(ctx, incidentId, todoId).Execute()

Get the details of an incident todo



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**todoId** | **string** | The ID of the incident todo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentTodoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IncidentTodoPayload**](IncidentTodoPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentTodos

> IncidentTodoListPayload GetIncidentTodos(ctx, incidentId).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get all todos for the given incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentTodosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withUsers** | **bool** | Include user data in api response. | [default to false]
 **pageLimit** | **int64** | maximum number of incidents to return. | [default to 1000]
 **pageOffset** | **int64** | index of the first element for the query. | [default to 0]

### Return type

[**IncidentTodoListPayload**](IncidentTodoListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidents

> IncidentListPayload GetIncidents(ctx).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get a list of incidents



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withUsers** | **bool** | Include user data in api response. | [default to false]
 **pageLimit** | **int64** | maximum number of incidents to return. | [default to 1000]
 **pageOffset** | **int64** | index of the first element for the query. | [default to 0]

### Return type

[**IncidentListPayload**](IncidentListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesForIncident

> ServiceListPayload GetServicesForIncident(ctx, incidentId).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get all services for the given incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesForIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withUsers** | **bool** | Include user data in api response. | [default to false]
 **pageLimit** | **int64** | maximum number of incidents to return. | [default to 1000]
 **pageOffset** | **int64** | index of the first element for the query. | [default to 0]

### Return type

[**ServiceListPayload**](ServiceListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamsForIncident

> TeamListPayload GetTeamsForIncident(ctx, incidentId).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get all teams for the given incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsForIncidentRequest struct via the builder pattern


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


## IncidentsSummary

> IncidentSummaryPayload IncidentsSummary(ctx).FilterCreatedBefore(filterCreatedBefore).FilterCreatedAfter(filterCreatedAfter).FilterLastModifiedBefore(filterLastModifiedBefore).FilterLastModifiedAfter(filterLastModifiedAfter).FilterDeletedBefore(filterDeletedBefore).FilterDeletedAfter(filterDeletedAfter).FilterResolvedBefore(filterResolvedBefore).FilterResolvedAfter(filterResolvedAfter).FilterServices(filterServices).FilterTeams(filterTeams).FilterText(filterText).Execute()

Incidents Summary



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncidentsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCreatedBefore** | **int64** | Scope incidents to created before epoch timestamp. | 
 **filterCreatedAfter** | **int64** | Scope incidents to created after epoch timestamp. | 
 **filterLastModifiedBefore** | **int64** | Scope incidents to last modified before epoch timestamp. | 
 **filterLastModifiedAfter** | **int64** | Scope incidents to last modified after epoch timestamp. | 
 **filterDeletedBefore** | **int64** | Scope incidents to deleted before epoch timestamp. | 
 **filterDeletedAfter** | **int64** | Scope incidents to deleted after epoch timestamp. | 
 **filterResolvedBefore** | **int64** | Scope incidents to deleted resolved epoch timestamp. | 
 **filterResolvedAfter** | **int64** | Scope incidents to resolved after epoch timestamp. | 
 **filterServices** | **string** | Scope incidents which were impacted by the selected services. A CSV of service names. | 
 **filterTeams** | **string** | Scope incidents which were impacted by the selected teams. A CSV of team names. | 
 **filterText** | **string** | Fuzzy text to use to search incident content. | 

### Return type

[**IncidentSummaryPayload**](IncidentSummaryPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIncident

> IncidentPayload PatchIncident(ctx, incidentId).Body(body).Execute()

Update an existing incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentPayload**](IncidentPayload.md) | Incident Payload. | 

### Return type

[**IncidentPayload**](IncidentPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIncidentPostmortem

> IncidentPostmortemPayload PatchIncidentPostmortem(ctx, incidentId, postmortemId).Body(body).Execute()

Update an existing postmortem



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**postmortemId** | **string** | The ID of the postmortem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIncidentPostmortemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IncidentPostmortemPayload**](IncidentPostmortemPayload.md) | Incident Postmortem Payload. | 

### Return type

[**IncidentPostmortemPayload**](IncidentPostmortemPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIncidentSelection

> IncidentSelectionPayload PatchIncidentSelection(ctx, incidentId, selectionId).Body(body).Execute()

Update an existing selection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**selectionId** | **string** | The ID of the incident field selection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIncidentSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IncidentSelectionPayload**](IncidentSelectionPayload.md) | Incident Selection Payload. | 

### Return type

[**IncidentSelectionPayload**](IncidentSelectionPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIncidentTodo

> IncidentTodoPayload PatchIncidentTodo(ctx, incidentId, todoId).Body(body).Execute()

Update an existing todo



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**todoId** | **string** | The ID of the incident todo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIncidentTodoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IncidentTodoPayload**](IncidentTodoPayload.md) | Incident Todo Payload. | 

### Return type

[**IncidentTodoPayload**](IncidentTodoPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveServiceFromIncident

> RemoveServiceFromIncident(ctx, incidentId, serviceId).Execute()

Remove a service from an incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**serviceId** | **string** | The ID of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServiceFromIncidentRequest struct via the builder pattern


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


## RemoveTeamFromIncident

> RemoveTeamFromIncident(ctx, incidentId, teamId).Execute()

Remove a team from an incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The ID(&#x60;incident_id&#x60;) of the incident. | 
**teamId** | **string** | The ID of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTeamFromIncidentRequest struct via the builder pattern


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


## SearchIncidents

> IncidentSearchPayload SearchIncidents(ctx).FilterCreatedBefore(filterCreatedBefore).FilterCreatedAfter(filterCreatedAfter).FilterLastModifiedBefore(filterLastModifiedBefore).FilterLastModifiedAfter(filterLastModifiedAfter).FilterDeletedBefore(filterDeletedBefore).FilterDeletedAfter(filterDeletedAfter).FilterResolvedBefore(filterResolvedBefore).FilterResolvedAfter(filterResolvedAfter).FilterCustomerImpacted(filterCustomerImpacted).FilterState(filterState).FilterCreatedByUuid(filterCreatedByUuid).FilterCreatedByHandle(filterCreatedByHandle).FilterCreatedByEmail(filterCreatedByEmail).FilterCreatedByName(filterCreatedByName).FilterLastModifiedByUuid(filterLastModifiedByUuid).FilterLastModifiedByHandle(filterLastModifiedByHandle).FilterLastModifiedByEmail(filterLastModifiedByEmail).FilterLastModifiedByName(filterLastModifiedByName).FilterCommanderUuid(filterCommanderUuid).FilterCommanderHandle(filterCommanderHandle).FilterCommanderEmail(filterCommanderEmail).FilterCommanderName(filterCommanderName).FilterPostmortemId(filterPostmortemId).FilterIncidentId(filterIncidentId).FilterPublicId(filterPublicId).FilterServices(filterServices).FilterTeams(filterTeams).FilterText(filterText).Execute()

Search incidents



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCreatedBefore** | **int64** | Scope incidents to created before epoch timestamp. | 
 **filterCreatedAfter** | **int64** | Scope incidents to created after epoch timestamp. | 
 **filterLastModifiedBefore** | **int64** | Scope incidents to last modified before epoch timestamp. | 
 **filterLastModifiedAfter** | **int64** | Scope incidents to last modified after epoch timestamp. | 
 **filterDeletedBefore** | **int64** | Scope incidents to deleted before epoch timestamp. | 
 **filterDeletedAfter** | **int64** | Scope incidents to deleted after epoch timestamp. | 
 **filterResolvedBefore** | **int64** | Scope incidents to deleted resolved epoch timestamp. | 
 **filterResolvedAfter** | **int64** | Scope incidents to resolved after epoch timestamp. | 
 **filterCustomerImpacted** | **bool** | Scope incidents to customer impacted status. | 
 **filterState** | **string** | Scope incidents to state of incident. | 
 **filterCreatedByUuid** | **string** | Scope incidents to created by uuid. | 
 **filterCreatedByHandle** | **string** | Scope incidents to created by handle. | 
 **filterCreatedByEmail** | [**string**](.md) | Scope incidents to created by email. | 
 **filterCreatedByName** | **string** | Scope incidents to created by name. | 
 **filterLastModifiedByUuid** | **string** | Scope incidents to last modified by uuid. | 
 **filterLastModifiedByHandle** | **string** | Scope incidents to last modified by handle. | 
 **filterLastModifiedByEmail** | [**string**](.md) | Scope incidents to last modified by email. | 
 **filterLastModifiedByName** | **string** | Scope incidents to last modified by name. | 
 **filterCommanderUuid** | **string** | Scope incidents to commander by uuid. | 
 **filterCommanderHandle** | **string** | Scope incidents to commander by handle. | 
 **filterCommanderEmail** | [**string**](.md) | Scope incidents to commander by email. | 
 **filterCommanderName** | **string** | Scope incidents to commander by name. | 
 **filterPostmortemId** | **int64** | Scope incidents with the following postmortem id. | 
 **filterIncidentId** | **int64** | Scope incidents with the following incident id. | 
 **filterPublicId** | **int64** | Scope incidents with the following public incident id. | 
 **filterServices** | **string** | Scope incidents which were impacted by the selected services. A CSV of service names. | 
 **filterTeams** | **string** | Scope incidents which were impacted by the selected teams. A CSV of team names. | 
 **filterText** | **string** | Fuzzy text to use to search incident content. | 

### Return type

[**IncidentSearchPayload**](IncidentSearchPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

