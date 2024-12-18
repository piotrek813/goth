package post

import (
	"log"
	"piotrek813/goth/components"
	alertboxcomponent "piotrek813/goth/components/alert_box_component"
	"piotrek813/goth/daos"
	"piotrek813/goth/models"
	postcomponents "piotrek813/goth/routes/post/post_components"
	"piotrek813/goth/session"
	"piotrek813/goth/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r *fiber.Router) {
	postGroup := (*r).Group("/post")

	postGroup.Post("/", func(c *fiber.Ctx) error {
		title := c.FormValue("title")
		body := c.FormValue("body")

		sess, err := session.Store.Get(c)
		userId := sess.Get("user_id").(int)

		err = daos.SavePost(&models.Post{UserId: userId, Title: title, Body: body})

		if err != nil {
			log.Println(err)
			return utils.Render(c, alertboxcomponent.AlertBox("wystąpił problem"))
		}

		return utils.Render(c, alertboxcomponent.WithAlertBox(components.AsDashboardContent(postcomponents.PostList(nil, userId)), "Dodano", alertboxcomponent.Success))
	})

	postGroup.Delete("/", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.FormValue("postId"))

		err := daos.DeletePost(id)

		if err != nil {
			return utils.Render(c, alertboxcomponent.AlertBox("Something went wrong", alertboxcomponent.Error))
		}

		return c.SendString("")
	})

	postGroup.Get("/", func(c *fiber.Ctx) error {
		posts, err := daos.FindAllPost()

		if err != nil {
			return utils.Render(c, alertboxcomponent.AlertBox("Nie udało się "))
		}

		if len(posts) == 0 {
			return c.SendString("Nic tu nie ma")
		}

		sess, _ := session.Store.Get(c)
		userId := sess.Get("user_id").(int)

		return utils.Render(c, postcomponents.PostList(posts, userId))
	})

	postGroup.Get("/form", func(c *fiber.Ctx) error {
		return utils.Render(c, components.AsDashboardContent(postcomponents.PostForm(), true))
	})
}
