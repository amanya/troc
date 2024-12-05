# Troc

A simple marketplace for bartering; trading by exchanging one commodity for
another.

This is an exercise to learn Go following the [Let's Go](https://lets-go.alexedwards.net/) book.

## Setting up the database

```sql
-- Create a new UTF-8 `troc` database.
CREATE DATABASE troc CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE troc;
```

```sql
-- Create a `trocs` table.
CREATE TABLE trocs (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	created DATETIME NOT NULL,
	expires DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_trocs_created ON trocs(created);
```

Some placeholder data:

```sql
INSERT INTO trocs (title, content, created, expires) VALUES (
	'An old silent pond',
	'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO trocs (title, content, created, expires) VALUES (
	'Over the wintry forest',
	'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO trocs (title, content, created, expires) VALUES (
	'First autumn morning',
	'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
```

Creating a user

```sql
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON trocs.* to 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
```

Create the sessions table for the session manager

```sql
CREATE TABLE sessions (
	token CHAR(43) PRIMARY KEY,
	data BLOB NOT NULL,
	expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```

Create the users table

```sql
CREATE TABLE users (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	hashed_password CHAR(60) NOT NULL,
	created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

## Configuring development TLS

```
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```