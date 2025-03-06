package port

import "github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"

type OperationTypesRepository interface {
	CreateOperationTypes(operationTypes entity.OperationsType) (entity.OperationsType, error)
	GetOperationTypesByID(operationTypesID int) (entity.OperationsType, error)
	GetAllOperationTypes() ([]entity.OperationsType, error)
	UpdateOperationTypes(operationTypes entity.OperationsType) (entity.OperationsType, error)
	DeleteOperationTypes(operationTypesID int) error
}
