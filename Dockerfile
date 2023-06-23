FROM golang:1.18

WORKDIR /app
ARG CONSUMER

RUN cd ..
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

ENTRYPOINT ["/usr/local/go/bin/go", "run", "./aapointments_scheduler/main.go"]
