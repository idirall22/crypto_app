CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar,
  "is_active" boolean,
  "confirmation_link" varchar,
  "password_hash" varchar,
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
  "currency_id" int NOT NULL,
  "user_id" int NOT NULL,
  "address" varchar UNIQUE NOT NULL,
  "amount" double
);

CREATE TABLE "currency" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "transaction" (
  "id" serial PRIMARY KEY,
  "amount" double NOT NULL,
  "commission" double NOT NULL,
  "currency_id" int NOT NULL,
  "sender_address" varchar NOT NULL,
  "recipient_address" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "user_metadata" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "wallet" ADD FOREIGN KEY ("currency_id") REFERENCES "currency" ("id");

ALTER TABLE "wallet" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("sender_address") REFERENCES "wallet" ("address");

ALTER TABLE "transaction" ADD FOREIGN KEY ("recipient_address") REFERENCES "wallet" ("address");
