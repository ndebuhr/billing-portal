db = connect("mongodb:27017/digitalai");
db.expenses.insertOne( { amount: 15.54, currency: "USD", reason: "Automated Finance Sync 7a79623a9ccd36c4" } );