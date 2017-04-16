CREATE TABLE events (
    id bigserial primary key,
    tutor integer REFERENCES users(id) NOT NULL,
    tutee integer REFERENCES users(id),
    start_time TIMESTAMPTZ,
    end_time TIMESTAMPTZ
);
