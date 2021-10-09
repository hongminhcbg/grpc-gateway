package erp

var ErrorInternalServer = NewApplicationError(500000, "internal server error")
var ErrorBadRequest = NewApplicationError(400000, "Bad request")
