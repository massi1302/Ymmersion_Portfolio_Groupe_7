
document.addEventListener('DOMContentLoaded', () => {
    const challengeCards = document.querySelectorAll('.challenge-card');
    const spotLight = document.querySelector('.spot-light');

    // Parallax sur les cartes de défi
    challengeCards.forEach(card => {
        card.addEventListener('mousemove', (e) => {
            const { offsetX, offsetY } = e;
            const { width, height } = card.getBoundingClientRect();
            
            const xRotation = -30 * ((offsetY - height / 2) / height);
            const yRotation = 30 * ((offsetX - width / 2) / width);

            card.style.transform = `perspective(500px) rotateX(${xRotation}deg) rotateY(${yRotation}deg)`;
        });

        card.addEventListener('mouseleave', () => {
            card.style.transform = 'perspective(500px) rotateX(0) rotateY(0)';
        });
    });

    // Préchargement et lazy loading des images
    const imageObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const img = entry.target;
                img.src = img.dataset.src;
                imageObserver.unobserve(img);
            }
        });
    });

    challengeCards.forEach(card => {
        const img = card.querySelector('img');
        img.dataset.src = img.src;
        img.src = ''; // Placeholder vide
        imageObserver.observe(img);
    });
});