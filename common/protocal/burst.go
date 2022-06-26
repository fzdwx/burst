package protocal

type (
	Burst struct {
		Type    string         `json:"type"`
		Headers map[string]any `json:"headers"`
		Data    []byte         `json:"data"`
	}
)

const (
	AddProxy = "addproxy"
)
