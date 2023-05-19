export async function attemptLogin(email: string, password: string) {
    const data = { 
        "data": {
            "type": "auth",
            "attributes" : {
                "email": email,
                "password": password
            }
        }
    }
    const response = await fetch(import.meta.env.VITE_API_SERVER + '/api/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })

    if (response.ok) {
        return await response.json();
    }

    switch (response.status) {
        case 401:
            throw new Error('Invalid email or password');
        case 500:
            throw new Error('Oops! Something went wrong on our side - please try again');
        default:
            throw new Error('Something went wrong - please try again');
    }
}