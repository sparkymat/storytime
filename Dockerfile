FROM alpine:3
COPY storytime /bin/storytime
RUN mkdir /app
WORKDIR /app
EXPOSE 8080
CMD ["/bin/storytime"]
