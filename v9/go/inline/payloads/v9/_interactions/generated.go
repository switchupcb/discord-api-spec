type AutoGenerated struct {
	Definitions []struct {
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
		} `json:"PartialAPIMessageInteractionGuildMember,omitempty"`
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
		} `json:"APIMessageInteraction,omitempty"`
		InteractionType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
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
		} `json:"APIUser,omitempty"`
		UserFlags struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"UserFlags,omitempty"`
		UserPremiumType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"UserPremiumType,omitempty"`
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
		} `json:"APIInteractionGuildMember,omitempty"`
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
		} `json:"APIGuildMember,omitempty"`
		APIInteractionResponse struct {
			AnyOf []struct {
				Ref string `json:"$ref"`
			} `json:"anyOf"`
			Description string `json:"description"`
		} `json:"APIInteractionResponse,omitempty"`
		APIInteractionResponsePong struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIInteractionResponsePong,omitempty"`
		APIInteractionResponseChannelMessageWithSource struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Data struct {
					Ref string `json:"$ref"`
				} `json:"data"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIInteractionResponseChannelMessageWithSource,omitempty"`
		APIInteractionResponseCallbackData struct {
			Type string `json:"type"`
			AdditionalProperties bool `json:"additionalProperties"`
			Properties struct {
				Flags struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"flags"`
				Content struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"content"`
				Tts struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"tts"`
				Embeds struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"embeds"`
				AllowedMentions struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"allowed_mentions"`
				Components struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"components"`
				Attachments struct {
					Type string `json:"type"`
					Items struct {
						Type string `json:"type"`
						AdditionalProperties bool `json:"additionalProperties"`
						Properties struct {
							Filename struct {
								Type string `json:"type"`
								Description string `json:"description"`
							} `json:"filename"`
							ID struct {
								Ref string `json:"$ref"`
								Description string `json:"description"`
							} `json:"id"`
							Description struct {
								Type string `json:"type"`
								Description string `json:"description"`
							} `json:"description"`
						} `json:"properties"`
						Required []string `json:"required"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"attachments"`
			} `json:"properties"`
			Description string `json:"description"`
		} `json:"APIInteractionResponseCallbackData,omitempty"`
		APIEmbed struct {
			Type string `json:"type"`
			Properties struct {
				Title struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"title"`
				Type struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
					Deprecated string `json:"deprecated"`
				} `json:"type"`
				Description struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"description"`
				URL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"url"`
				Timestamp struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"timestamp"`
				Color struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"color"`
				Footer struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"footer"`
				Image struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"image"`
				Thumbnail struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnail"`
				Video struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"video"`
				Provider struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"provider"`
				Author struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"author"`
				Fields struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"fields"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIEmbed,omitempty"`
		EmbedType struct {
			Type string `json:"type"`
			Enum []string `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"EmbedType,omitempty"`
		APIEmbedFooter struct {
			Type string `json:"type"`
			Properties struct {
				Text struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"text"`
				IconURL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"icon_url"`
				ProxyIconURL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"proxy_icon_url"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIEmbedFooter,omitempty"`
		APIEmbedImage struct {
			Type string `json:"type"`
			Properties struct {
				URL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"url"`
				ProxyURL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"proxy_url"`
				Height struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"height"`
				Width struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"width"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIEmbedImage,omitempty"`
		APIEmbedThumbnail struct {
			Type string `json:"type"`
			Properties struct {
				URL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"url"`
				ProxyURL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"proxy_url"`
				Height struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"height"`
				Width struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"width"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIEmbedThumbnail,omitempty"`
		APIEmbedVideo struct {
			Type string `json:"type"`
			Properties struct {
				URL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"url"`
				Height struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"height"`
				Width struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"width"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIEmbedVideo,omitempty"`
		APIEmbedProvider struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				URL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"url"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIEmbedProvider,omitempty"`
		APIEmbedAuthor struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				URL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"url"`
				IconURL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"icon_url"`
				ProxyIconURL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"proxy_icon_url"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIEmbedAuthor,omitempty"`
		APIEmbedField struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Value struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"value"`
				Inline struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"inline"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIEmbedField,omitempty"`
		APIAllowedMentions struct {
			Type string `json:"type"`
			Properties struct {
				Parse struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"parse"`
				Roles struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"roles"`
				Users struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"users"`
				RepliedUser struct {
					Type string `json:"type"`
					Description string `json:"description"`
					Default bool `json:"default"`
				} `json:"replied_user"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIAllowedMentions,omitempty"`
		AllowedMentionsTypes struct {
			Type string `json:"type"`
			Enum []string `json:"enum"`
			Description string `json:"description"`
		} `json:"AllowedMentionsTypes,omitempty"`
		APIActionRowComponentAPIMessageActionRowComponent struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Components struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"components"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIActionRowComponent<APIMessageActionRowComponent>,omitempty"`
		APIBaseComponentEnum12011286772870528755120112867728581288941201128677033636105957764 struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIBaseComponent<enum-1201128677-28705-28755-1201128677-28581-28894-1201128677-0-33636105957764>,omitempty"`
		APIMessageActionRowComponent struct {
			AnyOf []struct {
				Ref string `json:"$ref"`
			} `json:"anyOf"`
			Description string `json:"description"`
		} `json:"APIMessageActionRowComponent,omitempty"`
		APIButtonComponent struct {
			AnyOf []struct {
				Ref string `json:"$ref"`
			} `json:"anyOf"`
		} `json:"APIButtonComponent,omitempty"`
		APIButtonComponentWithCustomID struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Label struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"label"`
				Style struct {
					AnyOf []struct {
						Type string `json:"type"`
						Const int `json:"const"`
					} `json:"anyOf"`
					Description string `json:"description"`
				} `json:"style"`
				Emoji struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"emoji"`
				Disabled struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"disabled"`
				CustomID struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"custom_id"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIButtonComponentWithCustomId,omitempty"`
		APIButtonComponentBaseEnum12011286773053930552120112867730405305931201128677033636Enum12011286773055330564120112867730405305931201128677033636Enum12011286773056530574120112867730405305931201128677033636Enum12011286773057530583120112867730405305931201128677033636 struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Label struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"label"`
				Style struct {
					AnyOf []struct {
						Type string `json:"type"`
						Const int `json:"const"`
					} `json:"anyOf"`
					Description string `json:"description"`
				} `json:"style"`
				Emoji struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"emoji"`
				Disabled struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"disabled"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIButtonComponentBase<(enum-1201128677-30539-30552-1201128677-30405-30593-1201128677-0-33636|enum-1201128677-30553-30564-1201128677-30405-30593-1201128677-0-33636|enum-1201128677-30565-30574-1201128677-30405-30593-1201128677-0-33636|enum-1201128677-30575-30583-1201128677-30405-30593-1201128677-0-33636)>,omitempty"`
		APIBaseComponentEnum120112867728756287951201128677285812889412011286770336361131375322 struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIBaseComponent<enum-1201128677-28756-28795-1201128677-28581-28894-1201128677-0-336361131375322>,omitempty"`
		APIMessageComponentEmoji struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Animated struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"animated"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIMessageComponentEmoji,omitempty"`
		APIButtonComponentWithURL struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Label struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"label"`
				Style struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"style"`
				Emoji struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"emoji"`
				Disabled struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"disabled"`
				URL struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"url"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIButtonComponentWithURL,omitempty"`
		APIButtonComponentBaseEnum12011286773058430590120112867730405305931201128677033636 struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Label struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"label"`
				Style struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"style"`
				Emoji struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"emoji"`
				Disabled struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"disabled"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIButtonComponentBase<enum-1201128677-30584-30590-1201128677-30405-30593-1201128677-0-33636>,omitempty"`
		APISelectMenuComponent struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				CustomID struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"custom_id"`
				Options struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"options"`
				Placeholder struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"placeholder"`
				MinValues struct {
					Type string `json:"type"`
					Description string `json:"description"`
					Default int `json:"default"`
				} `json:"min_values"`
				MaxValues struct {
					Type string `json:"type"`
					Description string `json:"description"`
					Default int `json:"default"`
				} `json:"max_values"`
				Disabled struct {
					Type string `json:"type"`
					Description string `json:"description"`
					Default bool `json:"default"`
				} `json:"disabled"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APISelectMenuComponent,omitempty"`
		APIBaseComponentEnum12011286772879628844120112867728581288941201128677033636663308133 struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIBaseComponent<enum-1201128677-28796-28844-1201128677-28581-28894-1201128677-0-33636663308133>,omitempty"`
		APISelectMenuOption struct {
			Type string `json:"type"`
			Properties struct {
				Label struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"label"`
				Value struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"value"`
				Description struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"description"`
				Emoji struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"emoji"`
				Default struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"default"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APISelectMenuOption,omitempty"`
		MessageFlags struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"MessageFlags,omitempty"`
		APIInteractionResponseDeferredChannelMessageWithSource struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Data struct {
					Type string `json:"type"`
					Properties struct {
						Flags struct {
							Ref string `json:"$ref"`
							Description string `json:"description"`
						} `json:"flags"`
					} `json:"properties"`
					AdditionalProperties bool `json:"additionalProperties"`
				} `json:"data"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIInteractionResponseDeferredChannelMessageWithSource,omitempty"`
		APIInteractionResponseDeferredMessageUpdate struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIInteractionResponseDeferredMessageUpdate,omitempty"`
		APIInteractionResponseUpdateMessage struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Data struct {
					Ref string `json:"$ref"`
				} `json:"data"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIInteractionResponseUpdateMessage,omitempty"`
		APIApplicationCommandAutocompleteResponse struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Data struct {
					Ref string `json:"$ref"`
				} `json:"data"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIApplicationCommandAutocompleteResponse,omitempty"`
		APICommandAutocompleteInteractionResponseCallbackData struct {
			Type string `json:"type"`
			Properties struct {
				Choices struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
				} `json:"choices"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APICommandAutocompleteInteractionResponseCallbackData,omitempty"`
		APIApplicationCommandOptionChoice struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				Value struct {
					Type []string `json:"type"`
				} `json:"value"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIApplicationCommandOptionChoice,omitempty"`
		APIModalInteractionResponse struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Data struct {
					Ref string `json:"$ref"`
				} `json:"data"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
		} `json:"APIModalInteractionResponse,omitempty"`
		APIModalInteractionResponseCallbackData struct {
			Type string `json:"type"`
			Properties struct {
				CustomID struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"custom_id"`
				Title struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"title"`
				Components struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"components"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIModalInteractionResponseCallbackData,omitempty"`
		APIActionRowComponentAPIModalActionRowComponent struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Components struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"components"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIActionRowComponent<APIModalActionRowComponent>,omitempty"`
		APIModalActionRowComponent struct {
			Ref string `json:"$ref"`
		} `json:"APIModalActionRowComponent,omitempty"`
		APITextInputComponent struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
				Style struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"style"`
				CustomID struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"custom_id"`
				Label struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"label"`
				Placeholder struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"placeholder"`
				Value struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"value"`
				MinLength struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"min_length"`
				MaxLength struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"max_length"`
				Required struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"required"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APITextInputComponent,omitempty"`
		APIBaseComponentEnum120112867728845288911201128677285812889412011286770336361182532533 struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
					Description string `json:"description"`
				} `json:"type"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIBaseComponent<enum-1201128677-28845-28891-1201128677-28581-28894-1201128677-0-336361182532533>,omitempty"`
		TextInputStyle struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"TextInputStyle,omitempty"`
		InteractionResponseType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
		} `json:"InteractionResponseType,omitempty"`
	} `json:"definitions"`
}