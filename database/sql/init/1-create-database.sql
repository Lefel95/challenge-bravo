-- THIS IS THE INITIAL SCRIPT TO START THE DATABASE Conversion

CREATE DATABASE IF NOT EXISTS conversion 
CHARACTER SET utf8
COLLATE utf8_unicode_ci;

CREATE TABLE IF NOT EXISTS conversion.currency (
  id INT PRIMARY KEY NOT NULL,
  name VARCHAR (50) NOT NULL,
  ballast_to_dollar DECIMAL (10, 4) NOT NULL DEFAULT 0.0,
  created_at DATETIME NOT NULL DEFAULT NOW (),
  updated_at DATETIME NOT NULL DEFAULT NOW ()
) ENGINE = INNODB;