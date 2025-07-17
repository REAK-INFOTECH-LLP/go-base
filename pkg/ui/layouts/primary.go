package layouts

import (
	"reak/base/ent/admin"
	"reak/base/pkg/routenames"
	"reak/base/pkg/ui"
	. "reak/base/pkg/ui/components"
	"reak/base/pkg/ui/icons"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Primary(r *ui.Request, content Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Data("theme", "dark"),
			Head(
				Metatags(r),
				CSS(),
				JS(),
			),
			Body(
				Div(
					Class("drawer lg:drawer-open"),
					Input(
						ID("sidebar"),
						Type("checkbox"),
						Class("drawer-toggle"),
					),
					Div(
						Class("drawer-content flex flex-col p-7 prose-base"),
						If(len(r.Title) > 0, H1(Text(r.Title))),
						FlashMessages(r),
						content,
						Label(
							For("sidebar"),
							Class("btn btn-primary drawer-button lg:hidden"),
							Text("Open drawer"),
						),
					),
					sidebarMenu(r),
				),
				HtmxListeners(r),
			),
		),
	)
}

func sidebarMenu(r *ui.Request) Node {
	header := func(text string) Node {
		return Li(
			Class("menu-title mt-3 uppercase"),
			Span(Text(text)),
		)
	}

	adminSubMenu := func() Node {
		entityTypeNames := admin.GetEntityTypeNames()
		entityTypeLinks := make(Group, len(entityTypeNames))
		for _, n := range entityTypeNames {
			entityTypeLinks = append(entityTypeLinks, MenuLink(r, icons.PencilSquare(), n, routenames.AdminEntityList(n)))
		}

		return Group{
			header("Entities"),
			entityTypeLinks,
			header("Monitoring"),
			Li(
				A(
					icons.CircleStack(),
					Href(r.Path(routenames.AdminTasks)),
					Text("Tasks"),
					Target("_blank"),
				),
			),
		}
	}

	return Div(
		Class("drawer-side"),
		Label(
			For("sidebar"),
			Aria("label", "close sidebar"),
			Class("drawer-overlay"),
		),
		Div(
			Class("menu bg-base-200 text-base-content min-h-full w-80 p-4"),
			Div(
				Class("w-2/3 mx-auto mt-3 mb-10"),
				Img(
					Src(ui.StaticFile("logo.png")),
				),
			),
			Ul(
				HxBoost(),
				header("Account"),
				If(r.IsAuth, MenuLink(r, icons.Exit(), "Logout", routenames.Logout)),
				If(!r.IsAuth, MenuLink(r, icons.Enter(), "Login", routenames.Login)),
				If(!r.IsAuth, MenuLink(r, icons.UserPlus(), "Register", routenames.Register)),
				If(!r.IsAuth, MenuLink(r, icons.QuestionCircle(), "Forgot password", routenames.ForgotPasswordSubmit)),
				Iff(r.IsAdmin, adminSubMenu),
			),
		),
	)
}
