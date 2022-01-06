package main

import (
	"ChatApp/components"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/kingdevnl/GoLive"
	"log"
)

func main() {

	GoLive.RegisterComponent("messages", &components.Messages{})
	GoLive.RegisterComponent("inputs", &components.Inputs{})

	engine := html.New("./views", ".tpl")
	engine.Reload(true)
	GoLive.SetupEngine(engine)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/js", "./js", fiber.Static{
		CacheDuration: -1,
		Browse:        true,
	})

	store := session.New()

	GoLive.SetupHandler(GoLive.Config{Path: "/ws"}, app, func(c *fiber.Ctx) {
		sess, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return
		}
		if sess.Get("user") == nil {
			return
		}
		c.Locals("user", sess.Get("user").(string))
	})
	GoLive.SetupGarbageCollector()
	GoLive.InitComponents()

	app.Get("/", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("Internal Server Error")
		}

		//check if user is set on session
		if sess.Get("user") == nil {
			sess.Set("user", "User-"+GoLive.GenerateID(4))
			sess.Save()
		}

		return c.Render("index", fiber.Map{
			"title": "Chat App",
		})
	})
	log.Fatalln(app.Listen(":3000"))
}
