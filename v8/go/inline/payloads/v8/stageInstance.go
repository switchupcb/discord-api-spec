type AutoGenerated struct {
	Definitions struct {
		APIStageInstance struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				GuildID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"guild_id"`
				ChannelID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"channel_id"`
				Topic struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"topic"`
				PrivacyLevel struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"privacy_level"`
				DiscoverableDisabled struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"discoverable_disabled"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIStageInstance"`
		Snowflake struct {
			Type string `json:"type"`
			Description string `json:"description"`
		} `json:"Snowflake"`
		StageInstancePrivacyLevel struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"StageInstancePrivacyLevel"`
		APIInviteStageInstance struct {
			Type string `json:"type"`
			Properties struct {
				Topic struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"topic"`
				ParticipantCount struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"participant_count"`
				SpeakerCount struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"speaker_count"`
				Members struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"members"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIInviteStageInstance"`
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
	} `json:"definitions"`
}