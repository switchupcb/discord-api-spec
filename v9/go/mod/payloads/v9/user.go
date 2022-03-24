package v9

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
type ID struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Name struct {
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
type User struct {
	Ref         string `json:"$ref"`
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
type Icon struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
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
