CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL, 
    password VARCHAR(100) NOT NULL,
    gender_id SERIAL NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_gender FOREIGN KEY (gender_id) REFERENCES gender(id)
);

CREATE TABLE gender (
    id SERIAL PRIMARY KEY,
    gender VARCHAR(10)
);