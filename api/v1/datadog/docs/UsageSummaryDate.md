# UsageSummaryDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all agent hosts over all hours in the current date for all organizations. | [optional] 
**ApmAzureAppServiceHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all Azure app services using APM over all hours in the current date all organizations. | [optional] 
**ApmHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current date for all organizations. | [optional] 
**AwsHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all AWS hosts over all hours in the current date for all organizations. | [optional] 
**AwsLambdaFuncCount** | Pointer to **NullableInt64** | Shows the average of the number of functions that executed 1 or more times each hour in the current date for all organizations. | [optional] 
**AwsLambdaInvocationsSum** | Pointer to **NullableInt64** | Shows the sum of all AWS Lambda invocations over all hours in the current date for all organizations. | [optional] 
**AzureAppServiceTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all Azure app services over all hours in the current date for all organizations. | [optional] 
**BillableIngestedBytesSum** | Pointer to **NullableInt64** | Shows the sum of all log bytes ingested over all hours in the current date for all organizations. | [optional] 
**ComplianceContainerCountSum** | Pointer to **NullableInt64** | Shows the sum of compliance containers over all hours in the current date for all organizations. | [optional] 
**ComplianceHostCountSum** | Pointer to **NullableInt64** | Shows the sum of compliance hosts over all hours in the current date for all organizations. | [optional] 
**ContainerAvg** | Pointer to **NullableInt64** | Shows the average of all distinct containers over all hours in the current date for all organizations. | [optional] 
**ContainerHwm** | Pointer to **NullableInt64** | Shows the high-water mark of all distinct containers over all hours in the current date for all organizations. | [optional] 
**CustomTsAvg** | Pointer to **NullableInt64** | Shows the average number of distinct custom metrics over all hours in the current date for all organizations. | [optional] 
**Date** | Pointer to **time.Time** | The date for the usage. | [optional] 
**FargateTasksCountAvg** | Pointer to **NullableInt64** | Shows the high-watermark of all Fargate tasks over all hours in the current date for all organizations. | [optional] 
**FargateTasksCountHwm** | Pointer to **NullableInt64** | Shows the average of all Fargate tasks over all hours in the current date for all organizations. | [optional] 
**GcpHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all GCP hosts over all hours in the current date for all organizations. | [optional] 
**HerokuHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all Heroku dynos over all hours in the current date for all organizations. | [optional] 
**IncidentManagementMonthlyActiveUsersHwm** | Pointer to **NullableInt64** | Shows the high-water mark of incident management monthly active users over all hours in the current date for all organizations. | [optional] 
**IndexedEventsCountSum** | Pointer to **NullableInt64** | Shows the sum of all log events indexed over all hours in the current date for all organizations. | [optional] 
**InfraHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for all organizations. | [optional] 
**IngestedEventsBytesSum** | Pointer to **NullableInt64** | Shows the sum of all log bytes ingested over all hours in the current date for all organizations. | [optional] 
**IotDeviceSum** | Pointer to **NullableInt64** | Shows the sum of all IoT devices over all hours in the current date for all organizations. | [optional] 
**IotDeviceTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all IoT devices over all hours in the current date all organizations. | [optional] 
**MobileRumSessionCountAndroidSum** | Pointer to **NullableInt64** | Shows the sum of all mobile RUM Sessions on Android over all hours in the current date for all organizations. | [optional] 
**MobileRumSessionCountIosSum** | Pointer to **NullableInt64** | Shows the sum of all mobile RUM Sessions on iOS over all hours in the current date for all organizations. | [optional] 
**MobileRumSessionCountSum** | Pointer to **NullableInt64** | Shows the sum of all mobile RUM Sessions over all hours in the current date for all organizations | [optional] 
**NetflowIndexedEventsCountSum** | Pointer to **NullableInt64** | Shows the sum of all Network flows indexed over all hours in the current date for all organizations. | [optional] 
**NpmHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for all organizations. | [optional] 
**OpentelemetryHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for all organizations. | [optional] 
**Orgs** | Pointer to [**[]UsageSummaryDateOrg**](UsageSummaryDateOrg.md) | Organizations associated with a user. | [optional] 
**ProfilingHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all profiled hosts over all hours in the current date for all organizations. | [optional] 
**RumSessionCountSum** | Pointer to **NullableInt64** | Shows the sum of all browser RUM Sessions over all hours in the current date for all organizations | [optional] 
**RumTotalSessionCountSum** | Pointer to **NullableInt64** | Shows the sum of RUM Sessions (browser and mobile) over all hours in the current date for all organizations. | [optional] 
**SyntheticsBrowserCheckCallsCountSum** | Pointer to **NullableInt64** | Shows the sum of all Synthetic browser tests over all hours in the current date for all organizations. | [optional] 
**SyntheticsCheckCallsCountSum** | Pointer to **NullableInt64** | Shows the sum of all Synthetic API tests over all hours in the current date for all organizations. | [optional] 
**TraceSearchIndexedEventsCountSum** | Pointer to **NullableInt64** | Shows the sum of all Indexed Spans indexed over all hours in the current date for all organizations. | [optional] 
**TwolIngestedEventsBytesSum** | Pointer to **NullableInt64** | Shows the sum of all tracing without limits bytes ingested over all hours in the current date for all organizations. | [optional] 
**VsphereHostTop99p** | Pointer to **int64** | Shows the 99th percentile of all vSphere hosts over all hours in the current date for all organizations. | [optional] 

## Methods

### NewUsageSummaryDate

`func NewUsageSummaryDate() *UsageSummaryDate`

NewUsageSummaryDate instantiates a new UsageSummaryDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSummaryDateWithDefaults

`func NewUsageSummaryDateWithDefaults() *UsageSummaryDate`

NewUsageSummaryDateWithDefaults instantiates a new UsageSummaryDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentHostTop99p

`func (o *UsageSummaryDate) GetAgentHostTop99p() int64`

GetAgentHostTop99p returns the AgentHostTop99p field if non-nil, zero value otherwise.

### GetAgentHostTop99pOk

`func (o *UsageSummaryDate) GetAgentHostTop99pOk() (*int64, bool)`

GetAgentHostTop99pOk returns a tuple with the AgentHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentHostTop99p

`func (o *UsageSummaryDate) SetAgentHostTop99p(v int64)`

SetAgentHostTop99p sets AgentHostTop99p field to given value.

### HasAgentHostTop99p

`func (o *UsageSummaryDate) HasAgentHostTop99p() bool`

HasAgentHostTop99p returns a boolean if a field has been set.

### SetAgentHostTop99pNil

`func (o *UsageSummaryDate) SetAgentHostTop99pNil(b bool)`

 SetAgentHostTop99pNil sets the value for AgentHostTop99p to be an explicit nil

### UnsetAgentHostTop99p
`func (o *UsageSummaryDate) UnsetAgentHostTop99p()`

UnsetAgentHostTop99p ensures that no value is present for AgentHostTop99p, not even an explicit nil
### GetApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDate) GetApmAzureAppServiceHostTop99p() int64`

GetApmAzureAppServiceHostTop99p returns the ApmAzureAppServiceHostTop99p field if non-nil, zero value otherwise.

### GetApmAzureAppServiceHostTop99pOk

`func (o *UsageSummaryDate) GetApmAzureAppServiceHostTop99pOk() (*int64, bool)`

GetApmAzureAppServiceHostTop99pOk returns a tuple with the ApmAzureAppServiceHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDate) SetApmAzureAppServiceHostTop99p(v int64)`

SetApmAzureAppServiceHostTop99p sets ApmAzureAppServiceHostTop99p field to given value.

### HasApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDate) HasApmAzureAppServiceHostTop99p() bool`

HasApmAzureAppServiceHostTop99p returns a boolean if a field has been set.

### SetApmAzureAppServiceHostTop99pNil

`func (o *UsageSummaryDate) SetApmAzureAppServiceHostTop99pNil(b bool)`

 SetApmAzureAppServiceHostTop99pNil sets the value for ApmAzureAppServiceHostTop99p to be an explicit nil

### UnsetApmAzureAppServiceHostTop99p
`func (o *UsageSummaryDate) UnsetApmAzureAppServiceHostTop99p()`

UnsetApmAzureAppServiceHostTop99p ensures that no value is present for ApmAzureAppServiceHostTop99p, not even an explicit nil
### GetApmHostTop99p

`func (o *UsageSummaryDate) GetApmHostTop99p() int64`

GetApmHostTop99p returns the ApmHostTop99p field if non-nil, zero value otherwise.

### GetApmHostTop99pOk

`func (o *UsageSummaryDate) GetApmHostTop99pOk() (*int64, bool)`

GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostTop99p

`func (o *UsageSummaryDate) SetApmHostTop99p(v int64)`

SetApmHostTop99p sets ApmHostTop99p field to given value.

### HasApmHostTop99p

`func (o *UsageSummaryDate) HasApmHostTop99p() bool`

HasApmHostTop99p returns a boolean if a field has been set.

### SetApmHostTop99pNil

`func (o *UsageSummaryDate) SetApmHostTop99pNil(b bool)`

 SetApmHostTop99pNil sets the value for ApmHostTop99p to be an explicit nil

### UnsetApmHostTop99p
`func (o *UsageSummaryDate) UnsetApmHostTop99p()`

UnsetApmHostTop99p ensures that no value is present for ApmHostTop99p, not even an explicit nil
### GetAwsHostTop99p

`func (o *UsageSummaryDate) GetAwsHostTop99p() int64`

GetAwsHostTop99p returns the AwsHostTop99p field if non-nil, zero value otherwise.

### GetAwsHostTop99pOk

`func (o *UsageSummaryDate) GetAwsHostTop99pOk() (*int64, bool)`

GetAwsHostTop99pOk returns a tuple with the AwsHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostTop99p

`func (o *UsageSummaryDate) SetAwsHostTop99p(v int64)`

SetAwsHostTop99p sets AwsHostTop99p field to given value.

### HasAwsHostTop99p

`func (o *UsageSummaryDate) HasAwsHostTop99p() bool`

HasAwsHostTop99p returns a boolean if a field has been set.

### SetAwsHostTop99pNil

`func (o *UsageSummaryDate) SetAwsHostTop99pNil(b bool)`

 SetAwsHostTop99pNil sets the value for AwsHostTop99p to be an explicit nil

### UnsetAwsHostTop99p
`func (o *UsageSummaryDate) UnsetAwsHostTop99p()`

UnsetAwsHostTop99p ensures that no value is present for AwsHostTop99p, not even an explicit nil
### GetAwsLambdaFuncCount

`func (o *UsageSummaryDate) GetAwsLambdaFuncCount() int64`

GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field if non-nil, zero value otherwise.

### GetAwsLambdaFuncCountOk

`func (o *UsageSummaryDate) GetAwsLambdaFuncCountOk() (*int64, bool)`

GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaFuncCount

`func (o *UsageSummaryDate) SetAwsLambdaFuncCount(v int64)`

SetAwsLambdaFuncCount sets AwsLambdaFuncCount field to given value.

### HasAwsLambdaFuncCount

`func (o *UsageSummaryDate) HasAwsLambdaFuncCount() bool`

HasAwsLambdaFuncCount returns a boolean if a field has been set.

### SetAwsLambdaFuncCountNil

`func (o *UsageSummaryDate) SetAwsLambdaFuncCountNil(b bool)`

 SetAwsLambdaFuncCountNil sets the value for AwsLambdaFuncCount to be an explicit nil

### UnsetAwsLambdaFuncCount
`func (o *UsageSummaryDate) UnsetAwsLambdaFuncCount()`

UnsetAwsLambdaFuncCount ensures that no value is present for AwsLambdaFuncCount, not even an explicit nil
### GetAwsLambdaInvocationsSum

`func (o *UsageSummaryDate) GetAwsLambdaInvocationsSum() int64`

GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field if non-nil, zero value otherwise.

### GetAwsLambdaInvocationsSumOk

`func (o *UsageSummaryDate) GetAwsLambdaInvocationsSumOk() (*int64, bool)`

GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaInvocationsSum

`func (o *UsageSummaryDate) SetAwsLambdaInvocationsSum(v int64)`

SetAwsLambdaInvocationsSum sets AwsLambdaInvocationsSum field to given value.

### HasAwsLambdaInvocationsSum

`func (o *UsageSummaryDate) HasAwsLambdaInvocationsSum() bool`

HasAwsLambdaInvocationsSum returns a boolean if a field has been set.

### SetAwsLambdaInvocationsSumNil

`func (o *UsageSummaryDate) SetAwsLambdaInvocationsSumNil(b bool)`

 SetAwsLambdaInvocationsSumNil sets the value for AwsLambdaInvocationsSum to be an explicit nil

### UnsetAwsLambdaInvocationsSum
`func (o *UsageSummaryDate) UnsetAwsLambdaInvocationsSum()`

UnsetAwsLambdaInvocationsSum ensures that no value is present for AwsLambdaInvocationsSum, not even an explicit nil
### GetAzureAppServiceTop99p

`func (o *UsageSummaryDate) GetAzureAppServiceTop99p() int64`

GetAzureAppServiceTop99p returns the AzureAppServiceTop99p field if non-nil, zero value otherwise.

### GetAzureAppServiceTop99pOk

`func (o *UsageSummaryDate) GetAzureAppServiceTop99pOk() (*int64, bool)`

GetAzureAppServiceTop99pOk returns a tuple with the AzureAppServiceTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppServiceTop99p

`func (o *UsageSummaryDate) SetAzureAppServiceTop99p(v int64)`

SetAzureAppServiceTop99p sets AzureAppServiceTop99p field to given value.

### HasAzureAppServiceTop99p

`func (o *UsageSummaryDate) HasAzureAppServiceTop99p() bool`

HasAzureAppServiceTop99p returns a boolean if a field has been set.

### SetAzureAppServiceTop99pNil

`func (o *UsageSummaryDate) SetAzureAppServiceTop99pNil(b bool)`

 SetAzureAppServiceTop99pNil sets the value for AzureAppServiceTop99p to be an explicit nil

### UnsetAzureAppServiceTop99p
`func (o *UsageSummaryDate) UnsetAzureAppServiceTop99p()`

UnsetAzureAppServiceTop99p ensures that no value is present for AzureAppServiceTop99p, not even an explicit nil
### GetBillableIngestedBytesSum

`func (o *UsageSummaryDate) GetBillableIngestedBytesSum() int64`

GetBillableIngestedBytesSum returns the BillableIngestedBytesSum field if non-nil, zero value otherwise.

### GetBillableIngestedBytesSumOk

`func (o *UsageSummaryDate) GetBillableIngestedBytesSumOk() (*int64, bool)`

GetBillableIngestedBytesSumOk returns a tuple with the BillableIngestedBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableIngestedBytesSum

`func (o *UsageSummaryDate) SetBillableIngestedBytesSum(v int64)`

SetBillableIngestedBytesSum sets BillableIngestedBytesSum field to given value.

### HasBillableIngestedBytesSum

`func (o *UsageSummaryDate) HasBillableIngestedBytesSum() bool`

HasBillableIngestedBytesSum returns a boolean if a field has been set.

### SetBillableIngestedBytesSumNil

`func (o *UsageSummaryDate) SetBillableIngestedBytesSumNil(b bool)`

 SetBillableIngestedBytesSumNil sets the value for BillableIngestedBytesSum to be an explicit nil

### UnsetBillableIngestedBytesSum
`func (o *UsageSummaryDate) UnsetBillableIngestedBytesSum()`

UnsetBillableIngestedBytesSum ensures that no value is present for BillableIngestedBytesSum, not even an explicit nil
### GetComplianceContainerCountSum

`func (o *UsageSummaryDate) GetComplianceContainerCountSum() int64`

GetComplianceContainerCountSum returns the ComplianceContainerCountSum field if non-nil, zero value otherwise.

### GetComplianceContainerCountSumOk

`func (o *UsageSummaryDate) GetComplianceContainerCountSumOk() (*int64, bool)`

GetComplianceContainerCountSumOk returns a tuple with the ComplianceContainerCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceContainerCountSum

`func (o *UsageSummaryDate) SetComplianceContainerCountSum(v int64)`

SetComplianceContainerCountSum sets ComplianceContainerCountSum field to given value.

### HasComplianceContainerCountSum

`func (o *UsageSummaryDate) HasComplianceContainerCountSum() bool`

HasComplianceContainerCountSum returns a boolean if a field has been set.

### SetComplianceContainerCountSumNil

`func (o *UsageSummaryDate) SetComplianceContainerCountSumNil(b bool)`

 SetComplianceContainerCountSumNil sets the value for ComplianceContainerCountSum to be an explicit nil

### UnsetComplianceContainerCountSum
`func (o *UsageSummaryDate) UnsetComplianceContainerCountSum()`

UnsetComplianceContainerCountSum ensures that no value is present for ComplianceContainerCountSum, not even an explicit nil
### GetComplianceHostCountSum

`func (o *UsageSummaryDate) GetComplianceHostCountSum() int64`

GetComplianceHostCountSum returns the ComplianceHostCountSum field if non-nil, zero value otherwise.

### GetComplianceHostCountSumOk

`func (o *UsageSummaryDate) GetComplianceHostCountSumOk() (*int64, bool)`

GetComplianceHostCountSumOk returns a tuple with the ComplianceHostCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceHostCountSum

`func (o *UsageSummaryDate) SetComplianceHostCountSum(v int64)`

SetComplianceHostCountSum sets ComplianceHostCountSum field to given value.

### HasComplianceHostCountSum

`func (o *UsageSummaryDate) HasComplianceHostCountSum() bool`

HasComplianceHostCountSum returns a boolean if a field has been set.

### SetComplianceHostCountSumNil

`func (o *UsageSummaryDate) SetComplianceHostCountSumNil(b bool)`

 SetComplianceHostCountSumNil sets the value for ComplianceHostCountSum to be an explicit nil

### UnsetComplianceHostCountSum
`func (o *UsageSummaryDate) UnsetComplianceHostCountSum()`

UnsetComplianceHostCountSum ensures that no value is present for ComplianceHostCountSum, not even an explicit nil
### GetContainerAvg

`func (o *UsageSummaryDate) GetContainerAvg() int64`

GetContainerAvg returns the ContainerAvg field if non-nil, zero value otherwise.

### GetContainerAvgOk

`func (o *UsageSummaryDate) GetContainerAvgOk() (*int64, bool)`

GetContainerAvgOk returns a tuple with the ContainerAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerAvg

`func (o *UsageSummaryDate) SetContainerAvg(v int64)`

SetContainerAvg sets ContainerAvg field to given value.

### HasContainerAvg

`func (o *UsageSummaryDate) HasContainerAvg() bool`

HasContainerAvg returns a boolean if a field has been set.

### SetContainerAvgNil

`func (o *UsageSummaryDate) SetContainerAvgNil(b bool)`

 SetContainerAvgNil sets the value for ContainerAvg to be an explicit nil

### UnsetContainerAvg
`func (o *UsageSummaryDate) UnsetContainerAvg()`

UnsetContainerAvg ensures that no value is present for ContainerAvg, not even an explicit nil
### GetContainerHwm

`func (o *UsageSummaryDate) GetContainerHwm() int64`

GetContainerHwm returns the ContainerHwm field if non-nil, zero value otherwise.

### GetContainerHwmOk

`func (o *UsageSummaryDate) GetContainerHwmOk() (*int64, bool)`

GetContainerHwmOk returns a tuple with the ContainerHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHwm

`func (o *UsageSummaryDate) SetContainerHwm(v int64)`

SetContainerHwm sets ContainerHwm field to given value.

### HasContainerHwm

`func (o *UsageSummaryDate) HasContainerHwm() bool`

HasContainerHwm returns a boolean if a field has been set.

### SetContainerHwmNil

`func (o *UsageSummaryDate) SetContainerHwmNil(b bool)`

 SetContainerHwmNil sets the value for ContainerHwm to be an explicit nil

### UnsetContainerHwm
`func (o *UsageSummaryDate) UnsetContainerHwm()`

UnsetContainerHwm ensures that no value is present for ContainerHwm, not even an explicit nil
### GetCustomTsAvg

`func (o *UsageSummaryDate) GetCustomTsAvg() int64`

GetCustomTsAvg returns the CustomTsAvg field if non-nil, zero value otherwise.

### GetCustomTsAvgOk

`func (o *UsageSummaryDate) GetCustomTsAvgOk() (*int64, bool)`

GetCustomTsAvgOk returns a tuple with the CustomTsAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTsAvg

`func (o *UsageSummaryDate) SetCustomTsAvg(v int64)`

SetCustomTsAvg sets CustomTsAvg field to given value.

### HasCustomTsAvg

`func (o *UsageSummaryDate) HasCustomTsAvg() bool`

HasCustomTsAvg returns a boolean if a field has been set.

### SetCustomTsAvgNil

`func (o *UsageSummaryDate) SetCustomTsAvgNil(b bool)`

 SetCustomTsAvgNil sets the value for CustomTsAvg to be an explicit nil

### UnsetCustomTsAvg
`func (o *UsageSummaryDate) UnsetCustomTsAvg()`

UnsetCustomTsAvg ensures that no value is present for CustomTsAvg, not even an explicit nil
### GetDate

`func (o *UsageSummaryDate) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UsageSummaryDate) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UsageSummaryDate) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *UsageSummaryDate) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFargateTasksCountAvg

`func (o *UsageSummaryDate) GetFargateTasksCountAvg() int64`

GetFargateTasksCountAvg returns the FargateTasksCountAvg field if non-nil, zero value otherwise.

### GetFargateTasksCountAvgOk

`func (o *UsageSummaryDate) GetFargateTasksCountAvgOk() (*int64, bool)`

GetFargateTasksCountAvgOk returns a tuple with the FargateTasksCountAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountAvg

`func (o *UsageSummaryDate) SetFargateTasksCountAvg(v int64)`

SetFargateTasksCountAvg sets FargateTasksCountAvg field to given value.

### HasFargateTasksCountAvg

`func (o *UsageSummaryDate) HasFargateTasksCountAvg() bool`

HasFargateTasksCountAvg returns a boolean if a field has been set.

### SetFargateTasksCountAvgNil

`func (o *UsageSummaryDate) SetFargateTasksCountAvgNil(b bool)`

 SetFargateTasksCountAvgNil sets the value for FargateTasksCountAvg to be an explicit nil

### UnsetFargateTasksCountAvg
`func (o *UsageSummaryDate) UnsetFargateTasksCountAvg()`

UnsetFargateTasksCountAvg ensures that no value is present for FargateTasksCountAvg, not even an explicit nil
### GetFargateTasksCountHwm

`func (o *UsageSummaryDate) GetFargateTasksCountHwm() int64`

GetFargateTasksCountHwm returns the FargateTasksCountHwm field if non-nil, zero value otherwise.

### GetFargateTasksCountHwmOk

`func (o *UsageSummaryDate) GetFargateTasksCountHwmOk() (*int64, bool)`

GetFargateTasksCountHwmOk returns a tuple with the FargateTasksCountHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountHwm

`func (o *UsageSummaryDate) SetFargateTasksCountHwm(v int64)`

SetFargateTasksCountHwm sets FargateTasksCountHwm field to given value.

### HasFargateTasksCountHwm

`func (o *UsageSummaryDate) HasFargateTasksCountHwm() bool`

HasFargateTasksCountHwm returns a boolean if a field has been set.

### SetFargateTasksCountHwmNil

`func (o *UsageSummaryDate) SetFargateTasksCountHwmNil(b bool)`

 SetFargateTasksCountHwmNil sets the value for FargateTasksCountHwm to be an explicit nil

### UnsetFargateTasksCountHwm
`func (o *UsageSummaryDate) UnsetFargateTasksCountHwm()`

UnsetFargateTasksCountHwm ensures that no value is present for FargateTasksCountHwm, not even an explicit nil
### GetGcpHostTop99p

`func (o *UsageSummaryDate) GetGcpHostTop99p() int64`

GetGcpHostTop99p returns the GcpHostTop99p field if non-nil, zero value otherwise.

### GetGcpHostTop99pOk

`func (o *UsageSummaryDate) GetGcpHostTop99pOk() (*int64, bool)`

GetGcpHostTop99pOk returns a tuple with the GcpHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpHostTop99p

`func (o *UsageSummaryDate) SetGcpHostTop99p(v int64)`

SetGcpHostTop99p sets GcpHostTop99p field to given value.

### HasGcpHostTop99p

`func (o *UsageSummaryDate) HasGcpHostTop99p() bool`

HasGcpHostTop99p returns a boolean if a field has been set.

### SetGcpHostTop99pNil

`func (o *UsageSummaryDate) SetGcpHostTop99pNil(b bool)`

 SetGcpHostTop99pNil sets the value for GcpHostTop99p to be an explicit nil

### UnsetGcpHostTop99p
`func (o *UsageSummaryDate) UnsetGcpHostTop99p()`

UnsetGcpHostTop99p ensures that no value is present for GcpHostTop99p, not even an explicit nil
### GetHerokuHostTop99p

`func (o *UsageSummaryDate) GetHerokuHostTop99p() int64`

GetHerokuHostTop99p returns the HerokuHostTop99p field if non-nil, zero value otherwise.

### GetHerokuHostTop99pOk

`func (o *UsageSummaryDate) GetHerokuHostTop99pOk() (*int64, bool)`

GetHerokuHostTop99pOk returns a tuple with the HerokuHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHerokuHostTop99p

`func (o *UsageSummaryDate) SetHerokuHostTop99p(v int64)`

SetHerokuHostTop99p sets HerokuHostTop99p field to given value.

### HasHerokuHostTop99p

`func (o *UsageSummaryDate) HasHerokuHostTop99p() bool`

HasHerokuHostTop99p returns a boolean if a field has been set.

### SetHerokuHostTop99pNil

`func (o *UsageSummaryDate) SetHerokuHostTop99pNil(b bool)`

 SetHerokuHostTop99pNil sets the value for HerokuHostTop99p to be an explicit nil

### UnsetHerokuHostTop99p
`func (o *UsageSummaryDate) UnsetHerokuHostTop99p()`

UnsetHerokuHostTop99p ensures that no value is present for HerokuHostTop99p, not even an explicit nil
### GetIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDate) GetIncidentManagementMonthlyActiveUsersHwm() int64`

GetIncidentManagementMonthlyActiveUsersHwm returns the IncidentManagementMonthlyActiveUsersHwm field if non-nil, zero value otherwise.

### GetIncidentManagementMonthlyActiveUsersHwmOk

`func (o *UsageSummaryDate) GetIncidentManagementMonthlyActiveUsersHwmOk() (*int64, bool)`

GetIncidentManagementMonthlyActiveUsersHwmOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDate) SetIncidentManagementMonthlyActiveUsersHwm(v int64)`

SetIncidentManagementMonthlyActiveUsersHwm sets IncidentManagementMonthlyActiveUsersHwm field to given value.

### HasIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDate) HasIncidentManagementMonthlyActiveUsersHwm() bool`

HasIncidentManagementMonthlyActiveUsersHwm returns a boolean if a field has been set.

### SetIncidentManagementMonthlyActiveUsersHwmNil

`func (o *UsageSummaryDate) SetIncidentManagementMonthlyActiveUsersHwmNil(b bool)`

 SetIncidentManagementMonthlyActiveUsersHwmNil sets the value for IncidentManagementMonthlyActiveUsersHwm to be an explicit nil

### UnsetIncidentManagementMonthlyActiveUsersHwm
`func (o *UsageSummaryDate) UnsetIncidentManagementMonthlyActiveUsersHwm()`

UnsetIncidentManagementMonthlyActiveUsersHwm ensures that no value is present for IncidentManagementMonthlyActiveUsersHwm, not even an explicit nil
### GetIndexedEventsCountSum

`func (o *UsageSummaryDate) GetIndexedEventsCountSum() int64`

GetIndexedEventsCountSum returns the IndexedEventsCountSum field if non-nil, zero value otherwise.

### GetIndexedEventsCountSumOk

`func (o *UsageSummaryDate) GetIndexedEventsCountSumOk() (*int64, bool)`

GetIndexedEventsCountSumOk returns a tuple with the IndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCountSum

`func (o *UsageSummaryDate) SetIndexedEventsCountSum(v int64)`

SetIndexedEventsCountSum sets IndexedEventsCountSum field to given value.

### HasIndexedEventsCountSum

`func (o *UsageSummaryDate) HasIndexedEventsCountSum() bool`

HasIndexedEventsCountSum returns a boolean if a field has been set.

### SetIndexedEventsCountSumNil

`func (o *UsageSummaryDate) SetIndexedEventsCountSumNil(b bool)`

 SetIndexedEventsCountSumNil sets the value for IndexedEventsCountSum to be an explicit nil

### UnsetIndexedEventsCountSum
`func (o *UsageSummaryDate) UnsetIndexedEventsCountSum()`

UnsetIndexedEventsCountSum ensures that no value is present for IndexedEventsCountSum, not even an explicit nil
### GetInfraHostTop99p

`func (o *UsageSummaryDate) GetInfraHostTop99p() int64`

GetInfraHostTop99p returns the InfraHostTop99p field if non-nil, zero value otherwise.

### GetInfraHostTop99pOk

`func (o *UsageSummaryDate) GetInfraHostTop99pOk() (*int64, bool)`

GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostTop99p

`func (o *UsageSummaryDate) SetInfraHostTop99p(v int64)`

SetInfraHostTop99p sets InfraHostTop99p field to given value.

### HasInfraHostTop99p

`func (o *UsageSummaryDate) HasInfraHostTop99p() bool`

HasInfraHostTop99p returns a boolean if a field has been set.

### SetInfraHostTop99pNil

`func (o *UsageSummaryDate) SetInfraHostTop99pNil(b bool)`

 SetInfraHostTop99pNil sets the value for InfraHostTop99p to be an explicit nil

### UnsetInfraHostTop99p
`func (o *UsageSummaryDate) UnsetInfraHostTop99p()`

UnsetInfraHostTop99p ensures that no value is present for InfraHostTop99p, not even an explicit nil
### GetIngestedEventsBytesSum

`func (o *UsageSummaryDate) GetIngestedEventsBytesSum() int64`

GetIngestedEventsBytesSum returns the IngestedEventsBytesSum field if non-nil, zero value otherwise.

### GetIngestedEventsBytesSumOk

`func (o *UsageSummaryDate) GetIngestedEventsBytesSumOk() (*int64, bool)`

GetIngestedEventsBytesSumOk returns a tuple with the IngestedEventsBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEventsBytesSum

`func (o *UsageSummaryDate) SetIngestedEventsBytesSum(v int64)`

SetIngestedEventsBytesSum sets IngestedEventsBytesSum field to given value.

### HasIngestedEventsBytesSum

`func (o *UsageSummaryDate) HasIngestedEventsBytesSum() bool`

HasIngestedEventsBytesSum returns a boolean if a field has been set.

### SetIngestedEventsBytesSumNil

`func (o *UsageSummaryDate) SetIngestedEventsBytesSumNil(b bool)`

 SetIngestedEventsBytesSumNil sets the value for IngestedEventsBytesSum to be an explicit nil

### UnsetIngestedEventsBytesSum
`func (o *UsageSummaryDate) UnsetIngestedEventsBytesSum()`

UnsetIngestedEventsBytesSum ensures that no value is present for IngestedEventsBytesSum, not even an explicit nil
### GetIotDeviceSum

`func (o *UsageSummaryDate) GetIotDeviceSum() int64`

GetIotDeviceSum returns the IotDeviceSum field if non-nil, zero value otherwise.

### GetIotDeviceSumOk

`func (o *UsageSummaryDate) GetIotDeviceSumOk() (*int64, bool)`

GetIotDeviceSumOk returns a tuple with the IotDeviceSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceSum

`func (o *UsageSummaryDate) SetIotDeviceSum(v int64)`

SetIotDeviceSum sets IotDeviceSum field to given value.

### HasIotDeviceSum

`func (o *UsageSummaryDate) HasIotDeviceSum() bool`

HasIotDeviceSum returns a boolean if a field has been set.

### SetIotDeviceSumNil

`func (o *UsageSummaryDate) SetIotDeviceSumNil(b bool)`

 SetIotDeviceSumNil sets the value for IotDeviceSum to be an explicit nil

### UnsetIotDeviceSum
`func (o *UsageSummaryDate) UnsetIotDeviceSum()`

UnsetIotDeviceSum ensures that no value is present for IotDeviceSum, not even an explicit nil
### GetIotDeviceTop99p

`func (o *UsageSummaryDate) GetIotDeviceTop99p() int64`

GetIotDeviceTop99p returns the IotDeviceTop99p field if non-nil, zero value otherwise.

### GetIotDeviceTop99pOk

`func (o *UsageSummaryDate) GetIotDeviceTop99pOk() (*int64, bool)`

GetIotDeviceTop99pOk returns a tuple with the IotDeviceTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceTop99p

`func (o *UsageSummaryDate) SetIotDeviceTop99p(v int64)`

SetIotDeviceTop99p sets IotDeviceTop99p field to given value.

### HasIotDeviceTop99p

`func (o *UsageSummaryDate) HasIotDeviceTop99p() bool`

HasIotDeviceTop99p returns a boolean if a field has been set.

### SetIotDeviceTop99pNil

`func (o *UsageSummaryDate) SetIotDeviceTop99pNil(b bool)`

 SetIotDeviceTop99pNil sets the value for IotDeviceTop99p to be an explicit nil

### UnsetIotDeviceTop99p
`func (o *UsageSummaryDate) UnsetIotDeviceTop99p()`

UnsetIotDeviceTop99p ensures that no value is present for IotDeviceTop99p, not even an explicit nil
### GetMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDate) GetMobileRumSessionCountAndroidSum() int64`

GetMobileRumSessionCountAndroidSum returns the MobileRumSessionCountAndroidSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountAndroidSumOk

`func (o *UsageSummaryDate) GetMobileRumSessionCountAndroidSumOk() (*int64, bool)`

GetMobileRumSessionCountAndroidSumOk returns a tuple with the MobileRumSessionCountAndroidSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDate) SetMobileRumSessionCountAndroidSum(v int64)`

SetMobileRumSessionCountAndroidSum sets MobileRumSessionCountAndroidSum field to given value.

### HasMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDate) HasMobileRumSessionCountAndroidSum() bool`

HasMobileRumSessionCountAndroidSum returns a boolean if a field has been set.

### SetMobileRumSessionCountAndroidSumNil

`func (o *UsageSummaryDate) SetMobileRumSessionCountAndroidSumNil(b bool)`

 SetMobileRumSessionCountAndroidSumNil sets the value for MobileRumSessionCountAndroidSum to be an explicit nil

### UnsetMobileRumSessionCountAndroidSum
`func (o *UsageSummaryDate) UnsetMobileRumSessionCountAndroidSum()`

UnsetMobileRumSessionCountAndroidSum ensures that no value is present for MobileRumSessionCountAndroidSum, not even an explicit nil
### GetMobileRumSessionCountIosSum

`func (o *UsageSummaryDate) GetMobileRumSessionCountIosSum() int64`

GetMobileRumSessionCountIosSum returns the MobileRumSessionCountIosSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountIosSumOk

`func (o *UsageSummaryDate) GetMobileRumSessionCountIosSumOk() (*int64, bool)`

GetMobileRumSessionCountIosSumOk returns a tuple with the MobileRumSessionCountIosSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountIosSum

`func (o *UsageSummaryDate) SetMobileRumSessionCountIosSum(v int64)`

SetMobileRumSessionCountIosSum sets MobileRumSessionCountIosSum field to given value.

### HasMobileRumSessionCountIosSum

`func (o *UsageSummaryDate) HasMobileRumSessionCountIosSum() bool`

HasMobileRumSessionCountIosSum returns a boolean if a field has been set.

### SetMobileRumSessionCountIosSumNil

`func (o *UsageSummaryDate) SetMobileRumSessionCountIosSumNil(b bool)`

 SetMobileRumSessionCountIosSumNil sets the value for MobileRumSessionCountIosSum to be an explicit nil

### UnsetMobileRumSessionCountIosSum
`func (o *UsageSummaryDate) UnsetMobileRumSessionCountIosSum()`

UnsetMobileRumSessionCountIosSum ensures that no value is present for MobileRumSessionCountIosSum, not even an explicit nil
### GetMobileRumSessionCountSum

`func (o *UsageSummaryDate) GetMobileRumSessionCountSum() int64`

GetMobileRumSessionCountSum returns the MobileRumSessionCountSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountSumOk

`func (o *UsageSummaryDate) GetMobileRumSessionCountSumOk() (*int64, bool)`

GetMobileRumSessionCountSumOk returns a tuple with the MobileRumSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountSum

`func (o *UsageSummaryDate) SetMobileRumSessionCountSum(v int64)`

SetMobileRumSessionCountSum sets MobileRumSessionCountSum field to given value.

### HasMobileRumSessionCountSum

`func (o *UsageSummaryDate) HasMobileRumSessionCountSum() bool`

HasMobileRumSessionCountSum returns a boolean if a field has been set.

### SetMobileRumSessionCountSumNil

`func (o *UsageSummaryDate) SetMobileRumSessionCountSumNil(b bool)`

 SetMobileRumSessionCountSumNil sets the value for MobileRumSessionCountSum to be an explicit nil

### UnsetMobileRumSessionCountSum
`func (o *UsageSummaryDate) UnsetMobileRumSessionCountSum()`

UnsetMobileRumSessionCountSum ensures that no value is present for MobileRumSessionCountSum, not even an explicit nil
### GetNetflowIndexedEventsCountSum

`func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSum() int64`

GetNetflowIndexedEventsCountSum returns the NetflowIndexedEventsCountSum field if non-nil, zero value otherwise.

### GetNetflowIndexedEventsCountSumOk

`func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSumOk() (*int64, bool)`

GetNetflowIndexedEventsCountSumOk returns a tuple with the NetflowIndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetflowIndexedEventsCountSum

`func (o *UsageSummaryDate) SetNetflowIndexedEventsCountSum(v int64)`

SetNetflowIndexedEventsCountSum sets NetflowIndexedEventsCountSum field to given value.

### HasNetflowIndexedEventsCountSum

`func (o *UsageSummaryDate) HasNetflowIndexedEventsCountSum() bool`

HasNetflowIndexedEventsCountSum returns a boolean if a field has been set.

### SetNetflowIndexedEventsCountSumNil

`func (o *UsageSummaryDate) SetNetflowIndexedEventsCountSumNil(b bool)`

 SetNetflowIndexedEventsCountSumNil sets the value for NetflowIndexedEventsCountSum to be an explicit nil

### UnsetNetflowIndexedEventsCountSum
`func (o *UsageSummaryDate) UnsetNetflowIndexedEventsCountSum()`

UnsetNetflowIndexedEventsCountSum ensures that no value is present for NetflowIndexedEventsCountSum, not even an explicit nil
### GetNpmHostTop99p

`func (o *UsageSummaryDate) GetNpmHostTop99p() int64`

GetNpmHostTop99p returns the NpmHostTop99p field if non-nil, zero value otherwise.

### GetNpmHostTop99pOk

`func (o *UsageSummaryDate) GetNpmHostTop99pOk() (*int64, bool)`

GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostTop99p

`func (o *UsageSummaryDate) SetNpmHostTop99p(v int64)`

SetNpmHostTop99p sets NpmHostTop99p field to given value.

### HasNpmHostTop99p

`func (o *UsageSummaryDate) HasNpmHostTop99p() bool`

HasNpmHostTop99p returns a boolean if a field has been set.

### SetNpmHostTop99pNil

`func (o *UsageSummaryDate) SetNpmHostTop99pNil(b bool)`

 SetNpmHostTop99pNil sets the value for NpmHostTop99p to be an explicit nil

### UnsetNpmHostTop99p
`func (o *UsageSummaryDate) UnsetNpmHostTop99p()`

UnsetNpmHostTop99p ensures that no value is present for NpmHostTop99p, not even an explicit nil
### GetOpentelemetryHostTop99p

`func (o *UsageSummaryDate) GetOpentelemetryHostTop99p() int64`

GetOpentelemetryHostTop99p returns the OpentelemetryHostTop99p field if non-nil, zero value otherwise.

### GetOpentelemetryHostTop99pOk

`func (o *UsageSummaryDate) GetOpentelemetryHostTop99pOk() (*int64, bool)`

GetOpentelemetryHostTop99pOk returns a tuple with the OpentelemetryHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentelemetryHostTop99p

`func (o *UsageSummaryDate) SetOpentelemetryHostTop99p(v int64)`

SetOpentelemetryHostTop99p sets OpentelemetryHostTop99p field to given value.

### HasOpentelemetryHostTop99p

`func (o *UsageSummaryDate) HasOpentelemetryHostTop99p() bool`

HasOpentelemetryHostTop99p returns a boolean if a field has been set.

### SetOpentelemetryHostTop99pNil

`func (o *UsageSummaryDate) SetOpentelemetryHostTop99pNil(b bool)`

 SetOpentelemetryHostTop99pNil sets the value for OpentelemetryHostTop99p to be an explicit nil

### UnsetOpentelemetryHostTop99p
`func (o *UsageSummaryDate) UnsetOpentelemetryHostTop99p()`

UnsetOpentelemetryHostTop99p ensures that no value is present for OpentelemetryHostTop99p, not even an explicit nil
### GetOrgs

`func (o *UsageSummaryDate) GetOrgs() []UsageSummaryDateOrg`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *UsageSummaryDate) GetOrgsOk() (*[]UsageSummaryDateOrg, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *UsageSummaryDate) SetOrgs(v []UsageSummaryDateOrg)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *UsageSummaryDate) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetProfilingHostTop99p

`func (o *UsageSummaryDate) GetProfilingHostTop99p() int64`

GetProfilingHostTop99p returns the ProfilingHostTop99p field if non-nil, zero value otherwise.

### GetProfilingHostTop99pOk

`func (o *UsageSummaryDate) GetProfilingHostTop99pOk() (*int64, bool)`

GetProfilingHostTop99pOk returns a tuple with the ProfilingHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilingHostTop99p

`func (o *UsageSummaryDate) SetProfilingHostTop99p(v int64)`

SetProfilingHostTop99p sets ProfilingHostTop99p field to given value.

### HasProfilingHostTop99p

`func (o *UsageSummaryDate) HasProfilingHostTop99p() bool`

HasProfilingHostTop99p returns a boolean if a field has been set.

### SetProfilingHostTop99pNil

`func (o *UsageSummaryDate) SetProfilingHostTop99pNil(b bool)`

 SetProfilingHostTop99pNil sets the value for ProfilingHostTop99p to be an explicit nil

### UnsetProfilingHostTop99p
`func (o *UsageSummaryDate) UnsetProfilingHostTop99p()`

UnsetProfilingHostTop99p ensures that no value is present for ProfilingHostTop99p, not even an explicit nil
### GetRumSessionCountSum

`func (o *UsageSummaryDate) GetRumSessionCountSum() int64`

GetRumSessionCountSum returns the RumSessionCountSum field if non-nil, zero value otherwise.

### GetRumSessionCountSumOk

`func (o *UsageSummaryDate) GetRumSessionCountSumOk() (*int64, bool)`

GetRumSessionCountSumOk returns a tuple with the RumSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumSessionCountSum

`func (o *UsageSummaryDate) SetRumSessionCountSum(v int64)`

SetRumSessionCountSum sets RumSessionCountSum field to given value.

### HasRumSessionCountSum

`func (o *UsageSummaryDate) HasRumSessionCountSum() bool`

HasRumSessionCountSum returns a boolean if a field has been set.

### SetRumSessionCountSumNil

`func (o *UsageSummaryDate) SetRumSessionCountSumNil(b bool)`

 SetRumSessionCountSumNil sets the value for RumSessionCountSum to be an explicit nil

### UnsetRumSessionCountSum
`func (o *UsageSummaryDate) UnsetRumSessionCountSum()`

UnsetRumSessionCountSum ensures that no value is present for RumSessionCountSum, not even an explicit nil
### GetRumTotalSessionCountSum

`func (o *UsageSummaryDate) GetRumTotalSessionCountSum() int64`

GetRumTotalSessionCountSum returns the RumTotalSessionCountSum field if non-nil, zero value otherwise.

### GetRumTotalSessionCountSumOk

`func (o *UsageSummaryDate) GetRumTotalSessionCountSumOk() (*int64, bool)`

GetRumTotalSessionCountSumOk returns a tuple with the RumTotalSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumTotalSessionCountSum

`func (o *UsageSummaryDate) SetRumTotalSessionCountSum(v int64)`

SetRumTotalSessionCountSum sets RumTotalSessionCountSum field to given value.

### HasRumTotalSessionCountSum

`func (o *UsageSummaryDate) HasRumTotalSessionCountSum() bool`

HasRumTotalSessionCountSum returns a boolean if a field has been set.

### SetRumTotalSessionCountSumNil

`func (o *UsageSummaryDate) SetRumTotalSessionCountSumNil(b bool)`

 SetRumTotalSessionCountSumNil sets the value for RumTotalSessionCountSum to be an explicit nil

### UnsetRumTotalSessionCountSum
`func (o *UsageSummaryDate) UnsetRumTotalSessionCountSum()`

UnsetRumTotalSessionCountSum ensures that no value is present for RumTotalSessionCountSum, not even an explicit nil
### GetSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSum() int64`

GetSyntheticsBrowserCheckCallsCountSum returns the SyntheticsBrowserCheckCallsCountSum field if non-nil, zero value otherwise.

### GetSyntheticsBrowserCheckCallsCountSumOk

`func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSumOk() (*int64, bool)`

GetSyntheticsBrowserCheckCallsCountSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDate) SetSyntheticsBrowserCheckCallsCountSum(v int64)`

SetSyntheticsBrowserCheckCallsCountSum sets SyntheticsBrowserCheckCallsCountSum field to given value.

### HasSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDate) HasSyntheticsBrowserCheckCallsCountSum() bool`

HasSyntheticsBrowserCheckCallsCountSum returns a boolean if a field has been set.

### SetSyntheticsBrowserCheckCallsCountSumNil

`func (o *UsageSummaryDate) SetSyntheticsBrowserCheckCallsCountSumNil(b bool)`

 SetSyntheticsBrowserCheckCallsCountSumNil sets the value for SyntheticsBrowserCheckCallsCountSum to be an explicit nil

### UnsetSyntheticsBrowserCheckCallsCountSum
`func (o *UsageSummaryDate) UnsetSyntheticsBrowserCheckCallsCountSum()`

UnsetSyntheticsBrowserCheckCallsCountSum ensures that no value is present for SyntheticsBrowserCheckCallsCountSum, not even an explicit nil
### GetSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSum() int64`

GetSyntheticsCheckCallsCountSum returns the SyntheticsCheckCallsCountSum field if non-nil, zero value otherwise.

### GetSyntheticsCheckCallsCountSumOk

`func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSumOk() (*int64, bool)`

GetSyntheticsCheckCallsCountSumOk returns a tuple with the SyntheticsCheckCallsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDate) SetSyntheticsCheckCallsCountSum(v int64)`

SetSyntheticsCheckCallsCountSum sets SyntheticsCheckCallsCountSum field to given value.

### HasSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDate) HasSyntheticsCheckCallsCountSum() bool`

HasSyntheticsCheckCallsCountSum returns a boolean if a field has been set.

### SetSyntheticsCheckCallsCountSumNil

`func (o *UsageSummaryDate) SetSyntheticsCheckCallsCountSumNil(b bool)`

 SetSyntheticsCheckCallsCountSumNil sets the value for SyntheticsCheckCallsCountSum to be an explicit nil

### UnsetSyntheticsCheckCallsCountSum
`func (o *UsageSummaryDate) UnsetSyntheticsCheckCallsCountSum()`

UnsetSyntheticsCheckCallsCountSum ensures that no value is present for SyntheticsCheckCallsCountSum, not even an explicit nil
### GetTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSum() int64`

GetTraceSearchIndexedEventsCountSum returns the TraceSearchIndexedEventsCountSum field if non-nil, zero value otherwise.

### GetTraceSearchIndexedEventsCountSumOk

`func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSumOk() (*int64, bool)`

GetTraceSearchIndexedEventsCountSumOk returns a tuple with the TraceSearchIndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDate) SetTraceSearchIndexedEventsCountSum(v int64)`

SetTraceSearchIndexedEventsCountSum sets TraceSearchIndexedEventsCountSum field to given value.

### HasTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDate) HasTraceSearchIndexedEventsCountSum() bool`

HasTraceSearchIndexedEventsCountSum returns a boolean if a field has been set.

### SetTraceSearchIndexedEventsCountSumNil

`func (o *UsageSummaryDate) SetTraceSearchIndexedEventsCountSumNil(b bool)`

 SetTraceSearchIndexedEventsCountSumNil sets the value for TraceSearchIndexedEventsCountSum to be an explicit nil

### UnsetTraceSearchIndexedEventsCountSum
`func (o *UsageSummaryDate) UnsetTraceSearchIndexedEventsCountSum()`

UnsetTraceSearchIndexedEventsCountSum ensures that no value is present for TraceSearchIndexedEventsCountSum, not even an explicit nil
### GetTwolIngestedEventsBytesSum

`func (o *UsageSummaryDate) GetTwolIngestedEventsBytesSum() int64`

GetTwolIngestedEventsBytesSum returns the TwolIngestedEventsBytesSum field if non-nil, zero value otherwise.

### GetTwolIngestedEventsBytesSumOk

`func (o *UsageSummaryDate) GetTwolIngestedEventsBytesSumOk() (*int64, bool)`

GetTwolIngestedEventsBytesSumOk returns a tuple with the TwolIngestedEventsBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwolIngestedEventsBytesSum

`func (o *UsageSummaryDate) SetTwolIngestedEventsBytesSum(v int64)`

SetTwolIngestedEventsBytesSum sets TwolIngestedEventsBytesSum field to given value.

### HasTwolIngestedEventsBytesSum

`func (o *UsageSummaryDate) HasTwolIngestedEventsBytesSum() bool`

HasTwolIngestedEventsBytesSum returns a boolean if a field has been set.

### SetTwolIngestedEventsBytesSumNil

`func (o *UsageSummaryDate) SetTwolIngestedEventsBytesSumNil(b bool)`

 SetTwolIngestedEventsBytesSumNil sets the value for TwolIngestedEventsBytesSum to be an explicit nil

### UnsetTwolIngestedEventsBytesSum
`func (o *UsageSummaryDate) UnsetTwolIngestedEventsBytesSum()`

UnsetTwolIngestedEventsBytesSum ensures that no value is present for TwolIngestedEventsBytesSum, not even an explicit nil
### GetVsphereHostTop99p

`func (o *UsageSummaryDate) GetVsphereHostTop99p() int64`

GetVsphereHostTop99p returns the VsphereHostTop99p field if non-nil, zero value otherwise.

### GetVsphereHostTop99pOk

`func (o *UsageSummaryDate) GetVsphereHostTop99pOk() (*int64, bool)`

GetVsphereHostTop99pOk returns a tuple with the VsphereHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHostTop99p

`func (o *UsageSummaryDate) SetVsphereHostTop99p(v int64)`

SetVsphereHostTop99p sets VsphereHostTop99p field to given value.

### HasVsphereHostTop99p

`func (o *UsageSummaryDate) HasVsphereHostTop99p() bool`

HasVsphereHostTop99p returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


