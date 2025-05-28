from faster_whisper import WhisperModel
import warnings
warnings.filterwarnings("ignore", category=UserWarning, module="ctranslate2")


model = WhisperModel("ivrit-ai/faster-whisper-v2-d4",
                     device="cpu", compute_type="int8")

segments, _ = model.transcribe("audio_file.mp3", language="he")

text = " ".join([s.text for s in segments])
print(text)
