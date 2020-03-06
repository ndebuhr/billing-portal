db = connect("mongodb:27017/xebialabs");
db.expenses.remove( { reason: { $eq: "Automated Finance Sync 7a79623a9ccd36c4" } } )