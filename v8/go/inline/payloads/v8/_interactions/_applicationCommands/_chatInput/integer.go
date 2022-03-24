type AutoGenerated struct {
	Definitions struct {
		APIApplicationCommandIntegerOption struct {
			Ref string `json:"$ref"`
			Deprecated string `json:"deprecated"`
		} `json:"APIApplicationCommandIntegerOption"`
		APIApplicationCommandOptionWithAutocompleteOrChoicesWrapperInterface352328325254738352328325013671295383896APIApplicationCommandOptionChoiceNumber struct {
			AnyOf []struct {
				Type string `json:"type"`
				AdditionalProperties bool `json:"additionalProperties"`
				Properties struct {
					Autocomplete struct {
						Type string `json:"type"`
						Const bool `json:"const"`
					} `json:"autocomplete"`
					Type struct {
						Type string `json:"type"`
						Const int `json:"const"`
					} `json:"type"`
					Name struct {
						Type string `json:"type"`
					} `json:"name"`
					Description struct {
						Type string `json:"type"`
					} `json:"description"`
					Required struct {
						Type string `json:"type"`
					} `json:"required"`
					MinValue struct {
						Type string `json:"type"`
						Description string `json:"description"`
					} `json:"min_value"`
					MaxValue struct {
						Type string `json:"type"`
						Description string `json:"description"`
					} `json:"max_value"`
				} `json:"properties,omitempty"`
				Required []string `json:"required"`
				Properties0 struct {
					Autocomplete struct {
						Type string `json:"type"`
						Const bool `json:"const"`
					} `json:"autocomplete"`
					Choices struct {
						Type string `json:"type"`
						Items struct {
							Ref string `json:"$ref"`
						} `json:"items"`
					} `json:"choices"`
					Type struct {
						Type string `json:"type"`
						Const int `json:"const"`
					} `json:"type"`
					Name struct {
						Type string `json:"type"`
					} `json:"name"`
					Description struct {
						Type string `json:"type"`
					} `json:"description"`
					Required struct {
						Type string `json:"type"`
					} `json:"required"`
					MinValue struct {
						Type string `json:"type"`
						Description string `json:"description"`
					} `json:"min_value"`
					MaxValue struct {
						Type string `json:"type"`
						Description string `json:"description"`
					} `json:"max_value"`
				} `json:"properties,omitempty"`
			} `json:"anyOf"`
			Deprecated string `json:"deprecated"`
		} `json:"APIApplicationCommandOptionWithAutocompleteOrChoicesWrapper<interface-352328325-254-738-352328325-0-13671295383896,APIApplicationCommandOptionChoice<number>>"`
		APIApplicationCommandOptionBaseEnum20937207223453542093720722042720937207220818 struct {
			Type string `json:"type"`
			Properties struct {
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
				} `json:"type"`
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				Description struct {
					Type string `json:"type"`
				} `json:"description"`
				Required struct {
					Type string `json:"type"`
				} `json:"required"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Deprecated string `json:"deprecated"`
		} `json:"APIApplicationCommandOptionBase<enum-2093720722-345-354-2093720722-0-427-2093720722-0-818>"`
		APIApplicationCommandOptionChoiceNumber struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				Value struct {
					Type string `json:"type"`
				} `json:"value"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIApplicationCommandOptionChoice<number>"`
		APIApplicationCommandInteractionDataIntegerOption struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
				} `json:"type"`
				Value struct {
					Type string `json:"type"`
				} `json:"value"`
				Focused struct {
					Type string `json:"type"`
				} `json:"focused"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Required []string `json:"required"`
			Deprecated string `json:"deprecated"`
		} `json:"APIApplicationCommandInteractionDataIntegerOption"`
		APIInteractionDataOptionBaseEnum20937207223453542093720722042720937207220818Number struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				Type struct {
					Type string `json:"type"`
					Const int `json:"const"`
				} `json:"type"`
				Value struct {
					Type string `json:"type"`
				} `json:"value"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Deprecated string `json:"deprecated"`
		} `json:"APIInteractionDataOptionBase<enum-2093720722-345-354-2093720722-0-427-2093720722-0-818,number>"`
	} `json:"definitions"`
}