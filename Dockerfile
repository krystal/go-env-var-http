FROM scratch
COPY ./go-env-var-http /
ENV PORT 8080
EXPOSE 8080
WORKDIR /
ENTRYPOINT ["/go-env-var-http"]