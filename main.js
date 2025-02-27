let menu = document.querySelector('.menu-icon');
let navbar = document.querySelector('.navbar');

menu.onclick = () => {
    menu.classList.toggle('move');
    navbar.classList.toggle('open-menu');
};

// Close menu on page scroll
window.onscroll = () => {
    menu.classList.remove('move');
    navbar.classList.remove('open-menu');
};

//Input Form Date
window.onload = () => {
    let Today = new Date().ToISOString().Split('T')[0]
    document.getElementById('start-date').value = Today;
    document.getElementById('return-date').value = new Date(
        Date.now() + 7 * 86400000
    )
    .toISOString()
    .split("T")[0];
}
// Scroll Reveal animation
const animate = ScrollReveal({
    origin: 'top',    // Where the animation starts from (top of the screen)
    distance: '60px',  // How far the element will travel during the animation
    duration: 1000,    // Duration of the animation in milliseconds
    delay: 400,        // Delay before starting the animation
    easing: 'ease-in-out' // Optional: smooth the transition with easing
});

// Apply ScrollReveal to your navbar
//animate.reveal('.nav, .heading');
//animate.reveal(".home-img img",{origin:'right'});
//animate.reveal('.input-form'),{origin: 'buttom'};
//animate.reveal('.trend-box, .rental-box'),{interval: 100};