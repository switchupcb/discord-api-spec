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
type Definitions struct {
	APIApplicationCommandSubcommandOption APIApplicationCommandSubcommandOption `json:"APIApplicationCommandSubcommandOption"`
	APIApplicationCommandOptionBaseEnum20937207223013172093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223013172093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-301-317-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandBasicOption APIApplicationCommandBasicOption `json:"APIApplicationCommandBasicOption"`
	APIApplicationCommandStringOption APIApplicationCommandStringOption `json:"APIApplicationCommandStringOption"`
	APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperAPIApplicationCommandOptionBaseEnum20937207223363442093720722042720937207220818APIApplicationCommandOptionChoiceString APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperAPIApplicationCommandOptionBaseEnum20937207223363442093720722042720937207220818APIApplicationCommandOptionChoiceString `json:"APIApplicationCommandOptionWithAutocompleteOrChoicesWrapper<APIApplicationCommandOptionBase<enum-2093720722-336-344-2093720722-0-427-2093720722-0-818>,APIApplicationCommandOptionChoice<string>>"`
	APIApplicationCommandOptionBaseEnum20937207223363442093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223363442093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-336-344-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandOptionChoiceString APIApplicationCommandOptionChoiceString `json:"APIApplicationCommandOptionChoice<string>"`
	APIApplicationCommandIntegerOption APIApplicationCommandIntegerOption `json:"APIApplicationCommandIntegerOption"`
	APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface352328325254738352328325013671295383896APIApplicationCommandOptionChoiceNumber APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface352328325254738352328325013671295383896APIApplicationCommandOptionChoiceNumber `json:"APIApplicationCommandOptionWithAutocompleteOrChoicesWrapper<interface-352328325-254-738-352328325-0-13671295383896,APIApplicationCommandOptionChoice<number>>"`
	APIApplicationCommandOptionBaseEnum20937207223453542093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223453542093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-345-354-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandOptionChoiceNumber APIApplicationCommandOptionChoiceNumber `json:"APIApplicationCommandOptionChoice<number>"`
	APIApplicationCommandBooleanOption APIApplicationCommandBooleanOption `json:"APIApplicationCommandBooleanOption"`
	APIApplicationCommandOptionBaseEnum20937207223553642093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223553642093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-355-364-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandUserOption APIApplicationCommandUserOption `json:"APIApplicationCommandUserOption"`
	APIApplicationCommandOptionBaseEnum20937207223653712093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223653712093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-365-371-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandChannelOption APIApplicationCommandChannelOption `json:"APIApplicationCommandChannelOption"`
	APIApplicationCommandOptionBaseEnum20937207223723812093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223723812093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-372-381-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandRoleOption APIApplicationCommandRoleOption `json:"APIApplicationCommandRoleOption"`
	APIApplicationCommandOptionBaseEnum20937207223823882093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223823882093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-382-388-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandMentionableOption APIApplicationCommandMentionableOption `json:"APIApplicationCommandMentionableOption"`
	APIApplicationCommandOptionBaseEnum20937207223894022093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207223894022093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-389-402-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandNumberOption APIApplicationCommandNumberOption `json:"APIApplicationCommandNumberOption"`
	APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface10302492622547361030249262013611295383896APIApplicationCommandOptionChoiceNumber APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface10302492622547361030249262013611295383896APIApplicationCommandOptionChoiceNumber `json:"APIApplicationCommandOptionWithAutocompleteOrChoicesWrapper<interface-1030249262-254-736-1030249262-0-13611295383896,APIApplicationCommandOptionChoice<number>>"`
	APIApplicationCommandOptionBaseEnum20937207224034112093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207224034112093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-403-411-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandAttachmentOption APIApplicationCommandAttachmentOption `json:"APIApplicationCommandAttachmentOption"`
	APIApplicationCommandOptionBaseEnum20937207224124242093720722042720937207220818 APIApplicationCommandOptionBaseEnum20937207224124242093720722042720937207220818 `json:"APIApplicationCommandOptionBase<enum-2093720722-412-424-2093720722-0-427-2093720722-0-818>"`
	APIApplicationCommandInteractionDataSubcommandOption APIApplicationCommandInteractionDataSubcommandOption `json:"APIApplicationCommandInteractionDataSubcommandOption"`
	APIApplicationCommandInteractionDataBasicOption APIApplicationCommandInteractionDataBasicOption `json:"APIApplicationCommandInteractionDataBasicOption"`
	APIApplicationCommandInteractionDataStringOption APIApplicationCommandInteractionDataStringOption `json:"APIApplicationCommandInteractionDataStringOption"`
	APIInteractionDataOptionBaseEnum20937207223363442093720722042720937207220818String APIInteractionDataOptionBaseEnum20937207223363442093720722042720937207220818String `json:"APIInteractionDataOptionBase<enum-2093720722-336-344-2093720722-0-427-2093720722-0-818,string>"`
	APIApplicationCommandInteractionDataIntegerOption APIApplicationCommandInteractionDataIntegerOption `json:"APIApplicationCommandInteractionDataIntegerOption"`
	APIInteractionDataOptionBaseEnum20937207223453542093720722042720937207220818Number APIInteractionDataOptionBaseEnum20937207223453542093720722042720937207220818Number `json:"APIInteractionDataOptionBase<enum-2093720722-345-354-2093720722-0-427-2093720722-0-818,number>"`
	APIApplicationCommandInteractionDataBooleanOption APIApplicationCommandInteractionDataBooleanOption `json:"APIApplicationCommandInteractionDataBooleanOption"`
	APIInteractionDataOptionBaseEnum20937207223553642093720722042720937207220818Boolean APIInteractionDataOptionBaseEnum20937207223553642093720722042720937207220818Boolean `json:"APIInteractionDataOptionBase<enum-2093720722-355-364-2093720722-0-427-2093720722-0-818,boolean>"`
	APIApplicationCommandInteractionDataUserOption APIApplicationCommandInteractionDataUserOption `json:"APIApplicationCommandInteractionDataUserOption"`
	APIInteractionDataOptionBaseEnum20937207223653712093720722042720937207220818Snowflake APIInteractionDataOptionBaseEnum20937207223653712093720722042720937207220818Snowflake `json:"APIInteractionDataOptionBase<enum-2093720722-365-371-2093720722-0-427-2093720722-0-818,Snowflake>"`
	Snowflake Snowflake `json:"Snowflake"`
	APIApplicationCommandInteractionDataChannelOption APIApplicationCommandInteractionDataChannelOption `json:"APIApplicationCommandInteractionDataChannelOption"`
	APIInteractionDataOptionBaseEnum20937207223723812093720722042720937207220818Snowflake APIInteractionDataOptionBaseEnum20937207223723812093720722042720937207220818Snowflake `json:"APIInteractionDataOptionBase<enum-2093720722-372-381-2093720722-0-427-2093720722-0-818,Snowflake>"`
	APIApplicationCommandInteractionDataRoleOption APIApplicationCommandInteractionDataRoleOption `json:"APIApplicationCommandInteractionDataRoleOption"`
	APIInteractionDataOptionBaseEnum20937207223823882093720722042720937207220818Snowflake APIInteractionDataOptionBaseEnum20937207223823882093720722042720937207220818Snowflake `json:"APIInteractionDataOptionBase<enum-2093720722-382-388-2093720722-0-427-2093720722-0-818,Snowflake>"`
	APIApplicationCommandInteractionDataMentionableOption APIApplicationCommandInteractionDataMentionableOption `json:"APIApplicationCommandInteractionDataMentionableOption"`
	APIInteractionDataOptionBaseEnum20937207223894022093720722042720937207220818Snowflake APIInteractionDataOptionBaseEnum20937207223894022093720722042720937207220818Snowflake `json:"APIInteractionDataOptionBase<enum-2093720722-389-402-2093720722-0-427-2093720722-0-818,Snowflake>"`
	APIApplicationCommandInteractionDataNumberOption APIApplicationCommandInteractionDataNumberOption `json:"APIApplicationCommandInteractionDataNumberOption"`
	APIInteractionDataOptionBaseEnum20937207224034112093720722042720937207220818Number APIInteractionDataOptionBaseEnum20937207224034112093720722042720937207220818Number `json:"APIInteractionDataOptionBase<enum-2093720722-403-411-2093720722-0-427-2093720722-0-818,number>"`
	APIApplicationCommandInteractionDataAttachmentOption APIApplicationCommandInteractionDataAttachmentOption `json:"APIApplicationCommandInteractionDataAttachmentOption"`
	APIInteractionDataOptionBaseEnum20937207224124242093720722042720937207220818Snowflake APIInteractionDataOptionBaseEnum20937207224124242093720722042720937207220818Snowflake `json:"APIInteractionDataOptionBase<enum-2093720722-412-424-2093720722-0-427-2093720722-0-818,Snowflake>"`
}