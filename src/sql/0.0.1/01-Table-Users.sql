CREATE TABLE USERS(
    ID VARCHAR(36) PRIMARY KEY,
    EMAIL VARCHAR(100) NOT NULL,
    PASSWORD VARCHAR(20) NOT NULL
);

CREATE INDEX IX_USERS_EMAIL ON USERS (EMAIL);