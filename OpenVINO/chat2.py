
import sys
import time
import openvino_genai as ov_genai

device = sys.argv[1] if len(sys.argv) > 1 else "CPU"

print(f"Device: {device}")

start = time.perf_counter()
pipe = ov_genai.LLMPipeline("./models/tinyllama", device)
load_time = time.perf_counter() - start

prompt = "Explain Kubernetes in three sentences."

# First inference
start = time.perf_counter()
result = pipe.generate(
    prompt,
    max_new_tokens=128,
    do_sample=False
)
inference_time = time.perf_counter() - start

print(f"\nResponse:\n{result}")
print(f"\nPipeline initialization: {load_time:.2f}s")
print(f"First generation: {inference_time:.2f}s")

# Repeat with an initialized pipeline
start = time.perf_counter()
result = pipe.generate(
    prompt,
    max_new_tokens=128,
    do_sample=False
)
warm_time = time.perf_counter() - start

print(f"Second generation: {warm_time:.2f}s")
