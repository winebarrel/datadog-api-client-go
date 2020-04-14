/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

// GetType returns the facet type
func (o *IncidentFacetPercentilesAggregationDataSchema) GetType() string {
	return "aggregation_percentiles"
}

// GetType returns the facet type
func (o *IncidentFacetStatsAggregationDataSchema) GetType() string {
	return "aggregation_stats"
}

// GetType returns the facet type.
func (o *IncidentFacetTermsAggregationDataSchema) GetType() string {
	return "aggregation_terms"
}
