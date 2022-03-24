package v8

type ChannelID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Topic struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type PrivacyLevel struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Default     string `json:"default"`
}
type Properties struct {
	ChannelID    ChannelID    `json:"channel_id"`
	Topic        Topic        `json:"topic"`
	PrivacyLevel PrivacyLevel `json:"privacy_level"`
}
type RESTPostAPIStageInstanceJSONBody struct {
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
type StageInstancePrivacyLevel struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type RESTPostAPIStageInstanceResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type GuildID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PrivacyLevel struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type DiscoverableDisabled struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID                   ID                   `json:"id"`
	GuildID              GuildID              `json:"guild_id"`
	ChannelID            ChannelID            `json:"channel_id"`
	Topic                Topic                `json:"topic"`
	PrivacyLevel         PrivacyLevel         `json:"privacy_level"`
	DiscoverableDisabled DiscoverableDisabled `json:"discoverable_disabled"`
}
type APIStageInstance struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
	Deprecated           string     `json:"deprecated"`
}
type RESTGetAPIStageInstanceResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type RESTPatchAPIStageInstanceJSONBody struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Properties struct {
	Topic        Topic        `json:"topic"`
	PrivacyLevel PrivacyLevel `json:"privacy_level"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure163505282616441827163505282615901828163505282613281829163505282602376 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTPatchAPIStageInstanceResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
