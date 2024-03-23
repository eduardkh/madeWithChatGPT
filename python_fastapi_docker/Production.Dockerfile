# Use a slim Python image
FROM python:3.8-slim

# Set the working directory
WORKDIR /code

# Install dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the application source code
COPY app.py .

# Non-root user
RUN adduser --disabled-password --gecos '' webrunner
USER webrunner

# Start the application
CMD ["gunicorn", "-k", "uvicorn.workers.UvicornWorker", "app:app", "--bind", "0.0.0.0:8000"]
