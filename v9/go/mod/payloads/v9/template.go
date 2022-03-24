package v9

type Code struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Description struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type UsageCount struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type CreatorID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Creator struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type CreatedAt struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type UpdatedAt struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SourceGuildID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type SerializedSourceGuild struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type IsDirty struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Properties struct {
	Code                  Code                  `json:"code"`
	Name                  Name                  `json:"name"`
	Description           Description           `json:"description"`
	UsageCount            UsageCount            `json:"usage_count"`
	CreatorID             CreatorID             `json:"creator_id"`
	Creator               Creator               `json:"creator"`
	CreatedAt             CreatedAt             `json:"created_at"`
	UpdatedAt             UpdatedAt             `json:"updated_at"`
	SourceGuildID         SourceGuildID         `json:"source_guild_id"`
	SerializedSourceGuild SerializedSourceGuild `json:"serialized_source_guild"`
	IsDirty               IsDirty               `json:"is_dirty"`
}
type APITemplate struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Snowflake struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Username struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Discriminator struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Avatar struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Bot struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type System struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type MfaEnabled struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Banner struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type AccentColor struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Locale struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Verified struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Email struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Flags struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PremiumType struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PublicFlags struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Username      Username      `json:"username"`
	Discriminator Discriminator `json:"discriminator"`
	Avatar        Avatar        `json:"avatar"`
	Bot           Bot           `json:"bot"`
	System        System        `json:"system"`
	MfaEnabled    MfaEnabled    `json:"mfa_enabled"`
	Banner        Banner        `json:"banner"`
	AccentColor   AccentColor   `json:"accent_color"`
	Locale        Locale        `json:"locale"`
	Verified      Verified      `json:"verified"`
	Email         Email         `json:"email"`
	Flags         Flags         `json:"flags"`
	PremiumType   PremiumType   `json:"premium_type"`
	PublicFlags   PublicFlags   `json:"public_flags"`
}
type APIUser struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type UserFlags struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type UserPremiumType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Region struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type VerificationLevel struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type DefaultMessageNotifications struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type ExplicitContentFilter struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Roles struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Channels struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type AnyOf struct {
	Type string `json:"type,omitempty"`
	Ref  string `json:"$ref,omitempty"`
}
type AfkChannelID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type AfkTimeout struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SystemChannelID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type SystemChannelFlags struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PremiumProgressBarEnabled struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Description struct {
	Type []string `json:"type"`
}
type PreferredLocale struct {
	Type string `json:"type"`
}
type IconHash struct {
	Type []string `json:"type"`
}
type Properties struct {
	Name                        Name                        `json:"name"`
	Region                      Region                      `json:"region"`
	VerificationLevel           VerificationLevel           `json:"verification_level"`
	DefaultMessageNotifications DefaultMessageNotifications `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilter       `json:"explicit_content_filter"`
	Roles                       Roles                       `json:"roles"`
	Channels                    Channels                    `json:"channels"`
	AfkChannelID                AfkChannelID                `json:"afk_channel_id"`
	AfkTimeout                  AfkTimeout                  `json:"afk_timeout"`
	SystemChannelID             SystemChannelID             `json:"system_channel_id"`
	SystemChannelFlags          SystemChannelFlags          `json:"system_channel_flags"`
	PremiumProgressBarEnabled   PremiumProgressBarEnabled   `json:"premium_progress_bar_enabled"`
	Description                 Description                 `json:"description"`
	PreferredLocale             PreferredLocale             `json:"preferred_locale"`
	IconHash                    IconHash                    `json:"icon_hash"`
}
type APITemplateSerializedSourceGuild struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type GuildVerificationLevel struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type GuildDefaultMessageNotifications struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type GuildExplicitContentFilter struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Name struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
	Default     string   `json:"default"`
}
type AnyOf struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}
type Permissions struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
	Default     string  `json:"default"`
}
type Color struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
	Default     int      `json:"default"`
}
type Hoist struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
	Default     bool     `json:"default"`
}
type Icon struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type UnicodeEmoji struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Mentionable struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
	Default     bool     `json:"default"`
}
type ID struct {
	Type []string `json:"type"`
}
type Properties struct {
	Name         Name         `json:"name"`
	Permissions  Permissions  `json:"permissions"`
	Color        Color        `json:"color"`
	Hoist        Hoist        `json:"hoist"`
	Icon         Icon         `json:"icon"`
	UnicodeEmoji UnicodeEmoji `json:"unicode_emoji"`
	Mentionable  Mentionable  `json:"mentionable"`
	ID           ID           `json:"id"`
}
type APIGuildCreateRole struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type RESTPostAPIGuildRoleJSONBody struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	Name         Name         `json:"name"`
	Permissions  Permissions  `json:"permissions"`
	Color        Color        `json:"color"`
	Hoist        Hoist        `json:"hoist"`
	Icon         Icon         `json:"icon"`
	UnicodeEmoji UnicodeEmoji `json:"unicode_emoji"`
	Mentionable  Mentionable  `json:"mentionable"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure536029191157321651253602919115678165135360291911555316514536029191023031 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Name struct {
	Type string `json:"type"`
}
type ParentID struct {
	Type []string `json:"type"`
}
type PermissionOverwrites struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}
type Properties struct {
	Name                 Name                 `json:"name"`
	ID                   ID                   `json:"id"`
	ParentID             ParentID             `json:"parent_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
}
type APIGuildCreatePartialChannel struct {
	Type                 string     `json:"type"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
}
type StrictPartialAlias1589253880731877332815892538800215641TypeTopicNsfwBitrateUserLimitRateLimitPerUser struct {
	Ref string `json:"$ref"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceAlias1589253880728677297015892538800215641Alias1589253880731877332815892538800215641TypeTopicNsfwBitrateUserLimitRateLimitPerUser struct {
	Type                 string `json:"type"`
	AdditionalProperties bool   `json:"additionalProperties"`
	Description          string `json:"description"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure5360291911242137753602919111871378536029191102913785360291919861379536029191023031 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Allow struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
	Default     string  `json:"default"`
}
type Deny struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
	Default     string  `json:"default"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	Allow Allow `json:"allow"`
	Deny  Deny  `json:"deny"`
	Type  Type  `json:"type"`
	ID    ID    `json:"id"`
}
type APIGuildCreateOverwrite struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Properties struct {
	Allow Allow `json:"allow"`
	Deny  Deny  `json:"deny"`
	Type  Type  `json:"type"`
}
type RESTPutAPIChannelPermissionJSONBody struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type OverwriteType struct {
	Type string `json:"type"`
	Enum []int  `json:"enum"`
}
type GuildSystemChannelFlags struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
