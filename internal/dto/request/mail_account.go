package request

import "email-manage/internal/dto"

type MailAccountImportRequest struct {
	dto.BaseRequest
	Content    string `json:"content" binding:"required"`
	SourceType string `json:"source_type" binding:"required,oneof=paste file"`
	Filename   string `json:"filename,omitempty" binding:"omitempty,max=255"`
	Mode       string `json:"mode,omitempty" binding:"omitempty,oneof=fill_missing"`
}

func (r *MailAccountImportRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Content.required":    "导入内容不能为空",
		"SourceType.required": "导入来源不能为空",
		"SourceType.oneof":    "导入来源必须是 paste 或 file",
		"Filename.max":        "文件名长度不能超过255个字符",
		"Mode.oneof":          "导入模式仅支持 fill_missing",
	}
}

type CreateMailAccountRequest struct {
	dto.BaseRequest
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password,omitempty" binding:"omitempty,max=255"`
	ClientID     string `json:"client_id,omitempty" binding:"omitempty,max=255"`
	RefreshToken string `json:"refresh_token,omitempty" binding:"omitempty,max=4000"`
	Status       string `json:"status,omitempty" binding:"omitempty,oneof=unused used"`
	Remark       string `json:"remark,omitempty" binding:"omitempty,max=1000"`
	Tags         string `json:"tags,omitempty" binding:"omitempty,max=2000"`
	Folder       string `json:"folder,omitempty" binding:"omitempty,max=100"`
}

func (r *CreateMailAccountRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Email.required":       "邮箱不能为空",
		"Email.email":          "邮箱格式不正确",
		"Password.max":         "密码长度不能超过255个字符",
		"ClientID.max":         "client_id 长度不能超过255个字符",
		"RefreshToken.max":     "refresh_token 长度不能超过4000个字符",
		"Status.oneof":         "状态值仅支持 unused 或 used",
		"Remark.max":           "备注长度不能超过1000个字符",
		"Tags.max":             "标签长度不能超过2000个字符",
		"Folder.max":           "文件夹长度不能超过100个字符",
	}
}

type UpdateMailAccountRequest struct {
	dto.BaseRequest
	Password     string `json:"password,omitempty" binding:"omitempty,max=255"`
	ClientID     string `json:"client_id,omitempty" binding:"omitempty,max=255"`
	RefreshToken string `json:"refresh_token,omitempty" binding:"omitempty,max=4000"`
	Status       string `json:"status,omitempty" binding:"omitempty,oneof=unused used"`
	Remark       string `json:"remark,omitempty" binding:"omitempty,max=1000"`
	Tags         string `json:"tags,omitempty" binding:"omitempty,max=2000"`
	Folder       string `json:"folder,omitempty" binding:"omitempty,max=100"`
}

func (r *UpdateMailAccountRequest) GetValidationMessages() map[string]string {
	return (&CreateMailAccountRequest{}).GetValidationMessages()
}

type UpdateMailAccountStatusRequest struct {
	dto.BaseRequest
	Status string `json:"status" binding:"required,oneof=unused used"`
}

func (r *UpdateMailAccountStatusRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Status.required": "状态不能为空",
		"Status.oneof":    "状态值仅支持 unused 或 used",
	}
}

type FetchCodeRequest struct {
	dto.BaseRequest
	Keyword   string `json:"keyword,omitempty" binding:"omitempty,max=100"`
	SentAfter int64  `json:"sent_after,omitempty"` // Unix 时间戳（秒），表示触发验证码发送的时间
}

func (r *FetchCodeRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Keyword.max": "关键词长度不能超过100个字符",
	}
}

type BatchDeleteMailAccountRequest struct {
	dto.BaseRequest
	IDs []string `json:"ids" binding:"required,min=1"`
}

func (r *BatchDeleteMailAccountRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"IDs.required": "ID列表不能为空",
		"IDs.min":      "至少需要选择一个项目",
	}
}

type MailAccountListRequest struct {
	Page          int    `form:"page"`
	Size          int    `form:"size"`
	Keyword       string `form:"keyword"`
	Status        string `form:"status"`
	MailboxStatus string `form:"mailbox_status"`
	ProviderReady string `form:"provider_ready"`
}
