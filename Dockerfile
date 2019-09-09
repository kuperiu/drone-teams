FROM golang:1.12

LABEL maintainer="Drone.IO Community <drone-dev@googlegroups.com>" \
  org.label-schema.name="Drone Teams" \
  org.label-schema.vendor="Drone.IO Community" \
  org.label-schema.schema-version="1.0"

ADD drone-teams /bin/
ENTRYPOINT ["/bin/drone-teams"]
