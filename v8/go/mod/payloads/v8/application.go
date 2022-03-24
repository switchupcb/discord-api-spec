package v8

type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
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
type GuildID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
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
type Flags struct {
	Ref         string `json:"$ref"`
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
type APIApplication struct {
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
type Items struct {
	Ref string `json:"$ref"`
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
	Deprecated           string     `json:"deprecated"`
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
type User struct {
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
	Deprecated           string     `json:"deprecated"`
}
type TeamMemberMembershipState struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type ApplicationFlags struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
