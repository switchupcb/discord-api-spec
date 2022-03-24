type AutoGenerated struct {
	Definitions struct {
		GetAPIVoiceRegionsResult struct {
			Type string `json:"type"`
			Items struct {
				Ref string `json:"$ref"`
			} `json:"items"`
			Description string `json:"description"`
		} `json:"GetAPIVoiceRegionsResult"`
		APIVoiceRegion struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"id"`
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Optimal struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"optimal"`
				Deprecated struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"deprecated"`
				Custom struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"custom"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"APIVoiceRegion"`
	} `json:"definitions"`
}