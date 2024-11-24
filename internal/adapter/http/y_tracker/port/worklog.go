package port

import (
	"fmt"
	"time"

	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

func GetGetWorkLogsRequest(createdBy string, from, to time.Time) map[string][]string {
	params := map[string][]string{
		"createdBy": {createdBy},
		"createdAt": {"from:" + marshalTime(from), "to:" + marshalTime(to)},
	}
	return params
}

type WorkLogsResponse []WorkLogResponse

func (w WorkLogsResponse) ToCore() (ytrackercore.WorkLogs, error) {
	res := make(ytrackercore.WorkLogs, len(w))
	for i, workLog := range w {
		coreWorkLog, err := workLog.ToCore()
		if err != nil {
			return nil, err
		}
		res[i] = coreWorkLog
	}
	return res, nil
}

type WorkLogResponse struct {
	Self    string `json:"self"`
	Id      int    `json:"id"`
	Version int    `json:"version"`
	Comment string `json:"comment"`
	Issue   struct {
		Self    string `json:"self"`
		Id      string `json:"id"`
		Key     string `json:"key"`
		Display string `json:"display"`
	} `json:"issue"`
	CreatedBy struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Display     string `json:"display"`
		CloudUid    string `json:"cloudUid"`
		PassportUid int64  `json:"passportUid"`
	} `json:"createdBy"`
	UpdatedBy struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Display     string `json:"display"`
		CloudUid    string `json:"cloudUid"`
		PassportUid int64  `json:"passportUid"`
	} `json:"updatedBy"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Start     string `json:"start"`
	Duration  string `json:"duration"`
}

func (w WorkLogResponse) ToCore() (ytrackercore.WorkLog, error) {
	duration, err := unmarshalDuration(w.Duration)
	if err != nil {
		return ytrackercore.WorkLog{}, fmt.Errorf("parse duration:%w", err)
	}

	start, err := unmarshalTime(w.Start)
	if err != nil {
		return ytrackercore.WorkLog{}, fmt.Errorf("parse start:%w", err)
	}

	return ytrackercore.WorkLog{
		ID:        w.Id,
		TicketKey: w.Issue.Key,
		Comment:   w.Comment,
		Start:     start,
		Duration:  duration,
	}, nil
}

type AddWorkLogsRequest struct {
	Start    string `json:"start"`
	Duration string `json:"duration"`
	Comment  string `json:"comment,omitempty"`
}

func GetAddWorkLogsRequest(
	taskKey string,
	start time.Time,
	duration time.Duration,
	comment string,
) (map[string]string, AddWorkLogsRequest) {
	pathParams := map[string]string{
		"taskKey": taskKey,
	}
	body := AddWorkLogsRequest{
		Start:    marshalTime(start),
		Duration: marshalDuration(duration),
		Comment:  comment,
	}
	return pathParams, body
}
