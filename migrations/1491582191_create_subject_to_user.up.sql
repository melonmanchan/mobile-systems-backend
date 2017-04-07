CREATE TABLE user_to_subject (
    id bigserial primary key,
    user_id integer REFERENCES users(id),
    subject_id integer REFERENCES subjects(id)
);
