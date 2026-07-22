package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message string, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// ---> função que implementa a interface de erro para RestErr
func (r *RestErr) Error() string {
	return r.Message
}

// ----------------------------
// TRATAMENTO DE ERROS HTTP
// ----------------------------

// erro 400 - requisição inválida
func NewBadRequestValidationError(message string, err string, code int) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

// erro 401 - não autorizado
func NewUnauthorizedError(message string, err string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
  
// erro 404 - não encontrado
func NewNotFoundError(message string, err string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
		Causes:  causes,
	}
}

// erro 500 - erro interno do servidor
func NewInternalServerError(message string, err string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

// erro 503 - serviço indisponível
func NewServiceUnavailableError(message string, err string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "service_unavailable",
		Code:    http.StatusServiceUnavailable,
	}
}

// erro 504 - gateway timeout
func NewGatewayTimeoutError(message string, err string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "gateway_timeout",
		Code:    http.StatusGatewayTimeout,
	}
}
