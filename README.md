# hdfc-backend

Steps to run
1. Add .test-env in parent folder
2. Run "TESTING=true  go test -v ./..." to test all test cases
3. Run "go run main.go" to start server

Requests
1. Get Categories
```
curl --location 'localhost:6000/user/categories'
```

2. Get products by category
```
curl --location 'localhost:6000/user/products?category_id=6426abd3f74d410470a56d9c&page=2'
```

3. Get product by id
```
curl --location 'localhost:6000/user/product?id=6426b1c554ff40173fab6a24'
```

4. Create order
```
curl --location 'localhost:6000/user/order' \
--header 'Content-Type: application/json' \
--data '{
    "customer_id": "6426d9ec391b6f01556ebe64",
    "products": [
        {
            "product_id": "6426b1c554ff40173fab6a1b",
            "quantity": 1
        }
    ]
}'
```

5. Update order
```
curl --location --request PATCH 'localhost:6000/user/order?id=6426df2c688b48cf42ba81bd' \
--header 'Content-Type: application/json' \
--data '{
    "status": "1"
}'
```

6. Get orders of customer
```
curl --location 'localhost:6000/user/orders?customer_id=6426d9ec391b6f01556ebe64'
```
