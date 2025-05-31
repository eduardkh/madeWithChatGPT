from transformers import VitsModel, AutoTokenizer
import torch
import scipy.io.wavfile

# Load the model and tokenizer
model = VitsModel.from_pretrained("facebook/mms-tts-heb")
tokenizer = AutoTokenizer.from_pretrained("facebook/mms-tts-heb")

# Input Hebrew text
text = "שלום, קוראים לי פליקס"
inputs = tokenizer(text, return_tensors="pt")

# Generate speech waveform
with torch.no_grad():
    output = model(**inputs).waveform

# Save the output to a WAV file
scipy.io.wavfile.write(
    "output.wav", rate=model.config.sampling_rate, data=output.numpy())
