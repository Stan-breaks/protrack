CREATE TABLE users(
 id INTEGER PRIMARY KEY,
 userName text NOT NULL,
 email text NOT NULL,
 firstName text NOT NULL,
 lastName text NOT NULL,
 password text NOT NULL,
 created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
