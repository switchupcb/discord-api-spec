type AutoGenerated struct {
	Definitions struct {
		RESTGetAPIGuildEmojisResult struct {
			Type string `json:"type"`
			Items struct {
				Ref string `json:"$ref"`
			} `json:"items"`
			Description string `json:"description"`
		} `json:"RESTGetAPIGuildEmojisResult"`
		APIEmoji struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					AnyOf []struct {
						Ref string `json:"$ref,omitempty"`
						Type string `json:"type,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
				} `json:"id"`
				Name struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Animated struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"animated"`
				Roles struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
						Description string `json:"description"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"roles"`
				User struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"user"`
				RequireColons struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"require_colons"`
				Managed struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"managed"`
				Available struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"available"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Required []string `json:"required"`
			Description string `json:"description"`
		} `json:"APIEmoji"`
		APIPartialEmoji struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					AnyOf []struct {
						Ref string `json:"$ref,omitempty"`
						Type string `json:"type,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
				} `json:"id"`
				Name struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Animated struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"animated"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIPartialEmoji"`
		Snowflake struct {
			Type string `json:"type"`
			Description string `json:"description"`
		} `json:"Snowflake"`
		APIUser struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				Username struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"username"`
				Discriminator struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"discriminator"`
				Avatar struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"avatar"`
				Bot struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"bot"`
				System struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"system"`
				MfaEnabled struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"mfa_enabled"`
				Banner struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"banner"`
				AccentColor struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"accent_color"`
				Locale struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"locale"`
				Verified struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"verified"`
				Email struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"email"`
				Flags struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"flags"`
				PremiumType struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"premium_type"`
				PublicFlags struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"public_flags"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIUser"`
		UserFlags struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"UserFlags"`
		UserPremiumType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"UserPremiumType"`
		RESTGetAPIGuildEmojiResult struct {
			Ref string `json:"$ref"`
			Description string `json:"description"`
		} `json:"RESTGetAPIGuildEmojiResult"`
		RESTPostAPIGuildEmojiJSONBody struct {
			Ref string `json:"$ref"`
			Description string `json:"description"`
		} `json:"RESTPostAPIGuildEmojiJSONBody"`
		AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure193329939866391519332993986099161933299398470917193329939801648 struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Image struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"image"`
				Roles struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"roles"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"AddUndefinedToPossiblyUndefinedPropertiesOfInterface<structure-1933299398-663-915-1933299398-609-916-1933299398-470-917-1933299398-0-1648>"`
		RESTPostAPIGuildEmojiResult struct {
			Ref string `json:"$ref"`
			Description string `json:"description"`
		} `json:"RESTPostAPIGuildEmojiResult"`
		RESTPatchAPIGuildEmojiJSONBody struct {
			Ref string `json:"$ref"`
			Description string `json:"description"`
		} `json:"RESTPatchAPIGuildEmojiJSONBody"`
		AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure193329939812341375193329939811801376193329939810521377193329939801648 struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Roles struct {
					AnyOf []struct {
						Type string `json:"type"`
						Items struct {
							Ref string `json:"$ref"`
						} `json:"items,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
				} `json:"roles"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"AddUndefinedToPossiblyUndefinedPropertiesOfInterface<structure-1933299398-1234-1375-1933299398-1180-1376-1933299398-1052-1377-1933299398-0-1648>"`
		RESTPatchAPIGuildEmojiResult struct {
			Ref string `json:"$ref"`
			Description string `json:"description"`
		} `json:"RESTPatchAPIGuildEmojiResult"`
	} `json:"definitions"`
}