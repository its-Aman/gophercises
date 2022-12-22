package main

import (
	"CHOOSE_YOUR_OWN_ADVENTURE/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {

	port := flag.Int("port", 3000, "the port to start the CYOA application on")
	filename := flag.String("file", "../../gopher.json", "the json file with the CYOA story")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		log.Fatal("Error while opening the file... ", err)
	}
	defer file.Close()

	story, err := cyoa.JsonStory(file)
	if err != nil {
		log.Fatal("Error while parsing the file... ", err)
	}

	tpl := template.Must(template.New("").Parse(tmplWithStoryPath))
	mux := http.NewServeMux()
	hDefault := cyoa.NewHandler(story)
	hWithStory := cyoa.NewHandler(story, cyoa.WithTemplate(tpl), cyoa.WithPathFunc(storyPathFunc))
	mux.Handle("/", hDefault)
	mux.Handle("/story/", hWithStory)

	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func storyPathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	storyPath := "/story"
	if path == storyPath || path == storyPath+"/" {
		path = storyPath + "/intro"
	}

	return path[len(storyPath)+1:]
}

const tmplWithStoryPath = `
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>Choose Your Own Adventure</title>
</head>

<body>
    <section class="page">
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}

		<ul>
            {{range .Options}}
            <li>
                <a href="/story/{{.Arc}}">{{.Text}}</a>
            </li>
            {{end}}
        </ul>
    </section>
	
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>

</body>

</html>
`
