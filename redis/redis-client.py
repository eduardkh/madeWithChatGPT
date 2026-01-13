import redis
r = redis.Redis(host='127.0.0.1', port=6379, db=0)
r.set('x', '42')
print(r.get('x'))  # b'42'
