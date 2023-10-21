package factory

import (
	"github.com/PedroGuilhermeSilv/codepix/application/usecase"
	"github.com/PedroGuilhermeSilv/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{
		Db: database,
	}
	transactionRepository := repository.TransactionRepositoryDb{
		Db: database,
	}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}
	return transactionUseCase

}
