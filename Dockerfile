FROM node:alpine AS npm_builder
WORKDIR /app
COPY ui ./ui
RUN npm --prefix /app/ui install
RUN npm --prefix /app/ui run build

FROM golang:alpine AS go_builder
WORKDIR /app
RUN apk update && apk add --no-cache git
COPY cmd ./cmd
COPY go.mod ./go.mod
COPY internal ./internal
COPY --from=npm_builder /app/ui/dist ./cmd/shokku/dist
RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/shokku ./cmd/shokku

FROM gcr.io/distroless/static:nonroot
COPY --from=go_builder /go/bin/shokku /go/bin/shokku
ENTRYPOINT ["/go/bin/shokku"]