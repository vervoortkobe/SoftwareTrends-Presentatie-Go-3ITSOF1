document.getElementById('loginForm').addEventListener('submit', async function(e) {
    e.preventDefault();
    
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        console.log('Attempting login...');
        const response = await fetch('/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username,
                password
            }),
            credentials: 'include' // This is important for cookies
        });

        console.log('Login response status:', response.status);
        const data = await response.json();
        console.log('Login response data:', data);

        if (response.ok && data.success) {
            console.log('Login successful, redirecting to home page...');
            // Add a small delay to ensure cookie is set
            setTimeout(() => {
                window.location.href = '/';
            }, 100);
        } else {
            console.error('Login failed:', data.error);
            alert(data.error || 'Login failed. Please check your credentials.');
        }
    } catch (error) {
        console.error('Error during login:', error);
        alert('An error occurred during login.');
    }
});