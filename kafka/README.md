# Kafka

```bash
# spin-up the kafka node
docker-compose up -d

# check for errors
docker logs kafka-kafka-1

# send messages
python .\producer.py

# pull messages
python .\consumer.py

# turn off the kafka node
docker-compose down -v
```
