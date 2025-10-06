package main

import (
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
func NewGame(player1, player2, difficulty string) *Game {
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
func (g *Game) PlayMove(column int) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if column < 0 || column >= g.Columns {
		return false
	}

	// Gravité normale (de haut en bas)
	if !g.GravityInverted {
		for row := g.Rows - 1; row >= 0; row-- {
			if g.Board[row][column] == 0 {
				g.Board[row][column] = g.CurrentPlayer
				g.TurnCount++
				
				// Vérifier la victoire
				if g.checkWin(row, column) {
					g.Winner = g.CurrentPlayer
					return true
				}
				
				// Vérifier l'égalité
				if g.checkDraw() {
					g.IsDraw = true
					return true
				}
				
				// Changer de joueur
				if g.CurrentPlayer == 1 {
					g.CurrentPlayer = 2
				} else {
					g.CurrentPlayer = 1
				}
				
				// Inverser la gravité tous les 5 tours (bonus)
				if g.TurnCount%5 == 0 && g.TurnCount > 0 {
					g.GravityInverted = !g.GravityInverted
				}
				
				return true
			}
		}
	} else {
		// Gravité inversée (de bas en haut)
		for row := 0; row < g.Rows; row++ {
			if g.Board[row][column] == 0 {
				g.Board[row][column] = g.CurrentPlayer
				g.TurnCount++
				
				if g.checkWin(row, column) {
					g.Winner = g.CurrentPlayer
					return true
				}
				
				if g.checkDraw() {
					g.IsDraw = true
					return true
				}
				
				if g.CurrentPlayer == 1 {
					g.CurrentPlayer = 2
				} else {
					g.CurrentPlayer = 1
				}
				
				if g.TurnCount%5 == 0 && g.TurnCount > 0 {
					g.GravityInverted = !g.GravityInverted
				}
				
				return true
			}
		}
	}
	
	return false
}

// Vérifie si le joueur a gagné
func (g *Game) checkWin(row, col int) bool {
	player := g.Board[row][col]
	
	// Vérifier horizontal
	count := 1
	// Gauche
	for c := col - 1; c >= 0 && g.Board[row][c] == player; c-- {
		count++
	}
	// Droite
	for c := col + 1; c < g.Columns && g.Board[row][c] == player; c++ {
		count++
	}
	if count >= 4 {
		return true
	}
	
	// Vérifier vertical
	count = 1
	// Haut
	for r := row - 1; r >= 0 && g.Board[r][col] == player; r-- {
		count++
	}
	// Bas
	for r := row + 1; r < g.Rows && g.Board[r][col] == player; r++ {
		count++
	}
	if count >= 4 {
		return true
	}
	
	// Vérifier diagonale \
	count = 1
	// Haut-gauche
	for r, c := row-1, col-1; r >= 0 && c >= 0 && g.Board[r][c] == player; r, c = r-1, c-1 {
		count++
	}
	// Bas-droite
	for r, c := row+1, col+1; r < g.Rows && c < g.Columns && g.Board[r][c] == player; r, c = r+1, c+1 {
		count++
	}
	if count >= 4 {
		return true
	}
	
	// Vérifier diagonale /
	count = 1
	// Haut-droite
	for r, c := row-1, col+1; r >= 0 && c < g.Columns && g.Board[r][c] == player; r, c = r-1, c+1 {
		count++
	}
	// Bas-gauche
	for r, c := row+1, col-1; r < g.Rows && c >= 0 && g.Board[r][c] == player; r, c = r+1, c-1 {
		count++
	}
	if count >= 4 {
		return true
	}
	
	return false
}

// Vérifie si la grille est pleine (égalité)
func (g *Game) checkDraw() bool {
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Columns; col++ {
			if g.Board[row][col] == 0 {
				return false
			}
		}
	}
	return true
}

// Fonction utilitaire pour afficher un template
func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmplPath := filepath.Join("templates", tmplName)
	
	// Créer un FuncMap pour les fonctions personnalisées
	funcMap := template.FuncMap{
		"until": func(count int) []int {
			result := make([]int, count)
			for i := 0; i < count; i++ {
				result[i] = i
			}
			return result
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

	// Route d'accueil - Configuration du jeu
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			renderTemplate(w, "index.html", difficulties)
			return
		}
		
		if r.Method == http.MethodPost {
			player1 := r.FormValue("player1")
			player2 := r.FormValue("player2")
			difficulty := r.FormValue("difficulty")
			
			if player1 == "" || player2 == "" || difficulty == "" {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			
			currentGame = NewGame(player1, player2, difficulty)
			http.Redirect(w, r, "/game", http.StatusSeeOther)
		}
	})

	// Route du jeu
	mux.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		if currentGame == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		
		renderTemplate(w, "game.html", currentGame)
	})

	// Route pour jouer un coup
	mux.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || currentGame == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		
		columnStr := r.FormValue("column")
		column, err := strconv.Atoi(columnStr)
		if err != nil {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		}
		
		currentGame.PlayMove(column)
		
		// Rediriger vers la page appropriée
		if currentGame.Winner != 0 {
			http.Redirect(w, r, "/win", http.StatusSeeOther)
		} else if currentGame.IsDraw {
			http.Redirect(w, r, "/draw", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/game", http.StatusSeeOther)
		}
	})

	// Route victoire
	mux.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		if currentGame == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		renderTemplate(w, "win.html", currentGame)
	})

	// Route égalité
	mux.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		if currentGame == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		renderTemplate(w, "draw.html", currentGame)
	})

	// Route pour recommencer
	mux.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		currentGame = nil
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Println("✅ Serveur Power4 démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
