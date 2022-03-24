package _interactions
type InteractionType struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIInteractionResponse struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
}
type Type struct {
	Type string `json:"type"`
	Const int `json:"const"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIInteractionResponsePong struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Data struct {
	Ref string `json:"$ref"`
}
type Properties struct {
	Type Type `json:"type"`
	Data Data `json:"data"`
}
type APIInteractionResponseChannelMessageWithSource struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Flags struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Content struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Tts struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Embeds struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type AllowedMentions struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Components struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type Filename struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type ID struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Description struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Filename Filename `json:"filename"`
	ID ID `json:"id"`
	Description Description `json:"description"`
}
type Items struct {
	Type string `json:"type"`
	AdditionalProperties bool `json:"additionalProperties"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
}
type Attachments struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type Properties struct {
	Flags Flags `json:"flags"`
	Content Content `json:"content"`
	Tts Tts `json:"tts"`
	Embeds Embeds `json:"embeds"`
	AllowedMentions AllowedMentions `json:"allowed_mentions"`
	Components Components `json:"components"`
	Attachments Attachments `json:"attachments"`
}
type APIInteractionResponseCallbackData struct {
	Type string `json:"type"`
	AdditionalProperties bool `json:"additionalProperties"`
	Properties Properties `json:"properties"`
	Description string `json:"description"`
}
type Title struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Type struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type URL struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Timestamp struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Color struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Footer struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Image struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Thumbnail struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Video struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Provider struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Author struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Fields struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type Properties struct {
	Title Title `json:"title"`
	Type Type `json:"type"`
	Description Description `json:"description"`
	URL URL `json:"url"`
	Timestamp Timestamp `json:"timestamp"`
	Color Color `json:"color"`
	Footer Footer `json:"footer"`
	Image Image `json:"image"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Video Video `json:"video"`
	Provider Provider `json:"provider"`
	Author Author `json:"author"`
	Fields Fields `json:"fields"`
}
type APIEmbed struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type EmbedType struct {
	Type string `json:"type"`
	Enum []string `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type Text struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type IconURL struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type ProxyIconURL struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Text Text `json:"text"`
	IconURL IconURL `json:"icon_url"`
	ProxyIconURL ProxyIconURL `json:"proxy_icon_url"`
}
type APIEmbedFooter struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type ProxyURL struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Height struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Width struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	URL URL `json:"url"`
	ProxyURL ProxyURL `json:"proxy_url"`
	Height Height `json:"height"`
	Width Width `json:"width"`
}
type APIEmbedImage struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type APIEmbedThumbnail struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Properties struct {
	URL URL `json:"url"`
	Height Height `json:"height"`
	Width Width `json:"width"`
}
type APIEmbedVideo struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Name struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Name Name `json:"name"`
	URL URL `json:"url"`
}
type APIEmbedProvider struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Properties struct {
	Name Name `json:"name"`
	URL URL `json:"url"`
	IconURL IconURL `json:"icon_url"`
	ProxyIconURL ProxyIconURL `json:"proxy_icon_url"`
}
type APIEmbedAuthor struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Value struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Inline struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Name Name `json:"name"`
	Value Value `json:"value"`
	Inline Inline `json:"inline"`
}
type APIEmbedField struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Parse struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type Roles struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type Users struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type RepliedUser struct {
	Type string `json:"type"`
	Description string `json:"description"`
	Default bool `json:"default"`
}
type Properties struct {
	Parse Parse `json:"parse"`
	Roles Roles `json:"roles"`
	Users Users `json:"users"`
	RepliedUser RepliedUser `json:"replied_user"`
}
type APIAllowedMentions struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type AllowedMentionsTypes struct {
	Type string `json:"type"`
	Enum []string `json:"enum"`
	Description string `json:"description"`
}
type Snowflake struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Type struct {
	Type string `json:"type"`
	Const int `json:"const"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
	Components Components `json:"components"`
}
type APIActionRowComponentAPIMessageActionRowComponent struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum13545260102870528755135452601028581288941354526010033636105957764 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type APIMessageActionRowComponent struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
}
type APIButtonComponent struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type Label struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type AnyOf struct {
	Type string `json:"type"`
	Const int `json:"const"`
}
type Style struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
}
type Emoji struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Disabled struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type CustomID struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
	Label Label `json:"label"`
	Style Style `json:"style"`
	Emoji Emoji `json:"emoji"`
	Disabled Disabled `json:"disabled"`
	CustomID CustomID `json:"custom_id"`
}
type APIButtonComponentWithCustomID struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Properties struct {
	Type Type `json:"type"`
	Label Label `json:"label"`
	Style Style `json:"style"`
	Emoji Emoji `json:"emoji"`
	Disabled Disabled `json:"disabled"`
}
type APIButtonComponentBaseEnum13545260103053930552135452601030405305931354526010033636Enum13545260103055330564135452601030405305931354526010033636Enum13545260103056530574135452601030405305931354526010033636Enum13545260103057530583135452601030405305931354526010033636 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum135452601028756287951354526010285812889413545260100336361131375322 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Animated struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID ID `json:"id"`
	Name Name `json:"name"`
	Animated Animated `json:"animated"`
}
type APIMessageComponentEmoji struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Style struct {
	Type string `json:"type"`
	Const int `json:"const"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
	Label Label `json:"label"`
	Style Style `json:"style"`
	Emoji Emoji `json:"emoji"`
	Disabled Disabled `json:"disabled"`
	URL URL `json:"url"`
}
type APIButtonComponentWithURL struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Properties struct {
	Type Type `json:"type"`
	Label Label `json:"label"`
	Style Style `json:"style"`
	Emoji Emoji `json:"emoji"`
	Disabled Disabled `json:"disabled"`
}
type APIButtonComponentBaseEnum13545260103058430590135452601030405305931354526010033636 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Options struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type Placeholder struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type MinValues struct {
	Type string `json:"type"`
	Description string `json:"description"`
	Default int `json:"default"`
}
type MaxValues struct {
	Type string `json:"type"`
	Description string `json:"description"`
	Default int `json:"default"`
}
type Disabled struct {
	Type string `json:"type"`
	Description string `json:"description"`
	Default bool `json:"default"`
}
type Properties struct {
	Type Type `json:"type"`
	CustomID CustomID `json:"custom_id"`
	Options Options `json:"options"`
	Placeholder Placeholder `json:"placeholder"`
	MinValues MinValues `json:"min_values"`
	MaxValues MaxValues `json:"max_values"`
	Disabled Disabled `json:"disabled"`
}
type APISelectMenuComponent struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum13545260102879628844135452601028581288941354526010033636663308133 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Default struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Label Label `json:"label"`
	Value Value `json:"value"`
	Description Description `json:"description"`
	Emoji Emoji `json:"emoji"`
	Default Default `json:"default"`
}
type APISelectMenuOption struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type MessageFlags struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
}
type Properties struct {
	Flags Flags `json:"flags"`
}
type Data struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Properties struct {
	Type Type `json:"type"`
	Data Data `json:"data"`
}
type APIInteractionResponseDeferredChannelMessageWithSource struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIInteractionResponseDeferredMessageUpdate struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Data struct {
	Ref string `json:"$ref"`
}
type Properties struct {
	Type Type `json:"type"`
	Data Data `json:"data"`
}
type APIInteractionResponseUpdateMessage struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type APIApplicationCommandAutocompleteResponse struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Choices struct {
	Type string `json:"type"`
	Items Items `json:"items"`
}
type Properties struct {
	Choices Choices `json:"choices"`
}
type APICommandAutocompleteInteractionResponseCallbackData struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Name struct {
	Type string `json:"type"`
}
type Value struct {
	Type []string `json:"type"`
}
type Properties struct {
	Name Name `json:"name"`
	Value Value `json:"value"`
}
type APIApplicationCommandOptionChoice struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
	Data Data `json:"data"`
}
type APIModalInteractionResponse struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
}
type Properties struct {
	CustomID CustomID `json:"custom_id"`
	Title Title `json:"title"`
	Components Components `json:"components"`
}
type APIModalInteractionResponseCallbackData struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
	Components Components `json:"components"`
}
type APIActionRowComponentAPIModalActionRowComponent struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type APIModalActionRowComponent struct {
	Ref string `json:"$ref"`
}
type Style struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Value struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type MinLength struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type MaxLength struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Required struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
	Style Style `json:"style"`
	CustomID CustomID `json:"custom_id"`
	Label Label `json:"label"`
	Placeholder Placeholder `json:"placeholder"`
	Value Value `json:"value"`
	MinLength MinLength `json:"min_length"`
	MaxLength MaxLength `json:"max_length"`
	Required Required `json:"required"`
}
type APITextInputComponent struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum135452601028845288911354526010285812889413545260100336361182532533 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type TextInputStyle struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
}
type InteractionResponseType struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
}

