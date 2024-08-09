# ProxyTest

A proxy testing program for MiHoMo.

## Installation

```bash
pip install numpy click rich requests
git clone https://github.com/MarkIvory2973/ProxyTest.git
```

## Usage

```bash
cd ProxyTest
python ./main.py --help
python ./main.py --host 127.0.0.1 --port 9090 --https --excludes 剩余流量,官址 --group SELECT --k 0.3
```

## Parameters

|Parameter|Required|Default|Description|
|:-:|:-:|:-:|:-:|
|--host|-|127.0.0.1|MiHoMo API host|
|--port|-|9090|MiHoMo API port|
|--https|-|-|Use HTTPS|
|--excludes|-|-|Remove exclusions|
|--group|✓|-|Group name|
|--k|-|0.5|Weight (0~1)|

## The range of *k*

![The range of k](https://raw.githubusercontent.com/MarkIvory2973/ProxyTest/main/imgs/k.png)