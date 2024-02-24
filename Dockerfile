FROM golang:1.21 as build_image

WORKDIR /app

ADD Makefile go.mod go.sum /app/
RUN make dependencies
ADD . /app/
RUN make build

FROM ubuntu:22.04 as runner

COPY --from=build_image /app/cocom /bin/cocom

EXPOSE 8080
ENTRYPOINT ["/bin/cocom"]
CMD ["version"]
