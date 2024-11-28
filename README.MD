# Realtime Chat Application with Go and ELK Stack

## Overview
This project is a Realtime Chat Application built with Golang as the backend, integrated with the ELK Stack (Elasticsearch, Logstash, Kibana) for logging, monitoring, and visualization. The chat application supports real-time messaging using WebSockets and stores logs for analytics and debugging.

## Tech Stack
- Golang
- WebSocket
- Elastic Search
- Logstash
- Kibana

## Setup
1. Clone the project

```bash
git clone https://github.com/ivanrafli14/message-app.git
```

2. Copy the example env file over:
```bash
cp .env.example .env
```
3. Modify the env file you just copied .env with the correct credentials. 
4. Run the program
```bash
cd elk_stack
docker-compose up -d
```
