package setu

import (
	"context"
	"math/rand"
	"time"

	"github.com/0w0mewo/budong/pkg/domain/shetu"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ setuRepoProvider = &setuMongoDB{}

var stageNoImgBytes = bson.D{{"$project", bson.M{"data": 0}}}

func listInvenPipeline(skip, limit int64) mongo.Pipeline {
	return mongo.Pipeline{stageNoImgBytes, bson.D{{"$skip", skip}}, bson.D{{"$limit", limit}}}
}

func randSetuPipeline(n int64) mongo.Pipeline {
	if n <= 0 {
		n = 1
	}
	i := rand.Int63n(n)

	return mongo.Pipeline{stageNoImgBytes, bson.D{{"$skip", i}}, bson.D{{"$limit", 1}}}

}

// var matchAll = bson.D{{"$match", bson.D{}}}

type setuMongoDB struct {
	client  *mongo.Client
	table   *mongo.Collection
	timeout time.Duration
}

func newSetuMongoDB(dsn string) *setuMongoDB {
	opt := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	// collection with index of image id
	table := client.Database("setu-micro").Collection("setu")
	table.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"id": -1},
		Options: nil,
	})

	return &setuMongoDB{
		client:  client,
		table:   table,
		timeout: 10 * time.Second,
	}

}

func (s *setuMongoDB) Create(setu *shetu.SetuInfo) (*shetu.Setu, error) {
	// fetch it
	newRow, err := shetu.SetuFromSetuInfo(setu, true)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	_, err = s.table.InsertOne(ctx, newRow)

	return newRow, err
}

func (s *setuMongoDB) SelectById(id int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	res := &shetu.Setu{}
	matchIdStage := bson.D{{"$match", bson.M{"id": id}}}
	cur, err := s.table.Aggregate(ctx, mongo.Pipeline{matchIdStage})
	if err != nil {
		return nil, err
	}

	// select one record only
	cur.Next(ctx)
	err = cur.Decode(res)

	return res.Data, err
}

func (s *setuMongoDB) SelectByTitle(title string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	res := &shetu.Setu{}
	err := s.table.FindOne(ctx, bson.M{"title": title}).Decode(res)

	return res.Data, err
}

func (s *setuMongoDB) GetAmount() int64 {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	cnt, err := s.table.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0
	}

	return cnt
}

func (s *setuMongoDB) ListInventory(offset int64, limit int64) ([]*shetu.SetuInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	// paging result is excluded image bytes field
	dbres, err := s.table.Aggregate(ctx, listInvenPipeline(offset, limit))
	if err != nil {
		return nil, err
	}

	res := make([]*shetu.Setu, 0)

	err = dbres.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	ret := make([]*shetu.SetuInfo, 0)
	for _, r := range res {
		ret = append(ret, shetu.SetuToSetuInfo(r))
	}

	return ret, nil

}

func (s *setuMongoDB) SelectRandomId() (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	// exclude image bytes field from result
	cur, err := s.table.Aggregate(ctx, randSetuPipeline(s.GetAmount()))
	if err != nil {
		return -1, err
	}
	cur.Next(ctx)
	res := &shetu.Setu{}
	err = cur.Decode(res)

	return res.Id, err
}

func (s *setuMongoDB) IsInDB(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	err := s.table.FindOne(ctx, bson.M{"id": id})

	return err == nil
}

func (s *setuMongoDB) Close() error {
	return s.client.Disconnect(context.Background())
}
