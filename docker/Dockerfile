FROM alpine

ADD https://github.com/gohugoio/hugo/releases/download/v0.30.2/hugo_0.30.2_Linux-64bit.tar.gz /hugo.tar.gz
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh

RUN apk update && apk add nodejs-npm git && npm install -g firebase-tools \
  && chmod +x /usr/local/bin/docker-entrypoint.sh \
  && tar -xvzf hugo.tar.gz && mv hugo /usr/local/bin/hugo && rm hugo.tar.gz *.md

ENTRYPOINT ["docker-entrypoint.sh"]
