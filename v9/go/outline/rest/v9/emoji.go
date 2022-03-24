type AutoGenerated struct {
	Definitions Definitions `json:"definitions"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type RESTGetAPIGuildEmojisResult struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type AnyOf struct {
	Ref string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
}
type ID struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
}
type Name struct {
	Type []string `json:"type"`
	Description string `json:"description"`
}
type Animated struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Roles struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type User struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type RequireColons struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Managed struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Available struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID ID `json:"id"`
	Name Name `json:"name"`
	Animated Animated `json:"animated"`
	Roles Roles `json:"roles"`
	User User `json:"user"`
	RequireColons RequireColons `json:"require_colons"`
	Managed Managed `json:"managed"`
	Available Available `json:"available"`
}
type APIEmoji struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Required []string `json:"required"`
	Description string `json:"description"`
}
type Properties struct {
	ID ID `json:"id"`
	Name Name `json:"name"`
	Animated Animated `json:"animated"`
}
type APIPartialEmoji struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type Snowflake struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type ID struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Username struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Discriminator struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Avatar struct {
	Type []string `json:"type"`
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
}
type UserFlags struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
}
type UserPremiumType struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
}
type RESTGetAPIGuildEmojiResult struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type RESTPostAPIGuildEmojiJSONBody struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Image struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Name Name `json:"name"`
	Image Image `json:"image"`
	Roles Roles `json:"roles"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure193329939866391519332993986099161933299398470917193329939801648 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type RESTPostAPIGuildEmojiResult struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type RESTPatchAPIGuildEmojiJSONBody struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type AnyOf struct {
	Type string `json:"type"`
	Items Items `json:"items,omitempty"`
}
type Roles struct {
	AnyOf []AnyOf `json:"anyOf"`
	Description string `json:"description"`
}
type Properties struct {
	Name Name `json:"name"`
	Roles Roles `json:"roles"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure193329939812341375193329939811801376193329939810521377193329939801648 struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
}
type RESTPatchAPIGuildEmojiResult struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Definitions struct {
	RESTGetAPIGuildEmojisResult RESTGetAPIGuildEmojisResult `json:"RESTGetAPIGuildEmojisResult"`
	APIEmoji APIEmoji `json:"APIEmoji"`
	APIPartialEmoji APIPartialEmoji `json:"APIPartialEmoji"`
	Snowflake Snowflake `json:"Snowflake"`
	APIUser APIUser `json:"APIUser"`
	UserFlags UserFlags `json:"UserFlags"`
	UserPremiumType UserPremiumType `json:"UserPremiumType"`
	RESTGetAPIGuildEmojiResult RESTGetAPIGuildEmojiResult `json:"RESTGetAPIGuildEmojiResult"`
	RESTPostAPIGuildEmojiJSONBody RESTPostAPIGuildEmojiJSONBody `json:"RESTPostAPIGuildEmojiJSONBody"`
	AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure193329939866391519332993986099161933299398470917193329939801648 AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure193329939866391519332993986099161933299398470917193329939801648 `json:"AddUndefinedToPossiblyUndefinedPropertiesOfInterface<structure-1933299398-663-915-1933299398-609-916-1933299398-470-917-1933299398-0-1648>"`
	RESTPostAPIGuildEmojiResult RESTPostAPIGuildEmojiResult `json:"RESTPostAPIGuildEmojiResult"`
	RESTPatchAPIGuildEmojiJSONBody RESTPatchAPIGuildEmojiJSONBody `json:"RESTPatchAPIGuildEmojiJSONBody"`
	AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure193329939812341375193329939811801376193329939810521377193329939801648 AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure193329939812341375193329939811801376193329939810521377193329939801648 `json:"AddUndefinedToPossiblyUndefinedPropertiesOfInterface<structure-1933299398-1234-1375-1933299398-1180-1376-1933299398-1052-1377-1933299398-0-1648>"`
	RESTPatchAPIGuildEmojiResult RESTPatchAPIGuildEmojiResult `json:"RESTPatchAPIGuildEmojiResult"`
}