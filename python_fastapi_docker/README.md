# python_fastapi_docker

```bash
# Good enough Dockerfile
docker build . -t python_fastapi_docker
# Production ready Dockerfile
docker build -f Production.Dockerfile -t python_fastapi_docker .

docker run -p 8000:8000 python_fastapi_docker:latest
```