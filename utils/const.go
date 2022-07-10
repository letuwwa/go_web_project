package utils

const (
	DBOperationError   = "db_operation_error"
	DBOperationSuccess = "db_operation_success"
	DBEntityNotFund    = "db_entity_not_found"
	HashOperationError = "hash_operation_error"
	LoginError         = "login_error"
	LoginSuccess       = "login_success"
)

type Response struct {
	Status string `json:"status"`
}
