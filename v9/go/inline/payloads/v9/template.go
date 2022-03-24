type AutoGenerated struct {
	Definitions struct {
		APITemplate struct {
			Type string `json:"type"`
			Properties struct {
				Code struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"code"`
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Description struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"description"`
				UsageCount struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"usage_count"`
				CreatorID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"creator_id"`
				Creator struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"creator"`
				CreatedAt struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"created_at"`
				UpdatedAt struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"updated_at"`
				SourceGuildID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"source_guild_id"`
				SerializedSourceGuild struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"serialized_source_guild"`
				IsDirty struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"is_dirty"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APITemplate"`
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
		APITemplateSerializedSourceGuild struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Region struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"region"`
				VerificationLevel struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"verification_level"`
				DefaultMessageNotifications struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"default_message_notifications"`
				ExplicitContentFilter struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"explicit_content_filter"`
				Roles struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"roles"`
				Channels struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"channels"`
				AfkChannelID struct {
					AnyOf []struct {
						Type string `json:"type,omitempty"`
						Ref string `json:"$ref,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
				} `json:"afk_channel_id"`
				AfkTimeout struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"afk_timeout"`
				SystemChannelID struct {
					AnyOf []struct {
						Type string `json:"type,omitempty"`
						Ref string `json:"$ref,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
				} `json:"system_channel_id"`
				SystemChannelFlags struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"system_channel_flags"`
				PremiumProgressBarEnabled struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"premium_progress_bar_enabled"`
				Description struct {
					Type []string `json:"type"`
				} `json:"description"`
				PreferredLocale struct {
					Type string `json:"type"`
				} `json:"preferred_locale"`
				IconHash struct {
					Type []string `json:"type"`
				} `json:"icon_hash"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APITemplateSerializedSourceGuild"`
		GuildVerificationLevel struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"GuildVerificationLevel"`
		GuildDefaultMessageNotifications struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"GuildDefaultMessageNotifications"`
		GuildExplicitContentFilter struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"GuildExplicitContentFilter"`
		APIGuildCreateRole struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type []string `json:"type"`
					Description string `json:"description"`
					Default string `json:"default"`
				} `json:"name"`
				Permissions struct {
					AnyOf []struct {
						Type string `json:"type"`
						Description string `json:"description,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
					Default string `json:"default"`
				} `json:"permissions"`
				Color struct {
					Type []string `json:"type"`
					Description string `json:"description"`
					Default int `json:"default"`
				} `json:"color"`
				Hoist struct {
					Type []string `json:"type"`
					Description string `json:"description"`
					Default bool `json:"default"`
				} `json:"hoist"`
				Icon struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"icon"`
				UnicodeEmoji struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"unicode_emoji"`
				Mentionable struct {
					Type []string `json:"type"`
					Description string `json:"description"`
					Default bool `json:"default"`
				} `json:"mentionable"`
				ID struct {
					Type []string `json:"type"`
				} `json:"id"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIGuildCreateRole"`
		RESTPostAPIGuildRoleJSONBody struct {
			Ref string `json:"$ref"`
			Description string `json:"description"`
		} `json:"RESTPostAPIGuildRoleJSONBody"`
		AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure536029191157321651253602919115678165135360291911555316514536029191023031 struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type []string `json:"type"`
					Description string `json:"description"`
					Default string `json:"default"`
				} `json:"name"`
				Permissions struct {
					AnyOf []struct {
						Type string `json:"type"`
						Description string `json:"description,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
					Default string `json:"default"`
				} `json:"permissions"`
				Color struct {
					Type []string `json:"type"`
					Description string `json:"description"`
					Default int `json:"default"`
				} `json:"color"`
				Hoist struct {
					Type []string `json:"type"`
					Description string `json:"description"`
					Default bool `json:"default"`
				} `json:"hoist"`
				Icon struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"icon"`
				UnicodeEmoji struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"unicode_emoji"`
				Mentionable struct {
					Type []string `json:"type"`
					Description string `json:"description"`
					Default bool `json:"default"`
				} `json:"mentionable"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"AddUndefinedToPossiblyUndefinedPropertiesOfInterface<structure-536029191-15732-16512-536029191-15678-16513-536029191-15553-16514-536029191-0-23031>"`
		APIGuildCreatePartialChannel struct {
			Type string `json:"type"`
			AdditionalProperties bool `json:"additionalProperties"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				ID struct {
					Type []string `json:"type"`
				} `json:"id"`
				ParentID struct {
					Type []string `json:"type"`
				} `json:"parent_id"`
				PermissionOverwrites struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
				} `json:"permission_overwrites"`
			} `json:"properties"`
			Required []string `json:"required"`
		} `json:"APIGuildCreatePartialChannel"`
		StrictPartialAlias1589253880731877332815892538800215641TypeTopicNsfwBitrateUserLimitRateLimitPerUser struct {
			Ref string `json:"$ref"`
		} `json:"StrictPartial<alias-1589253880-73187-73328-1589253880-0-215641<,("type"|"topic"|"nsfw"|"bitrate"|"user_limit"|"rate_limit_per_user")>>"`
		AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceAlias1589253880728677297015892538800215641Alias1589253880731877332815892538800215641TypeTopicNsfwBitrateUserLimitRateLimitPerUser struct {
			Type string `json:"type"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"AddUndefinedToPossiblyUndefinedPropertiesOfInterface<alias-1589253880-72867-72970-1589253880-0-215641<alias-1589253880-73187-73328-1589253880-0-215641<,("type"|"topic"|"nsfw"|"bitrate"|"user_limit"|"rate_limit_per_user")>>>"`
		AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure5360291911242137753602919111871378536029191102913785360291919861379536029191023031 struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				ID struct {
					Type []string `json:"type"`
				} `json:"id"`
				ParentID struct {
					Type []string `json:"type"`
				} `json:"parent_id"`
				PermissionOverwrites struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
				} `json:"permission_overwrites"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"AddUndefinedToPossiblyUndefinedPropertiesOfInterface<structure-536029191-1242-1377-536029191-1187-1378-536029191-1029-1378-536029191-986-1379-536029191-0-23031>"`
		APIGuildCreateOverwrite struct {
			Type string `json:"type"`
			Properties struct {
				Allow struct {
					AnyOf []struct {
						Type string `json:"type"`
						Description string `json:"description,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
					Default string `json:"default"`
				} `json:"allow"`
				Deny struct {
					AnyOf []struct {
						Type string `json:"type"`
						Description string `json:"description,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
					Default string `json:"default"`
				} `json:"deny"`
				Type struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"type"`
				ID struct {
					Type []string `json:"type"`
				} `json:"id"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIGuildCreateOverwrite"`
		RESTPutAPIChannelPermissionJSONBody struct {
			Type string `json:"type"`
			Properties struct {
				Allow struct {
					AnyOf []struct {
						Type string `json:"type"`
						Description string `json:"description,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
					Default string `json:"default"`
				} `json:"allow"`
				Deny struct {
					AnyOf []struct {
						Type string `json:"type"`
						Description string `json:"description,omitempty"`
					} `json:"anyOf"`
					Description string `json:"description"`
					Default string `json:"default"`
				} `json:"deny"`
				Type struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"RESTPutAPIChannelPermissionJSONBody"`
		OverwriteType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
		} `json:"OverwriteType"`
		GuildSystemChannelFlags struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"GuildSystemChannelFlags"`
	} `json:"definitions"`
}