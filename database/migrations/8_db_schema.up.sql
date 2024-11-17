BEGIN;

CREATE TABLE IF NOT EXISTS Kafedrs 
(
  ID_Kaf BIGINT,  
  Kaf_Name VARCHAR(255)
);

DROP TABLE schema_migrations;

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,  
    user_name VARCHAR(255),
    chat_id BIGINT,
    ID_Kaf BIGINT
);

COMMIT;