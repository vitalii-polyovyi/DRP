package entity

import (
	"context"
	"drp/logger/helpers"
	"drp/logger/models"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pagination struct {
	Skip  int64
	Limit int64
}

func getPagination(r *http.Request) Pagination {
	var pageQ string = r.URL.Query().Get("page")
	var pageSizeQ string = r.URL.Query().Get("pageSize")

	var page int64 = 1
	if pageQ != "" {
		page, _ = strconv.ParseInt(pageQ, 0, 64)
	}
	skip := int64(math.Max(float64(page-1), 0)) * 20

	var pageSize int64 = 20
	if pageSizeQ != "" {
		pageSize, _ = strconv.ParseInt(pageSizeQ, 0, 64)
		if pageSize <= 0 {
			pageSize = 20
		}
	}

	return Pagination{
		Limit: pageSize,
		Skip:  skip,
	}
}

func getSort(r *http.Request) bson.D {
	var sortQ string = r.URL.Query().Get("sort")

	sortData := strings.Split(sortQ, ":")

	sort := bson.D{{Key: "created_at", Value: -1}}
	if len(sortData) == 1 && sortData[0] != "" {
		sort = bson.D{{Key: sortData[0], Value: 1}}
	} else if len(sortData) == 2 {
		dir := 1
		if strings.ToLower(sortData[1]) == "desc" {
			dir = -1
		}

		sort = bson.D{{Key: sortData[0], Value: dir}}
	}

	return sort
}

func getFilter(r *http.Request) bson.M {
	var filter bson.M

	query := r.URL.Query().Get("filter")
	err := json.Unmarshal([]byte(query), &filter)
	if err != nil {
		filter = bson.M{}
	}
	filter["app"] = r.Context().Value("app")

	return filter
}

func validate(model models.Model) []*helpers.ErrorResponse {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		log.Printf("validationError=%v\n", err)

		var errors []*helpers.ErrorResponse

		for _, err := range err.(validator.ValidationErrors) {
			var el helpers.ErrorResponse
			el.Error = err.Error()
			errors = append(errors, &el)
		}

		return errors
	}

	return nil
}

func getEntityQuery(objId primitive.ObjectID, r *http.Request) bson.D {
	return bson.D{
		{Key: "_id", Value: objId},
		{Key: "app", Value: r.Context().Value("app")},
	}
}

func List[T models.Model](r *http.Request) (map[string]interface{}, error) {
	var model T
	var records []T = make([]T, 0)

	sort := getSort(r)
	pagination := getPagination(r)
	filter := getFilter(r)

	log.Printf("pagination=%v\n", pagination)
	log.Printf("sort=%v\n", sort)
	log.Printf("filter=%v\n", filter)

	findOptions := options.Find()
	findOptions.SetSkip(pagination.Skip)
	findOptions.SetLimit(pagination.Limit)
	findOptions.SetSort(sort)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := model.GetCollection().Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	count, err := model.GetCollection().CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var elem T
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		records = append(records, elem)
	}

	return map[string]interface{}{
		"rows":  records,
		"total": count,
		"pages": int64(math.Ceil(float64(count) / float64(pagination.Limit))),
	}, nil
}

func Create[T models.Model](r *http.Request) (*mongo.InsertOneResult, []*helpers.ErrorResponse, error) {
	var model T

	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		return nil, nil, err
	}

	model.SetOnCreate()
	model.SetApp(fmt.Sprintf("%s", r.Context().Value("app")))

	log.Printf("%v\n", model)

	if err := validate(model); err != nil {
		return nil, err, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := model.GetCollection().InsertOne(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return result, nil, nil
}

func CreateBulk[T models.Model](r *http.Request) (*mongo.InsertManyResult, []*helpers.ErrorResponse, error) {
	var models []T
	var validModels []interface{}
	var validationErrors []*helpers.ErrorResponse = make([]*helpers.ErrorResponse, 0)

	if err := json.NewDecoder(r.Body).Decode(&models); err != nil {
		return nil, nil, err
	}

	for _, model := range models {
		model.SetOnCreate()
		model.SetApp(fmt.Sprintf("%s", r.Context().Value("app")))

		log.Printf("%v\n", model)

		if err := validate(model); err != nil {
			validationErrors = append(validationErrors, err...)
			continue
		}

		validModels = append(validModels, model)
	}

	if len(validModels) == 0 {
		return nil, validationErrors, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var m T
	result, err := m.GetCollection().InsertMany(ctx, validModels)
	if err != nil {
		return nil, nil, err
	}

	return result, nil, nil
}

func Update[T models.Model](r *http.Request) (*mongo.UpdateResult, []*helpers.ErrorResponse, error) {
	var model T

	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		return nil, nil, err
	}

	model.SetOnUpdate()
	model.SetApp(fmt.Sprintf("%s", r.Context().Value("app")))

	log.Printf("%v\n", model)

	if err := validate(model); err != nil {
		return nil, err, nil
	}

	var docId string = fmt.Sprint(r.Context().Value("id"))
	objId, _ := primitive.ObjectIDFromHex(docId)

	log.Println(objId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := getEntityQuery(objId, r)
	result, err := model.GetCollection().UpdateOne(ctx, query, bson.M{"$set": model})
	if err != nil {
		return nil, nil, err
	}

	return result, nil, nil
}

func Get[T models.Model](r *http.Request) (*T, error) {
	var docId string = fmt.Sprint(r.Context().Value("id"))
	objId, _ := primitive.ObjectIDFromHex(docId)

	var model T

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := getEntityQuery(objId, r)
	if err := model.GetCollection().FindOne(ctx, query).Decode(&model); err != nil {
		return nil, err
	}

	return &model, nil
}

func Delete[T models.Model](r *http.Request) (*mongo.DeleteResult, error) {
	var docId string = fmt.Sprint(r.Context().Value("id"))
	objId, _ := primitive.ObjectIDFromHex(docId)

	var model T

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := getEntityQuery(objId, r)
	result, err := model.GetCollection().DeleteOne(ctx, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}
