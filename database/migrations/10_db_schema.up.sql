BEGIN;

CREATE TABLE IF NOT EXISTS Kafedrs 
(
  ID_Kaf BIGINT,  
  Kaf_Name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,  
    user_name VARCHAR(255),
    chat_id BIGINT,
    user_login VARCHAR(255),
    user_password VARCHAR(255)
);

COMMIT;