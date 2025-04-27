FROM alpine:latest
WORKDIR /Initial
COPY ./target/ms-user .
COPY ../ms-user/config/config-docker.yaml .
RUN mkdir config && mv config-docker.yaml config/config.yaml
EXPOSE 8080 8881
ENTRYPOINT ["./ms-user"]