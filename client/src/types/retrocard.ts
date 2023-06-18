export type RetrocardType = {
    id: string | null;
    title: string;
    column: number;
    active: boolean;
    focus: boolean;
    discussedAt: string | null;
};

type serializedRetrocardType = {
    id: string;
    attributes: {
        title: string;
        column: number;
        active: boolean;
        focus: boolean;
        discussed_at: string;
    };
}

export function deserializeOne(json: serializedRetrocardType) : RetrocardType {
    const card: RetrocardType = {
        id: json.id,
        title: json.attributes.title,
        column: json.attributes.column,
        active: json.attributes.active,
        focus: json.attributes.focus,
        discussedAt: json.attributes.discussed_at
    };

    return card;
}

export function deserializeMany(json: serializedRetrocardType[]) : RetrocardType[] {
    return json.map(deserializeOne);
}