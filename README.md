# Partners Location API
This API simulates a delivery application where you can register and search partners by location.

### Requirements
* Docker 19.03.5 or later
* docker-compose 1.25.0 or later

### Installation
For install the project just run ``make install`` or if you prefer can run and install in the same time with ``make run``.
The application will run at ``localhost:3000``
Ps.: Observe the disponibility of ips at your docker network.

### Deployment

Just run ``make deploy``

### Technologies and Paterns

* GoLang
* Docker 
* MongoDB
* Lambda-local
* Hexagonal Architecture (Ports and Adapters)

Endpoints
---
#### Insert Partner

  _This endpoint saves a Partner._

* **URL**

  _/partners_

* **Method:**

  `POST`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
```json
{
	"id": "5e447f4f7cc7ecf5c2a900f1",
	"tradingName": "Adega da Cerveja - Pinheiros",
	"ownerName": "Zé da Silva",
	"document": "1432132123891/0001",
	"coverageArea": {
		"type": "MultiPolygon",
		"coordinates": [
			[
				[
					[30, 20],
					[45, 40],
					[10, 40],
					[30, 20]
				]
			],
			[
				[
					[15, 5],
					[40, 10],
					[10, 20],
					[5, 10],
					[15, 5]
				]
			]
		]
	},
	"address": {
		"type": "Point",
		"coordinates": [-46.57421, -21.785741]
	}
}
```
#### Get Partner

  _This endpoint return a Partner from a provided id._

* **URL**

  _/partners/{id}_

* **Method:**

  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
```json
{
	"id": "5e447f4f7cc7ecf5c2a900f1",
	"tradingName": "Adega da Cerveja - Pinheiros",
	"ownerName": "Zé da Silva",
	"document": "1432132123891/0001",
	"coverageArea": {
		"type": "MultiPolygon",
		"coordinates": [
			[
				[
					[30, 20],
					[45, 40],
					[10, 40],
					[30, 20]
				]
			],
			[
				[
					[15, 5],
					[40, 10],
					[10, 20],
					[5, 10],
					[15, 5]
				]
			]
		]
	},
	"address": {
		"type": "Point",
		"coordinates": [-46.57421, -21.785741]
	}
}
```
#### Search Partners

  _This endpoint return nearest Partner from a provided location._

* **URL**

  _/partners/lat/{lat}/lng/{lng}_

* **Method:**

  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
```json
{
	"id": "5e447f4f7cc7ecf5c2a900f1",
	"tradingName": "Adega da Cerveja - Pinheiros",
	"ownerName": "Zé da Silva",
	"document": "1432132123891/0001",
	"coverageArea": {
		"type": "MultiPolygon",
		"coordinates": [
			[
				[
					[30, 20],
					[45, 40],
					[10, 40],
					[30, 20]
				]
			],
			[
				[
					[15, 5],
					[40, 10],
					[10, 20],
					[5, 10],
					[15, 5]
				]
			]
		]
	},
	"address": {
		"type": "Point",
		"coordinates": [-46.57421, -21.785741]
	}
}
```

