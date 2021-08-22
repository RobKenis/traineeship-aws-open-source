FROM golang:rc-alpine AS build
ENV CGO_ENABLED=0
WORKDIR /src
COPY . .
RUN go build -o /out/traineeship *.go

FROM scratch
COPY --from=build /out/traineeship /
ENV PORT 8080
ENTRYPOINT ["/traineeship"]
