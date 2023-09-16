-- DROP DATABASE IF EXISTS final_backend;

-- CREATE DATABASE final_backend;

-- \c final_backend;

-- CREATE USER grupo8 WITH PASSWORD 'grupo8';
-- ALTER USER grupo8 WITH SUPERUSER;

-- GRANT ALL PRIVILEGES ON DATABASE final_backend TO grupo8;
-- GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO grupo8;

-- COMMIT;

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

DROP TABLE IF EXISTS appointments;

CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    patient_id INT NOT NULL,
    dentist_id INT NOT NULL,
    description TEXT,
    date VARCHAR(255) NOT NULL,
    hour VARCHAR(255) NOT NULL
);

ALTER TABLE appointments
ADD CONSTRAINT fk_patient
FOREIGN KEY (patient_id)
REFERENCES patients (id);

ALTER TABLE appointments
ADD CONSTRAINT fk_dentist
FOREIGN KEY (dentist_id)
REFERENCES dentists (id);

INSERT INTO patients (name, lastname, address, dni, discharge_date)
VALUES ('Fabricio', 'Montivero', 'direccion', '12345', '15/09/2023');