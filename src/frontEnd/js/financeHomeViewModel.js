function TransactionEntry(data){
    var self = this;

    self.details = ko.observable(data);
}

// finfanceHomeViewModel
function financeHomeViewModel(){

    self = this
    self.userTransactions = ko.observableArray();

    // call our golang backend engine
   $.ajax({
        method : 'GET',
        url : 'http://127.0.0.1:8000/backendEngine/getEntries/',
        dataType : 'json'
    }).then(function(result){
            for(index in result.Transactions){
                self.userTransactions.push(new TransactionEntry(result.Transactions[index]));
            }
        });

    }
    
// apply this view model to the items in the page and go for it
ko.applyBindings(new financeHomeViewModel());

