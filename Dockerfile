FROM scratch
COPY ./go-env-var-http /
WORKDIR /
ENTRYPOINT ["/go-env-var-http"]