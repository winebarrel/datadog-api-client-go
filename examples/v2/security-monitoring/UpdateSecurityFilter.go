// Update a security filter returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "security_filter" in the system
	SecurityFilterDataID := os.Getenv("SECURITY_FILTER_DATA_ID")

	body := datadog.SecurityFilterUpdateRequest{
		Data: datadog.SecurityFilterUpdateData{
			Attributes: datadog.SecurityFilterUpdateAttributes{
				ExclusionFilters: &[]datadog.SecurityFilterExclusionFilter{},
				FilteredDataType: datadog.SECURITYFILTERFILTEREDDATATYPE_LOGS.Ptr(),
				IsEnabled:        datadog.PtrBool(true),
				Name:             datadog.PtrString("Example-Update_a_security_filter_returns_OK_response"),
				Query:            datadog.PtrString("service:ExampleUpdateasecurityfilterreturnsOKresponse"),
				Version:          datadog.PtrInt32(1),
			},
			Type: datadog.SECURITYFILTERTYPE_SECURITY_FILTERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityMonitoringApi.UpdateSecurityFilter(ctx, SecurityFilterDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityFilter`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityFilter`:\n%s\n", responseContent)
}
