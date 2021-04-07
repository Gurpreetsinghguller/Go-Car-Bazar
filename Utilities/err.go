package Utilities

import "main/Models"

func SetError(err Models.Error, message string) Models.Error {
	err.IsError = true
	err.Message = message
	return err
}
