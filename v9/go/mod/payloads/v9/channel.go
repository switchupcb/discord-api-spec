package v9

type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID   ID   `json:"id"`
	Type Type `json:"type"`
	Name Name `json:"name"`
}
type APIPartialChannel struct {
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
type ChannelType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type AnyOf struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type TextChannelType struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type GuildChannelType struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type GuildTextChannelType struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type Type struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type GuildID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
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
type AnyOf struct {
	Ref  string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
}
type ParentID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Nsfw struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type LastMessageID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
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
type APIGuildCategoryChannel struct {
	Ref string `json:"$ref"`
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
type APIGuildStoreChannel struct {
	Ref string `json:"$ref"`
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
type Type struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Recipients struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
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
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIChannel struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type ChannelID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Author struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Content struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Timestamp struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type EditedTimestamp struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Tts struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type MentionEveryone struct {
	Type        string `json:"type"`
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
type Member struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Properties struct {
	Member        Member        `json:"member"`
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
type Items struct {
	Type                 string     `json:"type"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
}
type Mentions struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Items struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type MentionRoles struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type MentionChannels struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Attachments struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Embeds struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Reactions struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Nonce struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Pinned struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type WebhookID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Activity struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Description struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Items struct {
	Type string `json:"type"`
}
type RPCOrigins struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type BotPublic struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type BotRequireCodeGrant struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type TermsOfServiceURL struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type PrivacyPolicyURL struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Owner struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Summary struct {
	Type        string `json:"type"`
	Const       string `json:"const"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type VerifyKey struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type AnyOf struct {
	Ref  string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
}
type Team struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type PrimarySkuID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Slug struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type CoverImage struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID                  ID                  `json:"id"`
	Name                Name                `json:"name"`
	Icon                Icon                `json:"icon"`
	Description         Description         `json:"description"`
	RPCOrigins          RPCOrigins          `json:"rpc_origins"`
	BotPublic           BotPublic           `json:"bot_public"`
	BotRequireCodeGrant BotRequireCodeGrant `json:"bot_require_code_grant"`
	TermsOfServiceURL   TermsOfServiceURL   `json:"terms_of_service_url"`
	PrivacyPolicyURL    PrivacyPolicyURL    `json:"privacy_policy_url"`
	Owner               Owner               `json:"owner"`
	Summary             Summary             `json:"summary"`
	VerifyKey           VerifyKey           `json:"verify_key"`
	Team                Team                `json:"team"`
	GuildID             GuildID             `json:"guild_id"`
	PrimarySkuID        PrimarySkuID        `json:"primary_sku_id"`
	Slug                Slug                `json:"slug"`
	CoverImage          CoverImage          `json:"cover_image"`
	Flags               Flags               `json:"flags"`
}
type Application struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type MessageReference struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type ReferencedMessage struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Interaction struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Thread struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Components struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type StickerItems struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Stickers struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Properties struct {
	ID                ID                `json:"id"`
	ChannelID         ChannelID         `json:"channel_id"`
	GuildID           GuildID           `json:"guild_id"`
	Author            Author            `json:"author"`
	Member            Member            `json:"member"`
	Content           Content           `json:"content"`
	Timestamp         Timestamp         `json:"timestamp"`
	EditedTimestamp   EditedTimestamp   `json:"edited_timestamp"`
	Tts               Tts               `json:"tts"`
	MentionEveryone   MentionEveryone   `json:"mention_everyone"`
	Mentions          Mentions          `json:"mentions"`
	MentionRoles      MentionRoles      `json:"mention_roles"`
	MentionChannels   MentionChannels   `json:"mention_channels"`
	Attachments       Attachments       `json:"attachments"`
	Embeds            Embeds            `json:"embeds"`
	Reactions         Reactions         `json:"reactions"`
	Nonce             Nonce             `json:"nonce"`
	Pinned            Pinned            `json:"pinned"`
	WebhookID         WebhookID         `json:"webhook_id"`
	Type              Type              `json:"type"`
	Activity          Activity          `json:"activity"`
	Application       Application       `json:"application"`
	ApplicationID     ApplicationID     `json:"application_id"`
	MessageReference  MessageReference  `json:"message_reference"`
	Flags             Flags             `json:"flags"`
	ReferencedMessage ReferencedMessage `json:"referenced_message"`
	Interaction       Interaction       `json:"interaction"`
	Thread            Thread            `json:"thread"`
	Components        Components        `json:"components"`
	StickerItems      StickerItems      `json:"sticker_items"`
	Stickers          Stickers          `json:"stickers"`
}
type APIMessage struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type User struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
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
type Properties struct {
	ID      ID      `json:"id"`
	GuildID GuildID `json:"guild_id"`
	Type    Type    `json:"type"`
	Name    Name    `json:"name"`
}
type APIChannelMention struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Filename struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ContentType struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Size struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type URL struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ProxyURL struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Height struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Width struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Ephemeral struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID          ID          `json:"id"`
	Filename    Filename    `json:"filename"`
	Description Description `json:"description"`
	ContentType ContentType `json:"content_type"`
	Size        Size        `json:"size"`
	URL         URL         `json:"url"`
	ProxyURL    ProxyURL    `json:"proxy_url"`
	Height      Height      `json:"height"`
	Width       Width       `json:"width"`
	Ephemeral   Ephemeral   `json:"ephemeral"`
}
type APIAttachment struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Title struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Color struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Footer struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Image struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Thumbnail struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Video struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Provider struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Fields struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Properties struct {
	Title       Title       `json:"title"`
	Type        Type        `json:"type"`
	Description Description `json:"description"`
	URL         URL         `json:"url"`
	Timestamp   Timestamp   `json:"timestamp"`
	Color       Color       `json:"color"`
	Footer      Footer      `json:"footer"`
	Image       Image       `json:"image"`
	Thumbnail   Thumbnail   `json:"thumbnail"`
	Video       Video       `json:"video"`
	Provider    Provider    `json:"provider"`
	Author      Author      `json:"author"`
	Fields      Fields      `json:"fields"`
}
type APIEmbed struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type EmbedType struct {
	Type        string   `json:"type"`
	Enum        []string `json:"enum"`
	Description string   `json:"description"`
	Deprecated  string   `json:"deprecated"`
}
type Text struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type IconURL struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ProxyIconURL struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Text         Text         `json:"text"`
	IconURL      IconURL      `json:"icon_url"`
	ProxyIconURL ProxyIconURL `json:"proxy_icon_url"`
}
type APIEmbedFooter struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	URL      URL      `json:"url"`
	ProxyURL ProxyURL `json:"proxy_url"`
	Height   Height   `json:"height"`
	Width    Width    `json:"width"`
}
type APIEmbedImage struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type APIEmbedThumbnail struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	URL    URL    `json:"url"`
	Height Height `json:"height"`
	Width  Width  `json:"width"`
}
type APIEmbedVideo struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	Name Name `json:"name"`
	URL  URL  `json:"url"`
}
type APIEmbedProvider struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	Name         Name         `json:"name"`
	URL          URL          `json:"url"`
	IconURL      IconURL      `json:"icon_url"`
	ProxyIconURL ProxyIconURL `json:"proxy_icon_url"`
}
type APIEmbedAuthor struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Value struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Inline struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Name   Name   `json:"name"`
	Value  Value  `json:"value"`
	Inline Inline `json:"inline"`
}
type APIEmbedField struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Count struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Me struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Emoji struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	Count Count `json:"count"`
	Me    Me    `json:"me"`
	Emoji Emoji `json:"emoji"`
}
type APIReaction struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type ID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Animated struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID       ID       `json:"id"`
	Name     Name     `json:"name"`
	Animated Animated `json:"animated"`
}
type APIPartialEmoji struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type MessageType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PartyID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Type    Type    `json:"type"`
	PartyID PartyID `json:"party_id"`
}
type APIMessageActivity struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type MessageActivityType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Members struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type OwnerUserID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	Icon        Icon        `json:"icon"`
	ID          ID          `json:"id"`
	Members     Members     `json:"members"`
	Name        Name        `json:"name"`
	OwnerUserID OwnerUserID `json:"owner_user_id"`
}
type APITeam struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type MembershipState struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Items struct {
	Type  string `json:"type"`
	Const string `json:"const"`
}
type Permissions struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	MinItems    int    `json:"minItems"`
	MaxItems    int    `json:"maxItems"`
	Description string `json:"description"`
}
type TeamID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	MembershipState MembershipState `json:"membership_state"`
	Permissions     Permissions     `json:"permissions"`
	TeamID          TeamID          `json:"team_id"`
	User            User            `json:"user"`
}
type APITeamMember struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type TeamMemberMembershipState struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type ApplicationFlags struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type MessageID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	MessageID MessageID `json:"message_id"`
	ChannelID ChannelID `json:"channel_id"`
	GuildID   GuildID   `json:"guild_id"`
}
type APIMessageReference struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type MessageFlags struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Member struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID     ID     `json:"id"`
	Type   Type   `json:"type"`
	Name   Name   `json:"name"`
	User   User   `json:"user"`
	Member Member `json:"member"`
}
type APIMessageInteraction struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type InteractionType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Properties struct {
	Roles                      Roles                      `json:"roles"`
	PremiumSince               PremiumSince               `json:"premium_since"`
	Pending                    Pending                    `json:"pending"`
	Nick                       Nick                       `json:"nick"`
	Mute                       Mute                       `json:"mute"`
	JoinedAt                   JoinedAt                   `json:"joined_at"`
	Deaf                       Deaf                       `json:"deaf"`
	CommunicationDisabledUntil CommunicationDisabledUntil `json:"communication_disabled_until"`
	Avatar                     Avatar                     `json:"avatar"`
}
type PartialAPIMessageInteractionGuildMember struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Type struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Properties struct {
	Type       Type       `json:"type"`
	Components Components `json:"components"`
}
type APIActionRowComponentAPIMessageActionRowComponent struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum13545260102870528755135452601028581288941354526010033636105957764 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIMessageActionRowComponent struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type APIButtonComponent struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type Label struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type AnyOf struct {
	Type  string `json:"type"`
	Const int    `json:"const"`
}
type Style struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Disabled struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type CustomID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Type     Type     `json:"type"`
	Label    Label    `json:"label"`
	Style    Style    `json:"style"`
	Emoji    Emoji    `json:"emoji"`
	Disabled Disabled `json:"disabled"`
	CustomID CustomID `json:"custom_id"`
}
type APIButtonComponentWithCustomID struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Properties struct {
	Type     Type     `json:"type"`
	Label    Label    `json:"label"`
	Style    Style    `json:"style"`
	Emoji    Emoji    `json:"emoji"`
	Disabled Disabled `json:"disabled"`
}
type APIButtonComponentBaseEnum13545260103053930552135452601030405305931354526010033636Enum13545260103055330564135452601030405305931354526010033636Enum13545260103056530574135452601030405305931354526010033636Enum13545260103057530583135452601030405305931354526010033636 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum135452601028756287951354526010285812889413545260100336361131375322 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	ID       ID       `json:"id"`
	Name     Name     `json:"name"`
	Animated Animated `json:"animated"`
}
type APIMessageComponentEmoji struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Style struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Properties struct {
	Type     Type     `json:"type"`
	Label    Label    `json:"label"`
	Style    Style    `json:"style"`
	Emoji    Emoji    `json:"emoji"`
	Disabled Disabled `json:"disabled"`
	URL      URL      `json:"url"`
}
type APIButtonComponentWithURL struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Properties struct {
	Type     Type     `json:"type"`
	Label    Label    `json:"label"`
	Style    Style    `json:"style"`
	Emoji    Emoji    `json:"emoji"`
	Disabled Disabled `json:"disabled"`
}
type APIButtonComponentBaseEnum13545260103058430590135452601030405305931354526010033636 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Options struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Placeholder struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type MinValues struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     int    `json:"default"`
}
type MaxValues struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     int    `json:"default"`
}
type Disabled struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     bool   `json:"default"`
}
type Properties struct {
	Type        Type        `json:"type"`
	CustomID    CustomID    `json:"custom_id"`
	Options     Options     `json:"options"`
	Placeholder Placeholder `json:"placeholder"`
	MinValues   MinValues   `json:"min_values"`
	MaxValues   MaxValues   `json:"max_values"`
	Disabled    Disabled    `json:"disabled"`
}
type APISelectMenuComponent struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum13545260102879628844135452601028581288941354526010033636663308133 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Default struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Label       Label       `json:"label"`
	Value       Value       `json:"value"`
	Description Description `json:"description"`
	Emoji       Emoji       `json:"emoji"`
	Default     Default     `json:"default"`
}
type APISelectMenuOption struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type FormatType struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID         ID         `json:"id"`
	Name       Name       `json:"name"`
	FormatType FormatType `json:"format_type"`
}
type APIStickerItem struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type StickerFormatType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type PackID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Tags struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Asset struct {
	Type        string `json:"type"`
	Const       string `json:"const"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Available struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SortValue struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID          ID          `json:"id"`
	PackID      PackID      `json:"pack_id"`
	Name        Name        `json:"name"`
	Description Description `json:"description"`
	Tags        Tags        `json:"tags"`
	Asset       Asset       `json:"asset"`
	Type        Type        `json:"type"`
	FormatType  FormatType  `json:"format_type"`
	Available   Available   `json:"available"`
	GuildID     GuildID     `json:"guild_id"`
	User        User        `json:"user"`
	SortValue   SortValue   `json:"sort_value"`
}
type APISticker struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type StickerType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Properties struct {
	ChannelID ChannelID `json:"channel_id"`
	WebhookID WebhookID `json:"webhook_id"`
}
type APIFollowedChannel struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Threads struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type HasMore struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Threads Threads `json:"threads"`
	Members Members `json:"members"`
	HasMore HasMore `json:"has_more"`
}
type APIThreadList struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type AllowedMentionsTypes struct {
	Type        string   `json:"type"`
	Enum        []string `json:"enum"`
	Description string   `json:"description"`
}
type Parse struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Users struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type RepliedUser struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     bool   `json:"default"`
}
type Properties struct {
	Parse       Parse       `json:"parse"`
	Roles       Roles       `json:"roles"`
	Users       Users       `json:"users"`
	RepliedUser RepliedUser `json:"replied_user"`
}
type APIAllowedMentions struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type ComponentType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type ButtonStyle struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type TextInputStyle struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Type struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type Style struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type MinLength struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type MaxLength struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Required struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Type        Type        `json:"type"`
	Style       Style       `json:"style"`
	CustomID    CustomID    `json:"custom_id"`
	Label       Label       `json:"label"`
	Placeholder Placeholder `json:"placeholder"`
	Value       Value       `json:"value"`
	MinLength   MinLength   `json:"min_length"`
	MaxLength   MaxLength   `json:"max_length"`
	Required    Required    `json:"required"`
}
type APITextInputComponent struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Properties struct {
	Type Type `json:"type"`
}
type APIBaseComponentEnum135452601028845288911354526010285812889413545260100336361182532533 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIMessageComponent struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type APIModalComponent struct {
	AnyOf []AnyOf `json:"anyOf"`
}
type APIModalActionRowComponent struct {
	Ref string `json:"$ref"`
}
type Properties struct {
	Type       Type       `json:"type"`
	Components Components `json:"components"`
}
type APIActionRowComponentAPIModalActionRowComponent struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type APIActionRowComponentTypes struct {
	AnyOf []AnyOf `json:"anyOf"`
}
