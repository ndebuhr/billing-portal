// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ndebuhr/billing-service/models"
	"github.com/ndebuhr/billing-service/restapi/operations"
)

func configureFlags(api *operations.BillingPortalAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BillingPortalAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.AddExpenseHandler = operations.AddExpenseHandlerFunc(func(params operations.AddExpenseParams) middleware.Responder {
		db := getMongoClient()
		expense := params.Expense
		collection := db.Database("xebialabs").Collection("expenses")
		_, err := collection.InsertOne(context.TODO(), expense)
		if err != nil {
			log.Printf("failed mongodb insert operation: %v", err)
		}
		db.Disconnect(context.TODO())
		return operations.NewAddExpenseOK()
	})

	api.GetExpensesHandler = operations.GetExpensesHandlerFunc(func(params operations.GetExpensesParams) middleware.Responder {
		db := getMongoClient()
		findOptions := options.Find()
		findOptions.SetLimit(*params.Size)
		collection := db.Database("xebialabs").Collection("expenses")
		cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			log.Printf("failed mongodb find operation: %v", err)
		}
		var expenses []models.Expense
		for cur.Next(context.TODO()) {
			var expense models.Expense
			err := cur.Decode(&expense)
			if err != nil {
				log.Printf("failed mongodb result decoding: %v", err)
			}
			expenses = append(expenses, expense)
		}
		var payload models.Expenses
		for i := 0; i < len(expenses); i++ {
			payload = append(payload, &expenses[i])
		}
		return operations.NewGetExpensesOK().WithPayload(payload)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

func getMongoClient() *mongo.Client {
	dbOptions := options.Client().ApplyURI(
		"mongodb://xebialabs:xebialabs@mongodb:27017",
	)
	db, err := mongo.Connect(context.TODO(), dbOptions)
	if err != nil {
		log.Printf("database connection failed: %v", err)
	}
	// Check the database connection
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf("database ping failed: %v", err)
	}
	return db
}