CREATE TABLE sample_table (
  id serial NOT NULL PRIMARY KEY,
  name varchar(20),
  description varchar(200)
);

INSERT INTO sample_table (name, description)
  VALUES ('Bob', 'he is great');