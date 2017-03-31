FROM golang:1.8.0
EXPOSE 6061
COPY . xml2json-service
RUN go get github.com/basgys/goxml2json
RUN cd xml2json-service/main && go build main.go
ENTRYPOINT cd xml2json-service/main && ./main