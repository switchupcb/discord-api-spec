package _chatInput
type APIApplicationCommandIntegerOption struct {
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
type MinValue struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type MaxValue struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Autocomplete Autocomplete `json:"autocomplete"`
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
	MinValue MinValue `json:"min_value"`
	MaxValue MaxValue `json:"max_value"`
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
	MinValue MinValue `json:"min_value"`
	MaxValue MaxValue `json:"max_value"`
}
type AnyOf struct {
	Type string `json:"type"`
	AdditionalProperties bool `json:"additionalProperties"`
	Properties Properties `json:"properties,omitempty"`
	Required []string `json:"required"`
	Properties0 Properties0 `json:"properties,omitempty"`
}
type APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface157829078025461115782907800986APIApplicationCommandOptionChoiceNumber struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum15707879492262351570787949030815707879490580 struct {
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
type APIApplicationCommandOptionChoiceNumber struct {
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
type APIApplicationCommandInteractionDataIntegerOption struct {
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
type APIInteractionDataOptionBaseEnum15707879492262351570787949030815707879490580Number struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}

