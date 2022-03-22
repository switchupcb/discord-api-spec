type AutoGenerated struct {
	Definitions []Definitions `json:"definitions"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Roles struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type PremiumSince struct {
	Type []string `json:"type"`
	Description string `json:"description"`
}
type Pending struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Nick struct {
	Type []string `json:"type"`
	Description string `json:"description"`
}
type Mute struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type JoinedAt struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Deaf struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type CommunicationDisabledUntil struct {
	Type []string `json:"type"`
	Description string `json:"description"`
}
type Avatar struct {
	Type []string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Roles Roles `json:"roles"`
	PremiumSince PremiumSince `json:"premium_since"`
	Pending Pending `json:"pending"`
	Nick Nick `json:"nick"`
	Mute Mute `json:"mute"`
	JoinedAt JoinedAt `json:"joined_at"`
	Deaf Deaf `json:"deaf"`
	CommunicationDisabledUntil CommunicationDisabledUntil `json:"communication_disabled_until"`
	Avatar Avatar `json:"avatar"`
}
type PartialAPIMessageInteractionGuildMember struct {
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
type ID struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Type struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Name struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type User struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Member struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID ID `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
	User User `json:"user"`
	Member Member `json:"member"`
}
type APIMessageInteraction struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type InteractionType struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type Username struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Discriminator struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Bot struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type System struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type MfaEnabled struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Banner struct {
	Type []string `json:"type"`
	Description string `json:"description"`
}
type AccentColor struct {
	Type []string `json:"type"`
	Description string `json:"description"`
}
type Locale struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Verified struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Email struct {
	Type []string `json:"type"`
	Description string `json:"description"`
}
type Flags struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type PremiumType struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type PublicFlags struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID ID `json:"id"`
	Username Username `json:"username"`
	Discriminator Discriminator `json:"discriminator"`
	Avatar Avatar `json:"avatar"`
	Bot Bot `json:"bot"`
	System System `json:"system"`
	MfaEnabled MfaEnabled `json:"mfa_enabled"`
	Banner Banner `json:"banner"`
	AccentColor AccentColor `json:"accent_color"`
	Locale Locale `json:"locale"`
	Verified Verified `json:"verified"`
	Email Email `json:"email"`
	Flags Flags `json:"flags"`
	PremiumType PremiumType `json:"premium_type"`
	PublicFlags PublicFlags `json:"public_flags"`
}
type APIUser struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type UserFlags struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type UserPremiumType struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type Permissions struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	User User `json:"user"`
	Nick Nick `json:"nick"`
	Avatar Avatar `json:"avatar"`
	Roles Roles `json:"roles"`
	JoinedAt JoinedAt `json:"joined_at"`
	PremiumSince PremiumSince `json:"premium_since"`
	Deaf Deaf `json:"deaf"`
	Mute Mute `json:"mute"`
	Pending Pending `json:"pending"`
	CommunicationDisabledUntil CommunicationDisabledUntil `json:"communication_disabled_until"`
	Permissions Permissions `json:"permissions"`
}
type APIInteractionGuildMember struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	User User `json:"user"`
	Nick Nick `json:"nick"`
	Avatar Avatar `json:"avatar"`
	Roles Roles `json:"roles"`
	JoinedAt JoinedAt `json:"joined_at"`
	PremiumSince PremiumSince `json:"premium_since"`
	Deaf Deaf `json:"deaf"`
	Mute Mute `json:"mute"`
	Pending Pending `json:"pending"`
	CommunicationDisabledUntil CommunicationDisabledUntil `json:"communication_disabled_until"`
}
type APIGuildMember struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIInteractionResponse struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Content struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Tts struct {
	Type string `json:"type"`
	Description string `json:"description"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type APIEmbedThumbnail struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Parse struct {
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
	Deprecated string `json:"deprecated"`
}
type AllowedMentionsTypes struct {
	Type string `json:"type"`
	Enum []string `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum13302113962910629156133021139628863292951330211396035871105957764 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type APIMessageActionRowComponent struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type APIButtonComponent struct {
	AnyOf []AnyOf `json:"anyOf"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum133021139629157291961330211396288632929513302113960358711131375322 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum13302113962919729245133021139628863292951330211396035871663308133 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type MessageFlags struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIInteractionResponseDeferredMessageUpdate struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type APIApplicationCommandAutocompleteResponse struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type APIModalActionRowComponent struct {
	Ref string `json:"$ref"`
	Deprecated string `json:"deprecated"`
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
	Deprecated string `json:"deprecated"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum133021139629246292921330211396288632929513302113960358711182532533 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type TextInputStyle struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type InteractionResponseType struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type Definitions struct {
	PartialAPIMessageInteractionGuildMember PartialAPIMessageInteractionGuildMember `json:"PartialAPIMessageInteractionGuildMember,omitempty"`
	Snowflake Snowflake `json:"Snowflake"`
	APIMessageInteraction APIMessageInteraction `json:"APIMessageInteraction,omitempty"`
	InteractionType InteractionType `json:"InteractionType"`
	APIUser APIUser `json:"APIUser,omitempty"`
	UserFlags UserFlags `json:"UserFlags,omitempty"`
	UserPremiumType UserPremiumType `json:"UserPremiumType,omitempty"`
	APIInteractionGuildMember APIInteractionGuildMember `json:"APIInteractionGuildMember,omitempty"`
	APIGuildMember APIGuildMember `json:"APIGuildMember,omitempty"`
	APIInteractionResponse APIInteractionResponse `json:"APIInteractionResponse,omitempty"`
	APIInteractionResponsePong APIInteractionResponsePong `json:"APIInteractionResponsePong,omitempty"`
	APIInteractionResponseChannelMessageWithSource APIInteractionResponseChannelMessageWithSource `json:"APIInteractionResponseChannelMessageWithSource,omitempty"`
	APIInteractionResponseCallbackData APIInteractionResponseCallbackData `json:"APIInteractionResponseCallbackData,omitempty"`
	APIEmbed APIEmbed `json:"APIEmbed,omitempty"`
	EmbedType EmbedType `json:"EmbedType,omitempty"`
	APIEmbedFooter APIEmbedFooter `json:"APIEmbedFooter,omitempty"`
	APIEmbedImage APIEmbedImage `json:"APIEmbedImage,omitempty"`
	APIEmbedThumbnail APIEmbedThumbnail `json:"APIEmbedThumbnail,omitempty"`
	APIEmbedVideo APIEmbedVideo `json:"APIEmbedVideo,omitempty"`
	APIEmbedProvider APIEmbedProvider `json:"APIEmbedProvider,omitempty"`
	APIEmbedAuthor APIEmbedAuthor `json:"APIEmbedAuthor,omitempty"`
	APIEmbedField APIEmbedField `json:"APIEmbedField,omitempty"`
	APIAllowedMentions APIAllowedMentions `json:"APIAllowedMentions,omitempty"`
	AllowedMentionsTypes AllowedMentionsTypes `json:"AllowedMentionsTypes,omitempty"`
	APIActionRowComponentAPIMessageActionRowComponent APIActionRowComponentAPIMessageActionRowComponent `json:"APIActionRowComponent<APIMessageActionRowComponent>,omitempty"`
	APIBaseComponentEnum13302113962910629156133021139628863292951330211396035871105957764 APIBaseComponentEnum13302113962910629156133021139628863292951330211396035871105957764 `json:"APIBaseComponent<enum-1330211396-29106-29156-1330211396-28863-29295-1330211396-0-35871105957764>,omitempty"`
	APIMessageActionRowComponent APIMessageActionRowComponent `json:"APIMessageActionRowComponent,omitempty"`
	APIButtonComponent APIButtonComponent `json:"APIButtonComponent,omitempty"`
	APIButtonComponentWithCustomID APIButtonComponentWithCustomID `json:"APIButtonComponentWithCustomId,omitempty"`
	APIBaseComponentEnum133021139629157291961330211396288632929513302113960358711131375322 APIBaseComponentEnum133021139629157291961330211396288632929513302113960358711131375322 `json:"APIBaseComponent<enum-1330211396-29157-29196-1330211396-28863-29295-1330211396-0-358711131375322>,omitempty"`
	APIMessageComponentEmoji APIMessageComponentEmoji `json:"APIMessageComponentEmoji,omitempty"`
	APIButtonComponentWithURL APIButtonComponentWithURL `json:"APIButtonComponentWithURL,omitempty"`
	APISelectMenuComponent APISelectMenuComponent `json:"APISelectMenuComponent,omitempty"`
	APIBaseComponentEnum13302113962919729245133021139628863292951330211396035871663308133 APIBaseComponentEnum13302113962919729245133021139628863292951330211396035871663308133 `json:"APIBaseComponent<enum-1330211396-29197-29245-1330211396-28863-29295-1330211396-0-35871663308133>,omitempty"`
	APISelectMenuOption APISelectMenuOption `json:"APISelectMenuOption,omitempty"`
	MessageFlags MessageFlags `json:"MessageFlags,omitempty"`
	APIInteractionResponseDeferredChannelMessageWithSource APIInteractionResponseDeferredChannelMessageWithSource `json:"APIInteractionResponseDeferredChannelMessageWithSource,omitempty"`
	APIInteractionResponseDeferredMessageUpdate APIInteractionResponseDeferredMessageUpdate `json:"APIInteractionResponseDeferredMessageUpdate,omitempty"`
	APIInteractionResponseUpdateMessage APIInteractionResponseUpdateMessage `json:"APIInteractionResponseUpdateMessage,omitempty"`
	APIApplicationCommandAutocompleteResponse APIApplicationCommandAutocompleteResponse `json:"APIApplicationCommandAutocompleteResponse,omitempty"`
	APICommandAutocompleteInteractionResponseCallbackData APICommandAutocompleteInteractionResponseCallbackData `json:"APICommandAutocompleteInteractionResponseCallbackData,omitempty"`
	APIApplicationCommandOptionChoice APIApplicationCommandOptionChoice `json:"APIApplicationCommandOptionChoice,omitempty"`
	APIModalInteractionResponse APIModalInteractionResponse `json:"APIModalInteractionResponse,omitempty"`
	APIModalInteractionResponseCallbackData APIModalInteractionResponseCallbackData `json:"APIModalInteractionResponseCallbackData,omitempty"`
	APIActionRowComponentAPIModalActionRowComponent APIActionRowComponentAPIModalActionRowComponent `json:"APIActionRowComponent<APIModalActionRowComponent>,omitempty"`
	APIModalActionRowComponent APIModalActionRowComponent `json:"APIModalActionRowComponent,omitempty"`
	APITextInputComponent APITextInputComponent `json:"APITextInputComponent,omitempty"`
	APIBaseComponentEnum133021139629246292921330211396288632929513302113960358711182532533 APIBaseComponentEnum133021139629246292921330211396288632929513302113960358711182532533 `json:"APIBaseComponent<enum-1330211396-29246-29292-1330211396-28863-29295-1330211396-0-358711182532533>,omitempty"`
	TextInputStyle TextInputStyle `json:"TextInputStyle,omitempty"`
	InteractionResponseType InteractionResponseType `json:"InteractionResponseType,omitempty"`
}