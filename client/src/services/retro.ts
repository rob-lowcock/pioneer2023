export async function getLatestBoard() {
    const response = await fetch(import.meta.env.VITE_API_SERVER + '/api/retrocards?active=true');
    if (response.ok) {
        return await response.json();
    }

    switch (response.status) {
        case 400:
            throw new Error('Invalid request');
        case 500:
            throw new Error('Oops! Something went wrong on our side - please try again');
        default:
            throw new Error('Something went wrong - please try again');
    }
}