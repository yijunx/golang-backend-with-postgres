/* init generic db for development use */
CREATE DATABASE "db";

\connect "db";

CREATE SCHEMA IF NOT EXISTS "my-schema";

CREATE EXTENSION pg_trgm SCHEMA "my-schema";

CREATE USER "user" WITH PASSWORD 'password';

REVOKE ALL ON schema public FROM public;

GRANT ALL PRIVILEGES ON SCHEMA "my-schema" TO "user";

alter role "user" set search_path = "my-schema";