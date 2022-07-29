FROM scratch

COPY example /example
COPY fonts /fonts

ENV TZ=Asia/Taipei

ENTRYPOINT ["/example"]

