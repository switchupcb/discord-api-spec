package v8

type Items struct {
	Ref string `json:"$ref"`
}
type GetAPIVoiceRegionsResult struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type ID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Optimal struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Deprecated struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Custom struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID         ID         `json:"id"`
	Name       Name       `json:"name"`
	Optimal    Optimal    `json:"optimal"`
	Deprecated Deprecated `json:"deprecated"`
	Custom     Custom     `json:"custom"`
}
type APIVoiceRegion struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
