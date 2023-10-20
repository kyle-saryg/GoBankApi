# Parent image => Go runtime env.
FROM golang:latest

# Set environment variables
ENV POSTGRESDB_NAME=postgres
ENV POSTGRESDB_USER=postgres
ENV POSTGRESDB_PASSWORD=password
ENV JWT_SECRET=hellothere

WORKDIR /gobank

COPY . .

RUN make

EXPOSE 3000
EXPOSE 5432

CMD ["make", "run"]