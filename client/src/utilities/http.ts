import Cookies from "js-cookie";
import { UnauthorizedError } from "./errors";

export async function callAPI(url: string, method: string, body?: any) {
    const response = await attemptAPICall(url, method, body);

    if (response.ok) {
        return response.json();
    }

    switch (response.status) {
        // handle 401 unauthorized
        case 401:
            // if we don't even have an access token, then just get out of here
            if (!Cookies.get('access_token')) {
                throw new UnauthorizedError(response.statusText);
            }

            try {
                const response2 = await attemptWithNewTokens(url, method, body);
                return response2.json();
            } catch (error) {
                throw new UnauthorizedError(error.message);
            }
            break;
        case 400:
            throw new Error('Invalid request - is that me!');
        case 500:
            throw new Error('Oops! Something went wrong on our side - please try again');
        default:
            throw new Error("Failed to call API");
    }
}

async function attemptAPICall(url: string, method: string, body?: any) {
    const accessToken = Cookies.get('access_token');

    const response = await fetch(import.meta.env.VITE_API_SERVER + url, {
        method: method,
        body: body ? JSON.stringify(body) : null,
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`
        }
    })

    return response;
}

async function attemptWithNewTokens(url: string, method: string, body?: any) {
    const data = new FormData();
    const refreshToken = Cookies.get('refresh_token');
    if (!refreshToken) {
        throw new Error('No refresh token');
    }

    data.append('refresh_token', refreshToken);
    data.append('grant_type', 'refresh_token');
    data.append('client_id', import.meta.env.VITE_CLIENT_ID);
    data.append('scope', 'read');

    const getNewTokens = await fetch(import.meta.env.VITE_API_SERVER + '/api/token', {
        method: 'POST',
        body: data
    });

    if (getNewTokens.ok) {
        const newTokens = await getNewTokens.json();
        Cookies.set('access_token', newTokens.access_token);
        Cookies.set('refresh_token', newTokens.refresh_token);

        const reattempt = await attemptAPICall(url, method, body);
        if (reattempt.ok) {
            return reattempt;
        }

        throw new Error("Failed to call API after setting new tokens");
    }

    throw new Error("Failed to get new tokens");
}