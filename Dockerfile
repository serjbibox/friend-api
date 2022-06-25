FROM golang:latest AS build_stage
RUN mkdir -p go/src/friend-api
RUN mkdir -p go/bin/configs
WORKDIR /go/src/friend-api
COPY ./ ./

RUN go env -w GO111MODULE=auto
RUN go install .
COPY /configs /go/bin/configs
WORKDIR /go/bin/configs

FROM ubuntu:latest
RUN mkdir -p friend-api
WORKDIR /friend-api
COPY --from=build_stage /go/bin .
ENV FRIEND_DB_PASS "3ee99518a80428469b64240386d73adf85886639219572163c2b02d0d65c0803"
EXPOSE 8080 5432
ENTRYPOINT ./friend-api