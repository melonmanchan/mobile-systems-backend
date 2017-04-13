CREATE TABLE tutorships (
    id bigserial primary key,
    tutor_id integer REFERENCES users(id) NOT NULL,
    tutee_id integer REFERENCES users(id) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    unique(tutor_id, tutee_id)
);
