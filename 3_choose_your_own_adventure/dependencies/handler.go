package dependencies

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"html/template"
)

type MainHandler struct {
	story *Gopher
	tpl *template.Template
}

func NewMainHandler() *MainHandler{
	instance := &MainHandler{}

	gopherObject := &Gopher{}

	bytes, err := ioutil.ReadFile("3_choose_your_own_adventure/dependencies/gopher.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, gopherObject)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(defaultHandlerTmpl))

	instance.tpl = tpl
	instance.story = gopherObject
	return instance
}

func (mainHandler *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	multiplexer := mainHandler.httpMultiplexer()
	multiplexer.ServeHTTP(w, r)
}

func (mainHandler *MainHandler) httpMultiplexer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler.makeHandler(mainHandler.story.Intro))
	mux.HandleFunc("/new-york", mainHandler.makeHandler(mainHandler.story.NewYork))
	mux.HandleFunc("/denver", mainHandler.makeHandler(mainHandler.story.Denver))
	mux.HandleFunc("/debate", mainHandler.makeHandler(mainHandler.story.Debate))
	mux.HandleFunc("/home", mainHandler.makeHandler(mainHandler.story.Home))
	mux.HandleFunc("/mark-bates", mainHandler.makeHandler(mainHandler.story.MarkBates))
	mux.HandleFunc("/sean-kelly", mainHandler.makeHandler(mainHandler.story.SeanKelly))
	return mux
}

func (mainHandler *MainHandler) makeHandler(value interface {}) func(w http.ResponseWriter, r *http.Request){
	return func (w http.ResponseWriter, r *http.Request){
		mainHandler.tpl.Execute(w, value)
	}
}

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Choose your own adventure</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Story}}
			<p>{{.}}</p>
		{{end}}
		<ul>
		{{range .Options}}
			<li><a href="/{{.Arc}}">{{.Text}}</a></li>
		{{end}}
		</ul>
	</body>
</html>
`