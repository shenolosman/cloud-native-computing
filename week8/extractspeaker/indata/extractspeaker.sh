#!/bin/sh -xe
echo 'IN' . $1
echo 'OUT' . $2
rm -rf speaker.wav || true
ffmpeg -i $1 -ss 0 -to 60 -ab 160k -ac 2 -ar 44100 -vn $2
# ffmpeg -i $2 -af "afftdn=nf=-25" /indata/out1.wav
# ffmpeg -i /indata/out1.wav -af "afftdn=nf=-25" /indata/out2.wav
# ffmpeg -i /indata/out2.wav -af "highpass=f=200, lowpass=f=3000" /indata/out3.wav
# ffmpeg -i /indata/out3.wav -af "volume=4" /indata/out4.wav