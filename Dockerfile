FROM alpine:latest

WORKDIR /bin

COPY ZSTUCA_HardwareRepair .
COPY webapp webapp
COPY conf.json .

RUN apk add gcompat

EXPOSE 25555

CMD ["ZSTUCA_HardwareRepair"]