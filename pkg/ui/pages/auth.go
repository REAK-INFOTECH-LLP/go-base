package pages

import (
	"github.com/labstack/echo/v4"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"reak/base/pkg/ui"
	"reak/base/pkg/ui/forms"
	"reak/base/pkg/ui/layouts"
)

func Login(ctx echo.Context, form *forms.Login) error {
	r := ui.NewRequest(ctx)
	r.Title = "Login"

	return r.Render(layouts.Auth, form.Render(r))
}

func Register(ctx echo.Context, form *forms.Register) error {
	r := ui.NewRequest(ctx)
	r.Title = "Register"

	return r.Render(layouts.Auth, form.Render(r))
}

func ForgotPassword(ctx echo.Context, form *forms.ForgotPassword) error {
	r := ui.NewRequest(ctx)
	r.Title = "Forgot password"

	g := Group{
		Div(
			Class("content"),
			P(Text("Enter your email address and we'll email you a link that allows you to reset your password.")),
		),
		form.Render(r),
	}

	return r.Render(layouts.Auth, g)
}

func ResetPassword(ctx echo.Context, form *forms.ResetPassword) error {
	r := ui.NewRequest(ctx)
	r.Title = "Reset your password"

	return r.Render(layouts.Auth, form.Render(r))
}
