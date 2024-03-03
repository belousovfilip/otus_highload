CREATE TABLE users
(
    id         INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    first_name VARCHAR(255),
    last_name  VARCHAR(255),
    email      VARCHAR(255) UNIQUE            NOT NULL,
    password   VARCHAR(255)                   NOT NULL,
    age        INT8 UNSIGNED,
    gender     VARCHAR(10),
    city       VARCHAR(255),
    interests  TEXT,
    created_at DATETIME                       NOT NULL
);
