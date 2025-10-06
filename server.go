package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

// Structure Student
type Student struct {
	Name  string
	Age   int
	Quote string
	Hobby string
}

// Fonction utilitaire pour afficher un template
func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmplPath := filepath.Join("templates", tmplName)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur de lecture du template", http.StatusInternalServerError)
		log.Println("ParseFiles error:", err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erreur d’exécution du template", http.StatusInternalServerError)
		log.Println("Execute error:", err)
	}
}

func main() {
	mux := http.NewServeMux()

	// Fichiers statiques
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	// Route d'accueil (GET)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", nil)
	})

	// Route idCard (GET et POST)
	mux.HandleFunc("/idcard", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Si on vient directement par GET
			renderTemplate(w, "idCard.html", nil)
			return
		}

		if r.Method == http.MethodPost {
			// Récupération des données du formulaire
			name := r.FormValue("name")
			ageStr := r.FormValue("age")
			quote := r.FormValue("quote")
			hobby := r.FormValue("hobby")

			// Conversion de l'âge (string -> int)
			age, _ := strconv.Atoi(ageStr)

			// Création de la structure Student
			student := Student{
				Name:  name,
				Age:   age,
				Quote: quote,
				Hobby: hobby,
			}

			// Envoi au template
			renderTemplate(w, "idCard.html", student)
		}
	})

	log.Println("✅ Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

