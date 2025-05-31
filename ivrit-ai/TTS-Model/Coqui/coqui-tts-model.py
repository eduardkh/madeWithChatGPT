from torch.serialization import add_safe_globals
from TTS.api import TTS

from TTS.tts.configs.xtts_config import XttsConfig
from TTS.tts.models.xtts import XttsAudioConfig, XttsArgs
from TTS.config.shared_configs import BaseDatasetConfig

add_safe_globals([XttsConfig, XttsAudioConfig, BaseDatasetConfig, XttsArgs])

# Load XTTS v2
print("[+] Loading model...")
tts = TTS(model_name="tts_models/multilingual/multi-dataset/xtts_v2")

# Input text and speaker reference
text = "Hi, my name is Alon, and I'm a good boy who likes to do what good boys do, and it is to help my parents protect my sister and do my homework on time, and from time to time, I want to go to the gym and lift some weights"
speaker_wav = "alon.wav"  # required for voice cloning
language = "en"

print(f"[+] Synthesizing voice clone from '{speaker_wav}'...")
tts.tts_to_file(
    text=text,
    speaker_wav=speaker_wav,
    language=language,
    file_path="output.wav"
)

print("[✓] Done! Audio saved to output.wav")
