# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /invoices-manager

COPY go.mod go.sum Makefile ./  

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/invoices-manager cmd/*.go

RUN chmod a+x /bin/invoices-manager

CMD ["/bin/invoices-manager"]
