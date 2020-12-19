# ラズパイと学習リモコンを使ったIoTリモコン

## はじめに

このプログラムを動かすにはMQTTブローカーと、コントローラとなるアプリが必要です。
また、本プログラムはGo言語で動かしますので、Go言語がインストールされていることを確認して下さい。

- [MQTTブローカー MosquittoをUbuntu 18.04にインストール](https://qiita.com/Yosh0124/items/b94f114afe3d39fb285b)
- [リモコンアプリ(Android版)](https://github.com/Yosh0124/home-remocon-android)

## ハードウェア

- Raspberry Pi
- 学習リモコンモジュール : [ADRSIR](https://bit-trade-one.co.jp/product/module/adrsir/)

※学習リモコンの対応Raspberry Piに`Raspberry Pi Model B+／Raspberry Pi 2 Model B／Raspberry Pi 3 Model B`とありますが、Raspberry Pi 4でも可能です。

## インストール方法

### ダウンロード

まずはプログラムをダウンロード。

```bash
$ go get github.com/Yosh0124/home-remocon-raspberrypi
```

### 設定

ますは`sample.env`のファイル名を`.env`に変更し、ファイルの中身をご自身で立てたMQTT Brokerの設定に合わせて修正して下さい。
例えば、

```env
MQTT_USER=username
MQTT_PASSWD=password
MQTT_HOST="wss://mqtt.example.com:8083"
```

のような具合です。

次に、ビルドします。

### ビルド

```bash
$ cd $GOPATH/src/github.com/Yosh0124/home-remocon-raspberrypi
$ go build
```

```bash
$ go run
```

できちんと動作すれば成功です。

### 自動起動

以下のようにSystemdで自動起動するようにします。

```bash
$ sudo nano /etc/systemd/system/home-remocon.service
```

そして、ファイルを以下のように設定します。

```service
[Unit]
Description= launch home remote controller
After=syslog.target network.target

[Service]
WorkingDirectory=/home/pi/dev/src/github.com/Yosh0124/home-remocon-raspberrypi
ExecStart=/home/pi/dev/src/github.com/Yosh0124/home-remocon-raspberrypi/home-remocon-raspberrypi
Restart = always
Type = simple
User=pi
[Install]
WantedBy = multi-user.target
```

そして、サービスを起動します。

```bash
$ sudo systemctl start home-remocon.service
```

以下のコマンドで正常に起動していることを確認します。


```bash
$ sudo systemctl status home-remocon.service
```

問題なく動いてたら自動起動するように有効化します。

```bash
$ sudo systemctl enable home-remocon.service
```

以上です。
