# Temperature and Humidity sensor (AM2302)

[pinout](https://pinout.xyz/pinout/ground)

```bash
# install all the necessary packages in a virtual environment
python3 -m venv .venv
source .venv/bin/activate

pip install --upgrade pip setuptools wheel
pip install adafruit-circuitpython-dht
pip install gpiozero

# visualize the pins
pinout
```

| Sensor | GPIO pins |
| ------ | --------- |
| 1      | 1         |
| 2      | 7         |
| 4      | 9         |
