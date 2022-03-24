package v9

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
}
type Snowflake struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type StageInstancePrivacyLevel struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type RESTPostAPIStageInstanceResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
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
	Deprecated  string `json:"deprecated"`
}
type GuildScheduledEventID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID                    ID                    `json:"id"`
	GuildID               GuildID               `json:"guild_id"`
	ChannelID             ChannelID             `json:"channel_id"`
	Topic                 Topic                 `json:"topic"`
	PrivacyLevel          PrivacyLevel          `json:"privacy_level"`
	DiscoverableDisabled  DiscoverableDisabled  `json:"discoverable_disabled"`
	GuildScheduledEventID GuildScheduledEventID `json:"guild_scheduled_event_id"`
}
type APIStageInstance struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTGetAPIStageInstanceResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type RESTPatchAPIStageInstanceJSONBody struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	Topic        Topic        `json:"topic"`
	PrivacyLevel PrivacyLevel `json:"privacy_level"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure6379804731168135163798047311141352637980473971135363798047301662 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTPatchAPIStageInstanceResult struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}