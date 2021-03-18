FROM golang

# copy startscript

COPY docker-entrypoint.sh /

# download and compile vcsim
RUN go get -u github.com/vmware/govmomi/vcsim

# run start command
ENTRYPOINT ["/docker-entrypoint.sh"]
