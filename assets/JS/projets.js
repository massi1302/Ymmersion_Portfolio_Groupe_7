document.addEventListener('DOMContentLoaded', () => {
    const challengeCards = document.querySelectorAll('.challenge-card');
    const spotLight = document.querySelector('.spot-light');

    // Reveal animation for spot light
    spotLight.style.opacity = '0';
    spotLight.style.transform = 'translateY(20px)';
    setTimeout(() => {
        spotLight.style.transition = 'all 0.8s ease-out';
        spotLight.style.opacity = '1';
        spotLight.style.transform = 'translateY(0)';
    }, 300);

    // Interactive hover effects for challenge cards
    challengeCards.forEach(card => {
        const img = card.querySelector('img');
        const overlay = card.querySelector('.overlay');
        const title = card.querySelector('.challenge-title');

        card.addEventListener('mouseenter', () => {
            img.style.transform = 'scale(1.1)';
            overlay.style.backgroundColor = 'rgba(0, 0, 0, 0.6)';
            title.style.transform = 'translateY(0)';
            title.style.opacity = '1';
        });

        card.addEventListener('mouseleave', () => {
            img.style.transform = 'scale(1)';
            overlay.style.backgroundColor = 'transparent';
            title.style.transform = 'translateY(20px)';
            title.style.opacity = '0';
        });
    });

    // Preload images for smoother experience
    challengeCards.forEach(card => {
        const img = card.querySelector('img');
        const tempImage = new Image();
        tempImage.src = img.src;
    });
});