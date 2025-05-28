from fastapi import FastAPI, File, UploadFile, Request
from fastapi.middleware.cors import CORSMiddleware
import logging
import shutil
import tempfile
import os
from faster_whisper import WhisperModel

# --- Logging Setup ---
logging.basicConfig(level=logging.INFO,
                    format="%(asctime)s - %(levelname)s - %(message)s")
logger = logging.getLogger(__name__)

# --- FastAPI App ---
app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# --- Whisper Model Setup ---
model = WhisperModel("ivrit-ai/faster-whisper-v2-d4",
                     device="cpu", compute_type="int8")


@app.middleware("http")
async def log_requests(request: Request, call_next):
    logger.info(f"Incoming request: {request.method} {request.url}")
    response = await call_next(request)
    logger.info(f"Completed response: {response.status_code}")
    return response


@app.post("/transcribe")
async def transcribe_audio(file: UploadFile = File(...)):
    logger.info(f"Received file: {file.filename}")

    with tempfile.NamedTemporaryFile(delete=False, suffix=".wav") as tmp:
        shutil.copyfileobj(file.file, tmp)
        tmp_path = tmp.name

    try:
        segments, _ = model.transcribe(tmp_path, language="he")
        text = " ".join([s.text for s in segments])
        logger.info(f"Transcription completed: {text[:100]}...")
        return {"transcription": text}
    finally:
        os.unlink(tmp_path)
        logger.info(f"Temporary file deleted: {tmp_path}")
