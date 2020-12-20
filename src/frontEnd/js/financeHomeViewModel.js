(function(){

function Transaction(_description, _date, _amount, _balance){
    var self = this;
    self.Description = _description
    self.Date = _date
    self.Amount = _amount
    self.Balance = _balance
}

// finfanceHomeViewModel
function financeHomeViewModel(){

    var self = this;

    self.transactionList = ko.observableArray([]);
 
    $.ajax({
        method : 'GET',
        url : '/backendEngine/getEntries/',
        dataType : 'json',
        success : function(data){
            for(var entry in data) {
                transactionList.push(new Transaction(entry[0], entry[1], entry[2], entry[3]))
            }
        }
    });

// apply this view model to the items in the page and go for it
ko.applyBindings(new financeHomeViewModel());
}
})();