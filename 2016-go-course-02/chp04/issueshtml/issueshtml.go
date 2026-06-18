package main

import (
	"html/template"
	"log"
	"os"

	"github.com/aliwatters/go-course-02/chp04/github"
)

const templ = `<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align:left'>
	<th>#<th>State<th>User<th>Title
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a>
	<td>{{.State}}
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a>
</tr>
{{end}}
</table>`

var report = template.Must(template.New("issuelist").
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
