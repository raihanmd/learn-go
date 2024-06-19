# Golang Web `net/http` package

golang pkg `html/template` is safe for XSS, auto escaping char by default
for disable it u can do this

```go
func Tempate(w http.ResponseWriter, r *http.Request) {
  tmpl.ExecuteTemplate(w, "hello.gohtml", map[string]any{
    "SomeValue": template.HTML("<h1>Hello World</h1>"),
    //"SomeValue": template.CSS("h1{color:red;}"),
    //"SomeValue": template.JS("alert('Hello World');"),
  })
}
```
