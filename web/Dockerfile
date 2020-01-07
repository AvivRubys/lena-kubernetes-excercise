FROM golang:1.13-buster as build
WORKDIR /src

COPY go.mod go.sum /src/
RUN go mod download

COPY . /src/
RUN go build -o web .

FROM gcr.io/distroless/base-debian10
WORKDIR /service
COPY --from=build /src/web ./web
COPY --from=build /src/pages ./pages
EXPOSE 8080
CMD ["./web"]