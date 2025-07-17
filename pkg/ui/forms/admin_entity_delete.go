package forms

import (
	"net/http"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"reak/base/pkg/routenames"
	"reak/base/pkg/ui"
	. "reak/base/pkg/ui/components"
)

func AdminEntityDelete(r *ui.Request, entityTypeName string) Node {
	return Form(
		Method(http.MethodPost),
		P(
			Textf("Are you sure you want to delete this %s?", entityTypeName),
		),
		ControlGroup(
			FormButton(ColorError, "Delete"),
			ButtonLink(
				ColorNone,
				r.Path(routenames.AdminEntityList(entityTypeName)),
				"Cancel",
			),
		),
		CSRF(r),
	)
}
