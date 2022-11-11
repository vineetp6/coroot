package model

import "encoding/json"

const (
	UNKNOWN Status = iota
	OK
	INFO
	WARNING
	CRITICAL
)

type Status int

func (s Status) String() string {
	switch s {
	case OK:
		return "ok"
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case CRITICAL:
		return "critical"
	default:
		return "unknown"
	}
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

type Indicator struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
}

func CalcIndicators(app *Application) []Indicator {
	if app == nil {
		return nil
	}
	var res []Indicator
	for _, r := range app.Reports {
		if r.Status == UNKNOWN {
			continue
		}
		res = append(res, Indicator{
			Status:  r.Status,
			Message: string(r.Name),
		})
	}
	return res
}
