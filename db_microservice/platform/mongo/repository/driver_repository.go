package repository

import (
	"bitaksi_burakcanheyal/db_microservice/internal/domain/entity"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DriverRepository struct {
	collection *mongo.Collection
}

func NewDriverRepository(col *mongo.Collection) *DriverRepository {
	return &DriverRepository{collection: col}
}

// ──────────────────────────────────────────────
// CREATE
// ──────────────────────────────────────────────
func (r *DriverRepository) Create(ctx context.Context, d *entity.Driver) (string, error) {

	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	// ⭐ InsertOne -> GEO JSON location artık doğru formatta
	res, err := r.collection.InsertOne(ctx, d)
	if err != nil {
		return "", err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", mongo.ErrNilDocument
	}

	return oid.Hex(), nil
}

// ──────────────────────────────────────────────
// UPDATE
// ──────────────────────────────────────────────
func (r *DriverRepository) Update(ctx context.Context, id string, d *entity.Driver) error {

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	d.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"firstName": d.FirstName,
			"lastName":  d.LastName,
			"plate":     d.Plate,
			"taksiType": d.TaxiType,
			"carBrand":  d.CarBrand,
			"carModel":  d.CarModel,
			"location":  d.Location, // ⭐ Burada artık GeoJSON
			"updatedAt": d.UpdatedAt,
		},
	}

	_, err = r.collection.UpdateByID(ctx, oid, update)
	return err
}

// ──────────────────────────────────────────────
// LIST (Paged)
// ──────────────────────────────────────────────
func (r *DriverRepository) List(ctx context.Context, page, pageSize int) ([]entity.Driver, error) {

	skip := (page - 1) * pageSize

	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize))

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var drivers []entity.Driver
	if err := cursor.All(ctx, &drivers); err != nil {
		return nil, err
	}

	return drivers, nil
}

// ──────────────────────────────────────────────
// LIST ALL (No Pagination)
// ──────────────────────────────────────────────
func (r *DriverRepository) ListAll(ctx context.Context) ([]entity.Driver, error) {

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var drivers []entity.Driver
	if err := cursor.All(ctx, &drivers); err != nil {
		return nil, err
	}

	return drivers, nil
}
