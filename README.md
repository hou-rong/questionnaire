# Description

`questionnaire` is a RESTful web service which is part of the project called `Questionnaire`.

# Technical Details

`questionnaire` application implemented on `Golang` programming language (`1.11.5`). Application use several third-party libraries:

* [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang.
* [Gorilla Handlers](https://github.com/gorilla/handlers) - A collection of useful handlers for `"net/http"` package.
* [pq](https://github.com/lib/pq) - Pure Golang Postgres driver for `"database/sql"` package.
* [godotenv](https://github.com/joho/godotenv) - A Golang port of the Ruby dotenv project which loads environment variables from a `.env` file.
* [GORM](http://gorm.io/) - ORM library for Golang.

# APIs
RESTful web service consists of several URLs which we can group by categories.

A. *Category «Factors»*

1. `GET` request to the route `/api/factors` return list of all available factors.

2. `POST` request to the route `/api/factor` create the new factor.
    
3. `GET` request to the route `/api/factor/{factor_id:[0-9]+}` return information about the specific factor.

4. `DELETE` request to the route `/api/factor/{factor_id:[0-9]+}` delete the specific factor.

5. `PUT` request to the route `/api/factor/{factor_id:[0-9]+}` update information of the specific factor.

Set `Content-Type` as `application/json` in `POST` and `PUT` requests. In the body of these requests send JSON object with information about the specific factor.


# License
Copyright 2019, Kazakhstan, LLP «КаР-Тел» (Telecommunication Company «Beeline»). All rights reserved.