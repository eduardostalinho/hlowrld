FROM scratch
MAINTAINER Eduardo Carvalho "eduardooc.86@gmail.com"
WORKDIR /app
ADD hlowrld .
ENTRYPOINT ["./hlowrld"]
EXPOSE 25478
