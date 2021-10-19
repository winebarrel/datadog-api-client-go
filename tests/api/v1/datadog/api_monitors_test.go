/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"

	"gopkg.in/h2non/gock.v1"
)

func deleteMonitor(ctx context.Context, t *testing.T, monitorID int64) {
	_, httpresp, err := Client(ctx).MonitorsApi.DeleteMonitor(ctx, monitorID)
	if httpresp.StatusCode != 200 || err != nil {
		t.Logf("Deleting Monitor: %v failed with %v, Another test may have already deleted this monitor.", monitorID, httpresp.StatusCode)
	}
}
