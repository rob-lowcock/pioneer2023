CREATE TABLE retrocards (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    col INTEGER NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true
);