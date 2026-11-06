
import numpy as np
import openvino as ov

core = ov.Core()

# Small matrix multiplication model
a = ov.opset13.parameter([128, 128], np.float32)
b = ov.opset13.parameter([128, 128], np.float32)
result = ov.opset13.matmul(a, b, False, False)
model = ov.Model([result], [a, b])

inputs = {
    0: np.ones((128, 128), dtype=np.float32),
    1: np.ones((128, 128), dtype=np.float32),
}

for device in core.available_devices:
    print(f"\nTesting {device}")
    try:
        compiled = core.compile_model(model, device)
        output = compiled(inputs)
        print(f"SUCCESS: {device}")
        print(f"Output shape: {next(iter(output.values())).shape}")
    except Exception as e:
        print(f"FAILED: {type(e).__name__}: {e}")
