// Get all hosts with metadata deserializes successfully

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/winebarrel/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsApi.ListHosts(ctx, *datadog.NewListHostsOptionalParameters().WithIncludeHostsMetadata(true))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListHosts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `HostsApi.ListHosts`:\n%s\n", responseContent)
}
