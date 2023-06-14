import { Serializer } from "ts-japi";
import { callAPI } from "../utilities/http";

export type RetrocardType = {
    id: string | null;
    title: string;
    column: number;
    active: boolean;
}

export async function getLatestBoard() {
    const response = await callAPI('/api/retrocards?active=true', 'GET');
    
    return response;
}

export async function createRetrocard(retrocard: RetrocardType) {
    const cardSerializer = new Serializer('retrocard');
    const serializedCard = await cardSerializer.serialize(retrocard);

    await callAPI('/api/retrocards', 'POST', serializedCard);
}
