FROM scratch
COPY ./ddnsv6 /ddnsv6
ENTRYPOINT ["/ddnsv6"]