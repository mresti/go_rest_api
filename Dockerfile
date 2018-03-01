FROM golang:1.10-alpine

ENV PATH /api:$PATH

ARG API_PORT
ENV PORT $API_PORT

WORKDIR /api

# Copy binary titan
ADD ./bin/api_linux /api/api

# Modified files for titan
RUN chmod 555 /api/api

# Expose ports
EXPOSE $PORT

# Run Titan
CMD api -port $PORT
