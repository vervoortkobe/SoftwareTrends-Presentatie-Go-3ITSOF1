async function handleLogout() {
    try {
        const response = await fetch('/logout', {
            method: 'POST',
            credentials: 'include'
        });

        if (response.ok) {
            console.log('Logout successful');
            window.location.href = '/login';
        } else {
            console.error('Logout failed');
            alert('Failed to logout. Please try again.');
        }
    } catch (error) {
        console.error('Error during logout:', error);
        alert('An error occurred during logout.');
    }
}
