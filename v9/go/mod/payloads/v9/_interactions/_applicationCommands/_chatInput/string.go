package _chatInput
type APIApplicationCommandStringOption struct {
	Ref string `json:"$ref"`
}
type Autocomplete struct {
	Type string `json:"type"`
	Const bool `json:"const"`
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
	Autocomplete Autocomplete `json:"autocomplete"`
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Choices struct {
	Type string `json:"type"`
	Items Items `json:"items"`
}
type Properties0 struct {
	Autocomplete Autocomplete `json:"autocomplete"`
	Choices Choices `json:"choices"`
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type AnyOf struct {
	Type string `json:"type"`
	AdditionalProperties bool `json:"additionalProperties"`
	Properties Properties `json:"properties,omitempty"`
	Required []string `json:"required"`
	Properties0 Properties0 `json:"properties,omitempty"`
}
type APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperAPIApplicationCommandOptionBaseEnum15707879492172251570787949030815707879490580APIApplicationCommandOptionChoiceString struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum15707879492172251570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Value struct {
	Type string `json:"type"`
}
type Properties struct {
	Name Name `json:"name"`
	Value Value `json:"value"`
}
type APIApplicationCommandOptionChoiceString struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Focused struct {
	Type string `json:"type"`
}
type Properties struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Value Value `json:"value"`
	Focused Focused `json:"focused"`
}
type APIApplicationCommandInteractionDataStringOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Required []string `json:"required"`
}
type Properties struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Value Value `json:"value"`
}
type APIInteractionDataOptionBaseEnum15707879492172251570787949030815707879490580String struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}

