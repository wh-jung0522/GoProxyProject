# Go Proxy Project
- Public Korbit API 와 동일합니다.
- 데이터를 로컬서버에 저장할 수 있도록 작성하였습니다.
- 추가 : CandleStick의 경우 Upbit API에서 불러온 데이터입니다.

## Proxy 서버
    go run proxy.go
    
protocol.go:14의 경로를 변경하여 실행하여야 합니다.

    var dump_dir = "C:\\Users\\kwdg1\\Desktop\\GoProxyProject\\jsonDB"

## Client
    http://localhost:8887/ticker?currency_pair=eth_krw

    http://localhost:8887/ticker/detailed?currency_pair=eth_krw
    
    http://localhost:8887/ticker/detailed/all
    
    http://localhost:8887/transactions?currency_pair=eth_krw
    
    http://localhost:8887/orderbook?currency_pair=eth_krw
    
    http://localhost:8887/candles/minutes/1?currency_pair=eth_krw&to=2021-11-01T00%3A00%3A00Z&count=200

- 참고 : Public API 만 지원하도록 하였습니다. 
  - 따라서, GET Method만 지원합니다.

---
@whjung-0522