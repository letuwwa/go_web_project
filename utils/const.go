package utils

import "os"

const (
	DBOperationError   = "db_operation_error"
	DBOperationSuccess = "db_operation_success"
	DBEntityNotFund    = "db_entity_not_found"
	HashOperationError = "hash_operation_error"
	JWTOperationError  = "jwt_operation_error"
	LoginError         = "login_error"
	LoginSuccess       = "login_success"
	Forbidden          = "forbidden"
)

var JWTSalt = os.Getenv("JWTSalt")

type Response struct {
	Status string `json:"status"`
}

type ResponseWithString struct {
	Status string `json:"status"`
	Value  string `json:"value"`
}
