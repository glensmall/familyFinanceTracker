from django.shortcuts import render
from django.http import HttpResponse

# default handler for /getEntries/ request
def getEntries(request):
    # open the test data
    with open("testData.json", "r") as fileHandle:
    
        jsonResponse = fileHandle.readlines()

        return(HttpResponse(jsonResponse))


# TODO:  Add handlers in here for all the backend types we will get
#   addEntry
#   editEntry
#   removeEntry

# need to update the getEntries handler to work with mongo backend as well.
# will also need to define a dockerfile to run this app in a container
# and another one to run the web frontend


