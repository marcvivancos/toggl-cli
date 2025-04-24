package toggl

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
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
	Stop        *string  `json:"stop"`
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
	param["start"] = timeEntry.Start
	param["duration"] = timeEntry.Duration
	param["workspace_id"] = timeEntry.WorkspaceID
	param["description"] = timeEntry.Description
	param["created_with"] = "marcvivancos/toggl-cli"
	return param
}

func (timeEntry TimeEntry) ParsedStart() (time.Time, error) {
	if timeEntry.Start == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, timeEntry.Start)
}

func (cl *Client) GetCurrentTimeEntry() (*TimeEntry, error) {
	var response TimeEntry

	res, err := cl.do("GET", "me/time_entries/current", nil, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching current time entry: %s", res.Status)
	}

	enc := json.NewDecoder(res.Body)
	if err := enc.Decode(&response); err != nil {
		return nil, err
	}

	if response.ID == 0 {
		return nil, nil
	}

	return &response, nil
}

func (cl *Client) PostStartTimeEntry(
	timeEntry TimeEntry,
) (response TimeEntry, err error) {
	timeEntry.Duration = -1

	res, err := cl.do(
		"POST",
		fmt.Sprintf("/workspaces/%d/time_entries", timeEntry.WorkspaceID),
		nil,
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

	res, err := cl.do(
		"PUT",
		fmt.Sprintf("/workspaces/%d/time_entries/%d", workspaceID, timeEntryID),
		nil,
		param,
	)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("error stopping time entry: %s", res.Status)
	}

	return nil
}

func (cl *Client) PutResumeTimeEntry(workspaceID int, timeEntryID int) error {
	param := make(map[string]interface{})
	param["stop"] = nil
	param["duration"] = -1

	res, err := cl.do(
		"PUT",
		fmt.Sprintf("/workspaces/%d/time_entries/%d", workspaceID, timeEntryID),
		nil,
		param,
	)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("error resuming time entry: %s", res.Status)
	}

	return nil
}

func (cl *Client) GetLatestTimeEntry() (*TimeEntry, error) {
	// First check if there's a running time entry
	currentTimeEntry, err := cl.GetCurrentTimeEntry()
	if err != nil {
		return nil, err
	}

	if currentTimeEntry != nil {
		return currentTimeEntry, nil
	}

	params := make(map[string]string)
	params["start_date"] = time.Now().Add(-time.Hour * 24).UTC().Format(time.RFC3339)
	params["end_date"] = time.Now().UTC().Format(time.RFC3339)

	res, err := cl.do("GET", "/me/time_entries", params, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error fetching time entries: %s", string(bodyBytes))
	}

	var timeEntries []TimeEntry
	enc := json.NewDecoder(res.Body)
	if err := enc.Decode(&timeEntries); err != nil {
		return nil, err
	}

	if len(timeEntries) == 0 {
		return nil, nil
	}

	// Sort by latest stop time
	sort.Slice(timeEntries, func(i, j int) bool {
		if timeEntries[i].Stop == nil {
			return true
		}
		if timeEntries[j].Stop == nil {
			return false
		}
		stopTimeI, err := time.Parse(time.RFC3339, *timeEntries[i].Stop)
		if err != nil {
			return true
		}
		stopTimeJ, err := time.Parse(time.RFC3339, *timeEntries[j].Stop)
		if err != nil {
			return false
		}
		return stopTimeI.After(stopTimeJ)
	})

	return &timeEntries[0], nil
}
