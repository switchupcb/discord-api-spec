package v9

type WithUserCount struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	WithUserCount WithUserCount `json:"with_user_count"`
}
type RESTGetAPIGuildScheduledEventsQuery struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type RESTGetAPIGuildScheduledEventsResult struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type AnyOf struct {
	Ref string `json:"$ref"`
}
type APIGuildScheduledEvent struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type GuildID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type ChannelID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type AnyOf struct {
	Ref  string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
}
type CreatorID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Description struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type ScheduledStartTime struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ScheduledEndTime struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type PrivacyLevel struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Status struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type EntityType struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type EntityID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type EntityMetadata struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Creator struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type UserCount struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Image struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Properties struct {
	ID                 ID                 `json:"id"`
	GuildID            GuildID            `json:"guild_id"`
	ChannelID          ChannelID          `json:"channel_id"`
	CreatorID          CreatorID          `json:"creator_id"`
	Name               Name               `json:"name"`
	Description        Description        `json:"description"`
	ScheduledStartTime ScheduledStartTime `json:"scheduled_start_time"`
	ScheduledEndTime   ScheduledEndTime   `json:"scheduled_end_time"`
	PrivacyLevel       PrivacyLevel       `json:"privacy_level"`
	Status             Status             `json:"status"`
	EntityType         EntityType         `json:"entity_type"`
	EntityID           EntityID           `json:"entity_id"`
	EntityMetadata     EntityMetadata     `json:"entity_metadata"`
	Creator            Creator            `json:"creator"`
	UserCount          UserCount          `json:"user_count"`
	Image              Image              `json:"image"`
}
type APIStageInstanceGuildScheduledEvent struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Snowflake struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type GuildScheduledEventPrivacyLevel struct {
	Type        string `json:"type"`
	Const       int    `json:"const"`
	Description string `json:"description"`
}
type GuildScheduledEventStatus struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Location struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Location Location `json:"location"`
}
type APIGuildScheduledEventEntityMetadata struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
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
type Properties struct {
	ID                 ID                 `json:"id"`
	GuildID            GuildID            `json:"guild_id"`
	ChannelID          ChannelID          `json:"channel_id"`
	CreatorID          CreatorID          `json:"creator_id"`
	Name               Name               `json:"name"`
	Description        Description        `json:"description"`
	ScheduledStartTime ScheduledStartTime `json:"scheduled_start_time"`
	ScheduledEndTime   ScheduledEndTime   `json:"scheduled_end_time"`
	PrivacyLevel       PrivacyLevel       `json:"privacy_level"`
	Status             Status             `json:"status"`
	EntityType         EntityType         `json:"entity_type"`
	EntityID           EntityID           `json:"entity_id"`
	EntityMetadata     EntityMetadata     `json:"entity_metadata"`
	Creator            Creator            `json:"creator"`
	UserCount          UserCount          `json:"user_count"`
	Image              Image              `json:"image"`
}
type APIVoiceGuildScheduledEvent struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type APIExternalGuildScheduledEvent struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type RESTPostAPIGuildScheduledEventJSONBody struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type EntityType struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type EntityMetadata struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ChannelID          ChannelID          `json:"channel_id"`
	Name               Name               `json:"name"`
	PrivacyLevel       PrivacyLevel       `json:"privacy_level"`
	ScheduledStartTime ScheduledStartTime `json:"scheduled_start_time"`
	ScheduledEndTime   ScheduledEndTime   `json:"scheduled_end_time"`
	Description        Description        `json:"description"`
	EntityType         EntityType         `json:"entity_type"`
	EntityMetadata     EntityMetadata     `json:"entity_metadata"`
	Image              Image              `json:"image"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure14365639161057184714365639161003184814365639168411849143656391604111 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type GuildScheduledEventEntityType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type RESTPostAPIGuildScheduledEventResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	WithUserCount WithUserCount `json:"with_user_count"`
}
type RESTGetAPIGuildScheduledEventQuery struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTGetAPIGuildScheduledEventResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	Status             Status             `json:"status"`
	EntityMetadata     EntityMetadata     `json:"entity_metadata"`
	Description        Description        `json:"description"`
	ChannelID          ChannelID          `json:"channel_id"`
	Name               Name               `json:"name"`
	PrivacyLevel       PrivacyLevel       `json:"privacy_level"`
	ScheduledStartTime ScheduledStartTime `json:"scheduled_start_time"`
	ScheduledEndTime   ScheduledEndTime   `json:"scheduled_end_time"`
	EntityType         EntityType         `json:"entity_type"`
	Image              Image              `json:"image"`
}
type RESTPatchAPIGuildScheduledEventJSONBody struct {
	Type                 string     `json:"type"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Properties           Properties `json:"properties"`
	Description          string     `json:"description"`
}
type StrictPartialRESTPostAPIGuildScheduledEventJSONBody struct {
	Ref string `json:"$ref"`
}
type Properties struct {
	ChannelID          ChannelID          `json:"channel_id"`
	Name               Name               `json:"name"`
	PrivacyLevel       PrivacyLevel       `json:"privacy_level"`
	ScheduledStartTime ScheduledStartTime `json:"scheduled_start_time"`
	ScheduledEndTime   ScheduledEndTime   `json:"scheduled_end_time"`
	Description        Description        `json:"description"`
	EntityType         EntityType         `json:"entity_type"`
	EntityMetadata     EntityMetadata     `json:"entity_metadata"`
	Image              Image              `json:"image"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceAlias1589253880728677297015892538800215641DefAlias143656391684118491436563916041111035977792 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type EntityMetadata struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type Properties struct {
	Status         Status         `json:"status"`
	EntityMetadata EntityMetadata `json:"entity_metadata"`
	Description    Description    `json:"description"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure143656391627513052143656391626963053143656391626403053143656391624773054143656391604111 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTPatchAPIGuildScheduledEventResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Limit struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     int    `json:"default"`
}
type WithMember struct {
	Type        string `json:"type"`
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
type Properties struct {
	Limit      Limit      `json:"limit"`
	WithMember WithMember `json:"with_member"`
	Before     Before     `json:"before"`
	After      After      `json:"after"`
}
type RESTGetAPIGuildScheduledEventUsersQuery struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTGetAPIGuildScheduledEventUsersResult struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type GuildScheduledEventID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type User struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Member struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	GuildScheduledEventID GuildScheduledEventID `json:"guild_scheduled_event_id"`
	User                  User                  `json:"user"`
	Member                Member                `json:"member"`
}
type APIGuildScheduledEventUser struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Nick struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
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
