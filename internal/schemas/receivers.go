package schemas



type Receiver struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Balance float32 `json:"balance"`
}

type GetReceiverDTO struct {
	Name string `json:"name"`
	Balance float32 `json:"balance"`
	OperationHistory []Operation `json:"operation_history"`
}
