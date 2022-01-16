CREATE TABLE users
(
        id serial NOT NULL,
        email text UNIQUE,
        password_digest text,
        created_at timestamp NOT NULL,
        updated_at timestamp NOT NULL,
        PRIMARY KEY (id)
) WITHOUT OIDS;
