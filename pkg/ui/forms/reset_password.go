package forms

import (
	"net/http"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"reak/base/pkg/form"
	"reak/base/pkg/ui"
	. "reak/base/pkg/ui/components"
)

type ResetPassword struct {
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
	form.Submission
}

func (f *ResetPassword) Render(r *ui.Request) Node {
	return Form(
		ID("reset-password"),
		Method(http.MethodPost),
		HxBoost(),
		Action(r.CurrentPath),
		InputField(InputFieldParams{
			Form:        f,
			FormField:   "Password",
			Name:        "password",
			InputType:   "password",
			Label:       "Password",
			Placeholder: "******",
		}),
		InputField(InputFieldParams{
			Form:        f,
			FormField:   "PasswordConfirm",
			Name:        "password-confirm",
			InputType:   "password",
			Label:       "Confirm password",
			Placeholder: "******",
		}),
		ControlGroup(
			FormButton(ColorPrimary, "Update password"),
		),
		CSRF(r),
	)
}
