# Authentication
Authentication

CURL
Post Method
curl --location 'localhost:8080/user/v1/daijai/project' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYzNTMyMTMsInVzZXJJZCI6IjlhOGU4NmUyLTJlMmQtNGJhNy1hZTdhLTQwNGMyZDY4N2FmMyJ9.uanZxtyPJdP8ElI2TwCxrUWLbDey2WeiVOgOFl5GBLs' \
--header 'Content-Type: application/json' \
--data '{
"projectName":"MK_restaurant",
"createdBy":"sitikorn"
}'

Get Method Get projectId
curl --location 'localhost:8080/user/v1/daijai/Name' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYzNTMyMTMsInVzZXJJZCI6IjlhOGU4NmUyLTJlMmQtNGJhNy1hZTdhLTQwNGMyZDY4N2FmMyJ9.uanZxtyPJdP8ElI2TwCxrUWLbDey2WeiVOgOFl5GBLs' \
--data ''

Get Method All project
curl --location 'localhost:8080/user/v1/daijai/getall' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYzNTMyMTMsInVzZXJJZCI6IjlhOGU4NmUyLTJlMmQtNGJhNy1hZTdhLTQwNGMyZDY4N2FmMyJ9.uanZxtyPJdP8ElI2TwCxrUWLbDey2WeiVOgOFl5GBLs' \
--data ''

Update
curl --location --request PUT 'localhost:8080/user/v1/daijai/update-project' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYzNTcwMjQsInVzZXJJZCI6IjlhOGU4NmUyLTJlMmQtNGJhNy1hZTdhLTQwNGMyZDY4N2FmMyJ9.VD4-uxsmcbi-mpGZeXmvSzuO2pXv1pyWlae0DZa9Ndk' \
--header 'Content-Type: application/json' \
--data '{
"projectName":"MK_restaurant",
"status":"approved",
"createdBy":"pakornsit"
}'