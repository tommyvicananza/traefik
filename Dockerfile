# to use your binary please first create a statically linked
# run in your terminal
# `CGO_ENABLED=0 go build`
FROM scratch
COPY script/ca-certificates.crt /etc/ssl/certs/
COPY traefik /
EXPOSE 3333 8000
ENTRYPOINT ["/traefik"]
