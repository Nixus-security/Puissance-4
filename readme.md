# 🎮 Power4 Web

> Jeu de Puissance 4 moderne développé en Go avec interface web interactive

Un Connect Four (Puissance 4) jouable en ligne avec capture photo des joueurs, animations fluides et un système de gravité inversée unique.

---

## ✨ Fonctionnalités

### Core
- 🎯 **3 niveaux de difficulté** : Facile (6x7), Normal (6x9), Difficile (7x8)
- 👥 **Mode 2 joueurs** local avec noms personnalisés
- 📸 **Capture photo** via webcam pour afficher les joueurs
- 🏆 **Détection automatique** de victoire et égalité
- 🎨 **Interface moderne** avec animations CSS fluides
- 📱 **Design responsive** pour tous les écrans

### Bonus
- 🌀 **Gravité inversée** tous les 5 tours
- 🖥️ **Mode plein écran** automatique
- ✨ **Animations** de chute de pions réalistes
- 🎊 **Effets visuels** (bulles, confettis, transitions)

---

## 🚀 Installation

### Prérequis
- Go 1.25+ ([Télécharger](https://go.dev/dl/))
- Navigateur moderne avec support webcam

### Lancement rapide

```bash
# 1. Cloner le projet
git clone https://github.com/Nixus-security/Puissance-4.git
cd Puissance-4

# 2. Lancer le serveur
go run server.go

# 3. Ouvrir dans le navigateur
open http://localhost:8000
```

**C'est tout !** 🎉 Le serveur Go gère tout automatiquement.

---

## 🎮 Comment jouer

1. **Écran de démarrage** → Cliquez pour activer le plein écran
2. **Menu** → Entrez les noms et choisissez la difficulté
3. **Photos** → Capturez les photos des 2 joueurs
4. **Jeu** → Cliquez sur une cellule pour placer un pion
5. **Victoire** → Premier à aligner 4 pions gagne !

> ⚠️ **Attention** : La gravité s'inverse tous les 5 tours !

---

## 🎨 Design System

Une documentation complète du design est accessible via :

```
http://localhost:8000/design
```

Cette page contient :
- 🎨 **Palette de couleurs** complète avec codes hex
- ✍️ **Typographie** (Georgia pour titres, Poppins pour corps)
- 🧩 **Composants** (boutons, jetons, cartes glassmorphism)
- ✨ **Animations** (float, pulse, rise, drop) avec démos live
- 📱 **Tous les écrans** du jeu avec leurs routes
- 🛠️ **Stack technique** documentée

Parfait pour comprendre les choix de design et les réutiliser !

---

## 📁 Structure du projet

```
Puissance-4/
├── server.go           # Backend Go (serveur HTTP + logique)
├── go.mod              # Dépendances Go
│
├── templates/          # Templates HTML
│   ├── splash.html     # Écran de démarrage
│   ├── index.html      # Menu principal
│   ├── photo.html      # Capture photos
│   ├── game.html       # Plateau de jeu
│   ├── win.html        # Écran victoire
│   ├── draw.html       # Écran égalité
│   └── design.html     # 🆕 Design System & Documentation
│
└── static/             # Assets frontend
    ├── style.css       # Styles de base
    ├── game.css        # Styles du jeu
    ├── fullscreen.css  # Styles plein écran
    ├── game.js         # Logique frontend
    └── fullscreen.js   # Gestion plein écran
```

---

## 🛠️ Technologies

**Backend**
- **Go 1.25** - Serveur HTTP natif (`net/http`)
- **Templates Go** - Génération HTML dynamique

**Frontend**
- **HTML5** - Structure sémantique
- **CSS3** - Animations natives, gradients, responsive
- **JavaScript Vanilla** - Interactions, webcam, animations

**Aucune dépendance externe** - Projet 100% autonome !

---

## 🎯 Fonctionnalités techniques

### Backend Go
```go
// Gestion des routes HTTP
GET  /           → Splash screen
GET  /menu       → Menu principal
POST /menu       → Création partie
GET  /photo      → Capture photos
POST /create-game → Initialisation
GET  /game       → Plateau de jeu
POST /play       → Jouer un coup
GET  /win        → Victoire
GET  /draw       → Égalité
GET  /restart    → Recommencer
GET  /design     → 🆕 Documentation design
```

### Logique du jeu
- ✅ Placement des pions avec gravité (normale/inversée)
- ✅ Vérification victoire (4 directions)
- ✅ Détection égalité (grille pleine)
- ✅ Alternance des joueurs automatique
- ✅ Gestion d'état en mémoire (pas de DB)

---

## ⚙️ Configuration

### Changer le port
```go
// Dans server.go, ligne finale
http.ListenAndServe(":8000", mux)  // Modifier 8000
```

### Personnaliser les difficultés
```go
// Dans server.go
var difficulties = map[string]GameDifficulty{
    "easy":   {Name: "Easy", Rows: 6, Columns: 7},
    "custom": {Name: "Custom", Rows: 8, Columns: 10}, // Ajouter
}
```

### Ajouter la route design (si pas encore fait)
```go
// Dans server.go, fonction main()
mux.HandleFunc("/design", func(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "design.html", nil)
})
```

---

## 🐛 Dépannage

**Le serveur ne démarre pas**
```bash
# Vérifier que le port n'est pas utilisé
lsof -i :8000
# Ou changer de port dans server.go
```

**La webcam ne fonctionne pas**
- Vérifiez les permissions du navigateur
- Essayez un autre navigateur (Chrome recommandé)
- Utilisez HTTPS en production

**Les styles ne s'affichent pas**
```bash
# Vider le cache du navigateur
Ctrl + F5  (Windows/Linux)
Cmd + Shift + R  (Mac)
```

---

## 📊 Statistiques du projet

- **Langage** : 100% Go (backend)
- **Lignes de code** : ~500 Go, ~1200 HTML/CSS/JS
- **Dépendances** : 0 (packages standard uniquement)
- **Performance** : <1ms par coup
- **Taille** : ~50 KB (binaire compilé)

---

## 🚀 Améliorations futures

- [ ] Mode solo contre IA (algorithme minimax)
- [ ] Sauvegarde des scores (SQLite)
- [ ] Mode multijoueur en ligne (WebSockets)
- [ ] Effets sonores
- [ ] Thèmes personnalisables
- [ ] Historique des parties
- [ ] Classement des joueurs
- [x] Documentation design complète

---

## 📄 Licence

Projet éducatif - Libre d'utilisation et de modification

---

## 👨‍💻 Auteur

**Anthony Nagul**  
Développé dans le cadre du projet Power4 Web

---

## 🙏 Crédits

- Design inspiré des interfaces modernes
- Animations CSS natives
- Go standard library
- Documentation design intégrée

---

<div align="center">

**⭐ N'hésitez pas à mettre une étoile si vous aimez le projet ! ⭐**

[Signaler un bug](https://github.com/Nixus-security/Puissance-4/issues) • [Proposer une fonctionnalité](https://github.com/Nixus-security/Puissance-4/issues)

</div>

---

## 📸 Aperçu

### Menu principal
Interface moderne avec sélection de difficulté animée

### Capture photo
Système de capture webcam avec compte à rebours

### Jeu
Plateau interactif avec affichage des photos des joueurs

### Gravité inversée
Effet visuel unique tous les 5 tours

### 🆕 Design System
Documentation complète des couleurs, composants et animations

---

**Bon jeu ! 🎮**