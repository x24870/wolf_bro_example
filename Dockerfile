FROM scratch

COPY example /example

ENV TZ=Asia/Taipei

ENTRYPOINT ["/example"]

