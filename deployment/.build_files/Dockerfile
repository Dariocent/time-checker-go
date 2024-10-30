# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY src/cmd/cmd /time-checker-go

ENV ADDR=":8080"

EXPOSE 8080

# USER nonroot:nonroot

ENTRYPOINT ["/time-checker-go"]