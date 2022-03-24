type AutoGenerated struct {
	Definitions struct {
		PartialAPIMessageInteractionGuildMember struct {
			Type string `json:"type"`
			Properties struct {
				Roles struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"roles"`
				PremiumSince struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"premium_since"`
				Pending struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"pending"`
				Nick struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"nick"`
				Mute struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"mute"`
				JoinedAt struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"joined_at"`
				Deaf struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"deaf"`
				CommunicationDisabledUntil struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"communication_disabled_until"`
				Avatar struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"avatar"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Deprecated string `json:"deprecated"`
		} `json:"PartialAPIMessageInteractionGuildMember"`
		Snowflake struct {
			Type string `json:"type"`
			Description string `json:"description"`
		} `json:"Snowflake"`
		APIMessageInteraction struct {
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
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				User struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"user"`
				Member struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"member"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIMessageInteraction"`
		InteractionType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"InteractionType"`
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
			Deprecated string `json:"deprecated"`
		} `json:"APIUser"`
		UserFlags struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"UserFlags"`
		UserPremiumType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"UserPremiumType"`
		APIInteractionGuildMember struct {
			Type string `json:"type"`
			Properties struct {
				User struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"user"`
				Nick struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"nick"`
				Avatar struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"avatar"`
				Roles struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"roles"`
				JoinedAt struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"joined_at"`
				PremiumSince struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"premium_since"`
				Deaf struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"deaf"`
				Mute struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"mute"`
				Pending struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"pending"`
				CommunicationDisabledUntil struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"communication_disabled_until"`
				Permissions struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"permissions"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIInteractionGuildMember"`
		APIGuildMember struct {
			Type string `json:"type"`
			Properties struct {
				User struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"user"`
				Nick struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"nick"`
				Avatar struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"avatar"`
				Roles struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"roles"`
				JoinedAt struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"joined_at"`
				PremiumSince struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"premium_since"`
				Deaf struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"deaf"`
				Mute struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"mute"`
				Pending struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"pending"`
				CommunicationDisabledUntil struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"communication_disabled_until"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIGuildMember"`
	} `json:"definitions"`
}