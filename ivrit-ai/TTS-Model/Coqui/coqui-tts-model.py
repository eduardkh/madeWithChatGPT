from TTS.api import TTS

tts = TTS(model_name="tts_models/multilingual/multi-dataset/xtts_v2")
print(tts.speakers)

tts.tts_to_file(
    text="Shalom, kore'im li Felix",  # transliterated
    speaker=tts.speakers[-1],
    language=tts.languages[0],
    file_path="output.wav"
)
