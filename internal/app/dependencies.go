package app

import (
	uploadController "template/internal/controllers/upload"
	userController "template/internal/controllers/user"
	uploadRepo "template/internal/repositories/upload"
	uploadService "template/internal/services/upload"
	userService "template/internal/services/user"
	"template/pkg/cache"
	"template/pkg/database"
	uploadPkg "template/pkg/upload"
)

// Dependencies 应用运行时依赖容器
type Dependencies struct {
	UserController   *userController.Controller
	UploadController *uploadController.Controller
}

// NewDependencies 构建默认依赖集合
func NewDependencies() *Dependencies {
	db := database.GetDB()

	userSvc := userService.NewService(db, cache.GetCache(), userService.NewEmailMailer())

	uploadCfg := uploadPkg.NewDefaultConfig()
	uploadSvc := uploadService.NewService(
		uploadRepo.NewUploadRepository(db),
		uploadPkg.NewLocalStorage(uploadCfg.UploadDir),
		uploadCfg,
	)

	return &Dependencies{
		UserController:   userController.NewController(userSvc),
		UploadController: uploadController.NewController(uploadSvc),
	}
}
