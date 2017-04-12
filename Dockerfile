FROM scratch
ADD fizzbuzz /
ENTRYPOINT ["/fizzbuzz"]
CMD ["help"]
