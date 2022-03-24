package v8

type RESTGetAPICurrentUserResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
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
	Deprecated           string     `json:"deprecated"`
}
type Snowflake struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type UserFlags struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type UserPremiumType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type RESTGetAPIUserResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type RESTGetCurrentUserGuildMemberResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type User struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Nick struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Roles struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type JoinedAt struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type PremiumSince struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Deaf struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Mute struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Pending struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type CommunicationDisabledUntil struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Properties struct {
	User                       User                       `json:"user"`
	Nick                       Nick                       `json:"nick"`
	Avatar                     Avatar                     `json:"avatar"`
	Roles                      Roles                      `json:"roles"`
	JoinedAt                   JoinedAt                   `json:"joined_at"`
	PremiumSince               PremiumSince               `json:"premium_since"`
	Deaf                       Deaf                       `json:"deaf"`
	Mute                       Mute                       `json:"mute"`
	Pending                    Pending                    `json:"pending"`
	CommunicationDisabledUntil CommunicationDisabledUntil `json:"communication_disabled_until"`
}
type APIGuildMember struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type RESTPatchAPICurrentUserJSONBody struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Properties struct {
	Username Username `json:"username"`
	Avatar   Avatar   `json:"avatar"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure94726783813381532947267838128415339472678381036153494726783803987 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTPatchAPICurrentUserResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Before struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type After struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Limit struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     int    `json:"default"`
}
type Properties struct {
	Before Before `json:"before"`
	After  After  `json:"after"`
	Limit  Limit  `json:"limit"`
}
type RESTGetAPICurrentUserGuildsQuery struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type ID struct {
	Ref string `json:"$ref"`
}
type Name struct {
	Type string `json:"type"`
}
type Icon struct {
	Type []string `json:"type"`
}
type Owner struct {
	Type string `json:"type"`
}
type Features struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}
type Permissions struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID          ID          `json:"id"`
	Name        Name        `json:"name"`
	Icon        Icon        `json:"icon"`
	Owner       Owner       `json:"owner"`
	Features    Features    `json:"features"`
	Permissions Permissions `json:"permissions"`
}
type RESTAPIPartialCurrentUserGuild struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Deprecated           string     `json:"deprecated"`
}
type GuildFeature struct {
	Type        string   `json:"type"`
	Enum        []string `json:"enum"`
	Description string   `json:"description"`
	Deprecated  string   `json:"deprecated"`
}
type RESTGetAPICurrentUserGuildsResult struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type RecipientID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	RecipientID RecipientID `json:"recipient_id"`
}
type RESTPostAPICurrentUserCreateDMChannelJSONBody struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type RESTPostAPICurrentUserCreateDMChannelResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIChannel struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
	Deprecated  string  `json:"deprecated"`
}
type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Type struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type AnyOf struct {
	Ref  string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
}
type LastMessageID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Recipients struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type ApplicationID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Icon struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type OwnerID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
	Recipients    Recipients    `json:"recipients"`
	ApplicationID ApplicationID `json:"application_id"`
	Icon          Icon          `json:"icon"`
	OwnerID       OwnerID       `json:"owner_id"`
}
type APIGroupDMChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
}
type APITextBasedChannelEnum12254432917712777512254432917308847112254432910358711767421466 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum12254432917712777512254432917308847112254432910358711767421466 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type APIPartialChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type ChannelType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
	Recipients    Recipients    `json:"recipients"`
}
type APIDMChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
}
type APITextBasedChannelEnum1225443291760476531225443291730884711225443291035871102715350 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum1225443291760476531225443291730884711225443291035871102715350 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type GuildID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PermissionOverwrites struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Position struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ParentID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Nsfw struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Topic struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type LastPinTimestamp struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type RateLimitPerUser struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	Type                 Type                 `json:"type"`
	Name                 Name                 `json:"name"`
	GuildID              GuildID              `json:"guild_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
	Position             Position             `json:"position"`
	ParentID             ParentID             `json:"parent_id"`
	Nsfw                 Nsfw                 `json:"nsfw"`
	LastMessageID        LastMessageID        `json:"last_message_id"`
	Topic                Topic                `json:"topic"`
	LastPinTimestamp     LastPinTimestamp     `json:"last_pin_timestamp"`
	RateLimitPerUser     RateLimitPerUser     `json:"rate_limit_per_user"`
}
type APITextChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	Type                 Type                 `json:"type"`
	Name                 Name                 `json:"name"`
	GuildID              GuildID              `json:"guild_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
	Position             Position             `json:"position"`
	ParentID             ParentID             `json:"parent_id"`
	Nsfw                 Nsfw                 `json:"nsfw"`
	LastMessageID        LastMessageID        `json:"last_message_id"`
	Topic                Topic                `json:"topic"`
	LastPinTimestamp     LastPinTimestamp     `json:"last_pin_timestamp"`
}
type APIGuildTextChannelEnum1225443291754876031225443291730884711225443291035871732917173 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
}
type APITextBasedChannelEnum1225443291754876031225443291730884711225443291035871732917173 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum1225443291754876031225443291730884711225443291035871732917173 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	Type                 Type                 `json:"type"`
	Name                 Name                 `json:"name"`
	GuildID              GuildID              `json:"guild_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
	Position             Position             `json:"position"`
	ParentID             ParentID             `json:"parent_id"`
	Nsfw                 Nsfw                 `json:"nsfw"`
}
type APIGuildChannelEnum1225443291754876031225443291730884711225443291035871732917173 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Allow struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Deny struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID    ID    `json:"id"`
	Type  Type  `json:"type"`
	Allow Allow `json:"allow"`
	Deny  Deny  `json:"deny"`
}
type APIOverwrite struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type OverwriteType struct {
	Type       string `json:"type"`
	Enum       []int  `json:"enum"`
	Deprecated string `json:"deprecated"`
}
type APINewsChannel struct {
	Ref        string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type Type struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	Type                 Type                 `json:"type"`
	Name                 Name                 `json:"name"`
	GuildID              GuildID              `json:"guild_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
	Position             Position             `json:"position"`
	ParentID             ParentID             `json:"parent_id"`
	Nsfw                 Nsfw                 `json:"nsfw"`
	LastMessageID        LastMessageID        `json:"last_message_id"`
	Topic                Topic                `json:"topic"`
	LastPinTimestamp     LastPinTimestamp     `json:"last_pin_timestamp"`
}
type APIGuildTextChannelEnum12254432917959812212254432917308847112254432910358711915191889 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
}
type APITextBasedChannelEnum12254432917959812212254432917308847112254432910358711915191889 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum12254432917959812212254432917308847112254432910358711915191889 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	Type                 Type                 `json:"type"`
	Name                 Name                 `json:"name"`
	GuildID              GuildID              `json:"guild_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
	Position             Position             `json:"position"`
	ParentID             ParentID             `json:"parent_id"`
	Nsfw                 Nsfw                 `json:"nsfw"`
}
type APIGuildChannelEnum12254432917959812212254432917308847112254432910358711915191889 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type APIGuildStoreChannel struct {
	Ref        string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type APIGuildChannelEnum1225443291812383071225443291730884711225443291035871138797936 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum1225443291812383071225443291730884711225443291035871138797936 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type AnyOf struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Type struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Bitrate struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type UserLimit struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type RtcRegion struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type VideoQualityMode struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	Type                 Type                 `json:"type"`
	Name                 Name                 `json:"name"`
	GuildID              GuildID              `json:"guild_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
	Position             Position             `json:"position"`
	ParentID             ParentID             `json:"parent_id"`
	Nsfw                 Nsfw                 `json:"nsfw"`
	Bitrate              Bitrate              `json:"bitrate"`
	UserLimit            UserLimit            `json:"user_limit"`
	RtcRegion            RtcRegion            `json:"rtc_region"`
	VideoQualityMode     VideoQualityMode     `json:"video_quality_mode"`
}
type APIVoiceChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	Type                 Type                 `json:"type"`
	Name                 Name                 `json:"name"`
	GuildID              GuildID              `json:"guild_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
	Position             Position             `json:"position"`
	ParentID             ParentID             `json:"parent_id"`
	Nsfw                 Nsfw                 `json:"nsfw"`
}
type APIGuildChannelEnum12254432918308846812254432917308847112254432910358711830896565Enum1225443291765477111225443291730884711225443291035871121029024 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum12254432918308846812254432917308847112254432910358711830896565Enum1225443291765477111225443291730884711225443291035871121029024 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type VideoQualityMode struct {
	Type       string `json:"type"`
	Enum       []int  `json:"enum"`
	Deprecated string `json:"deprecated"`
}
type APIGuildCategoryChannel struct {
	Ref        string `json:"$ref"`
	Deprecated string `json:"deprecated"`
}
type Type struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	Type                 Type                 `json:"type"`
	Name                 Name                 `json:"name"`
	GuildID              GuildID              `json:"guild_id"`
	PermissionOverwrites PermissionOverwrites `json:"permission_overwrites"`
	Position             Position             `json:"position"`
	ParentID             ParentID             `json:"parent_id"`
	Nsfw                 Nsfw                 `json:"nsfw"`
}
type APIGuildChannelEnum12254432917776795812254432917308847112254432910358711025724393 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
	Deprecated           string     `json:"deprecated"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum12254432917776795812254432917308847112254432910358711025724393 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type RESTGetAPICurrentUserConnectionsResult struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type ID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Type struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Revoked struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Enabled struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Syncing struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type RoleID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type EnableEmoticons struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ExpireBehavior struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type ExpireGracePeriod struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Account struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type SyncedAt struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SubscriberCount struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Application struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID                ID                `json:"id"`
	Name              Name              `json:"name"`
	Type              Type              `json:"type"`
	Enabled           Enabled           `json:"enabled"`
	Syncing           Syncing           `json:"syncing"`
	RoleID            RoleID            `json:"role_id"`
	EnableEmoticons   EnableEmoticons   `json:"enable_emoticons"`
	ExpireBehavior    ExpireBehavior    `json:"expire_behavior"`
	ExpireGracePeriod ExpireGracePeriod `json:"expire_grace_period"`
	User              User              `json:"user"`
	Account           Account           `json:"account"`
	SyncedAt          SyncedAt          `json:"synced_at"`
	SubscriberCount   SubscriberCount   `json:"subscriber_count"`
	Revoked           Revoked           `json:"revoked"`
	Application       Application       `json:"application"`
}
type Items struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Integrations struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type FriendSync struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ShowActivity struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Visibility struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID           ID           `json:"id"`
	Name         Name         `json:"name"`
	Type         Type         `json:"type"`
	Revoked      Revoked      `json:"revoked"`
	Integrations Integrations `json:"integrations"`
	Verified     Verified     `json:"verified"`
	FriendSync   FriendSync   `json:"friend_sync"`
	ShowActivity ShowActivity `json:"show_activity"`
	Visibility   Visibility   `json:"visibility"`
}
type APIConnection struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type APIGuildInteractionType struct {
	Type       string   `json:"type"`
	Enum       []string `json:"enum"`
	Deprecated string   `json:"deprecated"`
}
type IntegrationExpireBehavior struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type ID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Name Name `json:"name"`
}
type APIIntegrationAccount struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Description struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Summary struct {
	Type        string `json:"type"`
	Const       string `json:"const"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Bot struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID          ID          `json:"id"`
	Name        Name        `json:"name"`
	Icon        Icon        `json:"icon"`
	Description Description `json:"description"`
	Summary     Summary     `json:"summary"`
	Bot         Bot         `json:"bot"`
}
type APIGuildIntegrationApplication struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type ConnectionVisibility struct {
	Type       string `json:"type"`
	Enum       []int  `json:"enum"`
	Deprecated string `json:"deprecated"`
}
