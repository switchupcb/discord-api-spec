type AutoGenerated struct {
	Definitions struct {
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
	} `json:"definitions"`
}