package usecase

import (
	"time"

	appInit "lolipad/boilerplate/init"
	"lolipad/boilerplate/module/base/store"
)

type BaseUsecase struct {
	config         *appInit.Config
	baseRepo       store.Repository
	contextTimeout time.Duration
}

// NewAutotextUsecase will create new an contactUsecase object representation of Base.Usecase
func NewAutoTextUsecase(br store.Repository, timeout time.Duration, config *appInit.Config) Usecase {
	return &BaseUsecase{
		config:         config,
		baseRepo:       br,
		contextTimeout: timeout,
	}
}
