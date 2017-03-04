CREATE TABLE authentication_methods (
    id bigserial primary key,
    type varchar(32) NOT NULL
);

CREATE TABLE user_types (
    id bigserial primary key,
    type varchar(32) NOT NULL
);

CREATE TABLE users (
    id bigserial primary key,
    first_name varchar(64) NOT NULL,
    last_name varchar(64) NOT NULL,
    email varchar(64) NOT NULL,
    password varchar(256),
    auth_method integer REFERENCES authentication_methods(id),
    user_type integer REFERENCES user_types(id)
);
