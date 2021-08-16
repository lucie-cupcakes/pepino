# pepino
Key-value database made with simplicity in mind

# compiling, configuring & running as service

    git clone https://github.com/lucie-cupcakes/pepino.git
    make
    cp config.json.default config.json
    ./pepino

Edit the config file to your likings before invoking the executable.
you can create a systemd unit, or [pm2](https://github.com/Unitech/pm2) app.json if you want to keep track of the service status.

# What? Why? What does pepino even mean?

When I was trying out ``NoSQL`` databases, I tried `mongodb`, don't get me wrong, it's an excellent Database.
But it had waay to much to offer from what I actually needed, and it wasnt justifying the usage of memory and resources.
I thought to myself: Well I'm gonna make a new database engine, It will listen to HTTP and I can make it store and restore bytes and that's all I need.
Later I realized that I needed a way to let the user filter the data, so I added stored procedures as an option to have Python programs inside as Entries that can handle and filter user data exactly as the user wants.

**pepino means cucumber in Spanish**
When I first started the Database I asked my friend how should I name this program and they say this, I just followed them.
**the default password is an alcoholic drink that uses cucumber**

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

Adding an entry (This will also create a Database if it doesn't exists!)

	curl -X POST -d "Hello World!" 'localhost:50200/?password=caipiroska&db=mydb&entry=Hello'

Getting an entry:
    
    curl 'localhost:50200/?password=caipiroska&db=mydb&entry=Hello'

Deleting an entry:

    curl -X DELETE 'localhost:50200/?password=caipiroska&db=mydb&entry=Hello'

# examples: Stored Procedures
pepino allows the user to run entries as Python 3 code.

Here is an example program that loads an entry and prints it back:
```py
from os import environ
import requests

dburi = environ["PEPINODB_LURI"]
dbname = environ["PEPINODB_DB"]

res = requests.get(f"{dburi}&db={dbname}&entry=Hello")
res_str = res.content.decode("utf-8")
print("Entry 'hello' got from the Python Script = " + res_str)
```
Save the adobe example as ``program.py``, now we are gonna store it in the database:

	curl -X POST --data @- 'localhost:50100/?password=caipiroska&db=mydb&entry=program.py' < program.py

Now here is the trick, to ask PepinoDB to actually run the program, you have to pass the argument ``exec=true`` while calling the GET method:

	curl 'localhost:50200/?password=caipiroska&db=mydb&entry=program.py&exec=true'
After running this command you should get this back:
	
	Entry 'hello' got from the Python Script = Hello World!
	
Pretty cool, right? 😃

``@TODO: I'm gonna add an example parsing and filtering JSON data``

# Stored Procedures: env vars

As you seen in the examples before, we were using `environ["PEPINODB_LURI"]` this is because PepinoDB sets some Environmental Variables before running the Stored procedure.

Here are the variables being sent:
| Variable | Description |  Example Value | 
|--|--|--|
| ``PEPINODB_LURI`` | A convenience URI containing localhost as the host, the port and the password of the DatabaseHTTPService. | ``http://localhost:50200/?password=caipiroska`` |
| ``PEPINODB_HOST`` | The host the DatabaseHTTPService is listening to. | ``0.0.0.0`` |
| ``PEPINODB_PORT`` | The port the DatabaseHTTPService is listening to. | ``50200`` |
| ``PEPINODB_TLS`` | ``True`` if the DatabaseHTTPService is listening as HTTPS ``False`` if not | ``False`` |
| ``PEPINODB_PWD`` | The password needed for accessing the DatabaseHTTPService | ``caipiroska`` |
| ``PEPINODB_DB`` | The database name where the StoredProcedure is being runned from. | ``mydb`` |
| ``PEPINODB_SCRIPT`` | The StoredProcedure (entryName) being called. | ``program.py`` |

