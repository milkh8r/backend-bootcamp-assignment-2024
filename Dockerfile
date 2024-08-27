FROM ubuntu:latest
LABEL authors="milkhater"

ENTRYPOINT ["top", "-b"]