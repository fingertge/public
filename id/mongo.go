package id

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

type MongoRenew struct {
	dbName string
	table  string
	client *mongo.Client
	opts   *options.CollectionOptions
}

func NewMongoRenew(client *mongo.Client, dbName, table string) *MongoRenew {
	return &MongoRenew{
		client: client,
		dbName: dbName,
		table:  table,
		opts: &options.CollectionOptions{
			ReadConcern:    readconcern.Majority(),
			WriteConcern:   writeconcern.Majority(),
			ReadPreference: readpref.Primary(),
		},
	}
}

func (m *MongoRenew) pingPrimary() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.client.Ping(ctx, readpref.Primary())
}

func (m *MongoRenew) Prepare() error {
	return m.pingPrimary()
}

func (m *MongoRenew) Renew(ctx context.Context, domain string, quantum, offset uint64) (uint64, error) {
	curr, err := m.renew(ctx, domain, quantum)
	if errors.Is(err, mongo.ErrNoDocuments) {
		c := m.client.Database(m.dbName).Collection(m.table, m.opts)
		_, err = c.InsertOne(ctx, bson.D{{Key: "_id", Value: domain}, {Key: "current", Value: offset}})
		if err != nil {
			return 0, err
		}
		return m.renew(ctx, domain, quantum)
	}
	return curr, nil
}

func (m *MongoRenew) renew(ctx context.Context, domain string, quantum uint64) (uint64, error) {
	c := m.client.Database(m.dbName).Collection(m.table, m.opts)

	filter := bson.D{{Key: "_id", Value: domain}}
	update := bson.D{{Key: "$incr", Value: bson.D{{Key: "value", Value: quantum}}}}

	var opts options.FindOneAndUpdateOptions
	opts.SetUpsert(false).SetReturnDocument(options.Before)

	var doc struct{ Current uint64 }
	err := c.FindOneAndUpdate(ctx, filter, update, &opts).Decode(&doc)
	if err != nil {
		return 0, err
	}
	return doc.Current, nil
}
