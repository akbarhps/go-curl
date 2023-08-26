# http

Basic Feature of Curl made on go

## Installation

```bash
go mod tidy
```

then

```bash
go build .
```

## Usage

```bash
http ./request.json
```

Output:

```
===== REQUEST =====
{
        "url": "https://64e9579299cf45b15fe09811.mockapi.io/test/api/users",
        "method": "POST",
        "headers": {
                "Accept": "application/json",
                "Content-Type": "application/json"
        },
        "query": {},
        "form": {},
        "body": {
                "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/548.jpg",
                "createdAt": "2023-08-25T07:19:20.300Z",
                "id": "53",
                "name": "Awekaweok"
        }
}


===== RESPONSE HEADER =====
Connection: [keep-alive]
X-Powered-By: [Express]
Access-Control-Allow-Origin: [*]
Access-Control-Allow-Headers: [X-Requested-With,Content-Type,Cache-Control,access_token]
Content-Length: [175]
Vary: [Accept-Encoding]
Via: [1.1 vegur]
Server: [Cowboy]
Access-Control-Allow-Methods: [GET,PUT,POST,DELETE,OPTIONS]
Content-Type: [application/json]
Date: [Sat, 26 Aug 2023 02:41:24 GMT]


===== RESPONSE BODY =====
{
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/548.jpg",
        "createdAt": "2023-08-25T07:19:20.300Z",
        "id": "61",
        "name": "Awekaweok"
}
```