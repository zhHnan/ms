FROM alpine:latest
WORKDIR /Initial
COPY ./target/ms-project .
COPY ../ms-project/config/config-docker.yaml .
RUN mkdir config && mv config-docker.yaml config/config.yaml
EXPOSE 8082 8882
ENTRYPOINT ["./ms-project"]