FROM python:3.10-slim

LABEL maintainer="votre_email@example.com"

RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /zoraxy

RUN git clone https://github.com/tobychui/zoraxy.git .

CMD ["python3", "zoraxy/main.py"]
