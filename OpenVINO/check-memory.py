
"""Inspect system and OpenVINO inference memory capacity."""

import platform

import openvino as ov
import psutil


GIB = 1024 ** 3


def format_memory(value):
    """Convert bytes to GiB."""
    return f"{value / GIB:.2f} GiB"


def get_property(core, device, name):
    """Read an OpenVINO device property if available."""
    try:
        return core.get_property(device, name)
    except (RuntimeError, ValueError, TypeError):
        return None


def print_system_memory():
    """Display system RAM and CPU information."""
    memory = psutil.virtual_memory()
    swap = psutil.swap_memory()

    print("\n=== SYSTEM / CPU ===")
    print(f"OS: {platform.platform()}")
    print(f"CPU: {platform.processor()}")
    print(f"Physical cores: {psutil.cpu_count(logical=False)}")
    print(f"Logical CPUs: {psutil.cpu_count(logical=True)}")
    print(f"Total RAM: {format_memory(memory.total)}")
    print(f"Available RAM: {format_memory(memory.available)}")
    print(f"Used RAM: {format_memory(memory.used)}")
    print(f"RAM utilization: {memory.percent}%")
    print(f"Page file / swap total: {format_memory(swap.total)}")

    print("\nCPU inference capacity:")
    print("Uses system RAM; no fixed OpenVINO model-size limit.")
    print("Practical capacity depends on available memory,")
    print("process limits, model weights, and runtime overhead.")


def print_gpu_memory(core, device):
    """Display GPU memory properties reported by OpenVINO."""
    name = get_property(core, device, "FULL_DEVICE_NAME")
    device_type = get_property(core, device, "DEVICE_TYPE")

    print(f"\n=== {device} ===")
    print(f"Name: {name}")
    print(f"Type: {device_type}")

    properties = {
        "GPU_DEVICE_TOTAL_MEM_SIZE": "Reported GPU memory",
        "GPU_DEVICE_MAX_ALLOC_MEM_SIZE": "Maximum single allocation",
    }

    for key, label in properties.items():
        value = get_property(core, device, key)

        if isinstance(value, (int, float)):
            print(f"{label}: {format_memory(value)}")
        else:
            print(f"{label}: Not reported")

    print("Note: Integrated GPUs share system RAM.")
    print("Reported capacity is not reserved or guaranteed.")
    print("Single-allocation limits are not model-size limits.")


def main():
    """Inspect system memory and OpenVINO devices."""
    print(f"OpenVINO version: {ov.__version__}")

    print_system_memory()

    core = ov.Core()

    for device in core.available_devices:
        if device.startswith("GPU"):
            print_gpu_memory(core, device)

    print("\n=== SUMMARY ===")
    print("CPU: Limited primarily by available system memory.")
    print("iGPU: Shared RAM plus GPU allocation constraints.")
    print("dGPU: Device-specific memory and driver constraints.")
    print("These are reported limits, not allocation tests.")


if __name__ == "__main__":
    main()
