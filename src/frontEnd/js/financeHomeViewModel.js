(function(){

// finfanceHomeViewModel
function financeHomeViewModel(){

    var self = this;

    // will hold the transactions from the backend
    self.transactionList = ko.observable();


    // we're going to need to make a call to our backend and get the data
    $.get('/getEntries/', {Date: range}, self.transactionList);

}


// apply this view model to the items in the page and go for it
ko.applyBindings(new financeHomeViewModel());
})();