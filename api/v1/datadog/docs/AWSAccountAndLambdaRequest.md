# AWSAccountAndLambdaRequest

## Properties

| Name          | Type       | Description                                                                                    | Notes |
| ------------- | ---------- | ---------------------------------------------------------------------------------------------- | ----- |
| **AccountId** | **string** | Your AWS Account ID without dashes.                                                            |
| **LambdaArn** | **string** | ARN of the Datadog Lambda created during the Datadog-Amazon Web services Log collection setup. |

## Methods

### NewAWSAccountAndLambdaRequest

`func NewAWSAccountAndLambdaRequest(accountId string, lambdaArn string) *AWSAccountAndLambdaRequest`

NewAWSAccountAndLambdaRequest instantiates a new AWSAccountAndLambdaRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAWSAccountAndLambdaRequestWithDefaults

`func NewAWSAccountAndLambdaRequestWithDefaults() *AWSAccountAndLambdaRequest`

NewAWSAccountAndLambdaRequestWithDefaults instantiates a new AWSAccountAndLambdaRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAccountId

`func (o *AWSAccountAndLambdaRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSAccountAndLambdaRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSAccountAndLambdaRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### GetLambdaArn

`func (o *AWSAccountAndLambdaRequest) GetLambdaArn() string`

GetLambdaArn returns the LambdaArn field if non-nil, zero value otherwise.

### GetLambdaArnOk

`func (o *AWSAccountAndLambdaRequest) GetLambdaArnOk() (*string, bool)`

GetLambdaArnOk returns a tuple with the LambdaArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaArn

`func (o *AWSAccountAndLambdaRequest) SetLambdaArn(v string)`

SetLambdaArn sets LambdaArn field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
