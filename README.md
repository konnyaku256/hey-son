# hey-son
Son calling system for mom

## About
### What is "Hey Son"?
Hey Sonはおかあさんのための息子を呼び出す魔法です。

あなたはいつも息子を呼ぶときどうしていますか？ 階段を登って、廊下を歩いて、息子の部屋まで呼びに行っていませんか？

それって、ちょっと面倒ですよね。 この魔法を使えばそんな悩みをかんたんに解決できます。

### How to use
このWebアプリを開いてボタンを押すだけです。 たったこれだけで、あなたの息子を呼び出せます。

### How magic works
このアプリは息子の部屋のRaspberry Piとつながっています。 ボタンが押されたら、Raspberry Piに接続されたスピーカーから呼び出しが掛かるようになっています。

## Getting started
```
git clone https://github.com/konnyaku256/hey-son.git
cd hey-son
docker-compose build
docker-compose up -d
```

## Features
### Frontend
Hey Son Web Client
- Vue.js
- axios

### Backend
Audio playback API server
- Golang
- [Command] aplay, alsa-utils

### Hardware
- Raspberry Pi 4 model B / 4GB(RAM)
- Logitech Z120 Mini Stereo Speakers