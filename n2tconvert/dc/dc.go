package dc

// TagPosition 表示一个标记范围
type TagPosition struct {
	Start  int
	Length int
}

// TagMarker 表示一个标记
type TagMarker struct {
	Position TagPosition
	Type     string
	Value    interface{} `json:"Value,omitempty"`
}

// Result 表示解析结果
type Result struct {
	Version     int
	DisplayText string
	Markers     []TagMarker
}
