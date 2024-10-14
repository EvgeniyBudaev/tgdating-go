FROM postgres:16.4

RUN apt-get update \
    && apt-get install wget -y \
    && apt-get install postgresql-16-postgis-3 -y \
    && apt-get install postgis -y

CMD ["postgres"]