from confluent_kafka import Producer

conf = {'bootstrap.servers': 'localhost:9092'}
producer = Producer(conf)


def delivery_report(err, msg):
    if err:
        print(f'❌ Delivery failed: {err}')
    else:
        print(f'✅ Message delivered to {msg.topic()} [{msg.partition()}]')


for i in range(10):
    data = f"Message #{i} suka"
    producer.produce('demo-topic', key=str(i),
                     value=data, callback=delivery_report)

producer.flush()
