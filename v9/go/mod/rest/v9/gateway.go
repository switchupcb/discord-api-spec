package v9

type RESTGetAPIGatewayResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type URL struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	URL URL `json:"url"`
}
type APIGatewayInfo struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTGetAPIGatewayBotResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Shards struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SessionStartLimit struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	URL               URL               `json:"url"`
	Shards            Shards            `json:"shards"`
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}
type APIGatewayBotInfo struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Total struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Remaining struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ResetAfter struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type MaxConcurrency struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Total          Total          `json:"total"`
	Remaining      Remaining      `json:"remaining"`
	ResetAfter     ResetAfter     `json:"reset_after"`
	MaxConcurrency MaxConcurrency `json:"max_concurrency"`
}
type APIGatewaySessionStartLimit struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
