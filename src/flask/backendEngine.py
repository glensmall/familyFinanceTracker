from flask import Flask

# lets create out flask app
app = Flask(__name__)

# load our config object
app.config.from_object('config.Config')


# handlers for incoming requests
@app.route('getEntries/')
def getEntries():
    
    # add some code here to handle that


@app.route('addEntry/')
def addEntry():

    # add code here

@app.route('editEntry/')
def editEntry():

    # add code here

@app.route('removeEntry/')
def removeEntry():

    # add code here