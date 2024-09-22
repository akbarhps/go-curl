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

```json
[
  {
    "_____url": "https://64e9579299cf45b15fe09811.mockapi.io/test/api/users",
    "____status": "200 OK",
    "___duration": "1.6396044s",
    "__headers": {
      "Access-Control-Allow-Headers": "X-Requested-With,Content-Type,Cache-Control,access_token",
      "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,OPTIONS",
      "Access-Control-Allow-Origin": "*",
      "Connection": "keep-alive",
      "Content-Length": "18087",
      "Content-Type": "application/json",
      "Date": "Sun, 22 Sep 2024 03:18:45 GMT",
      "Etag": "\"-749874500\"",
      "Nel": {
        "failure_fraction": 0.05,
        "max_age": 3600,
        "report_to": "heroku-nel",
        "response_headers": [
          "Via"
        ],
        "success_fraction": 0.005
      },
      "Report-To": {
        "endpoints": [
          {
            "url": "https://nel.heroku.com/reports?ts=1726975125\u0026sid=1b10b0ff-8a76-4548-befa-353fc6c6c045\u0026s=9dPzhqKVYp9FYudjWXaLgTJg32IQ5WuW57i%2FKNDdSc8%3D"
          }
        ],
        "group": "heroku-nel",
        "max_age": 3600
      },
      "Reporting-Endpoints": "heroku-nel=https://nel.heroku.com/reports?ts=1726975125\u0026sid=1b10b0ff-8a76-4548-befa-353fc6c6c045\u0026s=9dPzhqKVYp9FYudjWXaLgTJg32IQ5WuW57i%2FKNDdSc8%3D",
      "Server": "Cowboy",
      "Via": "1.1 vegur",
      "X-Powered-By": "Express"
    },
    "_body": [
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/193.jpg",
        "createdAt": "2023-08-25T19:44:48.509Z",
        "id": "1",
        "name": "Blake Bechtelar I"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/470.jpg",
        "createdAt": "2023-08-25T15:56:15.129Z",
        "id": "2",
        "name": "Ms. Herbert Kerluke"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/218.jpg",
        "createdAt": "2023-08-25T04:47:17.823Z",
        "id": "3",
        "name": "Ken Cole"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/155.jpg",
        "createdAt": "2023-08-25T09:40:35.062Z",
        "id": "4",
        "name": "Minnie Bayer"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/640.jpg",
        "createdAt": "2023-08-25T08:58:48.218Z",
        "id": "5",
        "name": "Eleanor Gibson II"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/503.jpg",
        "createdAt": "2023-08-25T16:52:48.509Z",
        "id": "6",
        "name": "Amanda Blanda"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/863.jpg",
        "createdAt": "2023-08-25T21:52:16.495Z",
        "id": "7",
        "name": "Vincent Huels"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/502.jpg",
        "createdAt": "2023-08-25T23:36:27.892Z",
        "id": "8",
        "name": "Wilbert Botsford"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1145.jpg",
        "createdAt": "2023-08-25T09:42:56.103Z",
        "id": "9",
        "name": "Kristen Hoeger PhD"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/876.jpg",
        "createdAt": "2023-08-25T06:11:18.077Z",
        "id": "10",
        "name": "Mr. Francis Okuneva"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1174.jpg",
        "createdAt": "2023-08-25T21:30:47.569Z",
        "id": "11",
        "name": "Mrs. Blake Thompson"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/103.jpg",
        "createdAt": "2023-08-25T10:56:21.469Z",
        "id": "12",
        "name": "Miss Garry Kiehn IV"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/38.jpg",
        "createdAt": "2023-08-25T04:42:25.018Z",
        "id": "13",
        "name": "Shawn Muller"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/827.jpg",
        "createdAt": "2023-08-25T07:29:40.599Z",
        "id": "14",
        "name": "Kristy Haley"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/596.jpg",
        "createdAt": "2023-08-25T12:48:00.885Z",
        "id": "15",
        "name": "Miss Eric Fahey"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1005.jpg",
        "createdAt": "2023-08-25T19:25:03.678Z",
        "id": "16",
        "name": "Delores Bayer Jr."
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1172.jpg",
        "createdAt": "2023-08-25T13:58:38.171Z",
        "id": "17",
        "name": "Becky Homenick"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/931.jpg",
        "createdAt": "2023-08-25T09:21:01.569Z",
        "id": "18",
        "name": "Tommy Hermann"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/284.jpg",
        "createdAt": "2023-08-26T00:48:54.608Z",
        "id": "19",
        "name": "Alberta Ebert"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1090.jpg",
        "createdAt": "2023-08-25T09:29:51.270Z",
        "id": "20",
        "name": "Maureen Jacobi"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/841.jpg",
        "createdAt": "2023-08-25T08:40:26.985Z",
        "id": "21",
        "name": "Mack Koepp III"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/38.jpg",
        "createdAt": "2023-08-25T08:41:29.341Z",
        "id": "22",
        "name": "Cornelius Langworth"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/652.jpg",
        "createdAt": "2023-08-25T05:58:53.742Z",
        "id": "23",
        "name": "Janis Schuppe"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1092.jpg",
        "createdAt": "2023-08-25T18:20:15.903Z",
        "id": "24",
        "name": "Kelley Thompson III"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/823.jpg",
        "createdAt": "2023-08-25T03:39:43.458Z",
        "id": "25",
        "name": "Delia Jast"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/227.jpg",
        "createdAt": "2023-08-25T22:17:32.269Z",
        "id": "26",
        "name": "Elisa Wehner"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/232.jpg",
        "createdAt": "2023-08-25T18:32:00.482Z",
        "id": "27",
        "name": "Rodney Ryan MD"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/788.jpg",
        "createdAt": "2023-08-25T19:11:55.242Z",
        "id": "28",
        "name": "Sharon Harris I"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/70.jpg",
        "createdAt": "2023-08-25T01:59:47.167Z",
        "id": "29",
        "name": "Faye Muller"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/310.jpg",
        "createdAt": "2023-08-25T07:11:17.453Z",
        "id": "30",
        "name": "Cody Beatty DDS"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/560.jpg",
        "createdAt": "2023-08-25T12:02:27.087Z",
        "id": "31",
        "name": "Jean Hoeger"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1026.jpg",
        "createdAt": "2023-08-25T13:28:56.707Z",
        "id": "32",
        "name": "Tom Ullrich"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/930.jpg",
        "createdAt": "2023-08-25T03:02:23.826Z",
        "id": "33",
        "name": "Ms. Lee Hahn"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1081.jpg",
        "createdAt": "2023-08-25T14:34:35.990Z",
        "id": "34",
        "name": "Edmund Upton"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1146.jpg",
        "createdAt": "2023-08-25T13:53:46.137Z",
        "id": "35",
        "name": "Rhonda Leannon"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/856.jpg",
        "createdAt": "2023-08-25T11:33:12.354Z",
        "id": "36",
        "name": "Lee Dare"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/677.jpg",
        "createdAt": "2023-08-25T18:29:43.282Z",
        "id": "37",
        "name": "Dallas Kilback"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/460.jpg",
        "createdAt": "2023-08-25T18:08:57.154Z",
        "id": "38",
        "name": "Cynthia Frami"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/136.jpg",
        "createdAt": "2023-08-25T14:00:45.086Z",
        "id": "39",
        "name": "Lucas Bradtke"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/322.jpg",
        "createdAt": "2023-08-25T23:46:01.617Z",
        "id": "40",
        "name": "Jared Hansen"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1123.jpg",
        "createdAt": "2023-08-25T08:12:59.352Z",
        "id": "41",
        "name": "Mattie Jast"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/690.jpg",
        "createdAt": "2023-08-25T13:33:02.964Z",
        "id": "42",
        "name": "Kellie Dietrich Jr."
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1028.jpg",
        "createdAt": "2023-08-25T12:55:13.564Z",
        "id": "43",
        "name": "Marcos Wiza"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/639.jpg",
        "createdAt": "2023-08-25T18:32:27.044Z",
        "id": "44",
        "name": "Samantha O'Keefe"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/950.jpg",
        "createdAt": "2023-08-26T00:12:09.113Z",
        "id": "45",
        "name": "Rafael Wilderman"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/934.jpg",
        "createdAt": "2023-08-25T11:26:53.468Z",
        "id": "46",
        "name": "Doreen Borer"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/388.jpg",
        "createdAt": "2023-08-25T21:02:19.038Z",
        "id": "47",
        "name": "Rodney Hilpert"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/858.jpg",
        "createdAt": "2023-08-25T04:44:24.262Z",
        "id": "48",
        "name": "Joey Leffler"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/737.jpg",
        "createdAt": "2023-08-25T03:35:36.917Z",
        "id": "49",
        "name": "Mr. Pedro Cassin"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/548.jpg",
        "createdAt": "2023-08-25T07:19:20.300Z",
        "id": "50",
        "name": "Lora Kautzer DVM"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/548.jpg",
        "createdAt": "2023-08-25T07:19:20.300Z",
        "id": "51",
        "name": "Lora Kautzer DVM"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/805.jpg",
        "createdAt": "2023-08-25T19:02:50.107Z",
        "id": "52",
        "name": "Cameron Barrows"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/171.jpg",
        "createdAt": "2023-08-26T01:12:28.086Z",
        "id": "53",
        "name": "Adam Thiel"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/525.jpg",
        "createdAt": "2023-08-25T19:31:19.141Z",
        "id": "54",
        "name": "Monique Lebsack"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1107.jpg",
        "createdAt": "2023-08-25T11:33:02.423Z",
        "id": "55",
        "name": "Dr. Sheldon Tremblay"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/548.jpg",
        "createdAt": "2023-08-25T07:19:20.300Z",
        "id": "56",
        "name": "Awekaweok"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/548.jpg",
        "createdAt": "2023-08-25T07:19:20.300Z",
        "id": "57",
        "name": "Awekaweok"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/548.jpg",
        "createdAt": "2023-08-25T07:19:20.300Z",
        "id": "58",
        "name": "Awekaweok"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/548.jpg",
        "createdAt": "2023-08-25T07:19:20.300Z",
        "id": "59",
        "name": "Awekaweok"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1140.jpg",
        "createdAt": "2024-08-14T09:28:24.679Z",
        "id": "60",
        "name": "Linda Beahan"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/235.jpg",
        "createdAt": "2024-08-15T05:58:04.646Z",
        "id": "61",
        "name": "Al White"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/171.jpg",
        "createdAt": "2024-08-15T09:19:41.159Z",
        "id": "62",
        "name": "Lindsay Yundt"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/856.jpg",
        "createdAt": "2024-08-15T04:00:41.330Z",
        "id": "63",
        "name": "Miss Dana Herman"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/430.jpg",
        "createdAt": "2024-08-15T07:59:48.625Z",
        "id": "64",
        "name": "Ken Feest"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/408.jpg",
        "createdAt": "2024-08-15T06:52:08.007Z",
        "id": "65",
        "name": "Nelson Kub"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/462.jpg",
        "createdAt": "2024-08-15T09:08:30.681Z",
        "id": "66",
        "name": "Marty Collins"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/161.jpg",
        "createdAt": "2024-08-14T20:28:27.278Z",
        "id": "67",
        "name": "Alexis Sporer"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/629.jpg",
        "createdAt": "2024-08-14T15:39:21.094Z",
        "id": "68",
        "name": "Lois Hyatt"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/578.jpg",
        "createdAt": "2024-08-15T02:55:29.248Z",
        "id": "69",
        "name": "Clyde Jenkins"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1073.jpg",
        "createdAt": "2024-08-14T17:46:17.134Z",
        "id": "70",
        "name": "Kelly Rohan"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/304.jpg",
        "createdAt": "2024-08-15T06:25:09.200Z",
        "id": "71",
        "name": "Kay Fritsch"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/484.jpg",
        "createdAt": "2024-08-15T00:42:38.810Z",
        "id": "72",
        "name": "Gabriel Witting"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/702.jpg",
        "createdAt": "2024-08-14T19:27:16.341Z",
        "id": "73",
        "name": "Mable Dicki IV"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/720.jpg",
        "createdAt": "2024-08-14T17:50:11.937Z",
        "id": "74",
        "name": "Sonja Torp IV"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/90.jpg",
        "createdAt": "2024-08-14T16:08:38.669Z",
        "id": "75",
        "name": "Meredith Stoltenberg"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/89.jpg",
        "createdAt": "2024-08-15T06:05:48.093Z",
        "id": "76",
        "name": "Stewart Green"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/864.jpg",
        "createdAt": "2024-08-14T15:35:15.140Z",
        "id": "77",
        "name": "Sophia Emard"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/751.jpg",
        "createdAt": "2024-08-14T23:22:06.210Z",
        "id": "78",
        "name": "Wesley Jakubowski"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/533.jpg",
        "createdAt": "2024-08-14T20:15:07.115Z",
        "id": "79",
        "name": "Jim Keebler"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/576.jpg",
        "createdAt": "2024-08-14T19:29:20.556Z",
        "id": "80",
        "name": "Randy King"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/72.jpg",
        "createdAt": "2024-08-14T15:03:37.677Z",
        "id": "81",
        "name": "Dustin Herzog"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/195.jpg",
        "createdAt": "2024-08-14T16:56:24.104Z",
        "id": "82",
        "name": "Marianne Nicolas"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/466.jpg",
        "createdAt": "2024-08-14T16:19:40.658Z",
        "id": "83",
        "name": "Domingo Pfeffer"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/950.jpg",
        "createdAt": "2024-08-15T06:16:28.331Z",
        "id": "84",
        "name": "Viola Reichel"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/665.jpg",
        "createdAt": "2024-08-15T06:13:12.587Z",
        "id": "85",
        "name": "Randall Stracke"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/756.jpg",
        "createdAt": "2024-08-14T11:51:50.479Z",
        "id": "86",
        "name": "Miss Frankie Dach"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/649.jpg",
        "createdAt": "2024-08-15T02:53:45.648Z",
        "id": "87",
        "name": "Sandra Erdman"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/620.jpg",
        "createdAt": "2024-08-15T05:14:13.670Z",
        "id": "88",
        "name": "Carol Little"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1030.jpg",
        "createdAt": "2024-08-15T03:46:19.721Z",
        "id": "89",
        "name": "Martin Cremin"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/87.jpg",
        "createdAt": "2024-08-14T21:13:41.268Z",
        "id": "90",
        "name": "Audrey Ruecker"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/751.jpg",
        "createdAt": "2024-08-14T23:38:40.845Z",
        "id": "91",
        "name": "Cornelius Treutel"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/834.jpg",
        "createdAt": "2024-08-14T21:14:13.261Z",
        "id": "92",
        "name": "Terrence Kilback"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1240.jpg",
        "createdAt": "2024-08-14T09:21:55.463Z",
        "id": "93",
        "name": "Jorge Zboncak"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1161.jpg",
        "createdAt": "2024-08-14T16:31:16.870Z",
        "id": "94",
        "name": "Tonya Mosciski"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/792.jpg",
        "createdAt": "2024-08-15T00:11:20.109Z",
        "id": "95",
        "name": "Miss Ella O'Connell"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/436.jpg",
        "createdAt": "2024-08-14T17:31:51.807Z",
        "id": "96",
        "name": "Lynn Sporer"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/194.jpg",
        "createdAt": "2024-08-15T07:39:13.143Z",
        "id": "97",
        "name": "Julie Morissette Jr."
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/326.jpg",
        "createdAt": "2024-08-15T01:18:49.330Z",
        "id": "98",
        "name": "Alexandra Will"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/663.jpg",
        "createdAt": "2024-08-14T12:33:46.978Z",
        "id": "99",
        "name": "Mrs. Willie Willms"
      },
      {
        "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/969.jpg",
        "createdAt": "2024-08-15T06:29:44.730Z",
        "id": "100",
        "name": "Charlotte Rutherford"
      }
    ]
  },
  {
    "_____url": "https://64e9579299cf45b15fe09811.mockapi.io/test/api/users/1",
    "____status": "200 OK",
    "___duration": "380.9211ms",
    "__headers": {
      "Access-Control-Allow-Headers": "X-Requested-With,Content-Type,Cache-Control,access_token",
      "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,OPTIONS",
      "Access-Control-Allow-Origin": "*",
      "Connection": "keep-alive",
      "Content-Length": "182",
      "Content-Type": "application/json",
      "Date": "Sun, 22 Sep 2024 03:18:45 GMT",
      "Etag": "\"-565830836\"",
      "Nel": {
        "failure_fraction": 0.05,
        "max_age": 3600,
        "report_to": "heroku-nel",
        "response_headers": [
          "Via"
        ],
        "success_fraction": 0.005
      },
      "Report-To": {
        "endpoints": [
          {
            "url": "https://nel.heroku.com/reports?ts=1726975125\u0026sid=1b10b0ff-8a76-4548-befa-353fc6c6c045\u0026s=9dPzhqKVYp9FYudjWXaLgTJg32IQ5WuW57i%2FKNDdSc8%3D"
          }
        ],
        "group": "heroku-nel",
        "max_age": 3600
      },
      "Reporting-Endpoints": "heroku-nel=https://nel.heroku.com/reports?ts=1726975125\u0026sid=1b10b0ff-8a76-4548-befa-353fc6c6c045\u0026s=9dPzhqKVYp9FYudjWXaLgTJg32IQ5WuW57i%2FKNDdSc8%3D",
      "Server": "Cowboy",
      "Via": "1.1 vegur",
      "X-Powered-By": "Express"
    },
    "_body": {
      "avatar": "https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/193.jpg",
      "createdAt": "2023-08-25T19:44:48.509Z",
      "id": "1",
      "name": "Blake Bechtelar I"
    }
  }
]
```