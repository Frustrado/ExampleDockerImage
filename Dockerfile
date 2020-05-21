FROM golang:latest as builder
COPY . /auiapp
WORKDIR /auiapp
RUN make modules
RUN make build

FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /auiapp/bin/auiapp /aui
CMD ["/aui"]