let targetNumber = Math.floor(Math.random() * 100) + 1;
let attempts = 0;

function checkGuess() {
    const guessInput = document.getElementById('guess-input');
    const resultDiv = document.getElementById('game-result');
    const userGuess = parseInt(guessInput.value);
    
    attempts++;

    if (isNaN(userGuess) || userGuess < 1 || userGuess > 100) {
        resultDiv.textContent = "Veuillez entrer un nombre entre 1 et 100.";
        return;
    }

    if (userGuess > targetNumber) {
        resultDiv.textContent = "Plus bas !";
    } else if (userGuess < targetNumber) {
        resultDiv.textContent = "Plus haut !";
    } else {
        resultDiv.textContent = `Bravo ! Vous avez deviné en ${attempts} tentatives.`;
        targetNumber = Math.floor(Math.random() * 100) + 1;
        attempts = 0;
    }

    guessInput.value = '';
}