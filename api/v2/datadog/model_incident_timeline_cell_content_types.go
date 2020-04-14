/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

// GetType returns the cell_type type.
func (o *IncidentTimelineCellChoiceChangeContent) GetCellType() string {
	return "choice_change"
}

// GetType returns the cell_type type.
func (o *IncidentTimelineCellMarkdownContent) GetCellType() string {
	return "markdown"
}

// GetType returns the cell_type type.
func (o *IncidentTimelineCellStatusChangeContent) GetCellType() string {
	return "status_change"
}
