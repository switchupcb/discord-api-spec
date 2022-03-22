type AutoGenerated struct {
	Definitions []Definitions `json:"definitions"`
}
type RPCErrorCodes struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
}
type RPCCloseEventCodes struct {
	Type string `json:"type"`
	Enum []int `json:"enum"`
	Description string `json:"description"`
}
type Definitions struct {
	RPCErrorCodes RPCErrorCodes `json:"RPCErrorCodes,omitempty"`
	RPCCloseEventCodes RPCCloseEventCodes `json:"RPCCloseEventCodes,omitempty"`
}