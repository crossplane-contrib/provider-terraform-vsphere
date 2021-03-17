FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash

ARG ARCH
ARG TINI_VERSION

ADD provider-terraform-vsphere /usr/local/bin/crossplane-provider-terraform-vsphere

EXPOSE 8080
USER 1001
ENTRYPOINT ["crossplane-provider-terraform-vsphere"]
