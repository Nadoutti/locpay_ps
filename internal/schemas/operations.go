package schemas



type Operation struct {
	ID string `json:"id"`
	Receiver_id string `json:"receiver_id"`
	GrossValue float32 `json:"gross_value"`
	Fee float32 `json:"fee"`
	NetValue float32 `json:"net_value"`
	Status string `json:"status"`
}

type OperationDTO struct {
	Netvalue float32 `json:"net_value"`
	Status string `json:"status"`
}


type CreateOperation struct {
	Receiver_id string `json:"receiver_id"`
	GrossValue float32 `json:"gross_value"`
}
