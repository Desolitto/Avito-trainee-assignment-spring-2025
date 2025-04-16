CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('client', 'moderator', 'pvz_employee')),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE pvz (
    id SERIAL PRIMARY KEY,
    city TEXT NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань')),
    registered_at TIMESTAMP DEFAULT now()
);

CREATE TABLE receptions (
    id SERIAL PRIMARY KEY,
    pvz_id INTEGER REFERENCES pvz(id),
    started_at TIMESTAMP DEFAULT now(),
    status TEXT NOT NULL CHECK (status IN ('in_progress', 'closed'))
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    reception_id INTEGER REFERENCES receptions(id),
    pvz_id INTEGER REFERENCES pvz(id),
    accepted_at TIMESTAMP DEFAULT now(),
    type TEXT NOT NULL CHECK (type IN ('электроника', 'одежда', 'обувь'))
);
