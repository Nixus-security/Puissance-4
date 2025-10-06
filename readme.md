# 🎮 Power4 Web - Puissance 4 en ligne

Un jeu de Puissance 4 (Connect Four) développé en Go avec une interface web moderne et animée.

## 📋 Fonctionnalités

### ✅ Fonctionnalités principales
- ✨ Jeu complet de Puissance 4 pour 2 joueurs locaux
- 🎯 3 niveaux de difficulté :
  - **Facile** : Grille 6x7 (classique)
  - **Normal** : Grille 6x9 (plus de colonnes)
  - **Difficile** : Grille 7x8 (plus de lignes et colonnes)
- 👥 Personnalisation des noms de joueurs
- 🏆 Détection automatique de victoire (horizontal, vertical, diagonal)
- 🤝 Détection d'égalité
- 🔄 Bouton pour rejouer
- 📱 Interface responsive

### 🎁 Fonctionnalités bonus
- 🌀 **Gravité inversée** : Tous les 5 tours, la gravité s'inverse !
  - Les pions tombent du bas vers le haut
  - Changement visuel du fond d'écran (gradient rose)
  - Boutons de colonnes qui se retournent
  - Animation de transition fluide
- 🎨 **Design moderne** :
  - Animations fluides des pions qui tombent
  - Effets de surbrillance sur les boutons
  - Confettis lors de la victoire
  - Transitions de couleurs douces
  - Indicateur visuel du joueur actif

## 🚀 Installation et lancement

### Prérequis
- Go 1.16 ou supérieur

### Structure du projet
```
Puissance-4/
├── go.mod
├── server.go
├── README.md
├── templates/
│   ├── index.html
│   ├── game.html
│   ├── win.html
│   └── draw.html
└── static/
    ├── style.css
    └── game.css
```

### Étapes d'installation

1. **Cloner le repository**
```bash
git clone https://github.com/Nixus-security/Puissance-4.git
cd Puissance-4
```

2. **Initialiser le module Go**
```bash
go mod init power4
```

3. **Créer les dossiers nécessaires**
```bash
mkdir -p templates static
```

4. **Placer les fichiers**
- Mettre `server.go` à la racine
- Mettre tous les fichiers `.html` dans `templates/`
- Mettre tous les fichiers `.css` dans `static/`

5. **Lancer le serveur**
```bash
go run server.go
```

6. **Accéder au jeu**
Ouvrir votre navigateur à l'adresse : **http://localhost:8080**

## 🎮 Comment jouer

1. **Page d'accueil**
   - Entrez les noms des deux joueurs
   - Choisissez la difficulté (Easy, Normal ou Hard)
   - Cliquez sur "Commencer la partie"

2. **Pendant le jeu**
   - Le joueur actif est indiqué en haut (carte surbrillante)
   - Cliquez sur un bouton de colonne (▼) pour placer votre pion
   - Le pion tombe automatiquement dans la colonne choisie
   - **Attention** : Tous les 5 tours, la gravité s'inverse !
   - Le premier à aligner 4 pions (horizontal, vertical ou diagonal) gagne

3. **Fin de partie**
   - Écran de victoire avec le nom du gagnant
   - Ou écran d'égalité si la grille est pleine
   - Options : Rejouer ou retourner au menu

## 🎨 Caractéristiques techniques

### Backend (Go)
- Serveur HTTP avec `net/http`
- Templates HTML dynamiques avec `html/template`
- Gestion d'état en mémoire (pas de base de données)
- Logique de jeu complète :
  - Placement des pions
  - Alternance des joueurs
  - Détection de victoire (4 directions)
  - Détection d'égalité
  - Gravité inversée

### Frontend
- **HTML5** : Structure sémantique
- **CSS3** : 
  - Animations CSS natives
  - Transitions fluides
  - Gradients modernes
  - Responsive design
- **Pas de framework JS** : HTML pur avec formulaires POST

### Routes HTTP
- `GET /` : Page d'accueil (configuration)
- `POST /` : Création d'une nouvelle partie
- `GET /game` : Affichage du plateau de jeu
- `POST /play` : Jouer un coup (envoyer une colonne)
- `GET /win` : Écran de victoire
- `GET /draw` : Écran d'égalité
- `GET /restart` : Recommencer une partie

## 🌟 Détails de la gravité inversée

La fonctionnalité bonus "gravité inversée" ajoute une dimension stratégique :

- **Déclenchement** : Automatique tous les 5 tours (tour 5, 10, 15, etc.)
- **Effets visuels** :
  - Le fond passe d'un dégradé violet à un dégradé rose/rouge
  - Un message "⚠️ GRAVITÉ INVERSÉE ⚠️" s'affiche
  - Les boutons de colonnes se retournent (180°)
  - Animation de secousse du plateau
- **Gameplay** :
  - Les pions tombent du bas vers le haut
  - Nécessite de repenser sa stratégie
  - Retour à la normale au tour suivant (tous les 5 tours)

## 🐛 Résolution de problèmes

### Le serveur ne démarre pas
```bash
# Vérifier que le port 8080 n'est pas utilisé
lsof -i :8080

# Ou lancer sur un autre port (modifier server.go)
http.ListenAndServe(":3000", mux)
```

### Les templates ne se chargent pas
- Vérifier que le dossier `templates/` existe
- Vérifier que tous les fichiers `.html` sont présents
- Vérifier les permissions des fichiers

### Les styles CSS ne s'appliquent pas
- Vérifier que le dossier `static/` existe
- Vérifier que les fichiers `.css` sont présents
- Vider le cache du navigateur (Ctrl+F5)

## 📚 Bonnes pratiques respectées

✅ Utilisation uniquement des packages standard Go  
✅ Séparation claire de la logique (backend) et présentation (frontend)  
✅ Code commenté et structuré  
✅ Gestion des erreurs  
✅ Templates Go pour génération dynamique  
✅ Routes RESTful claires  
✅ Design responsive  
✅ Animations performantes (CSS uniquement)  

## 🎯 Améliorations possibles

- 🤖 Mode solo contre l'ordinateur (IA)
- 💾 Sauvegarde des scores
- 🌐 Mode multijoueur en ligne (WebSockets)
- 🎵 Effets sonores
- 📊 Statistiques de parties
- 🏅 Système de classement
- ⏱️ Timer par tour
- 🎨 Thèmes de couleurs personnalisables

## 📄 Licence

Projet éducatif - Libre d'utilisation

## 👥 Auteurs

Développé dans le cadre du projet Power4 Web

---

**Bon jeu ! 🎮🎉**