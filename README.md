# SWRT

SWRT is a cli tool to parse swagger.json and show routing.

# Install

	go get -u github.com/decafe09/swrt

# Usage

	swrt -f swagger/swagger.json

# Help

	Usage of swrt:
	  --help
	    	show help
	  --version
	    	show version
	  -f string
	    	import path your swagger.json (default "swagger.json")
	  -h	show help
	  -v	show version

# Example

This example use [Swagger Petstore sample](http://petstore.swagger.io/v2/swagger.json)

```
$ curl -s http://petstore.swagger.io/v2/swagger.json > swagger.json
$ swrt
http://petstore.swagger.io/v2
   PUT	/pet
  POST	/pet
   GET	/pet/findByStatus
   GET	/pet/findByTags
   GET	/pet/{petId}
  POST	/pet/{petId}
DELETE	/pet/{petId}
  POST	/pet/{petId}/uploadImage
   GET	/store/inventory
  POST	/store/order
   GET	/store/order/{orderId}
DELETE	/store/order/{orderId}
  POST	/user
  POST	/user/createWithArray
  POST	/user/createWithList
   GET	/user/login
   GET	/user/logout
   GET	/user/{username}
   PUT	/user/{username}
DELETE	/user/{username}
```