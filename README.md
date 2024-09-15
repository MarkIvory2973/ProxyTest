# ProxyTest

A proxy testing program for MiHoMo.

## Installation

```bash
git clone https://github.com/MarkIvory2973/ProxyTest.git
```

## Usage

⚠ ***Test the delay in MiHoMo before running the command (>1 times)***

```bash
cd ProxyTest/src
go run ./main.go --help
go run ./main.go --group SELECT --weight 0.3
```

## Parameters

|Parameter|Required|Default|Description|
|:-|:-:|:-|:-|
|--host|-|127.0.0.1|MiHoMo API host|
|--port|-|9090|MiHoMo API port|
|--tls|-|-|Use TLS|
|--group|✓|-|Group name|
|--weight|-|0.5|Weight (0~1)|

## The range of *k*

![The range of k](https://raw.githubusercontent.com/MarkIvory2973/ProxyTest/main/imgs/k.png)
