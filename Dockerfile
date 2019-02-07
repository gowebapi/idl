#
# Used to verify that the public repository is buildable.
#
FROM golang:1.11-stretch
RUN go get github.com/gowebapi/webidlgenerator
RUN go install github.com/gowebapi/webidlgenerator
RUN mkdir -p output
ADD . idl
RUN find idl/webapi -type f -name "*.idl" -or -name "*.go.md" | xargs webidlgenerator -single-package webapiall -log-warning=false -output output
