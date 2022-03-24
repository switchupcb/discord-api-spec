package gateway

type V struct {
	Type string `json:"type"`
}
type Encoding struct {
	Type string   `json:"type"`
	Enum []string `json:"enum"`
}
type Compress struct {
	Type  string `json:"type"`
	Const string `json:"const"`
}
type Properties struct {
	Compress Compress `json:"compress"`
	V        V        `json:"v"`
	Encoding Encoding `json:"encoding"`
}
type GatewayURLQuery struct {
	Properties           Properties `json:"properties"`
	Type                 string     `json:"type"`
	Description          string     `json:"description"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
