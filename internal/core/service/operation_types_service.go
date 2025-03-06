package service

import (
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/port"
)

type OperationTypesService struct {
	repo port.OperationTypesRepository
}

func NewOperationTypesService(repo port.OperationTypesRepository) *OperationTypesService {
	return &OperationTypesService{repo: repo}
}

func (s *OperationTypesService) CreateOperationTypes(opTypesInput entity.OperationsType) (entity.OperationsType, error) {

	opType := entity.OperationsType{
		OperationTypeId: opTypesInput.OperationTypeId,
		Description0:    opTypesInput.Description0,
	}

	return s.repo.CreateOperationTypes(opType)
}

func (s *OperationTypesService) GetOperationTypesByID(opTypeId int) (entity.OperationsType, error) {
	return s.repo.GetOperationTypesByID(opTypeId)
}

func (s *OperationTypesService) GetAllOperationTypes() ([]entity.OperationsType, error) {
	return s.repo.GetAllOperationTypes()
}

func (s *OperationTypesService) UpdateOperationTypes(opTypesInput entity.OperationsType) (entity.OperationsType, error) {

	opType := entity.OperationsType{
		OperationTypeId: opTypesInput.OperationTypeId,
		Description0:    opTypesInput.Description0,
	}

	return s.repo.UpdateOperationTypes(opType)
}

func (s *OperationTypesService) DeleteOperationTypes(opTypeId int) error {
	return s.repo.DeleteOperationTypes(opTypeId)
}
