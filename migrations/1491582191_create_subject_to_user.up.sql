CREATE TABLE user_to_subject (
    id bigserial primary key,
    user_id integer REFERENCES users(id) NOT NULL,
    subject_id integer REFERENCES subjects(id) NOT NULL,
    unique(user_id, subject_id)
);
