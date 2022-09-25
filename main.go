package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err = template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)

}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "write", nil)

}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	var id := r.FormValue("id")
	var title := r.FormValue("title")
	var content := r.FormValue("content")
}

func main() {
	fmt.Println("Listening on portL: 5000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":5000", nil)
}
