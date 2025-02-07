document.addEventListener('DOMContentLoaded', () => {
    const cards = document.querySelectorAll('.challenge-card');
    
    // Animation d'apparition des cartes
    cards.forEach((card, index) => {
        card.style.opacity = '0';
        card.style.animation = `cardAppear 0.8s cubic-bezier(0.23, 1, 0.32, 1) forwards ${index * 0.2}s`;
    });

    // Animation 3D au survol
    cards.forEach(card => {
        card.addEventListener('mousemove', (e) => {
            const rect = card.getBoundingClientRect();
            const x = e.clientX - rect.left;
            const y = e.clientY - rect.top;
            
            const xRotation = 20 * ((y - rect.height / 2) / rect.height);
            const yRotation = -20 * ((x - rect.width / 2) / rect.width);
            
            card.style.transform = `
                perspective(1000px)
                rotateX(${xRotation}deg)
                rotateY(${yRotation}deg)
                scale3d(1.05, 1.05, 1.05)
            `;

            // Make sure the image is visible
            const img = card.querySelector('img');
            if (img) {
                img.style.opacity = '1';
                if (!img.src && img.dataset.src) {
                    img.src = img.dataset.src;
                }
            }
        });

        card.addEventListener('mouseleave', () => {
            card.style.transform = 'perspective(1000px) rotateX(0) rotateY(0) scale3d(1, 1, 1)';
        });
    });

    // Effet de particules en arrière-plan
    const createParticles = () => {
        const particlesContainer = document.createElement('div');
        particlesContainer.className = 'particles-container';
        document.querySelector('main').prepend(particlesContainer);

        for (let i = 0; i < 50; i++) {
            const particle = document.createElement('div');
            particle.className = 'particle';
            particle.style.left = `${Math.random() * 100}%`;
            particle.style.animationDelay = `${Math.random() * 5}s`;
            particlesContainer.appendChild(particle);
        }
    };

    createParticles();

    // Simple lazy loading
    const images = document.querySelectorAll('.challenge-card img');
    images.forEach(img => {
        img.style.opacity = '0';
        img.style.transition = 'opacity 0.3s ease-in';
        
        if (img.complete) {
            img.style.opacity = '1';
        } else {
            img.onload = () => {
                img.style.opacity = '1';
            };
        }
    });
});