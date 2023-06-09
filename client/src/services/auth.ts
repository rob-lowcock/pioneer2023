export async function attemptLogin(email: string, password: string) {
    const params = {
        "client_id": import.meta.env.VITE_CLIENT_ID,
        "client_secret": import.meta.env.VITE_CLIENT_SECRET,
        "grant_type": "password",
        "username": email,
        "password": password
    }
    const response = await fetch(import.meta.env.VITE_API_SERVER + '/api/token?' + new URLSearchParams(params), {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })

    if (response.ok) {
        return await response.json();
    }

    switch (response.status) {
        case 403:
            throw new Error('Invalid email or password');
        case 500:
            throw new Error('Oops! Something went wrong on our side - please try again');
        default:
            throw new Error('Something went wrong - please try again');
    }
}