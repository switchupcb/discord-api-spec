package _chatInput
type APIApplicationCommandMentionableOption struct {
	Ref string `json:"$ref"`
}
type Type struct {
	Type string `json:"type"`
	Const int `json:"const"`
}
type Name struct {
	Type string `json:"type"`
}
type Description struct {
	Type string `json:"type"`
}
type Required struct {
	Type string `json:"type"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum15707879492702831570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandInteractionDataMentionableOption struct {
	Ref string `json:"$ref"`
}
type Value struct {
	Ref string `json:"$ref"`
}
type Properties struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Value Value `json:"value"`
}
type APIInteractionDataOptionBaseEnum15707879492702831570787949030815707879490580Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Snowflake struct {
	Type string `json:"type"`
	Description string `json:"description"`
}

