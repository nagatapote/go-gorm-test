CREATE TABLE users
(
        id serial NOT NULL,
        email text UNIQUE NOT NULL,
        crypted_password text NOT NULL,
        created_at timestamp NOT NULL,
        updated_at timestamp NOT NULL,
        PRIMARY KEY (id)
) WITHOUT OIDS;

CREATE TABLE files
(
        id serial NOT NULL,
        upload_name text UNIQUE NOT NULL,
        files_name text NOT NULL,
        created_at timestamp NOT NULL,
        updated_at timestamp NOT NULL,
        PRIMARY KEY (id)
) WITHOUT OIDS;
