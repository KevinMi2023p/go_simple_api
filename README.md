# go_simple_api
Based off the take home assignment: https://github.com/fetch-rewards/receipt-processor-challenge?tab=readme-ov-file
Everything is in Golang so no need for other instalation.

I made it so that to run the program you can just run it from the root of the repository.
```sh
go run main.go
```

It's using the default settings of the Gin framework.

Visit [`0.0.0.0:8080/`](http://0.0.0.0:8080/) in either postman or your browser with the API paths to see a responce.

The example below

```sh
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```
Should return an ID of
```sh
bc96de1690f1481e4394552a216fe91ba26713598e14e7f3b7752378badc92c9
```

Calling this endpoint
```sh
localhost:8080/receipts/bc96de1690f1481e4394552a216fe91ba26713598e14e7f3b7752378badc92c9/points
```
Should return 
```sh
{
    "points": 109
}
```
