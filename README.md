# Run code

In terminal run below command:
```
go run .\main.go
```


# sample cURL

```
curl --location --request GET 'http://localhost:8080' \
--header 'Content-Type: application/json' \
--data-raw '{
"ev": "contact_form_submitted",
"et": "form_submit",
"id": "cl_app_id_001",
"uid": "cl_app_id_001-uid-001",
"mid": "cl_app_id_001-uid-001",
"t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"l": "en-US",
"sc": "1920 x 1080",
"atrk1": "form_varient",
"atrv1": "red_top",
"atrt1": "string",
"atrk2": "ref",
"atrv2": "XPOWJRICW993LKJD",
"atrt2": "string",
"uatrk1": "name",
"uatrv1": "iron man",
"uatrt1": "string",
"uatrk2": "email",
"uatrv2": "ironman@avengers.com",
"uatrt2": "string",
"uatrk3": "age",
"uatrv3": "32",
"uatrt3": "integer"
}
'
```