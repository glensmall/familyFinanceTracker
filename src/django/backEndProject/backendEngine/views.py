from django.shortcuts import render
from django.http import HttpResponse

# default handler for /getEntries/ request
def getEntries(request):
    # open the test data
    with open("testData.json", "r") as fileHandle:
    
        jsonResponse = fileHandle.readlines()

        return(HttpResponse(jsonResponse))


