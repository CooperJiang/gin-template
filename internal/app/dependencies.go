package app

import (
	compatController "email-manage/internal/controllers/compat"
	mailAccountController "email-manage/internal/controllers/mail_account"
	uploadController "email-manage/internal/controllers/upload"
	userController "email-manage/internal/controllers/user"
	"email-manage/internal/integrations/mail_provider"
	mailAccountRepo "email-manage/internal/repositories/mail_account"
	uploadRepo "email-manage/internal/repositories/upload"
	mailAccountService "email-manage/internal/services/mail_account"
	uploadService "email-manage/internal/services/upload"
	userService "email-manage/internal/services/user"
	"email-manage/pkg/cache"
	"email-manage/pkg/config"
	"email-manage/pkg/database"
	"email-manage/pkg/security"
	uploadPkg "email-manage/pkg/upload"
)

// Dependencies 应用运行时依赖容器
type Dependencies struct {
	UserController        *userController.Controller
	UploadController      *uploadController.Controller
	MailAccountController *mailAccountController.Controller
	CompatController      *compatController.Controller
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

	// mail account dependencies
	cfg := config.GetConfig()
	var fieldCipher *security.FieldCipher
	if cfg.Security.EncryptionKey != "" {
		if cipher, err := security.NewFieldCipher(cfg.Security.EncryptionKey); err == nil {
			fieldCipher = cipher
		}
	}

	mailProviderClient := mail_provider.NewClient(mail_provider.ClientConfig{
		BaseURL: cfg.MailProvider.BaseURL,
	})
	mailAccountRepo := mailAccountRepo.NewRepository(db)
	mailAccountSvc := mailAccountService.NewService(mailAccountRepo, mailProviderClient, fieldCipher)

	return &Dependencies{
		UserController:        userController.NewController(userSvc),
		UploadController:      uploadController.NewController(uploadSvc),
		MailAccountController: mailAccountController.NewController(mailAccountSvc),
		CompatController:      compatController.NewController(userSvc, mailAccountSvc),
	}
}
