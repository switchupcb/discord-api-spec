type AutoGenerated struct {
	Definitions []struct {
		APIGuildApplicationCommandPermissions struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				ApplicationID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"application_id"`
				GuildID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"guild_id"`
				Permissions struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"permissions"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIGuildApplicationCommandPermissions"`
		Snowflake struct {
			Type string `json:"type"`
			Description string `json:"description"`
		} `json:"Snowflake"`
		APIApplicationCommandPermission struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				Type struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"type"`
				Permission struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"permission"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIApplicationCommandPermission"`
		ApplicationCommandPermissionType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"ApplicationCommandPermissionType"`
	} `json:"definitions"`
}