document.addEventListener('DOMContentLoaded', () => {
    const teamSection = document.querySelector('.team-section');
    const teamCards = Array.from(teamSection.querySelectorAll('.team-card'));
    let currentIndex = 1; // Start with middle member
  
    // Add swipe navigation buttons to each card
    teamCards.forEach((card, index) => {
      const swipeControls = document.createElement('div');
      swipeControls.classList.add('member-swipe-controls');
      swipeControls.innerHTML = `
        <button class="swipe-left" data-index="${index}">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M15 18l-6-6 6-6"/>
          </svg>
        </button>
        <button class="swipe-right" data-index="${index}">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M9 18l6-6-6-6"/>
          </svg>
        </button>
      `;
      card.appendChild(swipeControls);
    });
  
    function updateMemberDisplay() {
      teamCards.forEach((card, index) => {
        card.classList.remove('active-member', 'side-left', 'side-right', 'hidden');
        
        if (index === currentIndex) {
          card.classList.add('active-member');
        } else if (index === currentIndex - 1) {
          card.classList.add('side-left');
        } else if (index === currentIndex + 1) {
          card.classList.add('side-right');
        } else {
          card.classList.add('hidden');
        }
      });
    }
  
    // Swipe event listeners
    teamSection.addEventListener('click', (e) => {
      const leftButton = e.target.closest('.swipe-left');
      const rightButton = e.target.closest('.swipe-right');
  
      if (leftButton) {
        currentIndex = Math.max(0, currentIndex - 1);
        updateMemberDisplay();
      }
  
      if (rightButton) {
        currentIndex = Math.min(teamCards.length - 1, currentIndex + 1);
        updateMemberDisplay();
      }
    });
  
    // Initial display
    updateMemberDisplay();
  });