CREATE TABLE cars (
    id serial NOT NULL,
    brand varchar(100) NOT NULL,
    model varchar(100) NOT NULL,
    image_url text,
    year integer NOT NULL,
    price varchar(50) NOT NULL,
    color varchar(50) NOT NULL
);

CREATE TABLE users (
    id serial NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar(20),
    password varchar(255) NOT NULL,
    role varchar(50) DEFAULT 'user'::character varying
);
