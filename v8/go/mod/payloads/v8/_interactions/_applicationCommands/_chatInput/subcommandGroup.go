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
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum20937207223183352093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum20937207223013172093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIApplicationCommandBasicOption struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandStringOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperAPIApplicationCommandOptionBaseEnum20937207223363442093720722042720937207220818APIApplicationCommandOptionChoiceString struct {
	AnyOf []AnyOf `json:"anyOf"`
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum20937207223363442093720722042720937207220818 struct {
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
type APIApplicationCommandOptionChoiceString struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandIntegerOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface352328325254738352328325013671295383896APIApplicationCommandOptionChoiceNumber struct {
	AnyOf []AnyOf `json:"anyOf"`
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum20937207223453542093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandBooleanOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
	Name Name `json:"name"`
	Description Description `json:"description"`
	Required Required `json:"required"`
}
type APIApplicationCommandOptionBaseEnum20937207223553642093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandUserOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandOptionBaseEnum20937207223653712093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandRoleOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandOptionBaseEnum20937207223823882093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandMentionableOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandOptionBaseEnum20937207223894022093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandNumberOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandAttachmentOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandOptionBaseEnum20937207224124242093720722042720937207220818 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandInteractionDataSubcommandOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIApplicationCommandInteractionDataBasicOption struct {
	AnyOf []AnyOf `json:"anyOf"`
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
type APIApplicationCommandInteractionDataStringOption struct {
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
type APIInteractionDataOptionBaseEnum20937207223363442093720722042720937207220818String struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Name Name `json:"name"`
	Type Type `json:"type"`
	Value Value `json:"value"`
}
type APIInteractionDataOptionBaseEnum20937207223453542093720722042720937207220818Number struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandInteractionDataBooleanOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIInteractionDataOptionBaseEnum20937207223553642093720722042720937207220818Boolean struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandInteractionDataUserOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIInteractionDataOptionBaseEnum20937207223653712093720722042720937207220818Snowflake struct {
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
type APIApplicationCommandInteractionDataChannelOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIInteractionDataOptionBaseEnum20937207223723812093720722042720937207220818Snowflake struct {
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
type APIInteractionDataOptionBaseEnum20937207223823882093720722042720937207220818Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandInteractionDataMentionableOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIInteractionDataOptionBaseEnum20937207223894022093720722042720937207220818Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
type APIApplicationCommandInteractionDataAttachmentOption struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIInteractionDataOptionBaseEnum20937207224124242093720722042720937207220818Snowflake struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
}

