# sudo docker build --tag es-data-collector .
# sudo docker rm -f es-data-collector-app
# sudo docker run --name es-data-collector-app -p 8080:8080 -p 3000:3000 -e ES_DATA_COLLECTOR_ELASTICSEARH_ADDRESS='http://192.168.1.52:9200' -e ES_DATA_COLLECTOR_APP_DB_CONNECTION_STRING='root:root@tcp(192.168.1.52:3306)/es-data-collector?parseTime=true' es-data-collector

FROM golang:1.21.6-alpine3.19

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
ADD . /app

RUN go mod download

RUN go env -w GO111MODULE=on

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /es_data_collector

RUN apk update
RUN apk add nodejs npm
RUN cd ui && npm install && npm run build && cd ..

RUN chmod +x start.sh

EXPOSE 8080
EXPOSE 3000

# Run
CMD ["sh","./start.sh"]