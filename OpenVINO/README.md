# OpenVINO Python API

```powershell
# SpinUp
cd .\OpenVINO\

py -3.13 -m venv .venv
.\.venv\Scripts\Activate.ps1

python -m pip install --upgrade pip
pip install openvino

# Check hardware
python -c "import openvino as ov; print('Version:', ov.__version__); print('Devices:', ov.Core().available_devices)"

py .\list-hardware.py

# Test inference
pip install numpy
py .\test-inference.py

# Teardown
deactivate
Remove-Item -Recurse -Force .venv
```

## run TinyLlama 1.1B Chat

```powershell
# 1. Install the dependencies
pip install openvino-genai "optimum-intel[openvino]"

# 2. Download and convert the model
optimum-cli export openvino `
  --model TinyLlama/TinyLlama-1.1B-Chat-v1.0 `
  --weight-format int4 `
  ./models/tinyllama

# 3. Run inference
python chat.py CPU # Intel i7-12700K
python chat.py GPU.0 # Intel UHD 770
python chat.py GPU.1 # NVIDIA RTX 3070, does not work

# check iGPU memory limit
python check-memory.py
```

```powershell
# serve with OVMS
curl.exe -fL `
  "https://github.com/openvinotoolkit/model_server/releases/download/v2026.4.1/ovms_windows_2026.4.1_python_on.zip" `
  -o ovms.zip

tar -xf ovms.zip

.\ovms\setupvars.ps1
.\ovms\ovms.exe --version

.\ovms\ovms.exe `
  --model_path "$((Resolve-Path .\models\tinyllama).Path)" `
  --model_name tinyllama `
  --task text_generation `
  --target_device CPU `
  --rest_port 8000

# Verify the API
curl.exe http://localhost:8000/v1/models
```
