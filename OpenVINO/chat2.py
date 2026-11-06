
import sys
import time

import openvino_genai as ov_genai


MODEL = "./models/tinyllama"
DEVICE = sys.argv[1] if len(sys.argv) > 1 else "CPU"
PROMPT = "Explain Kubernetes in three sentences."


def benchmark(pipe, config, label):
    """Run inference and report performance metrics."""
    start = time.perf_counter()

    result = pipe.generate([PROMPT], config)

    elapsed = time.perf_counter() - start
    metrics = result.perf_metrics

    print(f"\n--- {label} ---")
    print(f"Response: {result.texts[0]}")
    print(f"Wall time: {elapsed:.3f} s")

    print(f"Input tokens: {metrics.get_num_input_tokens()}")
    print(f"Generated tokens: {metrics.get_num_generated_tokens()}")
    print(f"TTFT: {metrics.get_ttft().mean:.2f} ms")
    print(f"TPOT: {metrics.get_tpot().mean:.2f} ms")
    print(
        f"Throughput: "
        f"{metrics.get_throughput().mean:.2f} tokens/s"
    )


def main():
    """Initialize the model and benchmark inference."""
    print(f"\nDevice: {DEVICE}")

    start = time.perf_counter()
    pipe = ov_genai.LLMPipeline(MODEL, DEVICE)
    load_time = time.perf_counter() - start

    print(f"Pipeline initialization: {load_time:.3f} s")

    config = ov_genai.GenerationConfig()
    config.max_new_tokens = 128
    config.do_sample = False

    benchmark(pipe, config, "First generation")
    benchmark(pipe, config, "Warm generation")


if __name__ == "__main__":
    main()
