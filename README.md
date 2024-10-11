# taco
- 가볍게 사용할 수 있는 HTTP/HTTPS 프록시입니다.
- Linux 외의 환경은 고려하지 않았습니다.

## 빌드

```bash
sh build.sh
```

## 배포 

```bash
sudo sh -x deploy.sh
```

## 기본 사용법
- 서버가 열리는 기본 포트는 8080 포트입니다.
- Proxy-Host 헤더는 `https://naver.com` 같은 scheme:host:port 값을 받는 특수한 전용 헤더입니다.
- Proxy-Host 헤더에 최종 목적지 호스트를 전달하면, 호스트를 Proxy-Host의 값으로 치환해서 요청을 대신 쏴줍니다. 
- query-parameter, header, body 등은 그대로 전달합니다. Proxy-Host 헤더만 빼고요.

curl exaple
```bash
curl -X POST http://IP:8080/api/v1/health \
     -H "Content-Type: application/json" \
     -H "Proxy-Host: https://...lambda-url.ap-northeast-2.on.aws" \
     -d '{"status": "UP"}'
```

wget example
```bash
wget --method=POST \
     --header="Content-Type: application/json" \
     --header="Proxy-Host: https://...lambda-url.ap-northeast-2.on.aws" \
     --body-data='{"status": "UP"}' \
     http://IP:8080/api/v1/health
```
