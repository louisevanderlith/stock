FROM alpine:3.12.0

COPY cmd/cmd .

EXPOSE 8101

ENTRYPOINT [ "./cmd" ]