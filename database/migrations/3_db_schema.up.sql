BEGIN;

CREATE TABLE Kafedrs (
  ID_Kaf INT PRIMARY KEY,  -- Теперь это первичный ключ
  Kaf_Name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users(
    ID SERIAL PRIMARY KEY,  -- Изменено с INT AUTO_INCREMENT на SERIAL
    user_name VARCHAR(255) NOT NULL,
    chatID INT UNIQUE NOT NULL,
    ID_Kaf INT,
    FOREIGN KEY (ID_Kaf) REFERENCES Kafedrs(ID_Kaf)  -- Внешний ключ теперь корректен
);

COMMIT;