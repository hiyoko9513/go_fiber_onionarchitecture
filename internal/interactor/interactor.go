package interactor

import (
	"hiyoko-fiber/internal/application/usecase"
	"hiyoko-fiber/internal/domain/services"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/infrastructure/persistence/repository"
	"hiyoko-fiber/internal/presentation/http/app/handler"
)

type Interactor interface {
	NewTableRepository() services.TableRepository
	NewUserRepository() services.UserRepository

	NewUserUseCase() usecase.UserUseCase
	NewFileUseCase() usecase.FileUseCase

	NewAppHandler() handler.AppHandler
	NewUserHandler() handler.UserHandler
	NewAuthHandler() handler.AuthHandler
	NewCliHandler() handler.CliHandler
}

type interactor struct {
	conn *database.MysqlEntClient
}

func NewInteractor(conn *database.MysqlEntClient) Interactor {
	return &interactor{conn}
}

func (i *interactor) NewTableRepository() services.TableRepository {
	return repository.NewTableRepository(i.conn)
}

func (i *interactor) NewUserRepository() services.UserRepository {
	return repository.NewUserRepository(i.conn)
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewFileUseCase() usecase.FileUseCase {
	return usecase.NewFileUseCase()
}

type appHandler struct {
	handler.UserHandler
	handler.AuthHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	appHandler.AuthHandler = i.NewAuthHandler()
	return appHandler
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}

func (i *interactor) NewAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(i.NewUserUseCase())
}

func (i *interactor) NewCliHandler() handler.CliHandler {
	return handler.NewCliHandler(i.NewFileUseCase())
}
