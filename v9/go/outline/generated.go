type AutoGenerated struct {
	Definitions []Definitions `json:"definitions"`
}
type Snowflake struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Definitions struct {
	Snowflake Snowflake `json:"Snowflake,omitempty"`
}