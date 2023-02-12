-- Create tables in database
-- To add dummy data run dummy_data.sql script

CREATE TABLE professors (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  active INTEGER NOT NULL
);

CREATE TABLE classes (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  weekday INTEGER NOT NULL,
  time TEXT NOT NULL,
  professor_id INTEGER NOT NULL,
  active INTEGER NOT NULL,
  FOREIGN KEY(professor_id)
    REFERENCES professors(id)
);

CREATE TABLE students (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  phone_number TEXT,
  email TEXT,
  nif INTEGER NOT NULL,
  active INTEGER NOT NULL
);

CREATE TABLE class_ocupancy (
  class_id INTEGER NOT NULL,
  student_id INTEGER NOT NULL,
  constraint PK_class_ocupancy PRIMARY KEY (class_id, student_id),
  FOREIGN KEY (class_id)
    REFERENCES classes (id),
  FOREIGN KEY (student_id)
    REFERENCES students (id)
);

-- internal configurations
CREATE TABLE subscription (
  id INTEGER PRIMARY KEY,
  value INTEGER,
  name TEXT NOT NULL
);

CREATE TABLE promo (
  id INTEGER PRIMARY KEY,
  value INTEGER,
  name TEXT NOT NULL
);