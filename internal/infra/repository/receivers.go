package repository

import (
	"locpay_api/internal/schemas"
	"locpay_api/pkg/utils"
)

func FindRecieverById(id string) (schemas.GetReceiverDTO, string, error) {

	connect := utils.GetSupabase()

	var receiver schemas.Receiver


	_, err := connect.From("receivers").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&receiver)


	if err != nil {
		return schemas.GetReceiverDTO{}, "erro no receiver", err
	}

	var operation_history []schemas.Operation

	_, err = connect.From("operations").
		Select("*", "", false).
		Eq("receiver_id", receiver.ID).
		ExecuteTo(&operation_history)

	if err != nil {
		return schemas.GetReceiverDTO{}, "erro no operation", err
	}
		

	op := schemas.GetReceiverDTO {
		Balance: receiver.Balance,
		Name: receiver.Name,
		OperationHistory: operation_history,


	}

	return op, "", nil
}
