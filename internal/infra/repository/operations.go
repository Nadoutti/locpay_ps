package repository

import (
	"locpay_api/internal/schemas"
	"locpay_api/pkg/utils"

	"github.com/google/uuid"
)

func FindOperationById(id string) (schemas.Operation, string, error) {

	connect := utils.GetSupabase()

	var operation schemas.Operation


	_, err := connect.From("operations").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&operation)


	if err != nil {
		return schemas.Operation{}, "", err
	}

	// op := schemas.OperationDTO {
	// 	Netvalue: operation.NetValue,
	// 	Status: operation.Status,
	//
	// }

	return operation, operation.ID, err
}

func CreateOperation(data *schemas.CreateOperation) (schemas.OperationDTO, string, error) {
	
	connect := utils.GetSupabase()

	// calculando a fee

	fee := data.GrossValue * 0.03

	netValue := data.GrossValue - fee

	newOperation := map[string]interface{}{
		"id": uuid.New().String(),
		"receiver_id": data.Receiver_id,
		"gross_value": data.GrossValue,
		"fee": fee,
		"net_value": netValue,
		"status": "pending",
	} 


	var createdOperation schemas.Operation

	_, err := connect.From("operations").
		Insert(newOperation, false, "", "", "").
		Single().
		ExecuteTo(&createdOperation)

	if (err != nil) {
		return schemas.OperationDTO{}, "", err
	}

	createdOp := schemas.OperationDTO {
		Netvalue: createdOperation.NetValue,
		Status: createdOperation.Status,
	}


	return createdOp, createdOperation.ID, nil

}

func ConfirmOperation(id string) (string, schemas.OperationDTO, error) {
	connect := utils.GetSupabase()

	var operation schemas.Operation


	// verificando se a operation existe
	_, err := connect.From("operations").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&operation)


	updateData := map[string]interface{}{
		"status": "confirmed",
	}

	// atualizando a operation
	_,_, err = connect.From("operations").
		Update(updateData, "", "minimal").
		Eq("id", id).
		Execute()


	if err != nil {
		return "", schemas.OperationDTO{}, err
	}

	var op schemas.Operation

	_, err = connect.From("operations").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&op)


	if err != nil {
		return "",schemas.OperationDTO{}, err
	}

	satinizedOperation := schemas.OperationDTO{
		Netvalue: op.NetValue,
		Status: op.Status,
	}


	return "operacao confirmada", satinizedOperation, nil
	
}
