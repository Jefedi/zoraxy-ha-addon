FROM python:3.10.12

WORKDIR /usr/src/app

# Copier les fichiers du dépôt Zoraxy
COPY data/ ./

# Installer les dépendances nécessaires pour Zoraxy
RUN apt-get update && \
    apt-get install -y git golang build-essential libssl-dev && \
    apt-get install -y libasound-dev libportaudio2 libportaudiocpp0 portaudio19-dev python3-opencv ffmpeg

# Installer les dépendances Python
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt

# Rendre le script exécutable
COPY --chmod=700 data/run.sh /run.sh
RUN chmod +x /run.sh

# Exposer les ports nécessaires
EXPOSE 6980 16045 8000

CMD ["/run.sh"]
