-- Add dummy data to tables.
-- To Create table run init_database.sql script

-- internal configurations
INSERT INTO subscription (value, name)
VALUES
(25, "1 class per week"),
(35, "2 class per week"),
(45, "3 class per week");

INSERT INTO promo (value, name)
VALUES
(0, "No Promo"),
(50, "Half Month"),
(25, "Family");

-- end internal configurations

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

--(name, phone_number, email, nif, active)
INSERT INTO students (name, phone_number, email, nif, health, year, subscription, promo, active, fee)
VALUES
("Ana Mirtilo", 918765432, "ana@mirtilo.com", 3749186593, "None", 2020, 1, 1, 1, 25),
("Maria Inês", null, "maria.ines@example.com", 150905729, "No problems", 2020, 1, 2, 1, 12.5),
("Joana Gomes de Mello", 924356718, null, 903325787, "lower back", 2020, 3, 1, 1, 6.25),
("Diana Silva", 911638960, null, 612118644, "health", 2020, 2, 2, 1, 17.5),
("Carolina Esteves", 919767466, "carol.est@example.com", 593063500, "healthy", 2020, 3, 1, 1, 45),
("Filipa Madeira", 917657745, "fpm.madeira@example.com", 494778454, "Lower back problems", 2020, 1, 2, 1, 35),
("Rita Maria Castelo", 919495950, null, 891634795, "-", 2020, 3, 3, 1, 11.25),
("Teresa Pessoa", 965871537, "", 519509302, "", 2020, 1, 1, 0, 25);

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
