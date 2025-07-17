package routenames

import (
	"fmt"
)

const (
	About                = "about"
	Login                = "login"
	LoginSubmit          = "login.submit"
	Register             = "register"
	RegisterSubmit       = "register.submit"
	ForgotPassword       = "forgot_password"
	ForgotPasswordSubmit = "forgot_password.submit"
	Logout               = "logout"
	VerifyEmail          = "verify_email"
	ResetPassword        = "reset_password"
	ResetPasswordSubmit  = "reset_password.submit"
	AdminTasks           = "admin:tasks"
)

func AdminEntityList(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_list", entityTypeName)
}

func AdminEntityAdd(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_add", entityTypeName)
}

func AdminEntityEdit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_edit", entityTypeName)
}

func AdminEntityDelete(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_delete", entityTypeName)
}

func AdminEntityAddSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_add.submit", entityTypeName)
}

func AdminEntityEditSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_edit.submit", entityTypeName)
}

func AdminEntityDeleteSubmit(entityTypeName string) string {
	return fmt.Sprintf("admin:%s_delete.submit", entityTypeName)
}
