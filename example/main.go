package main

import (
	"fmt"

	"github.com/alarbada/go-server-actions/actions"

	g "github.com/alarbada/gomponents"
	. "github.com/alarbada/gomponents/html"
	"github.com/alarbada/gomponents/hx"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("listening to :8080")

	e := router.Engine()
	e.Static("/public", "./public")

	e.Run()
}

var router = actions.NewRouter()

func Page(children ...g.Node) g.Node {
	return HTML(
		Head(
			Script(Src("https://unpkg.com/htmx.org@1.9.6")),
			Script(Src("https://cdn.tailwindcss.com")),
			Raw(`
				<link href="https://cdn.jsdelivr.net/npm/daisyui@3.9.4/dist/full.css" rel="stylesheet" type="text/css" />
				<script src="https://cdn.tailwindcss.com"></script>
				<script src="/public/main.js"></script>
			`),
		),
		Body(
			navbar(),
			g.Fragment(children...),
		),
	)
}

func HelloWorld() g.Node {
	return P(Text("Hello World"))
}

func navbar() g.Node {
	return Div(Class("flex gap-2"),
		Button(
			Class("btn btn-primary"),
			logHello.Hx(),
			hx.Swap("none"),
			Text("log hello")),
		Button(
			Class("btn btn-primary"),
			addHello.Hx(),
			hx.Swap("beforeend"),
			hx.Target("#hello-list"),
			Text("add hello")),

		Button(
			Class("btn btn-primary"),
			homePage.Hx(),
			hx.PushUrlT(),
			hx.Target("body"),
			Text("go to home page")),

		Button(
			Class("btn btn-primary"),
			action1.Hx(),
			hx.Target("body"),
			hx.PushUrlT(),
			Text("action 1")),

		Button(
			Class("btn btn-primary"),
			dashboardPage.Hx(),
			hx.PushUrlT(),
			hx.Target("body"),
			Text("go to dashboard page")),
	)
}

var homePage = router.GET("/")
var _ = homePage.Handle(func(c *gin.Context) g.Node {
	return Page(
		H1(Text("this is the home")),
		Div(ID("hello-list")),
	)
})

var dashboardRouter = router.Group("/dashboard")

var dashboardPage = dashboardRouter.GET("")
var _ = dashboardPage.Handle(func(c *gin.Context) g.Node {
	return Page(
		H1(Text("Dashboard")),
		Div(ID("hello-list")),
	)
})

var helloWorldPage = router.GET("/hello-world")
var _ = helloWorldPage.Handle(func(c *gin.Context) g.Node {
	return Page(HelloWorld())
})

var logHello = router.POST("/hello").Handle(func(c *gin.Context) g.Node {
	fmt.Println("hello")
	return nil
})

var addHello = router.POST("/add-hello").Handle(func(c *gin.Context) g.Node {
	return HelloWorld()
})

var counter = 0

func Counter() g.Node {
	counter++

	if counter <= 5 {
		return Div(
			ID("counter"),
			action1.Hx(),
			hx.Trigger("every 200ms"),
			hx.Target("this"),
			hx.Swap("outerHTML"),
			hx.Select("#counter"),
			Textf("%d", counter),
		)
	}

	return Div(
		ID("counter"),
		P(Text("counter reached")),
	)
}

var action1 = router.GET("/action1")
var _ = action1.Handle(func(c *gin.Context) g.Node {
	return Page(
		H1(Text("action1")),
		Div(ID("hello-list")),

		Counter(),
	)
})
