-- Add dummy data to tables.
-- To Create table run init_database.sql script

INSERT INTO professors (name, active)
VALUES
("Ana", 1),
("Maria", 1),
("Sofia", 1),
("Mafalda", 0);

INSERT INTO classes (name, weekday, time, professor_id, active)
VALUES
("Yoga", 0, "13:00", 1, 1),
("Yoga", 0, "19:00", 1, 1),
("Pilates", 1, "19:00", 2, 1),
("Pilates", 2, "19:00", 2, 1),
("Dance", 3, "18:45", 3, 1),
("Dance", 4, "18:45", 3, 1),
("All", 4, "10:00", 1, 0);

INSERT INTO students (name, phone_number, email, nif, active)
VALUES
("Ana Mirtilo", "918765432", "ana@mirtilo.com", 3749186593, 1),
("Maria Inês", null, "maria.ines@example.com", 150905729, 1),
("Joana Gomes de Mello", "924356718", null, 903325787, 1),
("Diana Silva", "911638960", null, 612118644, 1),
("Carolina Esteves", "919767466", "carol.est@example.com", 593063500, 1),
("Filipa Madeira", "917657745", "fpm.madeira@example.com", 494778454, 1),
("Rita Maria Castelo", "919495950", null, 891634795, 1),
("Teresa Pessoa", "965871537", "", 519509302, 0);

INSERT INTO class_ocupancy (class_id, student_id)
VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(2, 1),
(2, 2),
(2, 3),
(2, 4),
(3, 6),
(3, 7),
(3, 5),
(4, 5),
(4, 6),
(4, 7),
(4, 3),
(5, 1),
(5, 4),
(5, 7),
(5, 3),
(6, 3),
(6, 7),
(6, 4),
(6, 1);
