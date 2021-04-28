# UsageSummaryDateOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all agent hosts over all hours in the current date for the given org. | [optional] 
**ApmAzureAppServiceHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all Azure app services using APM over all hours in the current date for the given org. | [optional] 
**ApmHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all distinct APM hosts over all hours in the current date for the given org. | [optional] 
**AwsHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all AWS hosts over all hours in the current date for the given org. | [optional] 
**AwsLambdaFuncCount** | Pointer to **NullableInt64** | Shows the sum of all AWS Lambda invocations over all hours in the current date for the given org. | [optional] 
**AwsLambdaInvocationsSum** | Pointer to **NullableInt64** | Shows the sum of all AWS Lambda invocations over all hours in the current date for the given org. | [optional] 
**AzureAppServiceTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all Azure app services over all hours in the current date for the given org. | [optional] 
**BillableIngestedBytesSum** | Pointer to **NullableInt64** | Shows the sum of all log bytes ingested over all hours in the current date for the given org. | [optional] 
**ComplianceContainerAggSum** | Pointer to **NullableInt64** | Shows the sum of all compliance containers over all hours in the current date for the given org. | [optional] 
**ComplianceHostAggSum** | Pointer to **NullableInt64** | Shows the sum of all compliance hosts over all hours in the current date for the given org. | [optional] 
**ContainerAvg** | Pointer to **NullableInt64** | Shows the average of all distinct containers over all hours in the current date for the given org. | [optional] 
**ContainerHwm** | Pointer to **NullableInt64** | Shows the high-water mark of all distinct containers over all hours in the current date for the given org. | [optional] 
**CustomTsAvg** | Pointer to **NullableInt64** | Shows the average number of distinct custom metrics over all hours in the current date for the given org. | [optional] 
**FargateTasksCountAvg** | Pointer to **NullableInt64** | The average task count for Fargate. | [optional] 
**FargateTasksCountHwm** | Pointer to **NullableInt64** | Shows the high-water mark of all Fargate tasks over all hours in the current date for the given org. | [optional] 
**GcpHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all GCP hosts over all hours in the current date for the given org. | [optional] 
**HerokuHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all Heroku dynos over all hours in the current date for the given org. | [optional] 
**Id** | Pointer to **string** | The organization id. | [optional] 
**IncidentManagementMonthlyActiveUsersHwm** | Pointer to **NullableInt64** | Shows the high-water mark of incident management monthly active users over all hours in the current date for the given org. | [optional] 
**IndexedEventsCountSum** | Pointer to **NullableInt64** | Shows the sum of all log events indexed over all hours in the current date for the given org. | [optional] 
**InfraHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for the given org. | [optional] 
**IngestedEventsBytesSum** | Pointer to **NullableInt64** | Shows the sum of all log bytes ingested over all hours in the current date for the given org. | [optional] 
**IotDeviceAggSum** | Pointer to **NullableInt64** | Shows the sum of all IoT devices over all hours in the current date for the given org. | [optional] 
**IotDeviceTop99pSum** | Pointer to **NullableInt64** | Shows the 99th percentile of all IoT devices over all hours in the current date for the given org. | [optional] 
**MobileRumSessionCountAndroidSum** | Pointer to **NullableInt64** | Shows the sum of all mobile RUM Sessions on Android over all hours in the current date for the given org. | [optional] 
**MobileRumSessionCountIosSum** | Pointer to **NullableInt64** | Shows the sum of all mobile RUM Sessions on iOS over all hours in the current date for the given org. | [optional] 
**MobileRumSessionCountSum** | Pointer to **NullableInt64** | Shows the sum of all mobile RUM Sessions over all hours in the current date for the given org. | [optional] 
**Name** | Pointer to **string** | The organization name. | [optional] 
**NetflowIndexedEventsCountSum** | Pointer to **NullableInt64** | Shows the sum of all Network flows indexed over all hours in the current date for the given org. | [optional] 
**NpmHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all distinct Networks hosts over all hours in the current date for the given org. | [optional] 
**OpentelemetryHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for the given org. | [optional] 
**ProfilingHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all profiled hosts over all hours in the current date for the given org. | [optional] 
**PublicId** | Pointer to **string** | The organization public id. | [optional] 
**RumSessionCountSum** | Pointer to **NullableInt64** | Shows the sum of all browser RUM Sessions over all hours in the current date for the given org. | [optional] 
**RumTotalSessionCountSum** | Pointer to **NullableInt64** | Shows the sum of RUM Sessions (browser and mobile) over all hours in the current date for the given org. | [optional] 
**SyntheticsBrowserCheckCallsCountSum** | Pointer to **NullableInt64** | Shows the sum of all Synthetic browser tests over all hours in the current date for the given org. | [optional] 
**SyntheticsCheckCallsCountSum** | Pointer to **NullableInt64** | Shows the sum of all Synthetic API tests over all hours in the current date for the given org. | [optional] 
**TraceSearchIndexedEventsCountSum** | Pointer to **NullableInt64** | Shows the sum of all Indexed Spans indexed over all hours in the current date for the given org. | [optional] 
**TwolIngestedEventsBytesSum** | Pointer to **NullableInt64** | Shows the sum of all tracing without limits bytes ingested over all hours in the current date for the given org. | [optional] 
**VsphereHostTop99p** | Pointer to **NullableInt64** | Shows the 99th percentile of all vSphere hosts over all hours in the current date for the given org. | [optional] 

## Methods

### NewUsageSummaryDateOrg

`func NewUsageSummaryDateOrg() *UsageSummaryDateOrg`

NewUsageSummaryDateOrg instantiates a new UsageSummaryDateOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSummaryDateOrgWithDefaults

`func NewUsageSummaryDateOrgWithDefaults() *UsageSummaryDateOrg`

NewUsageSummaryDateOrgWithDefaults instantiates a new UsageSummaryDateOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentHostTop99p

`func (o *UsageSummaryDateOrg) GetAgentHostTop99p() int64`

GetAgentHostTop99p returns the AgentHostTop99p field if non-nil, zero value otherwise.

### GetAgentHostTop99pOk

`func (o *UsageSummaryDateOrg) GetAgentHostTop99pOk() (*int64, bool)`

GetAgentHostTop99pOk returns a tuple with the AgentHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentHostTop99p

`func (o *UsageSummaryDateOrg) SetAgentHostTop99p(v int64)`

SetAgentHostTop99p sets AgentHostTop99p field to given value.

### HasAgentHostTop99p

`func (o *UsageSummaryDateOrg) HasAgentHostTop99p() bool`

HasAgentHostTop99p returns a boolean if a field has been set.

### SetAgentHostTop99pNil

`func (o *UsageSummaryDateOrg) SetAgentHostTop99pNil(b bool)`

 SetAgentHostTop99pNil sets the value for AgentHostTop99p to be an explicit nil

### UnsetAgentHostTop99p
`func (o *UsageSummaryDateOrg) UnsetAgentHostTop99p()`

UnsetAgentHostTop99p ensures that no value is present for AgentHostTop99p, not even an explicit nil
### GetApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDateOrg) GetApmAzureAppServiceHostTop99p() int64`

GetApmAzureAppServiceHostTop99p returns the ApmAzureAppServiceHostTop99p field if non-nil, zero value otherwise.

### GetApmAzureAppServiceHostTop99pOk

`func (o *UsageSummaryDateOrg) GetApmAzureAppServiceHostTop99pOk() (*int64, bool)`

GetApmAzureAppServiceHostTop99pOk returns a tuple with the ApmAzureAppServiceHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDateOrg) SetApmAzureAppServiceHostTop99p(v int64)`

SetApmAzureAppServiceHostTop99p sets ApmAzureAppServiceHostTop99p field to given value.

### HasApmAzureAppServiceHostTop99p

`func (o *UsageSummaryDateOrg) HasApmAzureAppServiceHostTop99p() bool`

HasApmAzureAppServiceHostTop99p returns a boolean if a field has been set.

### SetApmAzureAppServiceHostTop99pNil

`func (o *UsageSummaryDateOrg) SetApmAzureAppServiceHostTop99pNil(b bool)`

 SetApmAzureAppServiceHostTop99pNil sets the value for ApmAzureAppServiceHostTop99p to be an explicit nil

### UnsetApmAzureAppServiceHostTop99p
`func (o *UsageSummaryDateOrg) UnsetApmAzureAppServiceHostTop99p()`

UnsetApmAzureAppServiceHostTop99p ensures that no value is present for ApmAzureAppServiceHostTop99p, not even an explicit nil
### GetApmHostTop99p

`func (o *UsageSummaryDateOrg) GetApmHostTop99p() int64`

GetApmHostTop99p returns the ApmHostTop99p field if non-nil, zero value otherwise.

### GetApmHostTop99pOk

`func (o *UsageSummaryDateOrg) GetApmHostTop99pOk() (*int64, bool)`

GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostTop99p

`func (o *UsageSummaryDateOrg) SetApmHostTop99p(v int64)`

SetApmHostTop99p sets ApmHostTop99p field to given value.

### HasApmHostTop99p

`func (o *UsageSummaryDateOrg) HasApmHostTop99p() bool`

HasApmHostTop99p returns a boolean if a field has been set.

### SetApmHostTop99pNil

`func (o *UsageSummaryDateOrg) SetApmHostTop99pNil(b bool)`

 SetApmHostTop99pNil sets the value for ApmHostTop99p to be an explicit nil

### UnsetApmHostTop99p
`func (o *UsageSummaryDateOrg) UnsetApmHostTop99p()`

UnsetApmHostTop99p ensures that no value is present for ApmHostTop99p, not even an explicit nil
### GetAwsHostTop99p

`func (o *UsageSummaryDateOrg) GetAwsHostTop99p() int64`

GetAwsHostTop99p returns the AwsHostTop99p field if non-nil, zero value otherwise.

### GetAwsHostTop99pOk

`func (o *UsageSummaryDateOrg) GetAwsHostTop99pOk() (*int64, bool)`

GetAwsHostTop99pOk returns a tuple with the AwsHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostTop99p

`func (o *UsageSummaryDateOrg) SetAwsHostTop99p(v int64)`

SetAwsHostTop99p sets AwsHostTop99p field to given value.

### HasAwsHostTop99p

`func (o *UsageSummaryDateOrg) HasAwsHostTop99p() bool`

HasAwsHostTop99p returns a boolean if a field has been set.

### SetAwsHostTop99pNil

`func (o *UsageSummaryDateOrg) SetAwsHostTop99pNil(b bool)`

 SetAwsHostTop99pNil sets the value for AwsHostTop99p to be an explicit nil

### UnsetAwsHostTop99p
`func (o *UsageSummaryDateOrg) UnsetAwsHostTop99p()`

UnsetAwsHostTop99p ensures that no value is present for AwsHostTop99p, not even an explicit nil
### GetAwsLambdaFuncCount

`func (o *UsageSummaryDateOrg) GetAwsLambdaFuncCount() int64`

GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field if non-nil, zero value otherwise.

### GetAwsLambdaFuncCountOk

`func (o *UsageSummaryDateOrg) GetAwsLambdaFuncCountOk() (*int64, bool)`

GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaFuncCount

`func (o *UsageSummaryDateOrg) SetAwsLambdaFuncCount(v int64)`

SetAwsLambdaFuncCount sets AwsLambdaFuncCount field to given value.

### HasAwsLambdaFuncCount

`func (o *UsageSummaryDateOrg) HasAwsLambdaFuncCount() bool`

HasAwsLambdaFuncCount returns a boolean if a field has been set.

### SetAwsLambdaFuncCountNil

`func (o *UsageSummaryDateOrg) SetAwsLambdaFuncCountNil(b bool)`

 SetAwsLambdaFuncCountNil sets the value for AwsLambdaFuncCount to be an explicit nil

### UnsetAwsLambdaFuncCount
`func (o *UsageSummaryDateOrg) UnsetAwsLambdaFuncCount()`

UnsetAwsLambdaFuncCount ensures that no value is present for AwsLambdaFuncCount, not even an explicit nil
### GetAwsLambdaInvocationsSum

`func (o *UsageSummaryDateOrg) GetAwsLambdaInvocationsSum() int64`

GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field if non-nil, zero value otherwise.

### GetAwsLambdaInvocationsSumOk

`func (o *UsageSummaryDateOrg) GetAwsLambdaInvocationsSumOk() (*int64, bool)`

GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambdaInvocationsSum

`func (o *UsageSummaryDateOrg) SetAwsLambdaInvocationsSum(v int64)`

SetAwsLambdaInvocationsSum sets AwsLambdaInvocationsSum field to given value.

### HasAwsLambdaInvocationsSum

`func (o *UsageSummaryDateOrg) HasAwsLambdaInvocationsSum() bool`

HasAwsLambdaInvocationsSum returns a boolean if a field has been set.

### SetAwsLambdaInvocationsSumNil

`func (o *UsageSummaryDateOrg) SetAwsLambdaInvocationsSumNil(b bool)`

 SetAwsLambdaInvocationsSumNil sets the value for AwsLambdaInvocationsSum to be an explicit nil

### UnsetAwsLambdaInvocationsSum
`func (o *UsageSummaryDateOrg) UnsetAwsLambdaInvocationsSum()`

UnsetAwsLambdaInvocationsSum ensures that no value is present for AwsLambdaInvocationsSum, not even an explicit nil
### GetAzureAppServiceTop99p

`func (o *UsageSummaryDateOrg) GetAzureAppServiceTop99p() int64`

GetAzureAppServiceTop99p returns the AzureAppServiceTop99p field if non-nil, zero value otherwise.

### GetAzureAppServiceTop99pOk

`func (o *UsageSummaryDateOrg) GetAzureAppServiceTop99pOk() (*int64, bool)`

GetAzureAppServiceTop99pOk returns a tuple with the AzureAppServiceTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppServiceTop99p

`func (o *UsageSummaryDateOrg) SetAzureAppServiceTop99p(v int64)`

SetAzureAppServiceTop99p sets AzureAppServiceTop99p field to given value.

### HasAzureAppServiceTop99p

`func (o *UsageSummaryDateOrg) HasAzureAppServiceTop99p() bool`

HasAzureAppServiceTop99p returns a boolean if a field has been set.

### SetAzureAppServiceTop99pNil

`func (o *UsageSummaryDateOrg) SetAzureAppServiceTop99pNil(b bool)`

 SetAzureAppServiceTop99pNil sets the value for AzureAppServiceTop99p to be an explicit nil

### UnsetAzureAppServiceTop99p
`func (o *UsageSummaryDateOrg) UnsetAzureAppServiceTop99p()`

UnsetAzureAppServiceTop99p ensures that no value is present for AzureAppServiceTop99p, not even an explicit nil
### GetBillableIngestedBytesSum

`func (o *UsageSummaryDateOrg) GetBillableIngestedBytesSum() int64`

GetBillableIngestedBytesSum returns the BillableIngestedBytesSum field if non-nil, zero value otherwise.

### GetBillableIngestedBytesSumOk

`func (o *UsageSummaryDateOrg) GetBillableIngestedBytesSumOk() (*int64, bool)`

GetBillableIngestedBytesSumOk returns a tuple with the BillableIngestedBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableIngestedBytesSum

`func (o *UsageSummaryDateOrg) SetBillableIngestedBytesSum(v int64)`

SetBillableIngestedBytesSum sets BillableIngestedBytesSum field to given value.

### HasBillableIngestedBytesSum

`func (o *UsageSummaryDateOrg) HasBillableIngestedBytesSum() bool`

HasBillableIngestedBytesSum returns a boolean if a field has been set.

### SetBillableIngestedBytesSumNil

`func (o *UsageSummaryDateOrg) SetBillableIngestedBytesSumNil(b bool)`

 SetBillableIngestedBytesSumNil sets the value for BillableIngestedBytesSum to be an explicit nil

### UnsetBillableIngestedBytesSum
`func (o *UsageSummaryDateOrg) UnsetBillableIngestedBytesSum()`

UnsetBillableIngestedBytesSum ensures that no value is present for BillableIngestedBytesSum, not even an explicit nil
### GetComplianceContainerAggSum

`func (o *UsageSummaryDateOrg) GetComplianceContainerAggSum() int64`

GetComplianceContainerAggSum returns the ComplianceContainerAggSum field if non-nil, zero value otherwise.

### GetComplianceContainerAggSumOk

`func (o *UsageSummaryDateOrg) GetComplianceContainerAggSumOk() (*int64, bool)`

GetComplianceContainerAggSumOk returns a tuple with the ComplianceContainerAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceContainerAggSum

`func (o *UsageSummaryDateOrg) SetComplianceContainerAggSum(v int64)`

SetComplianceContainerAggSum sets ComplianceContainerAggSum field to given value.

### HasComplianceContainerAggSum

`func (o *UsageSummaryDateOrg) HasComplianceContainerAggSum() bool`

HasComplianceContainerAggSum returns a boolean if a field has been set.

### SetComplianceContainerAggSumNil

`func (o *UsageSummaryDateOrg) SetComplianceContainerAggSumNil(b bool)`

 SetComplianceContainerAggSumNil sets the value for ComplianceContainerAggSum to be an explicit nil

### UnsetComplianceContainerAggSum
`func (o *UsageSummaryDateOrg) UnsetComplianceContainerAggSum()`

UnsetComplianceContainerAggSum ensures that no value is present for ComplianceContainerAggSum, not even an explicit nil
### GetComplianceHostAggSum

`func (o *UsageSummaryDateOrg) GetComplianceHostAggSum() int64`

GetComplianceHostAggSum returns the ComplianceHostAggSum field if non-nil, zero value otherwise.

### GetComplianceHostAggSumOk

`func (o *UsageSummaryDateOrg) GetComplianceHostAggSumOk() (*int64, bool)`

GetComplianceHostAggSumOk returns a tuple with the ComplianceHostAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceHostAggSum

`func (o *UsageSummaryDateOrg) SetComplianceHostAggSum(v int64)`

SetComplianceHostAggSum sets ComplianceHostAggSum field to given value.

### HasComplianceHostAggSum

`func (o *UsageSummaryDateOrg) HasComplianceHostAggSum() bool`

HasComplianceHostAggSum returns a boolean if a field has been set.

### SetComplianceHostAggSumNil

`func (o *UsageSummaryDateOrg) SetComplianceHostAggSumNil(b bool)`

 SetComplianceHostAggSumNil sets the value for ComplianceHostAggSum to be an explicit nil

### UnsetComplianceHostAggSum
`func (o *UsageSummaryDateOrg) UnsetComplianceHostAggSum()`

UnsetComplianceHostAggSum ensures that no value is present for ComplianceHostAggSum, not even an explicit nil
### GetContainerAvg

`func (o *UsageSummaryDateOrg) GetContainerAvg() int64`

GetContainerAvg returns the ContainerAvg field if non-nil, zero value otherwise.

### GetContainerAvgOk

`func (o *UsageSummaryDateOrg) GetContainerAvgOk() (*int64, bool)`

GetContainerAvgOk returns a tuple with the ContainerAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerAvg

`func (o *UsageSummaryDateOrg) SetContainerAvg(v int64)`

SetContainerAvg sets ContainerAvg field to given value.

### HasContainerAvg

`func (o *UsageSummaryDateOrg) HasContainerAvg() bool`

HasContainerAvg returns a boolean if a field has been set.

### SetContainerAvgNil

`func (o *UsageSummaryDateOrg) SetContainerAvgNil(b bool)`

 SetContainerAvgNil sets the value for ContainerAvg to be an explicit nil

### UnsetContainerAvg
`func (o *UsageSummaryDateOrg) UnsetContainerAvg()`

UnsetContainerAvg ensures that no value is present for ContainerAvg, not even an explicit nil
### GetContainerHwm

`func (o *UsageSummaryDateOrg) GetContainerHwm() int64`

GetContainerHwm returns the ContainerHwm field if non-nil, zero value otherwise.

### GetContainerHwmOk

`func (o *UsageSummaryDateOrg) GetContainerHwmOk() (*int64, bool)`

GetContainerHwmOk returns a tuple with the ContainerHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHwm

`func (o *UsageSummaryDateOrg) SetContainerHwm(v int64)`

SetContainerHwm sets ContainerHwm field to given value.

### HasContainerHwm

`func (o *UsageSummaryDateOrg) HasContainerHwm() bool`

HasContainerHwm returns a boolean if a field has been set.

### SetContainerHwmNil

`func (o *UsageSummaryDateOrg) SetContainerHwmNil(b bool)`

 SetContainerHwmNil sets the value for ContainerHwm to be an explicit nil

### UnsetContainerHwm
`func (o *UsageSummaryDateOrg) UnsetContainerHwm()`

UnsetContainerHwm ensures that no value is present for ContainerHwm, not even an explicit nil
### GetCustomTsAvg

`func (o *UsageSummaryDateOrg) GetCustomTsAvg() int64`

GetCustomTsAvg returns the CustomTsAvg field if non-nil, zero value otherwise.

### GetCustomTsAvgOk

`func (o *UsageSummaryDateOrg) GetCustomTsAvgOk() (*int64, bool)`

GetCustomTsAvgOk returns a tuple with the CustomTsAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTsAvg

`func (o *UsageSummaryDateOrg) SetCustomTsAvg(v int64)`

SetCustomTsAvg sets CustomTsAvg field to given value.

### HasCustomTsAvg

`func (o *UsageSummaryDateOrg) HasCustomTsAvg() bool`

HasCustomTsAvg returns a boolean if a field has been set.

### SetCustomTsAvgNil

`func (o *UsageSummaryDateOrg) SetCustomTsAvgNil(b bool)`

 SetCustomTsAvgNil sets the value for CustomTsAvg to be an explicit nil

### UnsetCustomTsAvg
`func (o *UsageSummaryDateOrg) UnsetCustomTsAvg()`

UnsetCustomTsAvg ensures that no value is present for CustomTsAvg, not even an explicit nil
### GetFargateTasksCountAvg

`func (o *UsageSummaryDateOrg) GetFargateTasksCountAvg() int64`

GetFargateTasksCountAvg returns the FargateTasksCountAvg field if non-nil, zero value otherwise.

### GetFargateTasksCountAvgOk

`func (o *UsageSummaryDateOrg) GetFargateTasksCountAvgOk() (*int64, bool)`

GetFargateTasksCountAvgOk returns a tuple with the FargateTasksCountAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountAvg

`func (o *UsageSummaryDateOrg) SetFargateTasksCountAvg(v int64)`

SetFargateTasksCountAvg sets FargateTasksCountAvg field to given value.

### HasFargateTasksCountAvg

`func (o *UsageSummaryDateOrg) HasFargateTasksCountAvg() bool`

HasFargateTasksCountAvg returns a boolean if a field has been set.

### SetFargateTasksCountAvgNil

`func (o *UsageSummaryDateOrg) SetFargateTasksCountAvgNil(b bool)`

 SetFargateTasksCountAvgNil sets the value for FargateTasksCountAvg to be an explicit nil

### UnsetFargateTasksCountAvg
`func (o *UsageSummaryDateOrg) UnsetFargateTasksCountAvg()`

UnsetFargateTasksCountAvg ensures that no value is present for FargateTasksCountAvg, not even an explicit nil
### GetFargateTasksCountHwm

`func (o *UsageSummaryDateOrg) GetFargateTasksCountHwm() int64`

GetFargateTasksCountHwm returns the FargateTasksCountHwm field if non-nil, zero value otherwise.

### GetFargateTasksCountHwmOk

`func (o *UsageSummaryDateOrg) GetFargateTasksCountHwmOk() (*int64, bool)`

GetFargateTasksCountHwmOk returns a tuple with the FargateTasksCountHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateTasksCountHwm

`func (o *UsageSummaryDateOrg) SetFargateTasksCountHwm(v int64)`

SetFargateTasksCountHwm sets FargateTasksCountHwm field to given value.

### HasFargateTasksCountHwm

`func (o *UsageSummaryDateOrg) HasFargateTasksCountHwm() bool`

HasFargateTasksCountHwm returns a boolean if a field has been set.

### SetFargateTasksCountHwmNil

`func (o *UsageSummaryDateOrg) SetFargateTasksCountHwmNil(b bool)`

 SetFargateTasksCountHwmNil sets the value for FargateTasksCountHwm to be an explicit nil

### UnsetFargateTasksCountHwm
`func (o *UsageSummaryDateOrg) UnsetFargateTasksCountHwm()`

UnsetFargateTasksCountHwm ensures that no value is present for FargateTasksCountHwm, not even an explicit nil
### GetGcpHostTop99p

`func (o *UsageSummaryDateOrg) GetGcpHostTop99p() int64`

GetGcpHostTop99p returns the GcpHostTop99p field if non-nil, zero value otherwise.

### GetGcpHostTop99pOk

`func (o *UsageSummaryDateOrg) GetGcpHostTop99pOk() (*int64, bool)`

GetGcpHostTop99pOk returns a tuple with the GcpHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpHostTop99p

`func (o *UsageSummaryDateOrg) SetGcpHostTop99p(v int64)`

SetGcpHostTop99p sets GcpHostTop99p field to given value.

### HasGcpHostTop99p

`func (o *UsageSummaryDateOrg) HasGcpHostTop99p() bool`

HasGcpHostTop99p returns a boolean if a field has been set.

### SetGcpHostTop99pNil

`func (o *UsageSummaryDateOrg) SetGcpHostTop99pNil(b bool)`

 SetGcpHostTop99pNil sets the value for GcpHostTop99p to be an explicit nil

### UnsetGcpHostTop99p
`func (o *UsageSummaryDateOrg) UnsetGcpHostTop99p()`

UnsetGcpHostTop99p ensures that no value is present for GcpHostTop99p, not even an explicit nil
### GetHerokuHostTop99p

`func (o *UsageSummaryDateOrg) GetHerokuHostTop99p() int64`

GetHerokuHostTop99p returns the HerokuHostTop99p field if non-nil, zero value otherwise.

### GetHerokuHostTop99pOk

`func (o *UsageSummaryDateOrg) GetHerokuHostTop99pOk() (*int64, bool)`

GetHerokuHostTop99pOk returns a tuple with the HerokuHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHerokuHostTop99p

`func (o *UsageSummaryDateOrg) SetHerokuHostTop99p(v int64)`

SetHerokuHostTop99p sets HerokuHostTop99p field to given value.

### HasHerokuHostTop99p

`func (o *UsageSummaryDateOrg) HasHerokuHostTop99p() bool`

HasHerokuHostTop99p returns a boolean if a field has been set.

### SetHerokuHostTop99pNil

`func (o *UsageSummaryDateOrg) SetHerokuHostTop99pNil(b bool)`

 SetHerokuHostTop99pNil sets the value for HerokuHostTop99p to be an explicit nil

### UnsetHerokuHostTop99p
`func (o *UsageSummaryDateOrg) UnsetHerokuHostTop99p()`

UnsetHerokuHostTop99p ensures that no value is present for HerokuHostTop99p, not even an explicit nil
### GetId

`func (o *UsageSummaryDateOrg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageSummaryDateOrg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageSummaryDateOrg) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsageSummaryDateOrg) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDateOrg) GetIncidentManagementMonthlyActiveUsersHwm() int64`

GetIncidentManagementMonthlyActiveUsersHwm returns the IncidentManagementMonthlyActiveUsersHwm field if non-nil, zero value otherwise.

### GetIncidentManagementMonthlyActiveUsersHwmOk

`func (o *UsageSummaryDateOrg) GetIncidentManagementMonthlyActiveUsersHwmOk() (*int64, bool)`

GetIncidentManagementMonthlyActiveUsersHwmOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDateOrg) SetIncidentManagementMonthlyActiveUsersHwm(v int64)`

SetIncidentManagementMonthlyActiveUsersHwm sets IncidentManagementMonthlyActiveUsersHwm field to given value.

### HasIncidentManagementMonthlyActiveUsersHwm

`func (o *UsageSummaryDateOrg) HasIncidentManagementMonthlyActiveUsersHwm() bool`

HasIncidentManagementMonthlyActiveUsersHwm returns a boolean if a field has been set.

### SetIncidentManagementMonthlyActiveUsersHwmNil

`func (o *UsageSummaryDateOrg) SetIncidentManagementMonthlyActiveUsersHwmNil(b bool)`

 SetIncidentManagementMonthlyActiveUsersHwmNil sets the value for IncidentManagementMonthlyActiveUsersHwm to be an explicit nil

### UnsetIncidentManagementMonthlyActiveUsersHwm
`func (o *UsageSummaryDateOrg) UnsetIncidentManagementMonthlyActiveUsersHwm()`

UnsetIncidentManagementMonthlyActiveUsersHwm ensures that no value is present for IncidentManagementMonthlyActiveUsersHwm, not even an explicit nil
### GetIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) GetIndexedEventsCountSum() int64`

GetIndexedEventsCountSum returns the IndexedEventsCountSum field if non-nil, zero value otherwise.

### GetIndexedEventsCountSumOk

`func (o *UsageSummaryDateOrg) GetIndexedEventsCountSumOk() (*int64, bool)`

GetIndexedEventsCountSumOk returns a tuple with the IndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) SetIndexedEventsCountSum(v int64)`

SetIndexedEventsCountSum sets IndexedEventsCountSum field to given value.

### HasIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) HasIndexedEventsCountSum() bool`

HasIndexedEventsCountSum returns a boolean if a field has been set.

### SetIndexedEventsCountSumNil

`func (o *UsageSummaryDateOrg) SetIndexedEventsCountSumNil(b bool)`

 SetIndexedEventsCountSumNil sets the value for IndexedEventsCountSum to be an explicit nil

### UnsetIndexedEventsCountSum
`func (o *UsageSummaryDateOrg) UnsetIndexedEventsCountSum()`

UnsetIndexedEventsCountSum ensures that no value is present for IndexedEventsCountSum, not even an explicit nil
### GetInfraHostTop99p

`func (o *UsageSummaryDateOrg) GetInfraHostTop99p() int64`

GetInfraHostTop99p returns the InfraHostTop99p field if non-nil, zero value otherwise.

### GetInfraHostTop99pOk

`func (o *UsageSummaryDateOrg) GetInfraHostTop99pOk() (*int64, bool)`

GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostTop99p

`func (o *UsageSummaryDateOrg) SetInfraHostTop99p(v int64)`

SetInfraHostTop99p sets InfraHostTop99p field to given value.

### HasInfraHostTop99p

`func (o *UsageSummaryDateOrg) HasInfraHostTop99p() bool`

HasInfraHostTop99p returns a boolean if a field has been set.

### SetInfraHostTop99pNil

`func (o *UsageSummaryDateOrg) SetInfraHostTop99pNil(b bool)`

 SetInfraHostTop99pNil sets the value for InfraHostTop99p to be an explicit nil

### UnsetInfraHostTop99p
`func (o *UsageSummaryDateOrg) UnsetInfraHostTop99p()`

UnsetInfraHostTop99p ensures that no value is present for InfraHostTop99p, not even an explicit nil
### GetIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) GetIngestedEventsBytesSum() int64`

GetIngestedEventsBytesSum returns the IngestedEventsBytesSum field if non-nil, zero value otherwise.

### GetIngestedEventsBytesSumOk

`func (o *UsageSummaryDateOrg) GetIngestedEventsBytesSumOk() (*int64, bool)`

GetIngestedEventsBytesSumOk returns a tuple with the IngestedEventsBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) SetIngestedEventsBytesSum(v int64)`

SetIngestedEventsBytesSum sets IngestedEventsBytesSum field to given value.

### HasIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) HasIngestedEventsBytesSum() bool`

HasIngestedEventsBytesSum returns a boolean if a field has been set.

### SetIngestedEventsBytesSumNil

`func (o *UsageSummaryDateOrg) SetIngestedEventsBytesSumNil(b bool)`

 SetIngestedEventsBytesSumNil sets the value for IngestedEventsBytesSum to be an explicit nil

### UnsetIngestedEventsBytesSum
`func (o *UsageSummaryDateOrg) UnsetIngestedEventsBytesSum()`

UnsetIngestedEventsBytesSum ensures that no value is present for IngestedEventsBytesSum, not even an explicit nil
### GetIotDeviceAggSum

`func (o *UsageSummaryDateOrg) GetIotDeviceAggSum() int64`

GetIotDeviceAggSum returns the IotDeviceAggSum field if non-nil, zero value otherwise.

### GetIotDeviceAggSumOk

`func (o *UsageSummaryDateOrg) GetIotDeviceAggSumOk() (*int64, bool)`

GetIotDeviceAggSumOk returns a tuple with the IotDeviceAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceAggSum

`func (o *UsageSummaryDateOrg) SetIotDeviceAggSum(v int64)`

SetIotDeviceAggSum sets IotDeviceAggSum field to given value.

### HasIotDeviceAggSum

`func (o *UsageSummaryDateOrg) HasIotDeviceAggSum() bool`

HasIotDeviceAggSum returns a boolean if a field has been set.

### SetIotDeviceAggSumNil

`func (o *UsageSummaryDateOrg) SetIotDeviceAggSumNil(b bool)`

 SetIotDeviceAggSumNil sets the value for IotDeviceAggSum to be an explicit nil

### UnsetIotDeviceAggSum
`func (o *UsageSummaryDateOrg) UnsetIotDeviceAggSum()`

UnsetIotDeviceAggSum ensures that no value is present for IotDeviceAggSum, not even an explicit nil
### GetIotDeviceTop99pSum

`func (o *UsageSummaryDateOrg) GetIotDeviceTop99pSum() int64`

GetIotDeviceTop99pSum returns the IotDeviceTop99pSum field if non-nil, zero value otherwise.

### GetIotDeviceTop99pSumOk

`func (o *UsageSummaryDateOrg) GetIotDeviceTop99pSumOk() (*int64, bool)`

GetIotDeviceTop99pSumOk returns a tuple with the IotDeviceTop99pSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotDeviceTop99pSum

`func (o *UsageSummaryDateOrg) SetIotDeviceTop99pSum(v int64)`

SetIotDeviceTop99pSum sets IotDeviceTop99pSum field to given value.

### HasIotDeviceTop99pSum

`func (o *UsageSummaryDateOrg) HasIotDeviceTop99pSum() bool`

HasIotDeviceTop99pSum returns a boolean if a field has been set.

### SetIotDeviceTop99pSumNil

`func (o *UsageSummaryDateOrg) SetIotDeviceTop99pSumNil(b bool)`

 SetIotDeviceTop99pSumNil sets the value for IotDeviceTop99pSum to be an explicit nil

### UnsetIotDeviceTop99pSum
`func (o *UsageSummaryDateOrg) UnsetIotDeviceTop99pSum()`

UnsetIotDeviceTop99pSum ensures that no value is present for IotDeviceTop99pSum, not even an explicit nil
### GetMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountAndroidSum() int64`

GetMobileRumSessionCountAndroidSum returns the MobileRumSessionCountAndroidSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountAndroidSumOk

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountAndroidSumOk() (*int64, bool)`

GetMobileRumSessionCountAndroidSumOk returns a tuple with the MobileRumSessionCountAndroidSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountAndroidSum(v int64)`

SetMobileRumSessionCountAndroidSum sets MobileRumSessionCountAndroidSum field to given value.

### HasMobileRumSessionCountAndroidSum

`func (o *UsageSummaryDateOrg) HasMobileRumSessionCountAndroidSum() bool`

HasMobileRumSessionCountAndroidSum returns a boolean if a field has been set.

### SetMobileRumSessionCountAndroidSumNil

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountAndroidSumNil(b bool)`

 SetMobileRumSessionCountAndroidSumNil sets the value for MobileRumSessionCountAndroidSum to be an explicit nil

### UnsetMobileRumSessionCountAndroidSum
`func (o *UsageSummaryDateOrg) UnsetMobileRumSessionCountAndroidSum()`

UnsetMobileRumSessionCountAndroidSum ensures that no value is present for MobileRumSessionCountAndroidSum, not even an explicit nil
### GetMobileRumSessionCountIosSum

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountIosSum() int64`

GetMobileRumSessionCountIosSum returns the MobileRumSessionCountIosSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountIosSumOk

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountIosSumOk() (*int64, bool)`

GetMobileRumSessionCountIosSumOk returns a tuple with the MobileRumSessionCountIosSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountIosSum

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountIosSum(v int64)`

SetMobileRumSessionCountIosSum sets MobileRumSessionCountIosSum field to given value.

### HasMobileRumSessionCountIosSum

`func (o *UsageSummaryDateOrg) HasMobileRumSessionCountIosSum() bool`

HasMobileRumSessionCountIosSum returns a boolean if a field has been set.

### SetMobileRumSessionCountIosSumNil

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountIosSumNil(b bool)`

 SetMobileRumSessionCountIosSumNil sets the value for MobileRumSessionCountIosSum to be an explicit nil

### UnsetMobileRumSessionCountIosSum
`func (o *UsageSummaryDateOrg) UnsetMobileRumSessionCountIosSum()`

UnsetMobileRumSessionCountIosSum ensures that no value is present for MobileRumSessionCountIosSum, not even an explicit nil
### GetMobileRumSessionCountSum

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountSum() int64`

GetMobileRumSessionCountSum returns the MobileRumSessionCountSum field if non-nil, zero value otherwise.

### GetMobileRumSessionCountSumOk

`func (o *UsageSummaryDateOrg) GetMobileRumSessionCountSumOk() (*int64, bool)`

GetMobileRumSessionCountSumOk returns a tuple with the MobileRumSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileRumSessionCountSum

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountSum(v int64)`

SetMobileRumSessionCountSum sets MobileRumSessionCountSum field to given value.

### HasMobileRumSessionCountSum

`func (o *UsageSummaryDateOrg) HasMobileRumSessionCountSum() bool`

HasMobileRumSessionCountSum returns a boolean if a field has been set.

### SetMobileRumSessionCountSumNil

`func (o *UsageSummaryDateOrg) SetMobileRumSessionCountSumNil(b bool)`

 SetMobileRumSessionCountSumNil sets the value for MobileRumSessionCountSum to be an explicit nil

### UnsetMobileRumSessionCountSum
`func (o *UsageSummaryDateOrg) UnsetMobileRumSessionCountSum()`

UnsetMobileRumSessionCountSum ensures that no value is present for MobileRumSessionCountSum, not even an explicit nil
### GetName

`func (o *UsageSummaryDateOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsageSummaryDateOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsageSummaryDateOrg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsageSummaryDateOrg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetflowIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) GetNetflowIndexedEventsCountSum() int64`

GetNetflowIndexedEventsCountSum returns the NetflowIndexedEventsCountSum field if non-nil, zero value otherwise.

### GetNetflowIndexedEventsCountSumOk

`func (o *UsageSummaryDateOrg) GetNetflowIndexedEventsCountSumOk() (*int64, bool)`

GetNetflowIndexedEventsCountSumOk returns a tuple with the NetflowIndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetflowIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) SetNetflowIndexedEventsCountSum(v int64)`

SetNetflowIndexedEventsCountSum sets NetflowIndexedEventsCountSum field to given value.

### HasNetflowIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) HasNetflowIndexedEventsCountSum() bool`

HasNetflowIndexedEventsCountSum returns a boolean if a field has been set.

### SetNetflowIndexedEventsCountSumNil

`func (o *UsageSummaryDateOrg) SetNetflowIndexedEventsCountSumNil(b bool)`

 SetNetflowIndexedEventsCountSumNil sets the value for NetflowIndexedEventsCountSum to be an explicit nil

### UnsetNetflowIndexedEventsCountSum
`func (o *UsageSummaryDateOrg) UnsetNetflowIndexedEventsCountSum()`

UnsetNetflowIndexedEventsCountSum ensures that no value is present for NetflowIndexedEventsCountSum, not even an explicit nil
### GetNpmHostTop99p

`func (o *UsageSummaryDateOrg) GetNpmHostTop99p() int64`

GetNpmHostTop99p returns the NpmHostTop99p field if non-nil, zero value otherwise.

### GetNpmHostTop99pOk

`func (o *UsageSummaryDateOrg) GetNpmHostTop99pOk() (*int64, bool)`

GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostTop99p

`func (o *UsageSummaryDateOrg) SetNpmHostTop99p(v int64)`

SetNpmHostTop99p sets NpmHostTop99p field to given value.

### HasNpmHostTop99p

`func (o *UsageSummaryDateOrg) HasNpmHostTop99p() bool`

HasNpmHostTop99p returns a boolean if a field has been set.

### SetNpmHostTop99pNil

`func (o *UsageSummaryDateOrg) SetNpmHostTop99pNil(b bool)`

 SetNpmHostTop99pNil sets the value for NpmHostTop99p to be an explicit nil

### UnsetNpmHostTop99p
`func (o *UsageSummaryDateOrg) UnsetNpmHostTop99p()`

UnsetNpmHostTop99p ensures that no value is present for NpmHostTop99p, not even an explicit nil
### GetOpentelemetryHostTop99p

`func (o *UsageSummaryDateOrg) GetOpentelemetryHostTop99p() int64`

GetOpentelemetryHostTop99p returns the OpentelemetryHostTop99p field if non-nil, zero value otherwise.

### GetOpentelemetryHostTop99pOk

`func (o *UsageSummaryDateOrg) GetOpentelemetryHostTop99pOk() (*int64, bool)`

GetOpentelemetryHostTop99pOk returns a tuple with the OpentelemetryHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentelemetryHostTop99p

`func (o *UsageSummaryDateOrg) SetOpentelemetryHostTop99p(v int64)`

SetOpentelemetryHostTop99p sets OpentelemetryHostTop99p field to given value.

### HasOpentelemetryHostTop99p

`func (o *UsageSummaryDateOrg) HasOpentelemetryHostTop99p() bool`

HasOpentelemetryHostTop99p returns a boolean if a field has been set.

### SetOpentelemetryHostTop99pNil

`func (o *UsageSummaryDateOrg) SetOpentelemetryHostTop99pNil(b bool)`

 SetOpentelemetryHostTop99pNil sets the value for OpentelemetryHostTop99p to be an explicit nil

### UnsetOpentelemetryHostTop99p
`func (o *UsageSummaryDateOrg) UnsetOpentelemetryHostTop99p()`

UnsetOpentelemetryHostTop99p ensures that no value is present for OpentelemetryHostTop99p, not even an explicit nil
### GetProfilingHostTop99p

`func (o *UsageSummaryDateOrg) GetProfilingHostTop99p() int64`

GetProfilingHostTop99p returns the ProfilingHostTop99p field if non-nil, zero value otherwise.

### GetProfilingHostTop99pOk

`func (o *UsageSummaryDateOrg) GetProfilingHostTop99pOk() (*int64, bool)`

GetProfilingHostTop99pOk returns a tuple with the ProfilingHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilingHostTop99p

`func (o *UsageSummaryDateOrg) SetProfilingHostTop99p(v int64)`

SetProfilingHostTop99p sets ProfilingHostTop99p field to given value.

### HasProfilingHostTop99p

`func (o *UsageSummaryDateOrg) HasProfilingHostTop99p() bool`

HasProfilingHostTop99p returns a boolean if a field has been set.

### SetProfilingHostTop99pNil

`func (o *UsageSummaryDateOrg) SetProfilingHostTop99pNil(b bool)`

 SetProfilingHostTop99pNil sets the value for ProfilingHostTop99p to be an explicit nil

### UnsetProfilingHostTop99p
`func (o *UsageSummaryDateOrg) UnsetProfilingHostTop99p()`

UnsetProfilingHostTop99p ensures that no value is present for ProfilingHostTop99p, not even an explicit nil
### GetPublicId

`func (o *UsageSummaryDateOrg) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageSummaryDateOrg) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageSummaryDateOrg) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageSummaryDateOrg) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetRumSessionCountSum

`func (o *UsageSummaryDateOrg) GetRumSessionCountSum() int64`

GetRumSessionCountSum returns the RumSessionCountSum field if non-nil, zero value otherwise.

### GetRumSessionCountSumOk

`func (o *UsageSummaryDateOrg) GetRumSessionCountSumOk() (*int64, bool)`

GetRumSessionCountSumOk returns a tuple with the RumSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumSessionCountSum

`func (o *UsageSummaryDateOrg) SetRumSessionCountSum(v int64)`

SetRumSessionCountSum sets RumSessionCountSum field to given value.

### HasRumSessionCountSum

`func (o *UsageSummaryDateOrg) HasRumSessionCountSum() bool`

HasRumSessionCountSum returns a boolean if a field has been set.

### SetRumSessionCountSumNil

`func (o *UsageSummaryDateOrg) SetRumSessionCountSumNil(b bool)`

 SetRumSessionCountSumNil sets the value for RumSessionCountSum to be an explicit nil

### UnsetRumSessionCountSum
`func (o *UsageSummaryDateOrg) UnsetRumSessionCountSum()`

UnsetRumSessionCountSum ensures that no value is present for RumSessionCountSum, not even an explicit nil
### GetRumTotalSessionCountSum

`func (o *UsageSummaryDateOrg) GetRumTotalSessionCountSum() int64`

GetRumTotalSessionCountSum returns the RumTotalSessionCountSum field if non-nil, zero value otherwise.

### GetRumTotalSessionCountSumOk

`func (o *UsageSummaryDateOrg) GetRumTotalSessionCountSumOk() (*int64, bool)`

GetRumTotalSessionCountSumOk returns a tuple with the RumTotalSessionCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumTotalSessionCountSum

`func (o *UsageSummaryDateOrg) SetRumTotalSessionCountSum(v int64)`

SetRumTotalSessionCountSum sets RumTotalSessionCountSum field to given value.

### HasRumTotalSessionCountSum

`func (o *UsageSummaryDateOrg) HasRumTotalSessionCountSum() bool`

HasRumTotalSessionCountSum returns a boolean if a field has been set.

### SetRumTotalSessionCountSumNil

`func (o *UsageSummaryDateOrg) SetRumTotalSessionCountSumNil(b bool)`

 SetRumTotalSessionCountSumNil sets the value for RumTotalSessionCountSum to be an explicit nil

### UnsetRumTotalSessionCountSum
`func (o *UsageSummaryDateOrg) UnsetRumTotalSessionCountSum()`

UnsetRumTotalSessionCountSum ensures that no value is present for RumTotalSessionCountSum, not even an explicit nil
### GetSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDateOrg) GetSyntheticsBrowserCheckCallsCountSum() int64`

GetSyntheticsBrowserCheckCallsCountSum returns the SyntheticsBrowserCheckCallsCountSum field if non-nil, zero value otherwise.

### GetSyntheticsBrowserCheckCallsCountSumOk

`func (o *UsageSummaryDateOrg) GetSyntheticsBrowserCheckCallsCountSumOk() (*int64, bool)`

GetSyntheticsBrowserCheckCallsCountSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDateOrg) SetSyntheticsBrowserCheckCallsCountSum(v int64)`

SetSyntheticsBrowserCheckCallsCountSum sets SyntheticsBrowserCheckCallsCountSum field to given value.

### HasSyntheticsBrowserCheckCallsCountSum

`func (o *UsageSummaryDateOrg) HasSyntheticsBrowserCheckCallsCountSum() bool`

HasSyntheticsBrowserCheckCallsCountSum returns a boolean if a field has been set.

### SetSyntheticsBrowserCheckCallsCountSumNil

`func (o *UsageSummaryDateOrg) SetSyntheticsBrowserCheckCallsCountSumNil(b bool)`

 SetSyntheticsBrowserCheckCallsCountSumNil sets the value for SyntheticsBrowserCheckCallsCountSum to be an explicit nil

### UnsetSyntheticsBrowserCheckCallsCountSum
`func (o *UsageSummaryDateOrg) UnsetSyntheticsBrowserCheckCallsCountSum()`

UnsetSyntheticsBrowserCheckCallsCountSum ensures that no value is present for SyntheticsBrowserCheckCallsCountSum, not even an explicit nil
### GetSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDateOrg) GetSyntheticsCheckCallsCountSum() int64`

GetSyntheticsCheckCallsCountSum returns the SyntheticsCheckCallsCountSum field if non-nil, zero value otherwise.

### GetSyntheticsCheckCallsCountSumOk

`func (o *UsageSummaryDateOrg) GetSyntheticsCheckCallsCountSumOk() (*int64, bool)`

GetSyntheticsCheckCallsCountSumOk returns a tuple with the SyntheticsCheckCallsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDateOrg) SetSyntheticsCheckCallsCountSum(v int64)`

SetSyntheticsCheckCallsCountSum sets SyntheticsCheckCallsCountSum field to given value.

### HasSyntheticsCheckCallsCountSum

`func (o *UsageSummaryDateOrg) HasSyntheticsCheckCallsCountSum() bool`

HasSyntheticsCheckCallsCountSum returns a boolean if a field has been set.

### SetSyntheticsCheckCallsCountSumNil

`func (o *UsageSummaryDateOrg) SetSyntheticsCheckCallsCountSumNil(b bool)`

 SetSyntheticsCheckCallsCountSumNil sets the value for SyntheticsCheckCallsCountSum to be an explicit nil

### UnsetSyntheticsCheckCallsCountSum
`func (o *UsageSummaryDateOrg) UnsetSyntheticsCheckCallsCountSum()`

UnsetSyntheticsCheckCallsCountSum ensures that no value is present for SyntheticsCheckCallsCountSum, not even an explicit nil
### GetTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) GetTraceSearchIndexedEventsCountSum() int64`

GetTraceSearchIndexedEventsCountSum returns the TraceSearchIndexedEventsCountSum field if non-nil, zero value otherwise.

### GetTraceSearchIndexedEventsCountSumOk

`func (o *UsageSummaryDateOrg) GetTraceSearchIndexedEventsCountSumOk() (*int64, bool)`

GetTraceSearchIndexedEventsCountSumOk returns a tuple with the TraceSearchIndexedEventsCountSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) SetTraceSearchIndexedEventsCountSum(v int64)`

SetTraceSearchIndexedEventsCountSum sets TraceSearchIndexedEventsCountSum field to given value.

### HasTraceSearchIndexedEventsCountSum

`func (o *UsageSummaryDateOrg) HasTraceSearchIndexedEventsCountSum() bool`

HasTraceSearchIndexedEventsCountSum returns a boolean if a field has been set.

### SetTraceSearchIndexedEventsCountSumNil

`func (o *UsageSummaryDateOrg) SetTraceSearchIndexedEventsCountSumNil(b bool)`

 SetTraceSearchIndexedEventsCountSumNil sets the value for TraceSearchIndexedEventsCountSum to be an explicit nil

### UnsetTraceSearchIndexedEventsCountSum
`func (o *UsageSummaryDateOrg) UnsetTraceSearchIndexedEventsCountSum()`

UnsetTraceSearchIndexedEventsCountSum ensures that no value is present for TraceSearchIndexedEventsCountSum, not even an explicit nil
### GetTwolIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) GetTwolIngestedEventsBytesSum() int64`

GetTwolIngestedEventsBytesSum returns the TwolIngestedEventsBytesSum field if non-nil, zero value otherwise.

### GetTwolIngestedEventsBytesSumOk

`func (o *UsageSummaryDateOrg) GetTwolIngestedEventsBytesSumOk() (*int64, bool)`

GetTwolIngestedEventsBytesSumOk returns a tuple with the TwolIngestedEventsBytesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwolIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) SetTwolIngestedEventsBytesSum(v int64)`

SetTwolIngestedEventsBytesSum sets TwolIngestedEventsBytesSum field to given value.

### HasTwolIngestedEventsBytesSum

`func (o *UsageSummaryDateOrg) HasTwolIngestedEventsBytesSum() bool`

HasTwolIngestedEventsBytesSum returns a boolean if a field has been set.

### SetTwolIngestedEventsBytesSumNil

`func (o *UsageSummaryDateOrg) SetTwolIngestedEventsBytesSumNil(b bool)`

 SetTwolIngestedEventsBytesSumNil sets the value for TwolIngestedEventsBytesSum to be an explicit nil

### UnsetTwolIngestedEventsBytesSum
`func (o *UsageSummaryDateOrg) UnsetTwolIngestedEventsBytesSum()`

UnsetTwolIngestedEventsBytesSum ensures that no value is present for TwolIngestedEventsBytesSum, not even an explicit nil
### GetVsphereHostTop99p

`func (o *UsageSummaryDateOrg) GetVsphereHostTop99p() int64`

GetVsphereHostTop99p returns the VsphereHostTop99p field if non-nil, zero value otherwise.

### GetVsphereHostTop99pOk

`func (o *UsageSummaryDateOrg) GetVsphereHostTop99pOk() (*int64, bool)`

GetVsphereHostTop99pOk returns a tuple with the VsphereHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereHostTop99p

`func (o *UsageSummaryDateOrg) SetVsphereHostTop99p(v int64)`

SetVsphereHostTop99p sets VsphereHostTop99p field to given value.

### HasVsphereHostTop99p

`func (o *UsageSummaryDateOrg) HasVsphereHostTop99p() bool`

HasVsphereHostTop99p returns a boolean if a field has been set.

### SetVsphereHostTop99pNil

`func (o *UsageSummaryDateOrg) SetVsphereHostTop99pNil(b bool)`

 SetVsphereHostTop99pNil sets the value for VsphereHostTop99p to be an explicit nil

### UnsetVsphereHostTop99p
`func (o *UsageSummaryDateOrg) UnsetVsphereHostTop99p()`

UnsetVsphereHostTop99p ensures that no value is present for VsphereHostTop99p, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


