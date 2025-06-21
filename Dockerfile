FROM golang:alpine

ARG APP_NAME
ARG MAIN_FILE

RUN echo "APP_NAME: ${APP_NAME}, MAIN_FILE: ${MAIN_FILE}"

WORKDIR /usr/service

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./bin/${APP_NAME} ./cmd/${MAIN_FILE}

RUN chmod +x ./bin/${APP_NAME}
