type AutoGenerated struct {
	Definitions struct {
		APITeam struct {
			Type string `json:"type"`
			Properties struct {
				Icon struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"icon"`
				ID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				Members struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"members"`
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				OwnerUserID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"owner_user_id"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APITeam"`
		Snowflake struct {
			Type string `json:"type"`
			Description string `json:"description"`
		} `json:"Snowflake"`
		APITeamMember struct {
			Type string `json:"type"`
			Properties struct {
				MembershipState struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"membership_state"`
				Permissions struct {
					Type string `json:"type"`
					Items struct {
						Type string `json:"type"`
						Const string `json:"const"`
					} `json:"items"`
					MinItems int `json:"minItems"`
					MaxItems int `json:"maxItems"`
					Description string `json:"description"`
				} `json:"permissions"`
				TeamID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"team_id"`
				User struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"user"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APITeamMember"`
		TeamMemberMembershipState struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"TeamMemberMembershipState"`
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