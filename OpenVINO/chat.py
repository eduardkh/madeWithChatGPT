
import sys
import openvino_genai as ov_genai

device = sys.argv[1] if len(sys.argv) > 1 else "CPU"

print(f"Loading model on {device}...")

pipe = ov_genai.LLMPipeline(
    "./models/tinyllama",
    device
)

prompt = "Explain Kubernetes in three sentences."

print("\nGenerating response:\n")
print(pipe.generate(prompt, max_new_tokens=128))
