Per request for basic CRUD operations using graphql and internal in memory database

Clone repository
Then run go mod tidy to download all required modules.

This project can be run using the following command: 
go run server.go

This will stand up the database and graphql engine.

To view documentation regarding the graphql end points please see the included graphql playground

To visit open web browser and go to http://localhost:8080/

End point is reachable via http://localhost:8080/book

example of curl command:
curl -skX POST -H "Content-Type: application/json" -d '{"query":"query{books { title author date_pub book_cvr_img}}"}' http://localhost:8080/book

The db folder contains the functions and generic setup of the in memory database

GORM was the selected library used to provide connection to the database. This was chosen due to GORM's ability
to be modular and tailored around the type of database by swapping out the connection and calls the database connector.

This satisfies making database management easy to handle with as little disruption to core code already in place.

The graphql database was stood up using the schema first method.