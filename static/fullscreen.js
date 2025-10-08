// 🖥️ Système de plein écran global pour Power4

function enterFullscreen() {
  const elem = document.documentElement;
  
  if (elem.requestFullscreen) {
    elem.requestFullscreen().catch(err => {
      console.log('Erreur plein écran:', err);
    });
  } else if (elem.webkitRequestFullscreen) { /* Safari */
    elem.webkitRequestFullscreen();
  } else if (elem.msRequestFullscreen) { /* IE11 */
    elem.msRequestFullscreen();
  }
}

function exitFullscreen() {
  if (document.exitFullscreen) {
    document.exitFullscreen();
  } else if (document.webkitExitFullscreen) { /* Safari */
    document.webkitExitFullscreen();
  } else if (document.msExitFullscreen) { /* IE11 */
    document.msExitFullscreen();
  }
}

// Vérifier si on est en plein écran
function isFullscreen() {
  return !!(document.fullscreenElement || document.webkitFullscreenElement || 
            document.msFullscreenElement);
}

// Créer le bouton de plein écran (si pas déjà en plein écran)
function createFullscreenButton() {
  if (isFullscreen()) return;
  
  // Ne pas créer si le bouton existe déjà
  if (document.getElementById('floating-fullscreen-btn')) return;
  
  const btn = document.createElement('button');
  btn.id = 'floating-fullscreen-btn';
  btn.className = 'floating-fullscreen-btn';
  btn.innerHTML = '🖥️';
  btn.title = 'Activer le plein écran';
  
  btn.addEventListener('click', () => {
    enterFullscreen();
    btn.style.display = 'none';
  });
  
  document.body.appendChild(btn);
}

// Créer l'indicateur de plein écran
function createFullscreenIndicator() {
  // Vérifier si l'indicateur existe déjà
  if (document.getElementById('fullscreen-hint')) return;
  
  const hint = document.createElement('div');
  hint.id = 'fullscreen-hint';
  hint.className = 'fullscreen-indicator';
  hint.innerHTML = '<span>Appuyez sur <kbd>Échap</kbd> pour quitter le plein écran</span>';
  
  // Afficher seulement si en plein écran
  hint.style.display = isFullscreen() ? 'flex' : 'none';
  
  document.body.appendChild(hint);
}

// Mettre à jour l'affichage en fonction de l'état du plein écran
function updateFullscreenUI() {
  const hint = document.getElementById('fullscreen-hint');
  const btn = document.getElementById('floating-fullscreen-btn');
  
  if (isFullscreen()) {
    // En plein écran : afficher l'indicateur, masquer le bouton
    if (hint) hint.style.display = 'flex';
    if (btn) btn.style.display = 'none';
  } else {
    // Pas en plein écran : masquer l'indicateur, afficher le bouton
    if (hint) hint.style.display = 'none';
    if (btn) btn.style.display = 'flex';
  }
}

// Gestionnaire pour la touche Échap (le navigateur gère déjà la sortie, on met juste à jour l'UI)
document.addEventListener('keydown', (e) => {
  if (e.key === 'Escape' || e.key === 'Esc') {
    // Le navigateur quitte automatiquement le plein écran
    // On met à jour l'UI après un court délai
    setTimeout(updateFullscreenUI, 100);
  }
});

// Détecter les changements d'état du plein écran
document.addEventListener('fullscreenchange', updateFullscreenUI);
document.addEventListener('webkitfullscreenchange', updateFullscreenUI);
document.addEventListener('mozfullscreenchange', updateFullscreenUI);
document.addEventListener('msfullscreenchange', updateFullscreenUI);

// Initialisation au chargement de la page
window.addEventListener('load', () => {
  // Créer l'indicateur
  createFullscreenIndicator();
  
  // Créer le bouton flottant seulement si pas en plein écran
  if (!isFullscreen()) {
    createFullscreenButton();
  }
  
  // Mettre à jour l'affichage
  updateFullscreenUI();
});

// Vérifier l'état au chargement de la page (pour les navigations)
window.addEventListener('pageshow', () => {
  updateFullscreenUI();
});
