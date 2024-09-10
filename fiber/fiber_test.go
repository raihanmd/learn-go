package test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/stretchr/testify/assert"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

var app = fiber.New(fiber.Config{
	StructValidator: &structValidator{validate: validator.New()},
})

func init() {
	app.Use(cors.New())
}

func TestFiberApp(t *testing.T) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Hello, World!", string(bytes))
}

func TestHeader(t *testing.T) {
	app.Get("/protected", func(c fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)
		if token != "admin" {
			return c.SendStatus(401)
		}

		return c.JSON(fiber.Map{"message": "Access granted"})
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Add("Authorization", "admin")

	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, _ := io.ReadAll(res.Body)
	var resJson fiber.Map

	json.Unmarshal(bytes, &resJson)

	assert.Equal(t, "Access granted", resJson["message"])
}

type User struct {
	Name string `uri:"name" json:"name" form:"name" validate:"required,min=2,max=50"`
	Age  int    `uri:"age" json:"age" form:"age" validate:"required,number,min=1"`
}

func TestParams(t *testing.T) {
	app.Get("/hello/:name/:age", func(c fiber.Ctx) error {
		var user User

		err := c.Bind().URI(&user)
		if err != nil {
			return err
		}

		t.Logf("%+v", user)

		return c.SendString("Hello " + user.Name + ", your age is " + strconv.Itoa(user.Age))
	})

	req := httptest.NewRequest(http.MethodGet, "/hello/adit/18", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Hello adit, your age is 18", string(bytes))
}

func TestJsonReq(t *testing.T) {
	app.Post("/json", func(c fiber.Ctx) error {
		var user User

		err := c.Bind().JSON(&user)
		if err != nil {
			return err
		}

		t.Logf("%+v", user)

		return c.JSON(fiber.Map{"message": "Hello " + user.Name + ", your age is " + strconv.Itoa(user.Age)})
	})

	body := strings.NewReader(`{"name":"Eko", "age":18}`)
	req := httptest.NewRequest(http.MethodPost, "/json", body)
	req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)

	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, _ := io.ReadAll(res.Body)
	var resJson fiber.Map

	json.Unmarshal(bytes, &resJson)

	assert.Equal(t, "Hello Eko, your age is 18", resJson["message"])
}

func TestForm(t *testing.T) {
	app.Post("/form", func(c fiber.Ctx) error {
		var user User

		err := c.Bind().Form(&user)
		if err != nil {
			return err
		}

		t.Logf("%+v", user)

		return c.JSON(fiber.Map{"message": "Hello " + user.Name + ", your age is " + strconv.Itoa(user.Age)})
	})

	body := strings.NewReader("name=Eko&age=18")
	req := httptest.NewRequest(http.MethodPost, "/form", body)
	req.Header.Set("Content-Type", fiber.MIMEApplicationForm)

	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, _ := io.ReadAll(res.Body)
	var resJson fiber.Map

	json.Unmarshal(bytes, &resJson)

	assert.Equal(t, "Hello Eko, your age is 18", resJson["message"])
}

type File struct {
	File *multipart.FileHeader `form:"file"`
}

//go:embed source/text.txt
var fileSource []byte

func TestMultipartForm(t *testing.T) {
	app.Post("/multipart", func(c fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			t.Logf("Error: %v", err)
			return err
		}

		err = c.SaveFile(file, fmt.Sprintf("./upload/%s", file.Filename))
		if err != nil {
			t.Logf("Error: %v", err)
			return err
		}

		return c.JSON(fiber.Map{"message": "Success"})
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, _ := writer.CreateFormFile("file", "text.txt")
	file.Write(fileSource)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/multipart", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, _ := io.ReadAll(res.Body)

	var resJson fiber.Map
	json.Unmarshal(bytes, &resJson)

	assert.Equal(t, "Success", resJson["message"])
}

func TestGroup(t *testing.T) {
	type Product struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	myProducts := []Product{
		{Name: "Product 1", Price: 100},
		{Name: "Product 2", Price: 200},
	}

	productRoute := app.Group("/products")
	productRoute.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"products": myProducts,
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, _ := io.ReadAll(res.Body)
	var resJson fiber.Map

	json.Unmarshal(bytes, &resJson)

	assert.Equal(t, "Product 1", resJson["products"].([]interface{})[0].(map[string]interface{})["name"])
}
