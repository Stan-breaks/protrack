CREATE TABLE users(
 id INTEGER PRIMARY KEY,
 userName text NOT NULL,
 email text NOT NULL,
 firstName text NOT NULL,
 lastName text NOT NULL,
 password text NOT NULL,
 created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
 
CREATE TABLE sqlite_schema (
  type TEXT NOT NULL,
  name TEXT NOT NULL,
  tbl_name TEXT NOT NULL,
  rootpage INTEGER NOT NULL,
  sql TEXT NOT NULL
);

