# Description

`questionnaire` is a web service which is part of the project called `Questionnaire`.

# Technical Details

`questionnaire` application implemented on `Golang` programming language (`1.11.5`). Application use several third-party libraries:

* [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang.
* [Gorilla Handlers](https://github.com/gorilla/handlers) - A collection of useful handlers for `"net/http"` package.
* [pq](https://github.com/lib/pq) - Pure Golang Postgres driver for `"database/sql"` package.
* [godotenv](https://github.com/joho/godotenv) - A Golang port of the Ruby dotenv project which loads environment variables from a `.env` file.
* [GORM](http://gorm.io/) - ORM library for Golang.
* [goracle](https://github.com/go-goracle/goracle) - Oracle driver for Go, using the ODPI-C driver.
* [crontab](https://github.com/mileusna/crontab) - Golang crontab ticker.

Instruction to cross compile Golang application with CGO packages from `windows/x64` to `linux/amd64`:
- Install [tdm-gcc](http://tdm-gcc.tdragon.net/download). A compiler suite for 32- and 64-bit Windows based on the GNU toolchain.
- In `POWERSHELL` execute commands:
  ```
  set GOOS=linux
  set GOARCH=amd64
  go build -v -o questionnaire -ldflags="-extld=$CC"
  ```

# APIs
RESTful web service consists of several URLs which we can group by categories.

A. *Category «Factors»*

1. `GET` request to the route `/api/factors` return list of all available factors.

2. `POST` request to the route `/api/factor` create the new factor.
    
3. `GET` request to the route `/api/factor/{factor_id:[0-9]+}` return information about the specific factor.

4. `DELETE` request to the route `/api/factor/{factor_id:[0-9]+}` delete the specific factor.

5. `PUT` request to the route `/api/factor/{factor_id:[0-9]+}` update information of the specific factor.

Set `Content-Type` as `application/json` in `POST` and `PUT` requests. In the body of these requests send JSON object with information about the specific factor.

List of all available routes of the project you can see in `routes/routes.go` file. 

# CRONTAB

Several scripts are used to update and synchronize data in the sources:
* `controllers/creator.go`
* `controllers/tracker.go`

Both of these scripts run every minute (`* * * * *`).

# Database Details

If you get `ERROR: duplicate key violates unique constraint` message when trying to insert data into a `questions` table of the `PostgreSQL` database, try to execute such sql statement to fix error:
```
SELECT MAX(id) FROM questions;

SELECT nextval('questions_id_seq');

SELECT setval('questions_id_seq', (SELECT MAX(id) FROM options) + 1);
```

`СHECK_FOR_UPDATES_IN_SURVEYS` trigger:
```
CREATE TRIGGER СHECK_FOR_UPDATES_IN_SURVEYS
BEFORE UPDATE ON SURVEYS
FOR EACH ROW
WHEN (
    OLD.UPDATED_AT IS DISTINCT FROM NEW.UPDATED_AT
) EXECUTE PROCEDURE CREATE_SURVEYS_QUESTIONS_RELATIONSHIP();
```

`CREATE_SURVEYS_QUESTIONS_RELATIONSHIP` procedure:
```
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
```

`tracker` procedure:
```
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
```

`creator` procedure:
```
CREATE OR REPLACE PROCEDURE creator(SURVEY_IDENTIFIER uuid, EMPLOYEES VARCHAR[]) AS $FUNCTION$
    BEGIN
        DELETE FROM SURVEYS_EMPLOYEES_RELATIONSHIP
        WHERE SURVEY_ID = SURVEY_IDENTIFIER
        AND EMPLOYEE <> ALL (EMPLOYEES);
        INSERT INTO SURVEYS_EMPLOYEES_RELATIONSHIP (SURVEY_ID, EMPLOYEE) 
        SELECT SURVEY_IDENTIFIER SURVEY_ID, EMPLOYEE FROM UNNEST(ARRAY[EMPLOYEES]) EMPLOYEE
        ON CONFLICT ON CONSTRAINT unique_key 
        DO NOTHING;
    END;
$FUNCTION$ LANGUAGE plpgsql;
```

`factorio` function:
```
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
        INSERT INTO QUESTIONS (TEXT, WIDGET, REQUIRED, POSITION)
        SELECT
            UNNEST(ARRAY[TEXT_ARRAY]) AS TEXT,
            UNNEST(ARRAY[WIDGET_ARRAY]) AS WIDGET,
            UNNEST(ARRAY[REQUIRED_ARRAY]) AS REQUIRED,
            UNNEST(ARRAY[POSITION_ARRAY]) AS POSITION
        RETURNING ID
    ),
    GENERATE_FACTORS_QUESTIONS_RELATIONSHIP AS (
        INSERT INTO FACTORS_QUESTIONS_RELATIONSHIP (FACTOR_ID, QUESTION_ID)
        SELECT FACTOR_IDENTIFIER, ID
        FROM RESULTS
        ON CONFLICT ON CONSTRAINT FACTORS_QUESTIONS_RELATIONSHIP_UNIQUE_KEY DO NOTHING
    )
    SELECT ID FROM RESULTS;
END;
$$ LANGUAGE plpgsql;
```

`alexa` function:
```
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
```

`proper` procedure:
```
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
```

List of all available sql queries of the project you can easily find in `queries.sql` file.

# Production

Inside `build.sh` file you can notice the instruction which you can use to create the `docker image`:
```
docker build -t questionnaire_image .
```

`Dockerfile`:
```
FROM golang:1.11

RUN echo 'Acquire::http::proxy "http://pxuser:Hejdxgh7265@172.28.59.42:3128";' >> /etc/apt/apt.conf
RUN echo 'Acquire::https::proxy "http://pxuser:Hejdxgh7265@172.28.59.42:3128";' >> /etc/apt/apt.conf
RUN echo 'Acquire::ftp::proxy "http://pxuser:Hejdxgh7265@172.28.59.42:3128";' >> /etc/apt/apt.conf

ENV HTTP_PROXY "http://pxuser:Hejdxgh7265@172.28.59.42:3128"
ENV HTTPS_PROXY "https://pxuser:Hejdxgh7265@172.28.59.42:3128"
ENV FTP_PROXY "http://pxuser:Hejdxgh7265@172.28.59.42:3128"

RUN go get github.com/gorilla/mux && \
  go get github.com/gorilla/handlers && \
  go get github.com/lib/pq && \
  go get github.com/joho/godotenv && \
  go get github.com/jinzhu/gorm && \
  go get github.com/pkg/errors && \
  go get github.com/mileusna/crontab

RUN apt-get update && apt-get install -y libaio1 build-essential unzip curl vim

COPY ./dependencies/oracle/instantclient-basic-linux.x64-12.2.0.1.0.zip .
COPY ./dependencies/oracle/instantclient-sdk-linux.x64-12.2.0.1.0.zip .
COPY ./dependencies/oracle/instantclient-sqlplus-linux.x64-12.2.0.1.0.zip .

RUN unzip -qq instantclient-basic-linux.x64-12.2.0.1.0.zip -d /opt/oracle
RUN unzip -qq instantclient-sdk-linux.x64-12.2.0.1.0.zip -d /opt/oracle
RUN unzip -qq instantclient-sqlplus-linux.x64-12.2.0.1.0.zip -d /opt/oracle

RUN mkdir -p /opt/oracle/instantclient_12_2/bin
RUN mv /opt/oracle/instantclient_12_2/sqlplus /opt/oracle/instantclient_12_2/bin
RUN echo /opt/oracle/instantclient_12_2 > /etc/ld.so.conf.d/oracle-instantclient.conf
ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_12_2:$LD_LIBRARY_PATH
ENV ORACLE_HOME=/opt/oracle/instantclient_12_2
ENV PATH=$PATH:$ORACLE_HOME/bin

WORKDIR /go/src/questionnaire

ADD . .

RUN go build -o questionnaire

ENV PORT=8000

CMD ["./questionnaire"]
```

Inside `run.sh` file you can notice the instruction which you can use to create the `docker container`:
```
docker run --name questionnaire_container -e TZ=Asia/Almaty -d -p 1000:8000 questionnaire_image
```

# License
Copyright 2019, Kazakhstan, LLP «КаР-Тел» (Telecommunication Company «Beeline»). All rights reserved.