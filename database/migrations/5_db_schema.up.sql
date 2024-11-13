BEGIN;

UPDATE schema_migrations SET dirty=false WHERE version=4;

CREATE TABLE IF NOT EXISTS Kafedrs 
(
  ID_Kaf BIGINT,  
  Kaf_Name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users
(
    ID SERIAL PRIMARY KEY,  
    user_name VARCHAR(255),
    chatID BIGINT,
    ID_Kaf BIGINT
);

COMMIT;