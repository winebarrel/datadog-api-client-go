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
)

func cancelDowntime(ctx context.Context, t *testing.T, downtimeID int64) {
	_, err := Client(ctx).DowntimesApi.CancelDowntime(ctx, downtimeID)
	if err != nil {
		t.Logf("Canceling Downtime: %v failed, Another test may have already canceled this downtime: %v", downtimeID, err)
	}
}
