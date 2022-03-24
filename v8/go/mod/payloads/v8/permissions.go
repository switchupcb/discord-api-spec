package v8

type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Color struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Hoist struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Icon struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type UnicodeEmoji struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Position struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Permissions struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Managed struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Mentionable struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Tags struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID           ID           `json:"id"`
	Name         Name         `json:"name"`
	Color        Color        `json:"color"`
	Hoist        Hoist        `json:"hoist"`
	Icon         Icon         `json:"icon"`
	UnicodeEmoji UnicodeEmoji `json:"unicode_emoji"`
	Position     Position     `json:"position"`
	Permissions  Permissions  `json:"permissions"`
	Managed      Managed      `json:"managed"`
	Mentionable  Mentionable  `json:"mentionable"`
	Tags         Tags         `json:"tags"`
}
type APIRole struct {
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
type BotID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PremiumSubscriber struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type IntegrationID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	BotID             BotID             `json:"bot_id"`
	PremiumSubscriber PremiumSubscriber `json:"premium_subscriber"`
	IntegrationID     IntegrationID     `json:"integration_id"`
}
type APIRoleTags struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
