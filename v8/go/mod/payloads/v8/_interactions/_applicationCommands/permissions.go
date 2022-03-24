package _applicationCommands
type ID struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type ApplicationID struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type GuildID struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type Permissions struct {
	Type string `json:"type"`
	Items Items `json:"items"`
	Description string `json:"description"`
}
type Properties struct {
	ID ID `json:"id"`
	ApplicationID ApplicationID `json:"application_id"`
	GuildID GuildID `json:"guild_id"`
	Permissions Permissions `json:"permissions"`
}
type APIGuildApplicationCommandPermissions struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type Snowflake struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Type struct {
	Ref string `json:"$ref"`
	Description string `json:"description"`
}
type Permission struct {
	Type string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID ID `json:"id"`
	Type Type `json:"type"`
	Permission Permission `json:"permission"`
}
type APIApplicationCommandPermission struct {
	Type string `json:"type"`
	Properties Properties `json:"properties"`
	Required []string `json:"required"`
	AdditionalProperties bool `json:"additionalProperties"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type ApplicationCommandPermissionType struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}

