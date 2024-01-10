FROM ubuntu:22.04

ADD ./bin/app /app
ADD ./.env /.env

ENV GIN_MODE release
CMD ["/app"]
