FROM ubuntu:22.04
ARG TARGETARCH

ENV CNI_BIN_SRC /opt/cni/bin

ADD /cni-bin-$TARGETARCH /opt/cni/bin
ADD /bin/copier-$TARGETARCH copier

ENTRYPOINT ["/copier"]
