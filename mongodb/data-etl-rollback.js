db = connect("mongodb:27017/digitalai");
db.expenses.remove( { reason: { $eq: "Automated Finance Sync 7a79623a9ccd36c4" } } )