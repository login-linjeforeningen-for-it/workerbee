CREATE TYPE "time_type_enum" AS ENUM (
    'default',
    'no_end',
    'whole_day',
    'to_be_determined'
);

CREATE TYPE "location_type" AS ENUM (
    'mazemap',
    'coords',
    'address',
    'digital'
);

CREATE TABLE "audiences" (
  "id" SERIAL PRIMARY KEY,
  "name_no" text NOT NULL,
  "name_en" text NOT NULL,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "categories" (
    "id" SERIAL PRIMARY KEY,
    "name_en" varchar NOT NULL,
    "name_no" varchar NOT NULL,
    "color" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "job_types" (
    "id" SERIAL PRIMARY KEY,
    "name_en" varchar NOT NULL,
    "name_no" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "events" (
    "id" SERIAL PRIMARY KEY,
    "visible" bool NOT NULL DEFAULT false,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "description_no" varchar NOT NULL,
    "description_en" varchar NOT NULL,
    "informational_no" varchar,
    "informational_en" varchar,
    "time_type" time_type_enum NOT NULL DEFAULT 'default',
    "time_start" timestamp NOT NULL,
    "time_end" timestamp NOT NULL,
    "time_publish" timestamp NOT NULL,
    "time_signup_release" timestamp,
    "time_signup_deadline" timestamp,
    "canceled" bool NOT NULL DEFAULT false,
    "category_id" int,
    "digital" bool NOT NULL DEFAULT false,
    "highlight" bool NOT NULL DEFAULT false,
    "image_small" varchar,
    "image_banner" varchar,
    "link_facebook" varchar,
    "link_discord" varchar,
    "link_signup" varchar,
    "link_stream" varchar,
    "capacity" int,
    "is_full" bool NOT NULL DEFAULT false,
    "organization_id" int,
    "location_id" int,
    "parent_id" int,
    "rule_id" int,
    "audience_id" int,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "rules" (
    "id" SERIAL PRIMARY KEY,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "description_no" varchar NOT NULL,
    "description_en" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "organizations" (
    "id" SERIAL PRIMARY KEY,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "description_no" varchar NOT NULL,
    "description_en" varchar NOT NULL,
    "link_homepage" varchar NOT NULL,
    "link_linkedin" varchar,
    "link_facebook" varchar,
    "link_instagram" varchar,
    "logo" varchar,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "locations" (
    "id" SERIAL PRIMARY KEY,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "type" location_type NOT NULL DEFAULT 'digital',
    "mazemap_campus_id" int,
    "mazemap_poi_id" int,
    "address_street" varchar,
    "address_postcode" int,
    "city_id" int,
    "coordinate_lat" float,
    "coordinate_lon" float,
    "url" varchar,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "jobs" (
    "id" SERIAL PRIMARY KEY,
    "visible" bool NOT NULL DEFAULT false,
    "highlight" bool NOT NULL DEFAULT false,
    "title_no" varchar NOT NULL,
    "title_en" varchar NOT NULL,
    "position_title_no" varchar NOT NULL,
    "position_title_en" varchar NOT NULL,
    "description_short_no" varchar NOT NULL,
    "description_short_en" varchar NOT NULL,
    "description_long_no" varchar NOT NULL,
    "description_long_en" varchar NOT NULL,
    "job_type_id" int,
    "time_publish" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "time_expire" timestamp NOT NULL,
    "banner_image" varchar,
    "organization_id" int,
    "application_url" varchar,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "ad_city_relation" (
    "job_id" int NOT NULL,
    "city_id" int NOT NULL,
    PRIMARY KEY ("job_id", "city_id")
);

CREATE TABLE "cities" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar UNIQUE
);

CREATE TABLE "ad_skill_relation" (
    "job_id" int NOT NULL,
    "skill_id" int NOT NULL,
    PRIMARY KEY ("job_id", "skill_id")
);

CREATE TABLE "skills" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL
);

CREATE TABLE "albums" (
    "id" SERIAL PRIMARY KEY,
    "name_no" TEXT NOT NULL,
    "name_en" TEXT NOT NULL,
    "description_no" TEXT NOT NULL,
    "description_en" TEXT NOT NULL,
    "year" INT NOT NULL,
    "event_id" INT,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "quotes" (
    "id" SERIAL PRIMARY KEY,
    "author" text NOT NULL,
    "quoted" text NOT NULL,
    "content" text NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

-- Insert default values
INSERT INTO "audiences" ("name_en", "name_no")
VALUES
('Login', 'Login'),
('Active members', 'Aktive medlemmer'),
('Students', 'Studenter'),
('Open', 'Åpen'),
('Bachelor', 'Bachelor'),
('Master', 'Master'),
('PhD', 'PhD'),
('First semester', 'Første semester'),
('Second semester', 'Andre semester'),
('Third semester', 'Tredje semester'),
('Fourth semester', 'Fjerde semester'),
('Fifth semester', 'Femte semester'),
('Sixth semester', 'Sjette semester'),
('Seventh semester', 'Sjuende semester');

INSERT INTO "categories" ("name_en", "name_no", "color")
VALUES
('Login', 'Login', '#fd8738'),
('TekKom', 'TekKom', '#a206c9'),
('CTFKom', 'CTFKom', '#2da62b'),
('EvntKom', 'EvntKom', '#d62f43'),
('PR', 'PR', '#ffff00'),
('BedKom', 'BedKom', '#1f56c5'),
('SatKom', 'SatKom', '#64ddd7'),
('BroomBroom', 'BroomBroom', '#cd53d8'),
('Pearlgroup', 'Perlegruppa', '#cd53d8'),
('Bookclub', 'Bokklubb', '#cd53d8'),
('Houseband', 'Husbandet', '#cd53d8'),
('Social', 'Sosialt', '#d62f43'),
('Cyberdays', 'Cyberdagene', 'linear-gradient(120deg, hsla(217, 100%, 50%, 1) 10%, hsla(186, 100%, 69%, 1) 100%)'),
('Buddyweek', 'Fadderuka', '#fa75a6'),
('BarKom', 'BarKom', '#f1c40f'),
('Other', 'Andre', '#545b5f');

INSERT INTO "job_types" ("name_en", "name_no") VALUES
('Full Time', 'Fulltid'),
('Part Time', 'Deltid'),
('Internship', 'Praksisplass'),
('Voluntairy', 'Verv'),
('Summer', 'Sommer');

-- Alerts

CREATE TABLE "alerts" (
  id SERIAL PRIMARY KEY,
  service TEXT NOT NULL,
  page TEXT NOT NULL,
  title_en TEXT NOT NULL,
  title_no TEXT NOT NULL,
  description_en TEXT NOT NULL,
  description_no TEXT NOT NULL,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(service, page)
);

-- Honey

CREATE TABLE IF NOT EXISTS "honey" (
    id SERIAL PRIMARY KEY,
    service TEXT NOT NULL,
    language TEXT NOT NULL,
    page TEXT NOT NULL,
    text TEXT NOT NULL,
    UNIQUE(service, page, language)
);

CREATE TABLE daily_history (
    insert_date DATE NOT NULL PRIMARY KEY,
    inserted_count INTEGER NOT NULL DEFAULT 0
);

CREATE OR REPLACE FUNCTION update_insert_history()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO daily_history (insert_date, inserted_count)
    VALUES (DATE(NEW.created_at AT TIME ZONE 'Europe/Oslo'), 1)
    ON CONFLICT (insert_date)
    DO UPDATE SET 
        inserted_count = daily_history.inserted_count + 1;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_update_history()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO daily_history (insert_date, inserted_count)
    VALUES (DATE(NEW.updated_at AT TIME ZONE 'Europe/Oslo'), 1)
    ON CONFLICT (insert_date)
    DO UPDATE SET 
        inserted_count = daily_history.inserted_count + 1;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER track_events_updates
    AFTER UPDATE ON events
    FOR EACH ROW
    EXECUTE FUNCTION update_update_history();

CREATE TRIGGER track_rules_updates
    AFTER UPDATE ON rules
    FOR EACH ROW
    EXECUTE FUNCTION update_update_history();

CREATE TRIGGER track_organizations_updates
    AFTER UPDATE ON organizations
    FOR EACH ROW
    EXECUTE FUNCTION update_update_history();

CREATE TRIGGER track_locations_updates
    AFTER UPDATE ON locations
    FOR EACH ROW
    EXECUTE FUNCTION update_update_history();

CREATE TRIGGER track_jobs_updates
    AFTER UPDATE ON jobs
    FOR EACH ROW
    EXECUTE FUNCTION update_update_history();

CREATE TRIGGER track_albums_updates
    AFTER UPDATE ON albums
    FOR EACH ROW
    EXECUTE FUNCTION update_update_history();

CREATE TRIGGER track_alerts_updates
    AFTER UPDATE ON alerts
    FOR EACH ROW
    EXECUTE FUNCTION update_update_history();

CREATE TRIGGER track_events_inserts
    AFTER INSERT ON events
    FOR EACH ROW
    EXECUTE FUNCTION update_insert_history();

CREATE TRIGGER track_rules_inserts
    AFTER INSERT ON rules
    FOR EACH ROW
    EXECUTE FUNCTION update_insert_history();

CREATE TRIGGER track_organizations_inserts
    AFTER INSERT ON organizations
    FOR EACH ROW
    EXECUTE FUNCTION update_insert_history();

CREATE TRIGGER track_locations_inserts
    AFTER INSERT ON locations
    FOR EACH ROW
    EXECUTE FUNCTION update_insert_history();

CREATE TRIGGER track_jobs_inserts
    AFTER INSERT ON jobs
    FOR EACH ROW
    EXECUTE FUNCTION update_insert_history();

CREATE TRIGGER track_albums_inserts
    AFTER INSERT ON albums
    FOR EACH ROW
    EXECUTE FUNCTION update_insert_history();

CREATE TRIGGER track_alerts_inserts
    AFTER INSERT ON alerts
    FOR EACH ROW
    EXECUTE FUNCTION update_insert_history();

CREATE INDEX ON "albums" ("year");
CREATE INDEX ON "albums" ("created_at");
CREATE INDEX ON "albums" ("updated_at");

CREATE INDEX ON "events" ("visible");
CREATE INDEX ON "events" ("highlight");
CREATE INDEX ON "events" ("time_start");
CREATE INDEX ON "events" ("time_end");
CREATE INDEX ON "events" ("updated_at");
CREATE INDEX ON "events" ("created_at");

CREATE INDEX ON "rules" ("updated_at");
CREATE INDEX ON "rules" ("created_at");

CREATE INDEX ON "organizations" ("updated_at");
CREATE INDEX ON "organizations" ("created_at");

CREATE INDEX ON "locations" ("updated_at");
CREATE INDEX ON "locations" ("created_at");

CREATE INDEX ON "jobs" ("updated_at");
CREATE INDEX ON "jobs" ("created_at");
CREATE INDEX ON "ad_city_relation" ("job_id");
CREATE INDEX ON "ad_city_relation" ("city_id");
CREATE INDEX ON "ad_skill_relation" ("job_id");
CREATE INDEX ON "ad_skill_relation" ("skill_id");

CREATE INDEX ON "honey"(service, page);
CREATE INDEX ON "honey"(service);

CREATE INDEX ON "alerts"(service);

ALTER TABLE "events" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "events" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "events" ADD FOREIGN KEY ("rule_id") REFERENCES "rules" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "events" ADD FOREIGN KEY ("parent_id") REFERENCES "events" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "events" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "events" ADD FOREIGN KEY ("audience_id") REFERENCES "audiences" ("id") ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE "jobs" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "jobs" ADD FOREIGN KEY ("job_type_id") REFERENCES "job_types" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("job_id") REFERENCES "jobs" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "ad_skill_relation" ADD FOREIGN KEY ("job_id") REFERENCES "jobs" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "ad_skill_relation" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("id") ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "locations" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id") ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "albums" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE CASCADE ON DELETE SET NULL;