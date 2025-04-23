package toggl

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeEntry struct {
	At          string   `json:"at"`
	Billable    bool     `json:"billable"`
	Description string   `json:"description"`
	Duration    int64    `json:"duration"`
	Duronly     bool     `json:"duronly"`
	ID          int      `json:"id"`
	Start       string   `json:"start"`
	Tags        []string `json:"tags"`
	UserID      int      `json:"user_id"`
	ProjectID   int      `json:"project_id"`
	WorkspaceID int      `json:"workspace_id"`
}

func (timeEntry TimeEntry) AddParam() interface{} {
	param := make(map[string]interface{})
	if timeEntry.ProjectID != 0 {
		param["project_id"] = timeEntry.ProjectID
	}
	param["start"] = time.Now().UTC().Format(time.RFC3339)
	param["duration"] = timeEntry.Duration
	param["workspace_id"] = timeEntry.WorkspaceID
	param["description"] = timeEntry.Description
	param["created_with"] = "sachaos/toggl"
	return param
}

func (timeEntry TimeEntry) ParsedStart() (time.Time, error) {
	if timeEntry.Start == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, timeEntry.Start)
}

func (cl *Client) GetCurrentTimeEntry() (TimeEntry, error) {
	var response TimeEntry

	res, err := cl.do("GET", "me/time_entries/current", nil)
	if err != nil {
		return TimeEntry{}, err
	}

	enc := json.NewDecoder(res.Body)
	if err := enc.Decode(&response); err != nil {
		return TimeEntry{}, err
	}

	return response, nil
}

func (cl *Client) PostStartTimeEntry(
	timeEntry TimeEntry,
) (response TimeEntry, err error) {
	timeEntry.Duration = -1

	res, err := cl.do(
		"POST",
		fmt.Sprintf("/workspaces/%d/time_entries", timeEntry.WorkspaceID),
		timeEntry.AddParam(),
	)
	if err != nil {
		return TimeEntry{}, err
	}

	enc := json.NewDecoder(res.Body)
	if err := enc.Decode(&response); err != nil {
		return TimeEntry{}, err
	}

	return response, nil
}

func (cl *Client) PutStopTimeEntry(workspaceID int, timeEntryID int, stopTime time.Time) error {
	param := make(map[string]interface{})
	param["stop"] = stopTime.Format(time.RFC3339)

	_, err := cl.do(
		"PUT",
		fmt.Sprintf("/workspaces/%d/time_entries/%d", workspaceID, timeEntryID),
		param,
	)

	return err
}


