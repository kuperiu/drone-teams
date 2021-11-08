FROM golang:1.15.0-alpine AS build
ENV CGO_ENABLED="0"
ENV GOOS="linux"
ENV GOARCH="amd64"
WORKDIR /src
COPY . .
RUN go build -o /out/app .
FROM scratch AS bin
COPY --from=build /out/app /bin/app
ENTRYPOINT [ "/bin/app" ]
