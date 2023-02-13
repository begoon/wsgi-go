FROM gcr.io/distroless/static-debian11
#COPY --from=build /go/bin/app /
COPY exe /exe
CMD ["/exe", "serve"]

