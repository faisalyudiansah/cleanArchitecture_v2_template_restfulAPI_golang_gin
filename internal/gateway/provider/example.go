package provider

import (
	controllerExample "server/internal/controller"
	repositoryExample "server/internal/repository"
	"server/internal/route"
	usecaseExample "server/internal/usecase"

	"github.com/gin-gonic/gin"
)

var (
	exampleRepository repositoryExample.ExampleRepositoryInterface
)

var (
	exampleUseCase usecaseExample.ExampleUsecaseInterface
)

var (
	exampleController *controllerExample.ExampleController
)

func ProvideExampleModule(router *gin.Engine) {
	injectExampleModuleRepository()
	injectExampleModuleUseCase()
	injectExampleModuleController()

	route.ExampleControllerRoute(exampleController, router)

}

func injectExampleModuleRepository() {
	exampleRepository = repositoryExample.NewExampleRepositoryImplementation(db)
}

func injectExampleModuleUseCase() {
	exampleUseCase = usecaseExample.NewExampleUsecaseImplementation(store, exampleRepository)
}

func injectExampleModuleController() {
	exampleController = controllerExample.NewExampleController(exampleUseCase)
}
