/*
 * Создать таблицу "TYPES" ("ТИПЫ").
 */
CREATE TABLE CATEGORIES (
  ID SERIAL NOT NULL PRIMARY KEY,
  NAME VARCHAR NOT NULL UNIQUE
);

/*
 * Добавить запись в таблицу "TYPES" ("ТИПЫ").
 */
INSERT INTO TYPES (NAME) VALUES ('e-NPS');

/*
 * Создать таблицу "CONDITIONS" ("СОСТОЯНИЯ").
 */
CREATE TABLE CONDITIONS (
  ID SERIAL NOT NULL PRIMARY KEY,
  NAME VARCHAR NOT NULL UNIQUE
);

/*
 * Добавить запись в таблицу "TYPES" ("ТИПЫ").
 */
INSERT INTO CONDITIONS (NAME) VALUES ('Завершённый');

/*
 * Чтобы использовать функцию "uuid_generate_v4()", необходимо выполнить SQL команду.
 */
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/*
 * Создать таблицу "SURVEYS" ("ОПРОСЫ").
 */
CREATE TABLE SURVEYS(
  ID UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  NAME VARCHAR,
  DESCRIPTION TEXT,
  CATEGORY INTEGER,
  FOREIGN KEY (CATEGORY) REFERENCES CATEGORIES (ID),
  CONDITION INTEGER NOT NULL DEFAULT 1,
  FOREIGN KEY (CONDITION) REFERENCES CONDITIONS (ID),
  MARK BOOLEAN NOT NULL DEFAULT FALSE,
  CONTROL BOOLEAN NOT NULL DEFAULT FALSE,
  START_PERIOD TIMESTAMP,
  END_PERIOD TIMESTAMP,
  CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATED_AT TIMESTAMP NOT NULL DEFAULT NOW(),
  EMAIL VARCHAR,
  BLOCKED BOOLEAN NOT NULL DEFAULT FALSE
);

/*
 * Создать функцию с названием "trigger_set_timestamp".
 */
CREATE OR REPLACE FUNCTION TRIGGER_SET_TIMESTAMP() RETURNS TRIGGER AS $$
  BEGIN
    NEW.UPDATED_AT = NOW();
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

/*
 * Удалить функцию "TRIGGER_SET_TIMESTAMP".
 */
DROP FUNCTION IF EXISTS TRIGGER_SET_TIMESTAMP;

/*
 * Создать триггер с названием "SET_TIMESTAMP".
 */
CREATE TRIGGER SET_TIMESTAMP
BEFORE UPDATE ON SURVEYS
FOR EACH ROW
EXECUTE PROCEDURE TRIGGER_SET_TIMESTAMP();

/*
 * Удалить триггер "SET_TIMESTAMP".
 */
DROP TRIGGER SET_TIMESTAMP ON SURVEYS;

/*
 * Создать таблицу "OPTIONS" ("ВАРИАНТЫ ОТВЕТОВ")
 */
CREATE TABLE OPTIONS (
  ID SERIAL NOT NULL PRIMARY KEY,
  TEXT TEXT NOT NULL UNIQUE
);

/*
 * Создать таблицу "WIDGETS" ("ВИДЖЕТЫ")
 */
CREATE TABLE WIDGETS (
  ID SERIAL NOT NULL PRIMARY KEY,
  NAME VARCHAR NOT NULL UNIQUE
);

/*
 * Создать таблицу "QUESTIONS" ("ВОПРОСЫ")
 */
CREATE TABLE QUESTIONS (
  ID SERIAL NOT NULL PRIMARY KEY,
  TEXT TEXT NOT NULL UNIQUE,
  WIDGET INT,
  FOREIGN KEY (WIDGET) REFERENCES WIDGETS (ID)
);

/*
 * Создать таблицу "FACTORS" ("ФАКТОРЫ").
 */
CREATE TABLE FACTORS (
  ID SERIAL NOT NULL PRIMARY KEY,
  NAME VARCHAR NOT NULL UNIQUE
);

/*
 * Создать таблицу связей между опросами и факторами.
 */
CREATE TABLE SURVEYS_FACTORS_RELATIONSHIP (
  SURVEY_ID UUID NOT NULL,
  FOREIGN KEY (SURVEY_ID) REFERENCES SURVEYS (ID),
  FACTOR_ID INT NOT NULL,
  FOREIGN KEY (FACTOR_ID) REFERENCES FACTORS (ID)
);

/*
 * Создать таблицу связей между опросами и организационной структурой
 */
CREATE TABLE SURVEYS_ORGANIZATIONS_RELATIONSHIP (
  SURVEY_ID UUID NOT NULL,
  FOREIGN KEY (SURVEY_ID) REFERENCES SURVEYS (ID),
  ORGANIZATION_ID INT NOT NULL
);

/*
 * Создать таблицу связей между факторами и вопросами.
 */
CREATE TABLE FACTORS_QUESTIONS_RELATIONSHIP (
  FACTOR_ID INT NOT NULL,
  FOREIGN KEY (FACTOR_ID) REFERENCES FACTORS (ID),
  QUESTION_ID INT NOT NULL,
  FOREIGN KEY (QUESTION_ID) REFERENCES QUESTIONS (ID)
);

/*
 * Создать таблицу связей между вопросами и вариантами ответов.
 */
CREATE TABLE QUESTIONS_OPTIONS_RELATIONSHIP (
  QUESTION_ID INT NOT NULL,
  FOREIGN KEY (QUESTION_ID) REFERENCES QUESTIONS (ID),
  OPTION_ID INT NOT NULL,
  FOREIGN KEY (OPTION_ID) REFERENCES OPTIONS (ID)
);

/*
 * Создать таблицу связей между опросами и вопросами.
 */
CREATE TABLE SURVEYS_QUESTIONS_RELATIONSHIP (
  SURVEY_ID UUID NOT NULL,
  FOREIGN KEY (SURVEY_ID) REFERENCES SURVEYS (ID),
  QUESTION_ID INT NOT NULL,
  FOREIGN KEY (QUESTION_ID) REFERENCES QUESTIONS (ID)
);

/*
 * Создать связь между опросами и сотрудниками.
 */
CREATE TABLE SURVEYS_EMPLOYEES_RELATIONSHIP (
  SURVEY_ID UUID NOT NULL,
  FOREIGN KEY (SURVEY_ID) REFERENCES SURVEYS (ID),
  EMPLOYEE VARCHAR NOT NULL,
  STATUS BOOLEAN NOT NULL DEFAULT FALSE
);

/*
 * Создать триггер отслеживающий изменения в таблице "SURVEYS".
 */
CREATE TRIGGER check_for_updates_in_surveys
BEFORE UPDATE ON SURVEYS
FOR EACH ROW
WHEN (OLD.UPDATED_AT IS DISTINCT FROM NEW.UPDATED_AT)
EXECUTE PROCEDURE CREATE_SURVEYS_QUESTIONS_RELATIONSHIP();

/*
 * Удаление триггера "CHECK_FOR_UPDATES_IN_SURVEYS" в таблице "SURVEYS".
 */
DROP TRIGGER check_for_updates_in_surveys ON SURVEYS;

/*
 * Создать функцию для создания записей в таблице "SURVEYS_QUESTIONS_RELATIONSHIP".
 */
CREATE OR REPLACE FUNCTION CREATE_SURVEYS_QUESTIONS_RELATIONSHIP() RETURNS TRIGGER AS $FUNCTION$
  BEGIN
    DELETE FROM SURVEYS_QUESTIONS_RELATIONSHIP WHERE SURVEY_ID = NEW.ID;
    INSERT INTO SURVEYS_QUESTIONS_RELATIONSHIP (SURVEY_ID, QUESTION_ID)
    SELECT NEW.ID SURVEY_ID, QUESTION_ID
    FROM FACTORS_QUESTIONS_RELATIONSHIP
    WHERE FACTOR_ID IN (
      SELECT FACTOR_ID FROM SURVEYS_FACTORS_RELATIONSHIP
      WHERE SURVEY_ID = NEW.ID
    );
    RETURN NEW;
  END;
$FUNCTION$ LANGUAGE plpgsql;

/*
 * Создать процедуру для проверки актуальности состояния и статуса блокировки в таблице "SURVEYS".
 */
CREATE OR REPLACE PROCEDURE tracker(CUSTOM_TIME TIMESTAMP WITHOUT TIME ZONE) AS $FUNCTION$
BEGIN
  UPDATE SURVEYS SET CONDITION = 3 WHERE CONDITION = 2 AND START_PERIOD IS NOT NULL AND END_PERIOD IS NOT NULL AND CUSTOM_TIME > END_PERIOD;
  UPDATE SURVEYS SET BLOCKED = TRUE WHERE CONDITION = 2 AND START_PERIOD IS NOT NULL AND END_PERIOD IS NOT NULL AND CUSTOM_TIME BETWEEN START_PERIOD AND END_PERIOD;
  UPDATE SURVEYS
  SET TOTAL_RESPONDENTS = EMPLOYEES
  FROM (
         SELECT
           SURVEY_ID,
           COUNT(EMPLOYEE) AS EMPLOYEES
         FROM
           SURVEYS_EMPLOYEES_RELATIONSHIP
         GROUP BY
           SURVEY_ID
       ) SURVEYS_EMPLOYEES_RELATIONSHIP
  WHERE
      SURVEYS.ID = SURVEYS_EMPLOYEES_RELATIONSHIP.SURVEY_ID;
  UPDATE SURVEYS
  SET PAST_RESPONDENTS = EMPLOYEES
  FROM (
         SELECT
           SURVEY_ID,
           COUNT(EMPLOYEE) AS EMPLOYEES
         FROM
           SURVEYS_EMPLOYEES_RELATIONSHIP
         WHERE
             STATUS = TRUE
         GROUP BY
           SURVEY_ID
       ) SURVEYS_EMPLOYEES_RELATIONSHIP
  WHERE
      SURVEYS.ID = SURVEYS_EMPLOYEES_RELATIONSHIP.SURVEY_ID;
END;
$FUNCTION$ LANGUAGE plpgsql;

/*
 * Удалить процедуру "tracker".
 */
DROP PROCEDURE IF EXISTS tracker(TIMESTAMP);

/*
 * Вызов процедуры "tracker".
 */
CALL tracker('2019-03-29 16:37:00');

/*
 * Создать UNIQUE CONSTRAINT для таблицы "SURVEYS_EMPLOYEES_RELATIONSHIP".
 */
ALTER TABLE SURVEYS_EMPLOYEES_RELATIONSHIP ADD CONSTRAINT SURVEYS_EMPLOYEES_RELATIONSHIP_UNIQUE_KEY UNIQUE (SURVEY_ID, EMPLOYEE);

/*
 * Создать процедуру для создания актуальных (динамичных) данных в таблице "SURVEYS_EMPLOYEES_RELATIONSHIP".
 */
CREATE OR REPLACE PROCEDURE creator(SURVEY_IDENTIFIER uuid, EMPLOYEES VARCHAR[]) AS $FUNCTION$
  BEGIN
    DELETE FROM SURVEYS_EMPLOYEES_RELATIONSHIP
    WHERE SURVEY_ID = SURVEY_IDENTIFIER
    AND EMPLOYEE <> ALL (EMPLOYEES);
    INSERT INTO SURVEYS_EMPLOYEES_RELATIONSHIP (SURVEY_ID, EMPLOYEE)
    SELECT SURVEY_IDENTIFIER SURVEY_ID, EMPLOYEE FROM UNNEST(ARRAY[EMPLOYEES]) EMPLOYEE
    ON CONFLICT ON CONSTRAINT SURVEYS_EMPLOYEES_RELATIONSHIP_UNIQUE_KEY
    DO NOTHING;
  END;
$FUNCTION$ LANGUAGE plpgsql;

/*
 * Удалить процедуру "creator".
 */
DROP PROCEDURE IF EXISTS creator(uuid, varchar[]);

/*
 * Вызов процедуры "creator".
 */
CALL creator('99c89a24-fff2-4cbc-a542-b1e956a352f9', ARRAY['NNogerbek@beeline.kz']);

/*
 * Создать таблицу для email шаблонов.
 */
CREATE TABLE TEMPLATES (
  ID SERIAL NOT NULL PRIMARY KEY,
  TEMPLATE TEXT NOT NULL
);

/*
 * Удалить из таблицы "SURVEYS_FACTORS_RELATIONSHIP" constraint.
 */
ALTER TABLE SURVEYS_FACTORS_RELATIONSHIP DROP CONSTRAINT surveys_factors_relationship_survey_id_fkey;

/*
 * Добавить первичный ключ в таблицу "SURVEYS_FACTORS_RELATIONSHIP" с каскадным удалением.
 */
ALTER TABLE SURVEYS_FACTORS_RELATIONSHIP
ADD CONSTRAINT surveys_factors_relationship_survey_id_fkey
FOREIGN KEY (SURVEY_ID)
REFERENCES SURVEYS(ID)
ON DELETE CASCADE;

/*
 * Удалить из таблицы "SURVEYS_ORGANIZATIONS_RELATIONSHIP" constraint.
 */
ALTER TABLE SURVEYS_ORGANIZATIONS_RELATIONSHIP DROP CONSTRAINT surveys_organizations_relationship_survey_id_fkey;

/*
 * Добавить первичный ключ в таблицу "SURVEYS_ORGANIZATIONS_RELATIONSHIP" с каскадным удалением.
 */
ALTER TABLE SURVEYS_ORGANIZATIONS_RELATIONSHIP
ADD CONSTRAINT surveys_organizations_relationship_survey_id_fkey
FOREIGN KEY (SURVEY_ID)
REFERENCES SURVEYS(ID)
ON DELETE CASCADE;

/*
 * Удалить из таблицы "SURVEYS_ORGANIZATIONS_RELATIONSHIP" constraint.
 */
ALTER TABLE SURVEYS_QUESTIONS_RELATIONSHIP DROP CONSTRAINT surveys_questions_relationship_survey_id_fkey;

/*
 * Добавить первичный ключ в таблицу "SURVEYS_QUESTIONS_RELATIONSHIP" с каскадным удалением.
 */
ALTER TABLE SURVEYS_QUESTIONS_RELATIONSHIP
ADD CONSTRAINT surveys_questions_relationship_survey_id_fkey
FOREIGN KEY (SURVEY_ID)
REFERENCES SURVEYS(ID)
ON DELETE CASCADE;

/*
 * Добавить в таблицу "SURVEYS" столбец "TOTAL_RESPONDENTS" (общее количество респондентов).
 */
ALTER TABLE SURVEYS ADD COLUMN TOTAL_RESPONDENTS INT DEFAULT 0;

/*
 * Добавить в таблицу "SURVEYS" столбец "PAST_RESPONDENTS" (количество прошедших респондентов).
 */
ALTER TABLE SURVEYS ADD COLUMN PAST_RESPONDENTS INT DEFAULT 0;

/*
 * Обновить общее количество респондентов в таблице "SURVEYS".
 */
UPDATE SURVEYS
SET TOTAL_RESPONDENTS = EMPLOYEES
FROM (
       SELECT
         SURVEY_ID,
         COUNT(EMPLOYEE) AS EMPLOYEES
       FROM
         SURVEYS_EMPLOYEES_RELATIONSHIP
       GROUP BY SURVEY_ID
     ) SURVEYS_EMPLOYEES_RELATIONSHIP
WHERE
    SURVEYS.ID = SURVEYS_EMPLOYEES_RELATIONSHIP.SURVEY_ID;

/*
 * Обновить количество прошедших респондентов в таблице "SURVEYS".
 */
UPDATE SURVEYS
SET PAST_RESPONDENTS = EMPLOYEES
FROM (
       SELECT
         SURVEY_ID,
         COUNT(EMPLOYEE) AS EMPLOYEES
       FROM
         SURVEYS_EMPLOYEES_RELATIONSHIP
       WHERE
           STATUS = true
       GROUP BY
         SURVEY_ID
     ) SURVEYS_EMPLOYEES_RELATIONSHIP
WHERE
    SURVEYS.ID = SURVEYS_EMPLOYEES_RELATIONSHIP.SURVEY_ID;

/*
 * В таблицу "questions" добавить новый столбец "REQUIRED".
 */
ALTER TABLE QUESTIONS ADD COLUMN REQUIRED BOOLEAN DEFAULT FALSE;

/*
 * В таблицу "questions" добавить новый столбец "REQUIRED".
 */
ALTER TABLE QUESTIONS ADD COLUMN REQUIRED BOOLEAN DEFAULT FALSE;

/*
 * Удалить из таблицы "QUESTIONS_OPTIONS_RELATIONSHIP" constraint.
 */
ALTER TABLE QUESTIONS_OPTIONS_RELATIONSHIP DROP CONSTRAINT questions_options_relationship_option_id_fkey;

/*
 * Добавить первичный ключ в таблицу "SURVEYS_ORGANIZATIONS_RELATIONSHIP" с каскадным удалением.
 */
ALTER TABLE QUESTIONS_OPTIONS_RELATIONSHIP
ADD CONSTRAINT questions_options_relationship_option_id_fkey
FOREIGN KEY (OPTION_ID)
REFERENCES OPTIONS(ID)
ON DELETE CASCADE;

/*
 * Добавить в таблицу "OPTIONS" столбец "POSITION".
 */
ALTER TABLE OPTIONS ADD COLUMN POSITION INT;

/*
 * Добавить в таблицу "QUESTIONS" столбец "POSITION".
 */
ALTER TABLE QUESTIONS ADD COLUMN POSITION INT;

/*
 * Функция "factorio" создает вопросы в таблице "SURVEYS" и связь между фактором и новыми созданными вопросами в таблице "FACTORS_QUESTIONS_RELATIONSHIP".
 */
CREATE OR REPLACE FUNCTION factorio (
  FACTOR_IDENTIFIER INT,
  TEXT_ARRAY VARCHAR[],
  WIDGET_ARRAY INT[],
  REQUIRED_ARRAY BOOLEAN[],
  POSITION_ARRAY INT[]
) RETURNS SETOF INT AS $$
BEGIN
  RETURN QUERY
    WITH RESULTS AS (
      INSERT INTO QUESTIONS (TEXT, WIDGET, REQUIRED, POSITION, CATEGORY)
        SELECT
          UNNEST(ARRAY[TEXT_ARRAY]) AS TEXT,
          UNNEST(ARRAY[WIDGET_ARRAY]) AS WIDGET,
          UNNEST(ARRAY[REQUIRED_ARRAY]) AS REQUIRED,
          UNNEST(ARRAY[POSITION_ARRAY]) AS POSITION,
          1 AS CATEGORY
        RETURNING ID
    ),
         GENERATE_FACTORS_QUESTIONS_RELATIONSHIP AS
           (
             INSERT INTO FACTORS_QUESTIONS_RELATIONSHIP (FACTOR_ID, QUESTION_ID)
             SELECT FACTOR_IDENTIFIER, ID
             FROM RESULTS
             ON CONFLICT ON CONSTRAINT FACTORS_QUESTIONS_RELATIONSHIP_UNIQUE_KEY DO NOTHING
           )
    SELECT ID FROM RESULTS;
END;
$$ LANGUAGE plpgsql;

/*
 * Удалить функцию "factorio".
 */
DROP FUNCTION factorio (INT, VARCHAR[], INT[], BOOLEAN[], INT[]);

/*
 * Функция "alexa" создает варианты ответов в таблице "OPTIONS" и связь между фактором и новыми созданными вопросами в таблице "QUESTIONS_OPTIONS_RELATIONSHIP".
 */
CREATE OR REPLACE FUNCTION alexa (
  QUESTION_IDENTIFIER INT,
  TEXT_ARRAY VARCHAR[],
  POSITION_ARRAY INT[]
) RETURNS SETOF INT AS $$
BEGIN
  RETURN QUERY
    WITH RESULTS AS (
      INSERT INTO OPTIONS (TEXT, POSITION)
        SELECT
          UNNEST(ARRAY[TEXT_ARRAY]) AS TEXT,
          UNNEST(ARRAY[POSITION_ARRAY]) AS POSITION
        RETURNING ID
    ),
         GENERATE_QUESTIONS_OPTIONS_RELATIONSHIP AS
           (
             INSERT INTO QUESTIONS_OPTIONS_RELATIONSHIP (QUESTION_ID, OPTION_ID)
             SELECT QUESTION_IDENTIFIER, ID
             FROM RESULTS
             ON CONFLICT ON CONSTRAINT QUESTIONS_OPTIONS_RELATIONSHIP_UNIQUE_KEY DO NOTHING
           )
    SELECT ID FROM RESULTS;
END;
$$ LANGUAGE plpgsql;

/*
 * Удалить функцию "alexa".
 */
DROP FUNCTION alexa (INT, VARCHAR[], INT[]);

/*
 * Удалить из таблицы "SURVEYS_ORGANIZATIONS_RELATIONSHIP" constraint.
 */
ALTER TABLE FACTORS_QUESTIONS_RELATIONSHIP DROP CONSTRAINT factors_questions_relationship_question_id_fkey;

/*
 * Добавить первичный ключ в таблицу "FACTORS_QUESTIONS_RELATIONSHIP" с каскадным удалением.
 */
ALTER TABLE FACTORS_QUESTIONS_RELATIONSHIP
ADD CONSTRAINT factors_questions_relationship_question_id_fkey
FOREIGN KEY (QUESTION_ID)
REFERENCES QUESTIONS(ID)
ON DELETE CASCADE;

/*
 * Добавить первичный ключ в таблицу "SURVEYS_QUESTIONS_RELATIONSHIP" с каскадным удалением.
 */
ALTER TABLE SURVEYS_QUESTIONS_RELATIONSHIP
ADD CONSTRAINT surveys_questions_relationship_question_id_fkey
FOREIGN KEY (QUESTION_ID)
REFERENCES QUESTIONS(ID)
ON DELETE CASCADE;

/*
 * Данная процедура удаляет вопросы и ответы, которые внутри фактора.
 */
CREATE OR REPLACE PROCEDURE proper(FACTOR_IDENTIFIER INT) AS $FUNCTION$
BEGIN
  DELETE FROM OPTIONS WHERE ID IN (
    SELECT OPTION_ID FROM QUESTIONS_OPTIONS_RELATIONSHIP WHERE QUESTION_ID IN (
      SELECT QUESTION_ID FROM FACTORS_QUESTIONS_RELATIONSHIP WHERE FACTOR_ID = FACTOR_IDENTIFIER
    )
  );
  DELETE FROM QUESTIONS WHERE ID IN (
    SELECT QUESTION_ID FROM FACTORS_QUESTIONS_RELATIONSHIP WHERE FACTOR_ID = FACTOR_IDENTIFIER
  );
END;
$FUNCTION$ LANGUAGE plpgsql;

/*
 * Вызвать процедуру "proper".
 */
CALL proper(15);

/*
 * Данная процедура удаляет вопросы и ответы, которые внутри фактора и сам фактор.
 */
CREATE OR REPLACE PROCEDURE tide(FACTOR_IDENTIFIER INT) AS $FUNCTION$
BEGIN
  DELETE FROM OPTIONS WHERE ID IN (
    SELECT OPTION_ID FROM QUESTIONS_OPTIONS_RELATIONSHIP WHERE QUESTION_ID IN (
      SELECT QUESTION_ID FROM FACTORS_QUESTIONS_RELATIONSHIP WHERE FACTOR_ID = FACTOR_IDENTIFIER
    )
  );
  DELETE FROM QUESTIONS WHERE ID IN (
    SELECT QUESTION_ID FROM FACTORS_QUESTIONS_RELATIONSHIP WHERE FACTOR_ID = FACTOR_IDENTIFIER
  );
  DELETE FROM FACTORS WHERE ID = FACTOR_IDENTIFIER;
END;
$FUNCTION$ LANGUAGE plpgsql;

/*
 * Вызвать процедуру "tide".
 */
CALL tide(15);

/*
 * Создать UNIQUE CONSTRAINT для таблицы "SURVEYS_FACTORS_RELATIONSHIP".
 */
ALTER TABLE SURVEYS_FACTORS_RELATIONSHIP ADD CONSTRAINT SURVEYS_FACTORS_RELATIONSHIP_UNIQUE_KEY UNIQUE (SURVEY_ID, FACTOR_ID);

/*
 * Создать UNIQUE CONSTRAINT для таблицы "SURVEYS_ORGANIZATIONS_RELATIONSHIP".
 */
ALTER TABLE SURVEYS_ORGANIZATIONS_RELATIONSHIP ADD CONSTRAINT SURVEYS_ORGANIZATIONS_RELATIONSHIP_UNIQUE_KEY UNIQUE (SURVEY_ID, ORGANIZATION_ID);

/*
 * Создать UNIQUE CONSTRAINT для таблицы "SURVEYS_QUESTIONS_RELATIONSHIP".
 */
ALTER TABLE SURVEYS_QUESTIONS_RELATIONSHIP ADD CONSTRAINT SURVEYS_QUESTIONS_RELATIONSHIP_UNIQUE_KEY UNIQUE (SURVEY_ID, QUESTION_ID);

/*
 * Создать UNIQUE CONSTRAINT для таблицы "FACTORS_QUESTIONS_RELATIONSHIP".
 */
ALTER TABLE FACTORS_QUESTIONS_RELATIONSHIP ADD CONSTRAINT FACTORS_QUESTIONS_RELATIONSHIP_UNIQUE_KEY UNIQUE (FACTOR_ID, QUESTION_ID);

/*
 * Создать UNIQUE CONSTRAINT для таблицы "QUESTIONS_OPTIONS_RELATIONSHIP".
 */
ALTER TABLE QUESTIONS_OPTIONS_RELATIONSHIP ADD CONSTRAINT QUESTIONS_OPTIONS_RELATIONSHIP_UNIQUE_KEY UNIQUE (QUESTION_ID, OPTION_ID);

/*
 * Создать таблицу с названием "ANSWERS".
 */
CREATE TABLE ANSWERS (
  SURVEY_ID UUID,
  CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW(),
  EMPLOYEE VARCHAR,
  QUESTION_ID INT,
  QUESTION_TEXT TEXT,
  OPTION_ID INT,
  OPTION_TEXT TEXT
);

/*
 * Создать новую базу данных "questionnaire_production" идентичную по структуре c базой данных "questionnaire_development".
 */
SELECT pg_terminate_backend(pid) from pg_stat_activity WHERE pid <> pg_backend_pid();
CREATE DATABASE questionnaire_production WITH TEMPLATE questionnaire_development;

/*
 * Добавить FOREIGN KEY столбец "CATEGORY" в таблицу "QUESTIONS".
 */
ALTER TABLE QUESTIONS ADD CATEGORY INT;
ALTER TABLE QUESTIONS ADD CONSTRAINT QUESTIONS_CATEGORY_FKEY FOREIGN KEY (CATEGORY) REFERENCES CATEGORIES(ID);