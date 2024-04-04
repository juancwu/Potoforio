FROM node:20-alpine AS styles
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
WORKDIR /app
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --prod --frozen-lockfile
RUN pnpm run tw:prod

FROM golang:1.22.0 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/potoforio .

FROM alpine as runner
WORKDIR /go
COPY --from=builder /app/build/potoforio ./
COPY --from=styles /app/static ./static
COPY --from=styles /app/views ./views

EXPOSE 5173

ENV GO_ENV="production"

ENTRYPOINT ["./potoforio"]
