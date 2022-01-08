package setu

import (
	"context"
	"strconv"
	"time"

	"github.com/0w0mewo/budong/pkg/domain/shetu"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ setuRepoProvider = &setuMongoDB{}

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

	table := client.Database("setu-micro").Collection("setu")

	return &setuMongoDB{
		client:  client,
		table:   table,
		timeout: 3 * time.Second,
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

	err := s.table.FindOne(ctx, bson.D{{"id", strconv.Itoa(id)}}).Decode(res)

	return res.Data, err
}

func (s *setuMongoDB) SelectByTitle(title string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	res := &shetu.Setu{}

	err := s.table.FindOne(ctx, bson.D{{"title", title}}).Decode(res)

	return res.Data, err
}

func (s *setuMongoDB) GetAmount() uint64 {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	cnt, err := s.table.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0
	}

	return uint64(cnt)
}

//TODO
func (s *setuMongoDB) ListInventory(offset uint64, limit uint64) ([]*shetu.SetuInfo, error) {
	return nil, nil
}

// TODO
func (s *setuMongoDB) SelectRandomId() (int, error) {
	return 0, nil
}

func (s *setuMongoDB) IsInDB(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	err := s.table.FindOne(ctx, bson.D{{"id", strconv.Itoa(id)}})

	return err == nil
}

func (s *setuMongoDB) Close() error {
	return s.client.Disconnect(context.Background())
}
