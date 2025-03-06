package entity

type OperationsType struct {
	OperationTypeId int    `json:"operation_type_id" gorm:"primaryKey"`
	Description0    string `json:"description0"`
}
