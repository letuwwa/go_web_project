package config

const (
	DBOperationError   = "db_operation_error"
	DBEntityNotFund    = "db_entity_not_found"
	HashOperationError = "hash_operation_error"
)

type Response struct {
	Message string `json:"message"`
}

type ResponseDeleted struct {
	IsDeleted bool `json:"is_deleted"`
}
