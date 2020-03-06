FROM buildpack-deps:18.04 AS build

# HUGO extended install
RUN curl -s https://api.github.com/repos/gohugoio/hugo/releases/latest | \
    grep "browser_download_url" | grep "hugo_extended" | grep "Linux-64bit.tar.gz" | \
    cut -d : -f 2,3 | tr -d \" | \
    wget -qi - && \
    tar xfvz hugo*.tar.gz hugo && \
    install hugo /usr/local/bin/

# Setup the source code for build
COPY . /site
WORKDIR /site

# Build
RUN /hugo --minify

# Copy over build artifacts to a caddy web server
FROM abiosoft/caddy:1.0.3

COPY --from=build /site/public /srv
COPY --from=build /site/Caddyfile /etc/Caddyfile