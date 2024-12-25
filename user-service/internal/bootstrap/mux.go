package configuration

import (
	"github.com/kingstonduy/go-core/pipeline"
	"github.com/kingstonduy/user-service/internal/domain"
)

func ResgisterPipeline(
	IGetUserInformationHandler domain.IGetUserInformationHandler,
	IRLoginHandler domain.IRLoginHandler,
	IRegisterHandler domain.IRegisterHandler,
	IUpdateUserInformationHandler domain.IUpdateUserInformationHandler,

) {
	pipeline.RegisterRequestHandler(IGetUserInformationHandler)
	pipeline.RegisterRequestHandler(IRLoginHandler)
	pipeline.RegisterRequestHandler(IRegisterHandler)
	pipeline.RegisterRequestHandler(IUpdateUserInformationHandler)
}
