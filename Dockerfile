FROM golang:1.18-alpine

WORKDIR /app

# expose available environment variables
ENV DISCORD_TOKEN="OTk5NTU1MzcwMTM5Nzg3MzI1.G4m23r.N4Bv50-sdUXTCN4j1I0YrBoervQefJVL4s7OiQ"
ENV ATERNOS_SESSION="uLrJoiSOVLhqvF24v6KyMFONe8UuHg2qUWAgmiQi9SsAdSpPjUuYezWqllLTszYo2CvxMdJWL0I3RdqEllELHC0t7FKzwoGYLzsH"
ENV ATERNOS_SERVER="SuNnbrhLfzGMsqkz"
ENV MONGO_DB_URI=""
ENV PROXY=""

# install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy files
COPY . ./

# build binary
RUN go build -o ./bin/aternos-discord-bot ./cmd/main.go

CMD [ "./bin/aternos-discord-bot" ]
