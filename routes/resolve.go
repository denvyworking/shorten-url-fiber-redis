package routes

import (
	"github.com/denvyworking/shorten-url-fiber-redis/database"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	// создаем базу по дефолтному шаблону в нашем файле shorten
	r := database.CreateClient(0)
	defer r.Close() // close db connection

	value, err := r.Get(database.Ctx, url).Result() // просто получить значение
	// не нашел значение в БД
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short URL not found in the database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	// создаем отдельную ДБ для подсчета перехода по сслыке
	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	// когда мы заходим в данную функцию, мы перенаправляем клиентов с их ссылки short
	// на long (defaul url)и таким образом корректно обращаемся к серверу.
	return c.Redirect(value, 301)
}
