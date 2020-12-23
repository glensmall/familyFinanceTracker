function TransactionEntry(data){
    var self = this;

    self.details = ko.observable(data);
}

// finfanceHomeViewModel
function financeHomeViewModel(){

    self = this

    self.Testing = "This is some Knockout Test Data";
    self.userTransactions = ko.observableArray();

    // call our golang backend engine
   $.ajax({
        method : 'GET',
        url : 'http://localhost:8080/getEntries/',
        dataType : 'json'
    }).then(function(result){
            for(index in result.Transactions){
                self.userTransactions.push(new TransactionEntry(result.Transactions[index]));
            }
        });

    }
    
// apply this view model to the items in the page and go for it
ko.applyBindings(new financeHomeViewModel());

