type AutoGenerated struct {
	Definitions Definitions `json:"definitions"`
}
type ApplicationCommandOptionType struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
	Deprecated string `json:"deprecated"`
}
type Definitions struct {
	ApplicationCommandOptionType ApplicationCommandOptionType `json:"ApplicationCommandOptionType"`
}