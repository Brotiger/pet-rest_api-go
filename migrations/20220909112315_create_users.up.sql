CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null UNIQUE,
    encrypted_password varchar not null
)