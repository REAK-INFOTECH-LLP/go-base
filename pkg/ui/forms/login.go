package forms

import (
	"net/http"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"reak/base/pkg/form"
	"reak/base/pkg/routenames"
	"reak/base/pkg/ui"
	. "reak/base/pkg/ui/components"
)

type Login struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
	form.Submission
}

func (f *Login) Render(r *ui.Request) Node {
	return Form(
		ID("login"),
		Method(http.MethodPost),
		HxBoost(),
		Action(r.Path(routenames.LoginSubmit)),
		FlashMessages(r),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Email",
			Name:      "email",
			InputType: "email",
			Label:     "Email address",
			Value:     f.Email,
		}),
		InputField(InputFieldParams{
			Form:        f,
			FormField:   "Password",
			Name:        "password",
			InputType:   "password",
			Label:       "Password",
			Placeholder: "******",
		}),
		Div(
			Class("text-right text-primary mt-2"),
			A(
				Href(r.Path(routenames.ForgotPassword)),
				Text("Forgot password?"),
			),
		),
		ControlGroup(
			FormButton(ColorPrimary, "Login"),
			ButtonLink(ColorLink, r.Path(routenames.Login), "Cancel"),
		),
		CSRF(r),
		Div(
			Class("text-center text-base-content/50 mt-4"),
			Text("Don't have an account? "),
			A(
				Href(r.Path(routenames.Register)),
				Text("Register"),
			),
		),
	)
}
