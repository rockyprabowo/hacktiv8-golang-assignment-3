package telemetry_view_models

type StateMap map[string]string

type TelemetryConfigVM struct {
	Interval    string   `json:"interval"`
	IntervalMs  int64    `json:"interval_ms"`
	WaterStates StateMap `json:"water_states"`
	WindStates  StateMap `json:"wind_states"`
}
