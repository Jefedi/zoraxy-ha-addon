FROM python:3.10.12

WORKDIR /usr/src/app

COPY data/ ./
RUN apt-get update && \
    apt-get install -y libasound-dev libportaudio2 libportaudiocpp0 portaudio19-dev python3-opencv ffmpeg
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt

COPY --chmod=700 data/run.sh /run.sh
RUN chmod +x /run.sh

EXPOSE 6980 16045 8000

CMD ["/run.sh"]