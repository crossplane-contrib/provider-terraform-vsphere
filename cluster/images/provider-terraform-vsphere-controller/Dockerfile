FROM golang:1.16 AS tf-builder

ADD tf-provider.yaml .
RUN go get github.com/crossplane-contrib/terraform-provider-dl
RUN $GOPATH/bin/terraform-provider-dl --config=tf-provider.yaml --output=/tf-plugin
RUN chmod -R +x /tf-plugin

FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash

ARG ARCH
ARG TINI_VERSION

ADD provider-terraform-vsphere /usr/local/bin/crossplane-provider-terraform-vsphere

COPY --from=tf-builder /tf-plugin ./tf-plugin

EXPOSE 8080
USER 1001
ENTRYPOINT ["crossplane-provider-terraform-vsphere", "--pluginPath", "./tf-plugin"]
