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