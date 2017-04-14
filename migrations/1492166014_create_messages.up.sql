CREATE TABLE messages (
    id bigserial primary key,
    sender integer REFERENCES users(id) NOT NULL,
    receiver integer REFERENCES users(id) NOT NULL,
    content TEXT NOT NULL,
    sent_at TIMESTAMPTZ DEFAULT now()
);
