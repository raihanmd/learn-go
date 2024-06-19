package main

import (
	"embed"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	// ? .gohtml or .gohtml or .tmpl is same
	tmpl := template.Must(template.ParseGlob("./templates/*.gohtml"))

	if auth, _ := r.Cookie("auth"); auth.Value == "admin" {
		tmpl.ExecuteTemplate(w, "admin.gohtml", "Lynx")
	} else {
		tmpl.ExecuteTemplate(w, "hello.gohtml", "Lynx")
	}
}

func TestTemplate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "auth", Value: "admin"})
	rec := httptest.NewRecorder()

	SimpleHTML(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	t.Log(string(body))
}

//go:embed templates/*.gohtml
var tmplResource embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(tmplResource, "templates/*.gohtml"))

	if auth, _ := r.Cookie("auth"); auth.Value == "admin" {
		tmpl.ExecuteTemplate(w, "admin.gohtml", "Lynx")
	} else {
		tmpl.ExecuteTemplate(w, "hello.gohtml", "Lynx")
	}
}

func TestEmbedTemplate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "auth", Value: "user"})
	rec := httptest.NewRecorder()

	SimpleHTML(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	t.Log(string(body))
}

func DataTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/multidata.gohtml"))

	tmpl.ExecuteTemplate(w, "multidata.gohtml", map[string]any{
		"Name": map[string]any{
			"First": "John",
			"Last":  "Doe",
		},
		"Age": 20,
	})

	// ? or with struct

	// tmpl.ExecuteTemplate(w, "multidata.gohtml", struct{ Name struct{ First, Last string }, Age any }{
	// 	Name: struct{ First, Last string }{
	// 		First: "John",
	// 		Last:  "Doe",
	// 	},
	// 	Age:       20,
	// })
}

func TestDataTemplate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	DataTemplate(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	t.Log(string(body))
}

func ActionTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("./templates/*.gohtml"))
	auth, _ := r.Cookie("auth")

	tmpl.ExecuteTemplate(w, "dashboard.gohtml", map[string]any{
		"Title":   "Dashboard",
		"IsAdmin": auth.Value == "admin",
		"Count":   rand.Int31(),
		"Hobbies": []string{"Music", "Football", "Cooking"},
		"Address": map[string]any{
			"Street": "123 Main St",
			"City":   "New York",
		},
	})
}

func TestActionTemplate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "auth", Value: "not_admin"})
	rec := httptest.NewRecorder()

	ActionTemplate(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	t.Log(string(body))
}

func FuncTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("func")
	tmpl = tmpl.Funcs(map[string]any{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	})
	//? Pipeline for parse to next fubnc as param
	tmpl = template.Must(tmpl.Parse("\nHello {{ upper .Name | lower }}\nAnd Hello {{ .Name | upper }}"))

	name, _ := r.Cookie("name")

	tmpl.ExecuteTemplate(w, "func", map[string]any{
		"Name": name.Value,
	})
}

func TestFuncTemplate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "name", Value: "hayato"})
	rec := httptest.NewRecorder()

	FuncTemplate(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	t.Log(string(body))
}

// ? Best practice for template

//go:embed templates/hello.gohtml
var templatesEmbed embed.FS

var tmpl = template.Must(template.ParseFS(templatesEmbed, "templates/hello.gohtml"))

func CachingTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "hello.gohtml", "Lynx")
}

func TestCachingTemplate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	CachingTemplate(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	t.Log(string(body))
}
