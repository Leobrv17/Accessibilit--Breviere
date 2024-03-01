document.addEventListener('DOMContentLoaded', function () {
    const toggleButton = document.getElementById('accessibiliteToggle');
    const menuAccessibilite = document.querySelector('.menu_accessibilite');
    const inputs = menuAccessibilite.querySelectorAll('input');

    function toggleMenu() {
        const isHidden = menuAccessibilite.hasAttribute('hidden');
        if (isHidden) {
            menuAccessibilite.removeAttribute('hidden');
            menuAccessibilite.style.display = 'block';
            inputs.forEach(input => {
                input.setAttribute('tabindex', '0'); // Rend les inputs tabulables
                input.nextElementSibling.setAttribute('tabindex', '0'); // Assure que les labels sont également tabulables
            });
        } else {
            menuAccessibilite.setAttribute('hidden', 'hidden');
            menuAccessibilite.style.display = 'none';
            inputs.forEach(input => {
                input.setAttribute('tabindex', '-1'); // Retire les inputs de la séquence de tabulation
                input.nextElementSibling.setAttribute('tabindex', '-1'); // Retire les labels de la séquence de tabulation
            });
        }
    }

    toggleButton.addEventListener('click', toggleMenu);
});

function startDictation(str) {
    if ('webkitSpeechRecognition' in window) {
        var recognition = new webkitSpeechRecognition();

        recognition.continuous = false;
        recognition.interimResults = false;
        recognition.lang = "fr-FR";
        recognition.start();

        recognition.onresult = function (e) {
            document.getElementById(str).value = e.results[0][0].transcript;
            recognition.stop();
        };

        recognition.onerror = function (e) {
            recognition.stop();
        }
    } else {
        alert("La reconnaissance vocale n'est pas prise en charge par ce navigateur.");
    }
}

document.addEventListener('DOMContentLoaded', function () {
    const dyslexieCheckbox = document.getElementById('dyslexie');

    // Ajouter un écouteur d'événements au changement de l'état de la case à cocher
    dyslexieCheckbox.addEventListener('change', function() {
        // Vérifier si la case à cocher est cochée
        if (this.checked) {
            // Activer le mode dyslexique
            activerModeDyslexique();
        } else {
            // Désactiver le mode dyslexique
            desactiverModeDyslexique();
        }
    });
});

function activerModeDyslexique() {
    // Récupérer tous les éléments de texte de la page
    const elementsTextuels = document.querySelectorAll('h1, h2, h3, p, li, label, button, input[type="text"], input[type="button"], input[type="submit"], input[type="reset"], textarea');

    // Parcourir chaque élément de texte
    elementsTextuels.forEach(function(element) {

        // Appliquer la police OpenDyslexic
        element.style.fontFamily = 'OpenDyslexic, Arial, sans-serif';
    });
}

function desactiverModeDyslexique() {
    // Récupérer tous les éléments de texte de la page
    const elementsTextuels = document.querySelectorAll('h1, h2, h3, p, li, label, button, input[type="text"], input[type="button"], input[type="submit"], input[type="reset"], textarea');

    // Parcourir chaque élément de texte
    elementsTextuels.forEach(function(element) {

        // Retirer la police OpenDyslexic
        element.style.fontFamily = 'Roboto, sans-serif';

    });
}

function augmenterInterligne() {
    // Récupérer tous les éléments de texte de la page
    const elementsTextuels = document.querySelectorAll('h1, h2, h3, p, li, label, button, input[type="text"], input[type="button"], input[type="submit"], input[type="reset"], textarea');

    // Augmenter l'interligne pour chaque élément de texte
    elementsTextuels.forEach(function(element) {
        element.style.lineHeight = '2'; // Modifier la valeur d'interligne selon vos besoins
    });
}

// Fonction pour réinitialiser l'interligne par défaut
function reinitialiserInterligne() {
    // Récupérer tous les éléments de texte de la page
    const elementsTextuels = document.querySelectorAll('h1, h2, h3, p, li, label, button, input[type="text"], input[type="button"], input[type="submit"], input[type="reset"], textarea');

    // Réinitialiser l'interligne pour chaque élément de texte
    elementsTextuels.forEach(function(element) {
        element.style.lineHeight = ''; // Réinitialiser l'interligne par défaut
    });
}

document.addEventListener('DOMContentLoaded', function () {
    const interligneCheckbox = document.getElementById('interligne-augmentation');

    // Ajouter un écouteur d'événements au changement de l'état de la case à cocher
    interligneCheckbox.addEventListener('change', function() {
        // Vérifier si la case à cocher est cochée
        if (this.checked) {
            // Augmenter l'interligne
            augmenterInterligne();
        } else {
            // Réinitialiser l'interligne par défaut
            reinitialiserInterligne();
        }
    });
});

function augmenterContraste() {
    // Modifier les couleurs pour un meilleur contraste
    document.body.style.backgroundColor = '#000'; // Fond noir
    document.body.style.color = '#fff'; // Texte blanc
}

// Fonction pour réinitialiser les couleurs par défaut
function reinitialiserContraste() {
    // Rétablir les couleurs par défaut
    document.body.style.backgroundColor = ''; // Couleur de fond par défaut
    document.body.style.color = ''; // Couleur de texte par défaut
}

document.addEventListener('DOMContentLoaded', function () {
    const contrastCheckbox = document.getElementById('contrast');

    // Ajouter un écouteur d'événements au changement de l'état de la case à cocher
    contrastCheckbox.addEventListener('change', function() {
        // Vérifier si la case à cocher est cochée
        if (this.checked) {
            // Augmenter le contraste
            augmenterContraste();
        } else {
            // Réinitialiser le contraste par défaut
            reinitialiserContraste();
        }
    });
});


let slideIndex = 0;
showSlide(slideIndex);

function moveSlide(step) {
    showSlide(slideIndex += step);
}

function showSlide(n) {
    let i;
    let slides = document.getElementsByClassName("carousel-item");
    if (n >= slides.length) {slideIndex = 0}
    if (n < 0) {slideIndex = slides.length - 1}
    for (i = 0; i < slides.length; i++) {
        slides[i].style.display = "none";
    }
    slides[slideIndex].style.display = "block";
}
