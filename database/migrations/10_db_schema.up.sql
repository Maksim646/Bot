BEGIN;

CREATE TABLE IF NOT EXISTS alerts 
(
  id   SERIAL PRIMARY KEY,
  chat_id BIGINT,  
  teacher VARCHAR(255),
  subject_of_study VARCHAR (255),
  data_alert data
);

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,  
    user_name VARCHAR(255),
    chat_id BIGINT
);

COMMIT;