CREATE TABLE users (
	"id" VARCHAR(36) PRIMARY KEY, -- uuidv4
	"email" VARCHAR(255) UNIQUE NOT NULL,
	"password_hash" TEXT NOT NULL
);
