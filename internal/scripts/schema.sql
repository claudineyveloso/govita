-- User
DROP TABLE IF EXISTS "users";
CREATE TABLE IF NOT EXISTS users (
  id          UUID PRIMARY KEY,
  email       varchar(100) not null,
  password    varchar(100) not null,
  is_active   boolean not null default true,
  user_type   varchar(50) not null,
  created_at  timestamp not null,
  updated_at  timestamp not null
);
create unique index email_idx on users (email);

-- Search
DROP TABLE IF EXISTS "searches";
CREATE TABLE IF NOT EXISTS searches (
  id            UUID PRIMARY KEY,
  description   varchar(100) not null,
  created_at    timestamp not null,
  updated_at    timestamp not null
);
create unique index description_idx on searches (description);

-- Result
DROP TABLE IF EXISTS "results";
CREATE TABLE IF NOT EXISTS results (
  id            UUID PRIMARY KEY,
  image_url     varchar(255) not null,
  description   varchar(255) not null,
  font          varchar(100) not null,
  price         float not null,
  promotion     boolean not null default false,
  search_id     UUID not null,
  created_at    timestamp not null,
  updated_at    timestamp not null
);
ALTER TABLE
   "results"
ADD
  FOREIGN KEY (search_id) REFERENCES searches(id);
create unique index description_search_idx on results (description);
 -- Parameter
DROP TABLE IF EXISTS "parameters";
CREATE TABLE IF NOT EXISTS parameters (
  id                  UUID PRIMARY KEY,
  delete_at_days      integer NOT null default 0,
  percentage_pricing  integer NOT null default 0,
  created_at          timestamp not null,
  updated_at          timestamp not null
);
