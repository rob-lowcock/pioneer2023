import { Serializer } from "ts-japi";

export type RetrocardType = {
    id: string | null;
    title: string;
    column: number;
    active: boolean;
}

export async function getLatestBoard() {
    const response = await fetch(import.meta.env.VITE_API_SERVER + '/api/retrocards?active=true');
    if (response.ok) {
        return response.json();
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

export async function createRetrocard(retrocard: RetrocardType) {
    const cardSerializer = new Serializer('retrocard');
    const serializedCard = await cardSerializer.serialize(retrocard);

    const response = await fetch(import.meta.env.VITE_API_SERVER + '/api/retrocards', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(serializedCard)
    });

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
