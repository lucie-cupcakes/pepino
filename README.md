# pepino
Simple key-value database made with HTTP protocol in mind

# compiling, configuring & running as service

	git clone https://github.com/lucie-cupcakes/pepino.git
	make
	cp config.json.default config.json
	./pepino

Edit the config file to your likings before invoking the executable.
you can create a systemd unit, or [pm2](https://github.com/Unitech/pm2) app.json if you want to keep track of the service status.

# examples: Code
I made a simple CRUD Notes command line program, in different programming languages.

|Programming language| Repo |
|--|--|
| Go | [simple-notes-go](https://github.com/lucie-cupcakes/simple-notes-go) |
| Python | [simple-notes-py](https://github.com/lucie-cupcakes/simple-notes-py) |
| JavaScript (Node.js) | [simple-notes-js](https://github.com/lucie-cupcakes/simple-notes-js) |
| C# (Dotnet Core) | [simple-notes-cs](https://github.com/lucie-cupcakes/simple-notes-cs) |

# examples: CURL
We are gonna use the program [curl](https://curl.se/) to invoke HTTP requests.
Note: examples uses the default password.

Adding an entry (This will also create a Database if it not exists!)

	curl -X POST -d "Mongo DB is also a friend, not an enemy" 'localhost:50200/?password=caipiroska&db=mydb&entry=Friend'

Getting an entry:
	
	curl 'localhost:50200/?password=caipiroska&db=mydb&entry=Friend'

Deleting an entry:

	curl -X DELETE 'localhost:50200/?password=caipiroska&db=mydb&entry=Friend'

