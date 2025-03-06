package repository

import (
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/port"
	"gorm.io/gorm"
)

type OperationTypesRepository struct {
	db *gorm.DB
}

func NewOperationTypesRepository(db *gorm.DB) port.OperationTypesRepository {
	return &OperationTypesRepository{db: db}
}

func (r *OperationTypesRepository) CreateOperationTypes(operationType entity.OperationsType) (entity.OperationsType, error) {
	if err := r.db.Create(&operationType).Error; err != nil {
		return entity.OperationsType{}, err
	}
	return operationType, nil
}

func (r *OperationTypesRepository) GetOperationTypesByID(operationTypeId int) (entity.OperationsType, error) {
	var operationType entity.OperationsType
	if err := r.db.Where("operation_type_id = ?", operationTypeId).First(&operationType).Error; err != nil {
		return entity.OperationsType{}, err
	}
	return operationType, nil
}

func (r *OperationTypesRepository) GetAllOperationTypes() ([]entity.OperationsType, error) {
	var operationType []entity.OperationsType
	if err := r.db.Find(&operationType).Error; err != nil {
		return []entity.OperationsType{}, err
	}
	return operationType, nil
}

func (r *OperationTypesRepository) UpdateOperationTypes(operationType entity.OperationsType) (entity.OperationsType, error) {
	if err := r.db.Save(&operationType).Error; err != nil {
		return entity.OperationsType{}, err
	}
	return operationType, nil
}

func (r *OperationTypesRepository) DeleteOperationTypes(operationTypeId int) error {
	var operationType entity.OperationsType
	if err := r.db.Delete(&operationType, operationTypeId).Error; err != nil {
		return err
	}
	return nil
}
