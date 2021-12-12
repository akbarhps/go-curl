# Go-Curl

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
go-curl get "https://jsonplaceholder.typicode.com/posts/1" -b="title=foo;body=bar;userId=1" -h="content-type=application/json;accept=application/json"
```

Output:

```bash
Age: [9218]
Alt-Svc: [h3=":443"; ma=86400, h3-29=":443"; ma=86400, h3-28=":443"; ma=86400, h3-27=":443"; ma=86400]
X-Ratelimit-Limit: [1000]
Access-Control-Allow-Credentials: [true]
Cache-Control: [max-age=43200]
Expires: [-1]
Expect-Ct: [max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"]
Report-To: [{"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v3?s=3QksgiuiD68mn3gsAH7sw3gkV%2BqES3j3tMdqfH7jAHGK23D6E4Wh49cD6EPd9M5le%2FGHnpc4wz1tTetmkHZSQIGZaPNsWqTwwMoeJecbkq24j9EjUPJ9Vh%2BLiUzn76u%2Fbfe2KhQSmTvOOO0h080L"}],"group":"cf-nel","max_age":604800}]
Date: [Sun, 12 Dec 2021 15:29:06 GMT]
Content-Type: [application/json; charset=utf-8]
Nel: [{"success_fraction":0,"report_to":"cf-nel","max_age":604800}]
Vary: [Origin, Accept-Encoding]
X-Content-Type-Options: [nosniff]
Via: [1.1 vegur]
Server: [cloudflare]
Cf-Ray: [6bc7fc826b207fca-IAD]
X-Powered-By: [Express]
X-Ratelimit-Reset: [1639313769]
Etag: [W/"124-yiKdLzqO5gfBrJFrcdJ8Yq0LGnU"]
Cf-Cache-Status: [HIT]
X-Ratelimit-Remaining: [999]
Pragma: [no-cache]


{
  "userId": 1,
  "id": 1,
  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
}
```