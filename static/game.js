document.addEventListener("DOMContentLoaded", () => {
  // Sélectionner tous les boutons de colonne
  const colButtons = document.querySelectorAll(".col-btn");

  colButtons.forEach(btn => {
    btn.addEventListener("click", (e) => {
      const col = e.currentTarget.getAttribute("data-col");
      jouerColonne(col, e.currentTarget);
    });
  });
});

async function jouerColonne(col, btnElement) {
  // Désactiver temporairement le bouton pour éviter le double clic
  btnElement.disabled = true;

  try {
    // Envoyer la requête au serveur (POST) pour jouer dans cette colonne
    const resp = await fetch("/play", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        // tu peux envoyer le nom du joueur, état, etc.
      },
      body: JSON.stringify({ col: col })
    });

    const data = await resp.json();

    if (!resp.ok) {
      console.error("Erreur du serveur :", data.message || resp.statusText);
      btnElement.disabled = false;
      return;
    }

    // data contient la nouvelle grille, état du jeu, etc.
    // mettre à jour l’interface visuellement
    miseAJourPlateau(data);

    if (data.finPartie) {
      afficherFin(data);
    } else {
      // réactiver les boutons (ou selon la logique du tour)
      remettreBoutonsActifs();
    }
  } catch (err) {
    console.error("Erreur fetch :", err);
    btnElement.disabled = false;
  }
}

// Exemple de fonction de mise à jour visuelle du plateau
function miseAJourPlateau(data) {
  // Supposons que data.grille soit un tableau 2D, data.derniereLigne, data.derniereCol
  const { grille, derniereLigne, derniereCol } = data;

  // Mettre à jour chaque cellule par son id ou attribut
  for (let i = 0; i < grille.length; i++) {
    for (let j = 0; j < grille[i].length; j++) {
      const cell = document.querySelector(`[data-row="${i}"][data-col="${j}"]`);
      if (cell) {
        cell.textContent = grille[i][j] || "";
        // ou appliquer une classe selon le joueur (rouge / jaune)
        cell.className = "cell"; 
        if (grille[i][j] === "R") cell.classList.add("rouge");
        else if (grille[i][j] === "J") cell.classList.add("jaune");
      }
    }
  }

  // Animation pour le dernier jeton
  if (derniereLigne != null && derniereCol != null) {
    const cell = document.querySelector(`[data-row="${derniereLigne}"][data-col="${derniereCol}"]`);
    if (cell) {
      cell.classList.add("anim-tombe");
      // retirer la classe après l’animation
      cell.addEventListener("animationend", () => {
        cell.classList.remove("anim-tombe");
      }, { once: true });
    }
  }
}

function afficherFin(data) {
  // Par exemple afficher une modal ou un overlay disant "Le joueur X gagne"
  const overlay = document.getElementById("overlay-fin");
  const texte = overlay.querySelector(".texte-fin");
  texte.textContent = data.message;  // "Joueur 1 gagne !" ou "Égalité"
  overlay.style.display = "block";
}

// Exemple pour réactiver les boutons
function remettreBoutonsActifs() {
  const colButtons = document.querySelectorAll(".col-btn");
  colButtons.forEach(b => b.disabled = false);
}
