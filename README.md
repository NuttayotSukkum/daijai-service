# Authentication
Authentication

CURL
Post Method 
curl --location 'localhost:8080/user/v1/daijai/project' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYwNTEwNDEsInVzZXJJZCI6IjlhOGU4NmUyLTJlMmQtNGJhNy1hZTdhLTQwNGMyZDY4N2FmMyJ9.kRALgFOJWuRrnDTfbjTaHAitpr1DoFiljPpKdvUGTBg' \
--header 'Content-Type: application/json' \
--data '{
"project_name":"name",
"created_by":"name1"
}'

Get Method Get projectId
curl --location 'localhost:8080/user/v1/daijai/9ee56e3f-e915-47f8-969f-bcb04f7f41ae' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYwNTE0MDcsInVzZXJJZCI6IjlhOGU4NmUyLTJlMmQtNGJhNy1hZTdhLTQwNGMyZDY4N2FmMyJ9.ybKDZTMeR6EnZWitrxgp4vjMoYzAjqHaUrcpi-EU1-Q' \
--data ''

Get Method All project
curl --location 'localhost:8080/user/v1/daijai/getall' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYwNTE0MDcsInVzZXJJZCI6IjlhOGU4NmUyLTJlMmQtNGJhNy1hZTdhLTQwNGMyZDY4N2FmMyJ9.ybKDZTMeR6EnZWitrxgp4vjMoYzAjqHaUrcpi-EU1-Q' \
--data ''