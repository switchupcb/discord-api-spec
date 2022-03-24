package _chatInput
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
	Ref string `json:"$ref"`
}
type Options struct {
	Type string `json:"type"`
	Items Items `json:"items"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
	Options Options `json:"options"`
}
type APIApplicationCommandSubcommandGroupOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Required []string `json:"required"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum15707879491992161570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
	Options Options `json:"options"`
}
type APIApplicationCommandSubcommandOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Required []string `json:"required"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum15707879491821981570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIApplicationCommandBasicOption struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
}
type APIApplicationCommandStringOption struct {
	Ref string `json:"$ref"`
}
type Autocomplete struct {
	Type string `json:"type"`
	Const bool `json:"const"`
}
type Properties struct {
	Autocomplete Autocomplete `json:"autocomplete"`
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
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
type APIApplicationCommandIntegerOption struct {
	Ref string `json:"$ref"`
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
type APIApplicationCommandBooleanOption struct {
	Ref string `json:"$ref"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum15707879492362451570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandUserOption struct {
	Ref string `json:"$ref"`
}
type APIApplicationCommandOptionBaseEnum15707879492462521570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Items struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
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
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum15707879492532621570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandRoleOption struct {
	Ref string `json:"$ref"`
}
type APIApplicationCommandOptionBaseEnum15707879492632691570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandMentionableOption struct {
	Ref string `json:"$ref"`
}
type APIApplicationCommandOptionBaseEnum15707879492702831570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandNumberOption struct {
	Ref string `json:"$ref"`
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
type AnyOf struct {
	Type string `json:"type"`
	AdditionalProperties bool `json:"additionalProperties"`
	Properties Properties `json:"properties,omitempty"`
	Required []string `json:"required"`
	Properties0 Properties0 `json:"properties,omitempty"`
}
type APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface166070788725460916607078870980APIApplicationCommandOptionChoiceNumber struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum15707879492842921570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandAttachmentOption struct {
	Ref string `json:"$ref"`
}
type APIApplicationCommandOptionBaseEnum15707879492933051570787949030815707879490580 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Properties struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Options Options `json:"options"`
}
type APIApplicationCommandInteractionDataSubcommandGroupOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandInteractionDataSubcommandOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIApplicationCommandInteractionDataBasicOption struct {
	AnyOf []AnyOf `json:"anyOf"`
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
type APIApplicationCommandInteractionDataBooleanOption struct {
	Ref string `json:"$ref"`
}
type APIInteractionDataOptionBaseEnum15707879492362451570787949030815707879490580Boolean struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandInteractionDataUserOption struct {
	Ref string `json:"$ref"`
}
type APIInteractionDataOptionBaseEnum15707879492462521570787949030815707879490580Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Snowflake struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type APIApplicationCommandInteractionDataChannelOption struct {
	Ref string `json:"$ref"`
}
type APIInteractionDataOptionBaseEnum15707879492532621570787949030815707879490580Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandInteractionDataRoleOption struct {
	Ref string `json:"$ref"`
}
type APIInteractionDataOptionBaseEnum15707879492632691570787949030815707879490580Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandInteractionDataMentionableOption struct {
	Ref string `json:"$ref"`
}
type APIInteractionDataOptionBaseEnum15707879492702831570787949030815707879490580Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
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
}
type Properties struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Value Value `json:"value"`
}
type APIInteractionDataOptionBaseEnum15707879492842921570787949030815707879490580Number struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandInteractionDataAttachmentOption struct {
	Ref string `json:"$ref"`
}
type APIInteractionDataOptionBaseEnum15707879492933051570787949030815707879490580Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}

