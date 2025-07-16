FROM python:3.12-slim
ARG APPDIR
WORKDIR /app
COPY ${APPDIR}/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY ${APPDIR}/app.py .
EXPOSE 8080
CMD ["python", "app.py"]

# docker build -t eduardkh/healthy-app:latest --build-arg APPDIR=healthy-app .
# docker build -t eduardkh/unhealthy-app:latest --build-arg APPDIR=unhealthy-app .