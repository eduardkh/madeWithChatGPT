from confluent_kafka import Consumer

conf = {
    'bootstrap.servers': 'localhost:9092',
    'group.id': 'demo-group',
    'auto.offset.reset': 'earliest'
}

consumer = Consumer(conf)
consumer.subscribe(['demo-topic'])

print("⏳ Waiting for messages...")
try:
    while True:
        msg = consumer.poll(1.0)
        if msg is None:
            continue
        if msg.error():
            print(f'⚠️ Error: {msg.error()}')
        else:
            print(
                f'📨 Received: {msg.key().decode()} -> {msg.value().decode()}')
except KeyboardInterrupt:
    pass
finally:
    consumer.close()
