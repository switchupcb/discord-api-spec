package v9

type RESTGetAPIStickerResult struct {
	Ref string `json:"$ref"`
}
type ID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PackID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Description struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Tags struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Asset struct {
	Type        string `json:"type"`
	Const       string `json:"const"`
	Description string `json:"description"`
	Deprecated  string `json:"deprecated"`
}
type Type struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type FormatType struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Available struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type GuildID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type User struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type SortValue struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	ID          ID          `json:"id"`
	PackID      PackID      `json:"pack_id"`
	Name        Name        `json:"name"`
	Description Description `json:"description"`
	Tags        Tags        `json:"tags"`
	Asset       Asset       `json:"asset"`
	Type        Type        `json:"type"`
	FormatType  FormatType  `json:"format_type"`
	Available   Available   `json:"available"`
	GuildID     GuildID     `json:"guild_id"`
	User        User        `json:"user"`
	SortValue   SortValue   `json:"sort_value"`
}
type APISticker struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type Snowflake struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type StickerType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type StickerFormatType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Username struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Discriminator struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Avatar struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Bot struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type System struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type MfaEnabled struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Banner struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type AccentColor struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Locale struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Verified struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Email struct {
	Type        []string `json:"type"`
	Description string   `json:"description"`
}
type Flags struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PremiumType struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type PublicFlags struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID            ID            `json:"id"`
	Username      Username      `json:"username"`
	Discriminator Discriminator `json:"discriminator"`
	Avatar        Avatar        `json:"avatar"`
	Bot           Bot           `json:"bot"`
	System        System        `json:"system"`
	MfaEnabled    MfaEnabled    `json:"mfa_enabled"`
	Banner        Banner        `json:"banner"`
	AccentColor   AccentColor   `json:"accent_color"`
	Locale        Locale        `json:"locale"`
	Verified      Verified      `json:"verified"`
	Email         Email         `json:"email"`
	Flags         Flags         `json:"flags"`
	PremiumType   PremiumType   `json:"premium_type"`
	PublicFlags   PublicFlags   `json:"public_flags"`
}
type APIUser struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type UserFlags struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type UserPremiumType struct {
	Type        string `json:"type"`
	Enum        []int  `json:"enum"`
	Description string `json:"description"`
}
type Items struct {
	Ref string `json:"$ref"`
}
type StickerPacks struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}
type Properties struct {
	StickerPacks StickerPacks `json:"sticker_packs"`
}
type RESTGetNitroStickerPacksResult struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type Stickers struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type SkuID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type CoverStickerID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type BannerAssetID struct {
	Ref         string `json:"$ref"`
	Description string `json:"description"`
}
type Properties struct {
	ID             ID             `json:"id"`
	Stickers       Stickers       `json:"stickers"`
	Name           Name           `json:"name"`
	SkuID          SkuID          `json:"sku_id"`
	CoverStickerID CoverStickerID `json:"cover_sticker_id"`
	Description    Description    `json:"description"`
	BannerAssetID  BannerAssetID  `json:"banner_asset_id"`
}
type APIStickerPack struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTGetAPIGuildStickersResult struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}
type RESTGetAPIGuildStickerResult struct {
	Ref string `json:"$ref"`
}
type File struct {
	Description string `json:"description"`
}
type Properties struct {
	Name        Name        `json:"name"`
	Description Description `json:"description"`
	Tags        Tags        `json:"tags"`
	File        File        `json:"file"`
}
type RESTPostAPIGuildStickerFormDataBody struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type RESTPostAPIGuildStickerResult struct {
	Ref string `json:"$ref"`
}
type RESTPatchAPIGuildStickerJSONBody struct {
	Ref string `json:"$ref"`
}
type Properties struct {
	Name        Name        `json:"name"`
	Description Description `json:"description"`
	Tags        Tags        `json:"tags"`
}
type AddUndefinedToPossiblyUndefinedPropertiesOfInterfaceStructure574792367103413185747923679801319574792367932132057479236701433 struct {
	Type                 string     `json:"type"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Description          string     `json:"description"`
}
type RESTPatchAPIGuildStickerResult struct {
	Ref string `json:"$ref"`
}
