function TransactionEntry(data){
    var self = this;

    self.details = ko.observable(data);
}

// finfanceHomeViewModel
function financeHomeViewModel(){

    self = this

    self.Testing = "This is some Knockout Test Data";
    self.userTransactions = ko.observableArray();

   $.ajax({
        method : 'GET',
        url : 'http://localhost:8000/testdata.json',
        dataType : 'json'
    }).then(function(result){
            for(index in result.Transactions){
                self.userTransactions.push(new TransactionEntry(result.Transactions[index]));
            }
        });

    }
    
// apply this view model to the items in the page and go for it
ko.applyBindings(new financeHomeViewModel());

