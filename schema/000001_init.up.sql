CREATE DATABASE dealer;

\c dealer;

CREATE TABLE cars (
    id serial NOT NULL,
    brand varchar(100) NOT NULL,
    model varchar(100) NOT NULL,
    image_url text,
    year integer NOT NULL,
    price varchar(50) NOT NULL,
    color varchar(50) NOT NULL,
    description varchar(255) NOT NULL
);

CREATE TABLE users (
    id serial NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar(20),
    password varchar(255) NOT NULL,
    role varchar(50) DEFAULT 'user'::character varying
);

INSERT INTO users (email, phone, password, role) VALUES ('root', '+7 701 143 3567', 'root', 'admin');

INSERT INTO cars (brand, model, image_url, year, price, color, description) VALUES ('Porsche', '911 GT3 RS', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQL0e70REmE1h-dnQiODFozdIGMuCFdoZIc4Q&s', 2022, '85 000 000', 'yellow', 'German premium sport car');
