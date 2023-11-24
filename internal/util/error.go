package util

import "errors"

var (
	ErrorInvalidRequest                 = errors.New("Um ou mais campos inválidos")
	ErrorAuthorizationHeaderNotProvided = errors.New("Authorization header is not provided")
	ErrorPlanExpired                    = errors.New("Plano expirado ou não encontrado")
	ErrorPermNotEnough                  = errors.New("Permissão insuficiente")
	ErrorTokenExpired                   = errors.New("Token expirado")
	ErrorItemDuplicate                  = errors.New("Item já adicionado")
	ErrorDatabaseCreate                 = errors.New("Error Ao Criar registro")
	ErrorDatabaseRead                   = errors.New("Error Ao Consultar registro")
	ErrorLoginUser                      = errors.New("Senha incorreta")
	ErrorUserNotFound                   = errors.New("Não foi possível encontrar sua conta")
	ErroEmailInUse                      = errors.New("Email já está em uso")
	ErrorDatabaseUpdate                 = errors.New("Error Ao Atualizar registro")
	ErrorLoginDatabase                  = errors.New("Email Ou Senha Incorretos")
	ErrorLoginInvalidDatabase           = errors.New("Email Ou Senha Inválidos")
	ErrorDatabaseDelete                 = errors.New("Error Ao Deletar registro")
	ErrorDatabaseRestore                = errors.New("Error Ao Recuperar registro")
	ErrorGeneratePDF                    = errors.New("Error Ao Gerar Pdf")
	ErrorReadFile                       = errors.New("Error Ao Ler arquivo")
	ErrorUploadFile                     = errors.New("Error Ao Fazer upload do arquivo")
	ErrorSendEmail                      = errors.New("Error Ao Enviar Email")
	ErrorSessionBlocked                 = errors.New("Sessão bloqueada")
	ErrEmptyCsvFile                     = errors.New("Csv file is empty")
)
