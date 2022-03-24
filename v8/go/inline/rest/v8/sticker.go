type AutoGenerated struct {
	Definitions struct {
		RESTGetAPIStickerResult struct {
			Ref string `json:"$ref"`
			Deprecated string `json:"deprecated"`
		} `json:"RESTGetAPIStickerResult"`
		APISticker struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				PackID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pack_id"`
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Description struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"description"`
				Tags struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"tags"`
				Asset struct {
					Type string `json:"type"`
					Const string `json:"const"`
					Description string `json:"description"`
					Deprecated string `json:"deprecated"`
				} `json:"asset"`
				Type struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"type"`
				FormatType struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"format_type"`
				Available struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"available"`
				GuildID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"guild_id"`
				User struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"user"`
				SortValue struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"sort_value"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APISticker"`
		Snowflake struct {
			Type string `json:"type"`
			Description string `json:"description"`
		} `json:"Snowflake"`
		StickerType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"StickerType"`
		StickerFormatType struct {
			Type string `json:"type"`
			Enum []int `json:"enum"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"StickerFormatType"`
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
		RESTGetNitroStickerPacksResult struct {
			Type string `json:"type"`
			Properties struct {
				StickerPacks struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
				} `json:"sticker_packs"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Deprecated string `json:"deprecated"`
		} `json:"RESTGetNitroStickerPacksResult"`
		APIStickerPack struct {
			Type string `json:"type"`
			Properties struct {
				ID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				Stickers struct {
					Type string `json:"type"`
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Description string `json:"description"`
				} `json:"stickers"`
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				SkuID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"sku_id"`
				CoverStickerID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"cover_sticker_id"`
				Description struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"description"`
				BannerAssetID struct {
					Ref string `json:"$ref"`
					Description string `json:"description"`
				} `json:"banner_asset_id"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
			Deprecated string `json:"deprecated"`
		} `json:"APIStickerPack"`
		RESTGetAPIGuildStickersResult struct {
			Type string `json:"type"`
			Items struct {
				Ref string `json:"$ref"`
			} `json:"items"`
			Deprecated string `json:"deprecated"`
		} `json:"RESTGetAPIGuildStickersResult"`
		RESTGetAPIGuildStickerResult struct {
			Ref string `json:"$ref"`
			Deprecated string `json:"deprecated"`
		} `json:"RESTGetAPIGuildStickerResult"`
		RESTPostAPIGuildStickerFormDataBody struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Description struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"description"`
				Tags struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"tags"`
				File struct {
					Description string `json:"description"`
				} `json:"file"`
			} `json:"properties"`
			Required []string `json:"required"`
			AdditionalProperties bool `json:"additionalProperties"`
			Deprecated string `json:"deprecated"`
		} `json:"RESTPostAPIGuildStickerFormDataBody"`
		RESTPostAPIGuildStickerResult struct {
			Ref string `json:"$ref"`
			Deprecated string `json:"deprecated"`
		} `json:"RESTPostAPIGuildStickerResult"`
		RESTPatchAPIGuildStickerJSONBody struct {
			Ref string `json:"$ref"`
			Deprecated string `json:"deprecated"`
		} `json:"RESTPatchAPIGuildStickerJSONBody"`
		AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure44570964819232207445709648186922084457096481694220944570964802576 struct {
			Type string `json:"type"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"name"`
				Description struct {
					Type []string `json:"type"`
					Description string `json:"description"`
				} `json:"description"`
				Tags struct {
					Type string `json:"type"`
					Description string `json:"description"`
				} `json:"tags"`
			} `json:"properties"`
			AdditionalProperties bool `json:"additionalProperties"`
			Description string `json:"description"`
		} `json:"AddUndefinedToPossiblyUndefinedPropertiesOfInterface<structure-445709648-1923-2207-445709648-1869-2208-445709648-1694-2209-445709648-0-2576>"`
		RESTPatchAPIGuildStickerResult struct {
			Ref string `json:"$ref"`
			Deprecated string `json:"deprecated"`
		} `json:"RESTPatchAPIGuildStickerResult"`
	} `json:"definitions"`
}