package config

const (
	DBOperationError = "db_operation_error"
	DBEntityNotFund  = "db_entity_not_found"
)

type Response struct {
	Message string `json:"message"`
}
