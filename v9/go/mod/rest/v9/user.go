package v9

type RESTGetAPICurrentUserResult struct {
	Ref         string `json:"$ref"`
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
type Snowflake struct {
	Type        string `json:"type"`
	Description string `json:"description"`
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
type RESTGetAPIUserResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type RESTGetCurrentUserGuildMemberResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
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
}
type RESTPatchAPICurrentUserJSONBody struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	Username Username `json:"username"`
	Avatar   Avatar   `json:"avatar"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure154024499586210561540244995808105715402449956791058154024499502551 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTPatchAPICurrentUserResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
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
}
type GuildFeature struct {
	Type        string   `json:"type"`
	Enum        []string `json:"enum"`
	Description string   `json:"description"`
}
type RESTGetAPICurrentUserGuildsResult struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
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
}
type RESTPostAPICurrentUserCreateDMChannelResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIChannel struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
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
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
}
type APITextBasedChannelEnum13545260107743780613545260107458876213545260100336361767421466 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum13545260107743780613545260107458876213545260100336361767421466 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type APIPartialChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type ChannelType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
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
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
}
type APITextBasedChannelEnum1354526010763576841354526010745887621354526010033636102715350 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum1354526010763576841354526010745887621354526010033636102715350 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
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
type DefaultAutoArchiveDuration struct {
	Ref         string `json:"$ref"`
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
	ID                         ID                         `json:"id"`
	Type                       Type                       `json:"type"`
	Name                       Name                       `json:"name"`
	GuildID                    GuildID                    `json:"guild_id"`
	PermissionOverwrites       PermissionOverwrites       `json:"permission_overwrites"`
	Position                   Position                   `json:"position"`
	ParentID                   ParentID                   `json:"parent_id"`
	Nsfw                       Nsfw                       `json:"nsfw"`
	LastMessageID              LastMessageID              `json:"last_message_id"`
	DefaultAutoArchiveDuration DefaultAutoArchiveDuration `json:"default_auto_archive_duration"`
	Topic                      Topic                      `json:"topic"`
	LastPinTimestamp           LastPinTimestamp           `json:"last_pin_timestamp"`
	RateLimitPerUser           RateLimitPerUser           `json:"rate_limit_per_user"`
}
type APITextChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID                         ID                         `json:"id"`
	Type                       Type                       `json:"type"`
	Name                       Name                       `json:"name"`
	GuildID                    GuildID                    `json:"guild_id"`
	PermissionOverwrites       PermissionOverwrites       `json:"permission_overwrites"`
	Position                   Position                   `json:"position"`
	ParentID                   ParentID                   `json:"parent_id"`
	Nsfw                       Nsfw                       `json:"nsfw"`
	LastMessageID              LastMessageID              `json:"last_message_id"`
	DefaultAutoArchiveDuration DefaultAutoArchiveDuration `json:"default_auto_archive_duration"`
	Topic                      Topic                      `json:"topic"`
	LastPinTimestamp           LastPinTimestamp           `json:"last_pin_timestamp"`
}
type APIGuildTextChannelEnum1354526010757976341354526010745887621354526010033636732917173 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
}
type APITextBasedChannelEnum1354526010757976341354526010745887621354526010033636732917173 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum1354526010757976341354526010745887621354526010033636732917173 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
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
type APIGuildChannelEnum1354526010757976341354526010745887621354526010033636732917173 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
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
}
type OverwriteType struct {
	Type string `json:"type"`
	Enum []int  `json:"enum"`
}
type ThreadAutoArchiveDuration struct {
	Type string `json:"type"`
	Enum []int  `json:"enum"`
}
type APINewsChannel struct {
	Ref string `json:"$ref"`
}
type Type struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Properties struct {
	ID                         ID                         `json:"id"`
	Type                       Type                       `json:"type"`
	Name                       Name                       `json:"name"`
	GuildID                    GuildID                    `json:"guild_id"`
	PermissionOverwrites       PermissionOverwrites       `json:"permission_overwrites"`
	Position                   Position                   `json:"position"`
	ParentID                   ParentID                   `json:"parent_id"`
	Nsfw                       Nsfw                       `json:"nsfw"`
	LastMessageID              LastMessageID              `json:"last_message_id"`
	DefaultAutoArchiveDuration DefaultAutoArchiveDuration `json:"default_auto_archive_duration"`
	Topic                      Topic                      `json:"topic"`
	LastPinTimestamp           LastPinTimestamp           `json:"last_pin_timestamp"`
}
type APIGuildTextChannelEnum13545260107990815313545260107458876213545260100336361915191889 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Type          Type          `json:"type"`
	Name          Name          `json:"name"`
	LastMessageID LastMessageID `json:"last_message_id"`
}
type APITextBasedChannelEnum13545260107990815313545260107458876213545260100336361915191889 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum13545260107990815313545260107458876213545260100336361915191889 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
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
type APIGuildChannelEnum13545260107990815313545260107458876213545260100336361915191889 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type APIGuildStoreChannel struct {
	Ref string `json:"$ref"`
}
type APIGuildChannelEnum1354526010815483381354526010745887621354526010033636138797936 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum1354526010815483381354526010745887621354526010033636138797936 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
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
type APIGuildChannelEnum13545260108604875913545260107458876213545260100336361830896565Enum1354526010768577421354526010745887621354526010033636121029024 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum13545260108604875913545260107458876213545260100336361830896565Enum1354526010768577421354526010745887621354526010033636121029024 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type VideoQualityMode struct {
	Type string `json:"type"`
	Enum []int  `json:"enum"`
}
type APIGuildCategoryChannel struct {
	Ref string `json:"$ref"`
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
type APIGuildChannelEnum13545260107807798913545260107458876213545260100336361025724393 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum13545260107807798913545260107458876213545260100336361025724393 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Type struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Member struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type ThreadMetadata struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type MessageCount struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type MemberCount struct {
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
	Member               Member               `json:"member"`
	ThreadMetadata       ThreadMetadata       `json:"thread_metadata"`
	MessageCount         MessageCount         `json:"message_count"`
	MemberCount          MemberCount          `json:"member_count"`
	RateLimitPerUser     RateLimitPerUser     `json:"rate_limit_per_user"`
	OwnerID              OwnerID              `json:"owner_id"`
	LastMessageID        LastMessageID        `json:"last_message_id"`
}
type APIThreadChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
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
type APIGuildChannelEnum1354526010843085151354526010745887621354526010033636612297701Enum13545260108516860313545260107458876213545260100336361507172815Enum13545260108339842913545260107458876213545260100336361372201880 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Required             []string   `json:"required"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIChannelBaseEnum1354526010843085151354526010745887621354526010033636612297701Enum13545260108516860313545260107458876213545260100336361507172815Enum13545260108339842913545260107458876213545260100336361372201880 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type UserID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type JoinTimestamp struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID            ID            `json:"id"`
	UserID        UserID        `json:"user_id"`
	JoinTimestamp JoinTimestamp `json:"join_timestamp"`
	Flags         Flags         `json:"flags"`
}
type APIThreadMember struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type ThreadMemberFlags struct {
	Type []interface{} `json:"type"`
	Enum []interface{} `json:"enum"`
}
type Archived struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type AutoArchiveDuration struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type ArchiveTimestamp struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Locked struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Invitable struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type CreateTimestamp struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Archived            Archived            `json:"archived"`
	AutoArchiveDuration AutoArchiveDuration `json:"auto_archive_duration"`
	ArchiveTimestamp    ArchiveTimestamp    `json:"archive_timestamp"`
	Locked              Locked              `json:"locked"`
	Invitable           Invitable           `json:"invitable"`
	CreateTimestamp     CreateTimestamp     `json:"create_timestamp"`
}
type APIThreadMetadata struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTGetAPICurrentUserConnectionsResult struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
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
}
type APIGuildIntegrationType struct {
	Type string   `json:"type"`
	Enum []string `json:"enum"`
}
type IntegrationExpireBehavior struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
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
}
type ConnectionVisibility struct {
	Type string `json:"type"`
	Enum []int  `json:"enum"`
}
