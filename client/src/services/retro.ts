import { Serializer } from "ts-japi";
import { callAPI } from "../utilities/http";
import { RetrocardType, deserializeMany } from "../types/retrocard";

export async function getLatestBoard() {
    const response = await callAPI('/api/retrocards?active=true', 'GET');
    
    const cards = deserializeMany(response.data);

    return cards;
}

export async function createRetrocard(retrocard: RetrocardType) {
    console.log(retrocard)
    const cardSerializer = new Serializer('retrocard');
    const serializedCard = await cardSerializer.serialize(retrocard);

    await callAPI('/api/retrocards', 'POST', serializedCard);
}

export async function highlightCard(highlightedCard: RetrocardType | null, newCard: RetrocardType) {
    const cardSerializer = new Serializer('retrocard');
    
    if (highlightedCard) {
        highlightedCard.focus = false;
        const serializedOldCard = await cardSerializer.serialize(highlightedCard);
        await callAPI('/api/retrocards/' + highlightedCard.id, 'PUT', serializedOldCard);
    }

    newCard.focus = true;

    const serializedNewCard = await cardSerializer.serialize(newCard);

    await callAPI('/api/retrocards/' + newCard.id, 'PUT', serializedNewCard);
}