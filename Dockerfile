FROM scratch

COPY cmd/cmd .

EXPOSE 8101

ENTRYPOINT [ "./cmd" ]