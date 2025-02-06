document.addEventListener('DOMContentLoaded', () => {
  const teamSection = document.querySelector('.team-section');
  const teamCards = document.querySelectorAll('.team-card');
  let currentIndex = 0;
  const totalCards = teamCards.length;

  // Ajout des classes initiales
  function initializeCards() {
    teamCards.forEach((card, index) => {
      if (index === currentIndex) {
        card.classList.add('active');
      } else if (index === (currentIndex + 1) % totalCards) {
        card.classList.add('next');
      } else if (index === (currentIndex - 1 + totalCards) % totalCards) {
        card.classList.add('prev');
      } else {
        card.classList.add('hidden');
      }
    });
  }

  // Gestion du déplacement des cartes
  function moveCards(direction) {
    // Supprimer toutes les classes actuelles
    teamCards.forEach(card => {
      card.classList.remove('active', 'prev', 'next', 'hidden');
    });

    // Mettre à jour l'index
    if (direction === 'next') {
      currentIndex = (currentIndex + 1) % totalCards;
    } else {
      currentIndex = (currentIndex - 1 + totalCards) % totalCards;
    }

    // Appliquer les nouvelles classes
    teamCards.forEach((card, index) => {
      if (index === currentIndex) {
        card.classList.add('active');
      } else if (index === (currentIndex + 1) % totalCards) {
        card.classList.add('next');
      } else if (index === (currentIndex - 1 + totalCards) % totalCards) {
        card.classList.add('prev');
      } else {
        card.classList.add('hidden');
      }
    });
  }

  // Création des boutons de navigation
  const navigationButtons = document.createElement('div');
  navigationButtons.className = 'navigation-buttons';
  navigationButtons.innerHTML = `
    <button class="prev-btn">←</button>
    <button class="next-btn">→</button>
  `;
  teamSection.appendChild(navigationButtons);

  // Event listeners pour les boutons
  const prevBtn = navigationButtons.querySelector('.prev-btn');
  const nextBtn = navigationButtons.querySelector('.next-btn');

  prevBtn.addEventListener('click', () => moveCards('prev'));
  nextBtn.addEventListener('click', () => moveCards('next'));

  // Initialisation
  initializeCards();
});