-- -------------------------------------------------------------
-- Database: pereval
-- Generation Time: 2022-02-22 14:19:15.6730
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS pereval_id_seq;

-- Table Definition
CREATE TABLE IF NOT EXISTS "public"."pereval_added" (
    "id" int4 NOT NULL DEFAULT nextval('pereval_id_seq'::regclass),
    "date_added" timestamp,
    "raw_data" json,
    "images" json,
    "status" varchar(20) NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS pereval_areas_id_seq;

-- Table Definition
CREATE TABLE IF NOT EXISTS "public"."pereval_areas" (
    "id" int8 NOT NULL DEFAULT nextval('pereval_areas_id_seq'::regclass),
    "id_parent" int8 NOT NULL,
    "title" text,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS pereval_added_id_seq;

-- Table Definition
CREATE TABLE IF NOT EXISTS "public"."pereval_images" (
    "id" int4 NOT NULL DEFAULT nextval('pereval_added_id_seq'::regclass),
    "date_added" timestamp DEFAULT now(),
    "img" bytea NOT NULL,
    "title" text NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS untitled_table_200_id_seq;

-- Table Definition
CREATE TABLE IF NOT EXISTS "public"."spr_activities_types" (
    "id" int4 NOT NULL DEFAULT nextval('untitled_table_200_id_seq'::regclass),
    "title" text,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."pereval_areas" ("id", "id_parent", "title") VALUES
(0, 0, 'Планета Земля'),
(1, 0, 'Памиро-Алай'),
(65, 0, 'Алтай'),
(66, 65, 'Северо-Чуйский хребет'),
(88, 65, 'Южно-Чуйский хребет'),
(92, 65, 'Катунский хребет'),
(105, 1, 'Фанские горы'),
(106, 1, 'Гиссарский хребет (участок западнее перевала Анзоб)'),
(131, 1, 'Матчинский горный узел'),
(133, 1, 'Горный узел Такали, Туркестанский хребет'),
(137, 1, 'Высокий Алай'),
(147, 1, 'Кичик-Алай и Восточный Алай'),
(367, 375, 'Аладаглар'),
(375, 0, 'Тавр'),
(384, 0, 'Саяны'),
(386, 65, 'Хребет Листвяга'),
(387, 65, 'Ивановский хребет'),
(388, 65, 'Массив Мунгун-Тайга'),
(389, 65, 'Хребет Цаган-Шибэту'),
(390, 65, 'Хребет Чихачева (Сайлюгем)'),
(391, 65, 'Шапшальский хребет'),
(392, 65, 'Хребет Южный Алтай'),
(393, 65, 'Хребет Монгольский Алтай'),
(398, 384, 'Западный Саян'),
(399, 384, 'Восточный Саян'),
(402, 384, 'Кузнецкий Алатау'),
(459, 65, 'Курайский хребет');

INSERT INTO "public"."spr_activities_types" ("id", "title") VALUES
(1, 'пешком'),
(2, 'лыжи'),
(3, 'катамаран'),
(4, 'байдарка'),
(5, 'плот'),
(6, 'сплав'),
(7, 'велосипед'),
(8, 'автомобиль'),
(9, 'мотоцикл'),
(10, 'парус'),
(11, 'верхом');

