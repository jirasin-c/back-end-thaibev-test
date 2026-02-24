CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS occupations (
  code TEXT PRIMARY KEY,
  name TEXT NOT NULL
);

INSERT INTO occupations(code, name) VALUES
('DEV', 'Developer'),
('BA', 'Business Analyst'),
('QA', 'QA/Tester'),
('PM', 'Project Manager')
ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS people_profiles (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  first_name TEXT NOT NULL,
  last_name  TEXT NOT NULL,
  email      TEXT NOT NULL,
  phone      TEXT NOT NULL,
  birth_day  DATE NOT NULL,
  sex        TEXT NOT NULL CHECK (sex IN ('M','F')),
  occupation_code TEXT NOT NULL REFERENCES occupations(code),
  profile_file_name TEXT NOT NULL,
  profile_base64 TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX IF NOT EXISTS ux_people_email ON people_profiles(email);