import openvino as ov

core = ov.Core()

for device in core.available_devices:
    print(f"\nDevice: {device}")
    for property_name in [
        "FULL_DEVICE_NAME",
        "DEVICE_TYPE",
        "OPTIMIZATION_CAPABILITIES",
    ]:
        try:
            value = core.get_property(device, property_name)
            print(f"  {property_name}: {value}")
        except Exception as e:
            print(f"  {property_name}: unavailable ({e})")