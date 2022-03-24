package _chatInput
type APIApplicationCommandNumberOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface10302492622547361030249262013611295383896APIApplicationCommandOptionChoiceNumber struct {
	AnyOf []AnyOf `json:"anyOf"`
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum20937207224034112093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandInteractionDataNumberOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Required []string `json:"required"`
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Value Value `json:"value"`
}
type APIInteractionDataOptionBaseEnum20937207224034112093720722042720937207220818Number struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}

