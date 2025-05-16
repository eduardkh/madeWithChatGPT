import adafruit_dht
import board

sensor = adafruit_dht.DHT22(board.D4, use_pulseio=False)

temperature = sensor.temperature
humidity = sensor.humidity

print(f"Temp: {temperature:.1f}°C | Humidity: {humidity:.1f}%")
