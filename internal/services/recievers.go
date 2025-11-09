package services

import (
	"locpay_api/internal/infra/repository"
)


func GetRecieverById(id string) (map[string]interface{}, error) {

	response, str, err := repository.FindRecieverById(id)

	if err != nil {
		return map[string]interface{}{"error": "nao foi possivel encontrar o usuario", "response": str}, err
	}

	return map[string]interface{}{"message": response, "response": str}, err
}
