CREATE TABLE "role" (
  "id" serial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_active" boolean,
  "confirmation_link" varchar,
  "password_hash" varchar,
  "role" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_metadata" (
  "id" serial PRIMARY KEY,
  "user_id" int,
  "ip_address" varchar,
  "user_agent" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "wallet" (
  "id" serial PRIMARY KEY,
  "currency" varchar NOT NULL,
  "user_id" int NOT NULL,
  "address" varchar UNIQUE NOT NULL,
  "amount" decimal
);

CREATE TABLE "currency" (
  "id" serial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "transaction" (
  "id" serial PRIMARY KEY,
  "amount" decimal NOT NULL,
  "commission" decimal NOT NULL,
  "currency" varchar NOT NULL,
  "sender_address" varchar NOT NULL,
  "recipient_address" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("role") REFERENCES "role" ("name");

ALTER TABLE "user_metadata" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "wallet" ADD FOREIGN KEY ("currency") REFERENCES "currency" ("name");

ALTER TABLE "wallet" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("sender_address") REFERENCES "wallet" ("address");

ALTER TABLE "transaction" ADD FOREIGN KEY ("recipient_address") REFERENCES "wallet" ("address");

ALTER TABLE "transaction" ADD FOREIGN KEY ("currency") REFERENCES "currency" ("name");
