CREATE USER orders_rep WITH PASSWORD 'amScD6paIx9EcKT3jeBw';

CREATE DATABASE midas_rep;

ALTER DATABASE midas_rep owner to midas_rep;

GRANT ALL PRIVILEGES ON DATABASE midas_rep TO midas_rep;
