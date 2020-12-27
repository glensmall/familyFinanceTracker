from flask import Flask

# lets create out flask app
app = Flask(__name__)

# load our config object
app.config.from_object('config.DevConfig')


##################################################
# Handleer for the getEntries request
# will return all the transactions in JSON format
##################################################
@app.route('/getEntries/', methods=['GET', 'POST'])
def getEntries():
    
    # is this a get
    if app.config['TESTING'] == True:
    
        with open(app.config['TEST_RESPONSE'], 'r') as fileHandle:
            fileContents = fileHandle.readlines()

            return(str(fileContents))


##################################################
# Handleer for the addEntry request
# will add a transaction to the database
##################################################
@app.route('/addEntry/', methods=['GET', 'POST'])
def addEntry():

    # add code here
    return()

##################################################
# Handleer for the editEntry request
# will update an existing transaction in the database
##################################################
@app.route('/editEntry/', methods=['GET', 'POST'])
def editEntry():

    # add code here
    return()

##################################################
# Handleer for the removeEntry request
# will remove an transacion from the database
##################################################
@app.route('/removeEntry/', methods=['GET', 'POST'])
def removeEntry():

    # add code here
    return()




# Entrypoint for this app
if __name__ == '__main__':
    app.run()