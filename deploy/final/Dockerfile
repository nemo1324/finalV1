# ---------- Build stage ----------
ARG GO_VER=1.21
ARG ALPINE_VER=3.18

FROM golang:${GO_VER}-alpine${ALPINE_VER} AS builder

WORKDIR /src
COPY . .

RUN apk add --no-cache git

# Настройки для приватных реп (если нужны)
# ARG CI_JOB_TOKEN
# RUN go env -w GOPRIVATE='gitlab.wildberries.ru' && \
#     printf "machine gitlab.wildberries.ru\nlogin gitlab-ci-token\npassword ${CI_JOB_TOKEN}\n" > ~/.netrc && \
#     git config --global --add url."https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.wildberries.ru/".insteadOf "https://gitlab.wildberries.ru/"

# Сборка
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -installsuffix cgo -o app ./cmd/auth

# ---------- Runtime stage ----------
FROM alpine:${ALPINE_VER}

WORKDIR /root/
COPY --from=builder /src/app .

EXPOSE 9007

CMD ["./app"]
