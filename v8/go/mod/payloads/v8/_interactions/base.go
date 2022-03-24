package _interactions
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

