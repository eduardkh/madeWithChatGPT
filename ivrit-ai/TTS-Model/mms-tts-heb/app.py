from transformers import VitsModel, AutoTokenizer
import torch
import scipy.io.wavfile
import numpy as np

# Load model and tokenizer
print("[+] Loading model and tokenizer...")
model = VitsModel.from_pretrained("facebook/mms-tts-heb")
tokenizer = AutoTokenizer.from_pretrained("facebook/mms-tts-heb")

# Hebrew input text
text = " שלום, קוראים לי אלון. אני מדבר עברית."

# Tokenize input
print("[+] Tokenizing input text...")
inputs = tokenizer(text, return_tensors="pt")

# Generate waveform
print("[+] Generating waveform...")
with torch.no_grad():
    output = model(**inputs).waveform

# Convert to 16-bit PCM
print("[+] Converting to int16 and saving to output.wav...")
waveform = output.squeeze().numpy()
waveform_int16 = np.int16(waveform / np.max(np.abs(waveform)) * 32767)

# Save as WAV file
scipy.io.wavfile.write("output.wav", rate=int(model.config.sampling_rate), data=waveform_int16)

print("[✓] Done! Output saved to output.wav")
