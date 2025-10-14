package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
)

// GameDifficulty représente les niveaux de difficulté
type GameDifficulty struct {
	Name    string
	Rows    int
	Columns int
}

var difficulties = map[string]GameDifficulty{
	"easy":   {Name: "Easy", Rows: 6, Columns: 7},
	"normal": {Name: "Normal", Rows: 6, Columns: 9},
	"hard":   {Name: "Hard", Rows: 7, Columns: 8},
}

// Game représente l'état du jeu
type Game struct {
	Board           [][]int
	CurrentPlayer   int
	Player1Name     string
	Player2Name     string
	Player1Photo    template.URL // ✅ Changé de string à template.URL
	Player2Photo    template.URL // ✅ Changé de string à template.URL
	Rows            int
	Columns         int
	Winner          int
	IsDraw          bool
	TurnCount       int
	GravityInverted bool
	Difficulty      string
	mu              sync.Mutex
}

var currentGame *Game

// Initialise une nouvelle partie
func NewGame(player1, player2, difficulty, photo1, photo2 string) *Game {
	diff := difficulties[difficulty]
	board := make([][]int, diff.Rows)
	for i := range board {
		board[i] = make([]int, diff.Columns)
	}

	return &Game{
		Board:           board,
		CurrentPlayer:   1,
		Player1Name:     player1,
		Player2Name:     player2,
		Player1Photo:    template.URL(photo1), // ✅ Conversion en template.URL
		Player2Photo:    template.URL(photo2), // ✅ Conversion en template.URL
		Rows:            diff.Rows,
		Columns:         diff.Columns,
		Winner:          0,
		IsDraw:          false,
		TurnCount:       0,
		GravityInverted: false,
		Difficulty:      difficulty,
	}
}

// Place un pion dans une colonne
func (g *Game) PlayMove(column int) (int, int, bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if column < 0 || column >= g.Columns {
		return -1, -1, false
	}

	if !g.GravityInverted {
		for row := g.Rows - 1; row >= 0; row-- {
			if g.Board[row][column] == 0 {
				g.Board[row][column] = g.CurrentPlayer
				g.TurnCount++
				if g.checkWin(row, column) {
					g.Winner = g.CurrentPlayer
				}
				if g.checkDraw() {
					g.IsDraw = true
				}
				if g.Winner == 0 && !g.IsDraw {
					g.CurrentPlayer = 3 - g.CurrentPlayer
				}
				if g.TurnCount%5 == 0 && g.TurnCount > 0 {
					g.GravityInverted = !g.GravityInverted
				}
				return row, column, true
			}
		}
	} else {
		for row := 0; row < g.Rows; row++ {
			if g.Board[row][column] == 0 {
				g.Board[row][column] = g.CurrentPlayer
				g.TurnCount++
				if g.checkWin(row, column) {
					g.Winner = g.CurrentPlayer
				}
				if g.checkDraw() {
					g.IsDraw = true
				}
				if g.Winner == 0 && !g.IsDraw {
					g.CurrentPlayer = 3 - g.CurrentPlayer
				}
				if g.TurnCount%5 == 0 && g.TurnCount > 0 {
					g.GravityInverted = !g.GravityInverted
				}
				return row, column, true
			}
		}
	}

	return -1, -1, false
}

// Vérifie la victoire
func (g *Game) checkWin(row, col int) bool {
	player := g.Board[row][col]
	count := 1

	// Horizontal
	for c := col - 1; c >= 0 && g.Board[row][c] == player; c-- {
		count++
	}
	for c := col + 1; c < g.Columns && g.Board[row][c] == player; c++ {
		count++
	}
	if count >= 4 {
		return true
	}

	// Vertical
	count = 1
	for r := row - 1; r >= 0 && g.Board[r][col] == player; r-- {
		count++
	}
	for r := row + 1; r < g.Rows && g.Board[r][col] == player; r++ {
		count++
	}
	if count >= 4 {
		return true
	}

	// Diagonales
	count = 1
	for r, c := row-1, col-1; r >= 0 && c >= 0 && g.Board[r][c] == player; r, c = r-1, c-1 {
		count++
	}
	for r, c := row+1, col+1; r < g.Rows && c < g.Columns && g.Board[r][c] == player; r, c = r+1, c+1 {
		count++
	}
	if count >= 4 {
		return true
	}

	count = 1
	for r, c := row-1, col+1; r >= 0 && c < g.Columns && g.Board[r][c] == player; r, c = r-1, c+1 {
		count++
	}
	for r, c := row+1, col-1; r < g.Rows && c >= 0 && g.Board[r][c] == player; r, c = r+1, c-1 {
		count++
	}
	return count >= 4
}

// Vérifie l'égalité
func (g *Game) checkDraw() bool {
	for _, row := range g.Board {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}
	return true
}

// Template utilitaire
func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmplPath := filepath.Join("templates", tmplName)
	funcMap := template.FuncMap{
		"until": func(count int) []int {
			result := make([]int, count)
			for i := 0; i < count; i++ {
				result[i] = i
			}
			return result
		},
		"len": func(s string) int {
			return len(s)
		},
	}

	tmpl, err := template.New(tmplName).Funcs(funcMap).ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur de lecture du template", http.StatusInternalServerError)
		log.Println("ParseFiles error:", err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erreur d'exécution du template", http.StatusInternalServerError)
		log.Println("Execute error:", err)
	}
}

func main() {
	mux := http.NewServeMux()

	// Fichiers statiques
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	// Page splash
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		renderTemplate(w, "splash.html", nil)
	})

	// Menu - saisie des noms et difficulté
	mux.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			renderTemplate(w, "index.html", difficulties)
			return
		}

		if r.Method == http.MethodPost {
			player1 := r.FormValue("player1")
			player2 := r.FormValue("player2")
			difficulty := r.FormValue("difficulty")

			if player1 == "" || player2 == "" || difficulty == "" {
				http.Redirect(w, r, "/menu", http.StatusSeeOther)
				return
			}

			// Rediriger vers la page photo avec les paramètres
			http.Redirect(w, r, "/photo?player1="+player1+"&player2="+player2+"&difficulty="+difficulty, http.StatusSeeOther)
		}
	})

	// Page de capture photo
	mux.HandleFunc("/photo", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "photo.html", nil)
	})

	// Créer le jeu avec les photos
	mux.HandleFunc("/create-game", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Lire le JSON
		var gameData struct {
			Player1    string `json:"player1"`
			Player2    string `json:"player2"`
			Difficulty string `json:"difficulty"`
			Photo1     string `json:"photo1"`
			Photo2     string `json:"photo2"`
		}

		err := json.NewDecoder(r.Body).Decode(&gameData)
		if err != nil {
			log.Println("❌ Erreur décodage JSON:", err)
			http.Error(w, "Erreur parsing JSON", http.StatusBadRequest)
			return
		}

		// Debug
		log.Println("📥 Données reçues:")
		log.Println("  Player1:", gameData.Player1)
		log.Println("  Player2:", gameData.Player2)
		log.Println("  Difficulty:", gameData.Difficulty)
		log.Println("  Photo1 length:", len(gameData.Photo1))
		log.Println("  Photo2 length:", len(gameData.Photo2))

		if gameData.Player1 == "" || gameData.Player2 == "" || gameData.Difficulty == "" {
			log.Println("❌ Paramètres manquants")
			http.Error(w, "Paramètres manquants", http.StatusBadRequest)
			return
		}

		currentGame = NewGame(gameData.Player1, gameData.Player2, gameData.Difficulty, gameData.Photo1, gameData.Photo2)

		log.Println("✅ Game créé avec photos:")
		log.Println("  Photo1 stockée:", len(currentGame.Player1Photo), "caractères")
		log.Println("  Photo2 stockée:", len(currentGame.Player2Photo), "caractères")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Page du jeu
	mux.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		if currentGame == nil {
			http.Redirect(w, r, "/menu", http.StatusSeeOther)
			return
		}
		log.Println("📺 Affichage game.html")
		log.Println("  Photo1:", len(currentGame.Player1Photo), "caractères")
		log.Println("  Photo2:", len(currentGame.Player2Photo), "caractères")
		renderTemplate(w, "game.html", currentGame)
	})

	// Jouer un coup
	mux.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || currentGame == nil {
			http.Redirect(w, r, "/menu", http.StatusSeeOther)
			return
		}

		// Cas JSON (JavaScript Fetch)
		if r.Header.Get("Content-Type") == "application/json" {
			var payload struct {
				Col int `json:"col"`
			}
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				http.Error(w, `{"error":"Requête JSON invalide"}`, http.StatusBadRequest)
				return
			}

			row, col, ok := currentGame.PlayMove(payload.Col)
			if !ok {
				http.Error(w, `{"error":"Coup invalide"}`, http.StatusBadRequest)
				return
			}

			resp := map[string]interface{}{
				"grille":          currentGame.Board,
				"derniereLigne":   row,
				"derniereCol":     col,
				"finPartie":       currentGame.Winner != 0 || currentGame.IsDraw,
				"joueurActuel":    currentGame.CurrentPlayer,
				"gravityInverted": currentGame.GravityInverted,
			}

			if currentGame.Winner != 0 {
				resp["message"] = "Victoire de " + currentGame.PlayerName(currentGame.Winner)
			} else if currentGame.IsDraw {
				resp["message"] = "Match nul !"
			} else {
				resp["message"] = "À " + currentGame.PlayerName(currentGame.CurrentPlayer) + " de jouer"
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Cas HTML (formulaire classique)
		columnStr := r.FormValue("column")
		column, err := strconv.Atoi(columnStr)
		if err != nil {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		}

		currentGame.PlayMove(column)

		if currentGame.Winner != 0 {
			http.Redirect(w, r, "/win", http.StatusSeeOther)
		} else if currentGame.IsDraw {
			http.Redirect(w, r, "/draw", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
		}
	})

	// Routes win/draw/restart
	mux.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		if currentGame == nil {
			http.Redirect(w, r, "/menu", http.StatusSeeOther)
			return
		}
		log.Println("🏆 Affichage win.html")
		log.Println("  Winner:", currentGame.Winner)
		if currentGame.Winner == 1 {
			log.Println("  Photo gagnant:", len(currentGame.Player1Photo), "caractères")
		} else {
			log.Println("  Photo gagnant:", len(currentGame.Player2Photo), "caractères")
		}
		renderTemplate(w, "win.html", currentGame)
	})

	mux.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		if currentGame == nil {
			http.Redirect(w, r, "/menu", http.StatusSeeOther)
			return
		}
		renderTemplate(w, "draw.html", currentGame)
	})

	mux.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		currentGame = nil
		http.Redirect(w, r, "/menu", http.StatusSeeOther)
	})

	log.Println("✅ Serveur Power4 démarré sur http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}

// Méthode utilitaire
func (g *Game) PlayerName(id int) string {
	if id == 1 {
		return g.Player1Name
	} else if id == 2 {
		return g.Player2Name
	}
	return "?"
}

// Méthode pour obtenir la photo d'un joueur
func (g *Game) PlayerPhoto(id int) template.URL {
	if id == 1 {
		return g.Player1Photo
	} else if id == 2 {
		return g.Player2Photo
	}
	return ""
}