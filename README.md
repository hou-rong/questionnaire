# Description

`questionnaire` is a RESTful web service which is part of the project called `Questionnaire`.

# Technical Details

`questionnaire` application implemented on `Golang` programming language (`1.11.5`). Application use several third-party libraries:

* [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang.
* [Gorilla Handlers](https://github.com/gorilla/handlers) - A collection of useful handlers for `"net/http"` package.
* [pq](https://github.com/lib/pq) - Pure Golang Postgres driver for `"database/sql"` package.
* [godotenv](https://github.com/joho/godotenv) - A Golang port of the Ruby dotenv project which loads environment variables from a `.env` file.
* [GORM](http://gorm.io/) - ORM library for Golang.
* [goracle](https://github.com/go-goracle/goracle) - Oracle driver for Go, using the ODPI-C driver.

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


# Database Details

If you get `ERROR: duplicate key violates unique constraint` message when trying to insert data into a `questions` table of the `PostgreSQL` database, try to execute such sql statement to fix error:
```
SELECT MAX(id) FROM questions;

SELECT nextval('questions_id_seq');

SELECT setval('questions_id_seq', (SELECT MAX(id) FROM options) + 1);
```

# Production

Inside `build.sh` file you can notice the instruction which you can use to create the `docker image`:
```
docker build -t saturn_backend_image .
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
  go get github.com/pkg/errors

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
sudo docker run --name questionnaire_container -d -p 1000:8000 questionnaire_image
```

# License
Copyright 2019, Kazakhstan, LLP «КаР-Тел» (Telecommunication Company «Beeline»). All rights reserved.