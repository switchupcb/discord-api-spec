package v9

type GuildID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type AnyOf struct {
	Ref  string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
}
type ChannelID struct {
	AnyOf       []AnyOf `json:"anyOf"`
	Description string  `json:"description"`
}
type UserID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Member struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type SessionID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Deaf struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Mute struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SelfDeaf struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SelfMute struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SelfStream struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SelfVideo struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Suppress struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type RequestToSpeakTimestamp struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Properties struct {
	GuildID                 GuildID                 `json:"guild_id"`
	ChannelID               ChannelID               `json:"channel_id"`
	UserID                  UserID                  `json:"user_id"`
	Member                  Member                  `json:"member"`
	SessionID               SessionID               `json:"session_id"`
	Deaf                    Deaf                    `json:"deaf"`
	Mute                    Mute                    `json:"mute"`
	SelfDeaf                SelfDeaf                `json:"self_deaf"`
	SelfMute                SelfMute                `json:"self_mute"`
	SelfStream              SelfStream              `json:"self_stream"`
	SelfVideo               SelfVideo               `json:"self_video"`
	Suppress                Suppress                `json:"suppress"`
	RequestToSpeakTimestamp RequestToSpeakTimestamp `json:"request_to_speak_timestamp"`
}
type GatewayVoiceState struct {
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
type User struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Nick struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Avatar struct {
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
type ID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Optimal struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Deprecated struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Custom struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID         ID         `json:"id"`
	Name       Name       `json:"name"`
	Optimal    Optimal    `json:"optimal"`
	Deprecated Deprecated `json:"deprecated"`
	Custom     Custom     `json:"custom"`
}
type APIVoiceRegion struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
