package mongodbadapter

import (
	"context"
	"time"

	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBPartnerAdapter struct {
	tableName  string
	collection *mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewMongoDBPartnerAdapter(dbName, tableName string, conn *mongo.Client) *MongoDBPartnerAdapter {
	if conn == nil {
		conn = CreateSession()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return &MongoDBPartnerAdapter{
		tableName,
		conn.Database(dbName).Collection(tableName),
		ctx,
		cancel,
	}
}

func (m *MongoDBPartnerAdapter) SavePartner(partner *partners.Partner) (*partners.Partner, error) {
	defer m.cancel()
	p, err := bson.Marshal(partner)
	if err != nil {
		return nil, err
	}
	res, err := m.collection.InsertOne(m.ctx, p)
	if err != nil {
		return nil, err
	}
	partner.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return partner, nil
}

func (m *MongoDBPartnerAdapter) GetPartner(id string) (*partners.Partner, error) {
	defer m.cancel()
	partner := new(partners.Partner)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objId}
	err = m.collection.FindOne(m.ctx, filter).Decode(partner)
	if err != nil {
		return nil, err
	}
	return partner, nil
}

func (m *MongoDBPartnerAdapter) SearchPartners(point *partners.Point) ([]partners.Partner, error) {

	defer m.cancel()
	filter := bson.D{
		{"coverageArea",
			bson.D{
				{"MultiPolygon", []float64{point.Latitude, point.Longitude}},
			}},
	}
	cur, err := m.collection.Find(m.ctx, filter)
	if err != nil {
		return nil, err
	}
	return m.getNearestPartner(cur, point)
}

func (m *MongoDBPartnerAdapter) getNearestPartner(cur *mongo.Cursor, point *partners.Point) ([]partners.Partner, error) {
	partnerRes := make([]partners.Partner, 0)
	defer closeCursor(cur, m.ctx)
	for cur.Next(m.ctx) {
		tempPartner := new(partners.Partner)
		err := cur.Decode(tempPartner)
		if err != nil {
			return nil, err
		}
		partnerRes = append(partnerRes, *tempPartner)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return partnerRes, nil
}

func closeCursor(cur *mongo.Cursor, ctx context.Context) {
	if err := cur.Close(ctx); err != nil {
		panic(err.Error())
	}
}
