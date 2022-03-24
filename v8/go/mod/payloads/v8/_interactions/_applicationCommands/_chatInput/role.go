package _chatInput
type APIApplicationCommandRoleOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandOptionBaseEnum20937207223823882093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandInteractionDataRoleOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type Value struct {
	Ref string `json:"$ref"`
}
type Properties struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Value Value `json:"value"`
}
type APIInteractionDataOptionBaseEnum20937207223823882093720722042720937207220818Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type Snowflake struct {
	Type string `json:"type"`
	Description string `json:"description"`
}

