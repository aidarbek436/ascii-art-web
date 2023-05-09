package servers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	asciiart "student/ascii-art-web/ascii-art"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "File not found: index.html", http.StatusInternalServerError)
		return
	}
	if err = t.Execute(w, nil); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func AsciiPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "File not found: index.html", http.StatusInternalServerError)
		return
	}
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("ascii-art-input")
	banner := r.FormValue("font_types")
	res := asciiart.Converter(text, banner)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	Array_Body := DivideRequestBody(string(b))
	if RequestBodyErrorcheck(Array_Body) == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = t.Execute(w, res); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func DivideRequestBody(b string) []string {
	res := []string{}
	arr := []rune(b)
	text := ""
RestartLoop:
	for i := 0; i < len(arr); i++ {
		if arr[i] == '&' {
			text = string(arr[:i])
			res = append(res, text)
			arr = arr[i+1:]
			text = ""
			goto RestartLoop
		}
	}
	if arr != nil {
		res = append(res, string(arr))
	}
	return res
}

func RequestBodyErrorcheck(arr []string) bool {
	if len(arr) > 2 {
		return true
	} else if len(arr) < 2 {
		return true
	}
	input := []rune(string(arr[0]))
	banner := []rune(string(arr[1]))
	if len(input) < 15 || string(input[:16]) != "ascii-art-input" {
		return true
	}
	if len(banner) < 10 || string(banner[:11]) != "font_types" {
		return true
	}
	return false
}
