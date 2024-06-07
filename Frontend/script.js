document.addEventListener('DOMContentLoaded', () => {
    const registerForm = document.getElementById('registerForm');
    const loginForm = document.getElementById('loginForm');
    const searchForm = document.getElementById('searchForm');

    if (registerForm) {
        registerForm.addEventListener('submit', (e) => {
            e.preventDefault();
            //registro
            alert('¡Usuario registrado con éxito!');
            window.location.href = 'login.html';
        });
    }

    if (loginForm) {
        loginForm.addEventListener('submit', (e) => {
            e.preventDefault();
            //autenticación
            alert('¡Usuario autenticado con éxito!');
            window.location.href = 'reservations.html';
        });
    }

    if (searchForm) {
        searchForm.addEventListener('submit', (e) => {
            e.preventDefault();
            //búsqueda
            alert('¡Búsqueda realizada!');
            const availableCars = document.getElementById('availableCars');
            availableCars.innerHTML = '<p>Lista de autos disponibles según los criterios de búsqueda</p>';
        });
    }
});
