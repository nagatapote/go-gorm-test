CREATE TABLE users
(
        id serial NOT NULL,
        email text UNIQUE NOT NULL,
        crypted_password text NOT NULL,
        created_at timestamp NOT NULL,
        updated_at timestamp NOT NULL,
        PRIMARY KEY (id)
) WITHOUT OIDS;
