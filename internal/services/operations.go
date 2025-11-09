package services

import (
	"errors"
	"locpay_api/internal/infra/repository"
	"locpay_api/internal/schemas"
)


func CreateOperation(data *schemas.CreateOperation) (map[string]interface{}, error) {

	operation, operationId, err := repository.CreateOperation(data)

	if err != nil {
		return nil, err 
	}

	return map[string]interface{}{"message": operation, "operation_id": operationId}, nil
}

func GetOperationById(id string) (map[string]interface{}, error) {

	operation, operation_id, err := repository.FindOperationById(id)

	if err != nil {
		return map[string]interface{}{"error": "operacao ja existe"}, errors.New("operacao nao existe")
	}

	return map[string]interface{}{"message": operation, "id": operation_id}, err
}


func ConfirmOperation(id string) (map[string]interface{}, error) {

	response, operation, err := repository.ConfirmOperation(id)

	if response == "" {
		return map[string]interface{}{"error": "nao foi possivel confirmar a operacao"}, errors.New("nao foi possivel confirmar a operacao")
	}

	return map[string]interface{}{"message": "operacao confirmada", "confirmation": operation}, err
}

