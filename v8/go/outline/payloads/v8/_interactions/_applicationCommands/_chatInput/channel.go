type AutoGenerated struct {
	Definitions Definitions `json:"definitions"`
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
type Items struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type ChannelTypes struct {
	Type string `json:"type"`
	Items Items `json:"items"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
	ChannelTypes ChannelTypes `json:"channel_types"`
}
type APIApplicationCommandChannelOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Required []string `json:"required"`
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum20937207223723812093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandInteractionDataChannelOption struct {
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
type APIInteractionDataOptionBaseEnum20937207223723812093720722042720937207220818Snowflake struct {
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
type Definitions struct {
	APIApplicationCommandChannelOption APIApplicationCommandChannelOption `json:"APIApplicationCommandChannelOption"`
	APIApplicationCommandOptionBaseEnum20937207223723812093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223723812093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-372-381-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandInteractionDataChannelOption APIApplicationCommandInteractionDataChannelOption `json:"APIApplicationCommandInteractionDataChannelOption"`
	APIInteractionDataOptionBaseEnum20937207223723812093720722042720937207220818Snowflake APIInteractionDataOptionBaseEnum20937207223723812093720722042720937207220818Snowflake `json:"APIInteractionDataOptionBase<enum-2093720722-372-381-2093720722-0-427-2093720722-0-818,Snowflake>"`
	Snowflake Snowflake `json:"Snowflake"`
}