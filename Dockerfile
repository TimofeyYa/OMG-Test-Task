FROM ubuntu:22.04

ADD ./bin/app /app
ADD ./.env /.env

CMD ["/app"]
