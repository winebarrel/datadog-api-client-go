# \IncidentsConfigApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncidentConfigField**](IncidentsConfigApi.md#CreateIncidentConfigField) | **Post** /api/v1/incidents/config/fields | Create a new incident configuration field
[**CreateIncidentConfigFieldChoice**](IncidentsConfigApi.md#CreateIncidentConfigFieldChoice) | **Post** /api/v1/incidents/config/fields/{field_id}/relationships/choices | Create a new choice
[**DeleteIncidentConfigField**](IncidentsConfigApi.md#DeleteIncidentConfigField) | **Delete** /api/v1/incidents/config/fields/{field_id} | Deletes an existing field
[**DeleteIncidentConfigFieldChoice**](IncidentsConfigApi.md#DeleteIncidentConfigFieldChoice) | **Delete** /api/v1/incidents/config/fields/{field_id}/relationships/choices/{choice_id} | Update an existing choice for the given incident
[**GetIncidentConfigField**](IncidentsConfigApi.md#GetIncidentConfigField) | **Get** /api/v1/incidents/config/fields/{field_id} | Get the details of a field
[**GetIncidentConfigFieldChoice**](IncidentsConfigApi.md#GetIncidentConfigFieldChoice) | **Get** /api/v1/incidents/config/fields/{field_id}/relationships/choices/{choice_id} | Get the details of a field choice
[**GetIncidentConfigFieldChoices**](IncidentsConfigApi.md#GetIncidentConfigFieldChoices) | **Get** /api/v1/incidents/config/fields/{field_id}/relationships/choices | Get all choices for the given field
[**GetIncidentConfigFields**](IncidentsConfigApi.md#GetIncidentConfigFields) | **Get** /api/v1/incidents/config/fields | Get a list of all configuration fields
[**PatchIncidentConfigField**](IncidentsConfigApi.md#PatchIncidentConfigField) | **Patch** /api/v1/incidents/config/fields/{field_id} | Updates an existing configuration field
[**PatchIncidentConfigFieldChoice**](IncidentsConfigApi.md#PatchIncidentConfigFieldChoice) | **Patch** /api/v1/incidents/config/fields/{field_id}/relationships/choices/{choice_id} | Update an existing field choice



## CreateIncidentConfigField

> IncidentConfigFieldPayload CreateIncidentConfigField(ctx).Body(body).Execute()

Create a new incident configuration field



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentConfigFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IncidentConfigFieldPayload**](IncidentConfigFieldPayload.md) | Incident Configuration Field Payload. | 

### Return type

[**IncidentConfigFieldPayload**](IncidentConfigFieldPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncidentConfigFieldChoice

> IncidentConfigFieldChoicePayload CreateIncidentConfigFieldChoice(ctx, fieldId).Body(body).Execute()

Create a new choice



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the incident field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentConfigFieldChoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentConfigFieldChoicePayload**](IncidentConfigFieldChoicePayload.md) | Incident Configuration Field Choice Payload. | 

### Return type

[**IncidentConfigFieldChoicePayload**](IncidentConfigFieldChoicePayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncidentConfigField

> DeleteIncidentConfigField(ctx, fieldId).Execute()

Deletes an existing field



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the incident field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentConfigFieldRequest struct via the builder pattern


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


## DeleteIncidentConfigFieldChoice

> DeleteIncidentConfigFieldChoice(ctx, fieldId, choiceId).Execute()

Update an existing choice for the given incident



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the incident field. | 
**choiceId** | **string** | The ID of the incident field choice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentConfigFieldChoiceRequest struct via the builder pattern


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


## GetIncidentConfigField

> IncidentConfigFieldPayload GetIncidentConfigField(ctx, fieldId).Execute()

Get the details of a field



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the incident field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentConfigFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IncidentConfigFieldPayload**](IncidentConfigFieldPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentConfigFieldChoice

> IncidentConfigFieldChoicePayload GetIncidentConfigFieldChoice(ctx, fieldId, choiceId).Execute()

Get the details of a field choice



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the incident field. | 
**choiceId** | **string** | The ID of the incident field choice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentConfigFieldChoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IncidentConfigFieldChoicePayload**](IncidentConfigFieldChoicePayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentConfigFieldChoices

> IncidentConfigFieldChoiceListPayload GetIncidentConfigFieldChoices(ctx, fieldId).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get all choices for the given field



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the incident field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentConfigFieldChoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withUsers** | **bool** | Include user data in api response. | [default to false]
 **pageLimit** | **int64** | maximum number of incidents to return. | [default to 1000]
 **pageOffset** | **int64** | index of the first element for the query. | [default to 0]

### Return type

[**IncidentConfigFieldChoiceListPayload**](IncidentConfigFieldChoiceListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentConfigFields

> IncidentConfigFieldListPayload GetIncidentConfigFields(ctx).WithUsers(withUsers).PageLimit(pageLimit).PageOffset(pageOffset).Execute()

Get a list of all configuration fields



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentConfigFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withUsers** | **bool** | Include user data in api response. | [default to false]
 **pageLimit** | **int64** | maximum number of incidents to return. | [default to 1000]
 **pageOffset** | **int64** | index of the first element for the query. | [default to 0]

### Return type

[**IncidentConfigFieldListPayload**](IncidentConfigFieldListPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIncidentConfigField

> IncidentConfigFieldPayload PatchIncidentConfigField(ctx, fieldId).Body(body).Execute()

Updates an existing configuration field



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the incident field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIncidentConfigFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentConfigFieldPayload**](IncidentConfigFieldPayload.md) | Incident Configuration Field Payload. | 

### Return type

[**IncidentConfigFieldPayload**](IncidentConfigFieldPayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIncidentConfigFieldChoice

> IncidentConfigFieldChoicePayload PatchIncidentConfigFieldChoice(ctx, fieldId, choiceId).Body(body).Execute()

Update an existing field choice



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the incident field. | 
**choiceId** | **string** | The ID of the incident field choice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIncidentConfigFieldChoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IncidentConfigFieldChoicePayload**](IncidentConfigFieldChoicePayload.md) | Incident Configuration Field Choice Payload. | 

### Return type

[**IncidentConfigFieldChoicePayload**](IncidentConfigFieldChoicePayload.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

