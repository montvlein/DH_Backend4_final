CREATE USER grupo8 WITH PASSWORD 'grupo8';
ALTER USER grupo8 WITH SUPERUSER;

GRANT ALL PRIVILEGES ON DATABASE final_backend TO grupo8;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO grupo8;

COMMIT;

DROP TABLE IF EXISTS dentists;

CREATE TABLE dentists (
    id serial PRIMARY KEY,
    name VARCHAR(255),
    lastname VARCHAR(255),
    license VARCHAR(255)
);

DROP TABLE IF EXISTS patients;

CREATE TABLE patients (
    id serial PRIMARY KEY,
    name VARCHAR(255),
    lastname VARCHAR(255),
    address VARCHAR(255),
    dni VARCHAR(255),
    discharge_date VARCHAR(255)
);