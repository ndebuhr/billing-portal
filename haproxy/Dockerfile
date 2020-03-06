FROM haproxytech/haproxy-debian:2.0

# Setup the haproxy configuration
COPY ./haproxy.cfg /etc/haproxy/haproxy.cfg

CMD []
ENTRYPOINT ["/usr/local/sbin/haproxy", "-d", "-f", "/etc/haproxy/haproxy.cfg"]