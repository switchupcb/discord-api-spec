package rpc

type RPCErrorCodes struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Enum        []int  `json:"enum"`
}
type RPCCloseEventCodes struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Enum        []int  `json:"enum"`
}
