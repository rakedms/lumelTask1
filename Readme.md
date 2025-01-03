lumel task 1

remark:

started working in the apis initially so created a single model to handle data
after that worked into loading csv so didn't ocnecmtrate on product and customer model, all controllers will be based on order model

unable to complete all requirements and test the apis, just exposing the endpoints in the documentation

cmd:

 go run .\cmd\server\main.go


API Documentation

POST http://localhost:8080/refresh-data?file=salesData/new_sales.csv  (new file)

GET http://localhost:8080/api/revenue?start_date=2022-01-01&end_date=2023-01-01

GET http://localhost:8080/api/revenue-by-product?start_date=2022-01-01&end_date=2023-01-01


GET http://localhost:8080/api/revenue-by-category?start_date=2022-01-01&end_date=2023-01-01

GET http://localhost:8080/api/revenue-by-region?start_date=2022-01-01&end_date=2023-01-01


GET http://localhost:8080/api/top-products?start_date=2022-01-01&end_date=2023-01-01&limit=5

GET http://localhost:8080/api/top-products-by-category?start_date=2022-01-01&end_date=2023-01-01&category=Electronics&limit=5


GET http://localhost:8080/api/top-products-by-region?start_date=2022-01-01&end_date=2023-01-01&region=North&limit=5


GET http://localhost:8080/api/total-customers?start_date=2022-01-01&end_date=2023-01-01


GET http://localhost:8080/api/total-orders?start_date=2022-01-01&end_date=2023-01-01


GET http://localhost:8080/api/average-order-value?start_date=2022-01-01&end_date=2023-01-01

