package query

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryBuilder struct {
	db      *gorm.DB
	filters []Filter
	sorts   []Sort
}

type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}

type Sort struct {
	Field string
	Desc  bool
}

func NewQueryBuilder(db *gorm.DB) *QueryBuilder {
	return &QueryBuilder{
		db:      db,
		filters: make([]Filter, 0),
		sorts:   make([]Sort, 0),
	}
}

func (qb *QueryBuilder) Where(field string, value interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: "=", Value: value})
	return qb
}

func (qb *QueryBuilder) WhereNot(field string, value interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: "!=", Value: value})
	return qb
}

func (qb *QueryBuilder) WhereIn(field string, values interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: "in", Value: values})
	return qb
}

func (qb *QueryBuilder) WhereNotIn(field string, values interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: "not in", Value: values})
	return qb
}

func (qb *QueryBuilder) WhereLike(field string, value string) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: "like", Value: fmt.Sprintf("%%%s%%", value)})
	return qb
}

func (qb *QueryBuilder) WhereGreaterThan(field string, value interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: ">", Value: value})
	return qb
}

func (qb *QueryBuilder) WhereLessThan(field string, value interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: "<", Value: value})
	return qb
}

func (qb *QueryBuilder) WhereGreaterThanOrEqual(field string, value interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: ">=", Value: value})
	return qb
}

func (qb *QueryBuilder) WhereLessThanOrEqual(field string, value interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: "<=", Value: value})
	return qb
}

func (qb *QueryBuilder) WhereBetween(field string, start, end interface{}) *QueryBuilder {
	qb.filters = append(qb.filters, Filter{Field: field, Operator: "between", Value: []interface{}{start, end}})
	return qb
}

func (qb *QueryBuilder) OrderBy(field string) *QueryBuilder {
	qb.sorts = append(qb.sorts, Sort{Field: field, Desc: false})
	return qb
}

func (qb *QueryBuilder) OrderByDesc(field string) *QueryBuilder {
	qb.sorts = append(qb.sorts, Sort{Field: field, Desc: true})
	return qb
}

func (qb *QueryBuilder) Build() *gorm.DB {
	query := qb.db

	for _, filter := range qb.filters {
		switch filter.Operator {
		case "=":
			query = query.Where(fmt.Sprintf("%s = ?", filter.Field), filter.Value)
		case "!=":
			query = query.Where(fmt.Sprintf("%s != ?", filter.Field), filter.Value)
		case ">":
			query = query.Where(fmt.Sprintf("%s > ?", filter.Field), filter.Value)
		case "<":
			query = query.Where(fmt.Sprintf("%s < ?", filter.Field), filter.Value)
		case ">=":
			query = query.Where(fmt.Sprintf("%s >= ?", filter.Field), filter.Value)
		case "<=":
			query = query.Where(fmt.Sprintf("%s <= ?", filter.Field), filter.Value)
		case "like":
			query = query.Where(fmt.Sprintf("%s LIKE ?", filter.Field), filter.Value)
		case "in":
			query = query.Where(fmt.Sprintf("%s IN ?", filter.Field), filter.Value)
		case "not in":
			query = query.Where(fmt.Sprintf("%s NOT IN ?", filter.Field), filter.Value)
		case "between":
			values := filter.Value.([]interface{})
			query = query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", filter.Field), values[0], values[1])
		}
	}

	for _, sort := range qb.sorts {
		order := "ASC"
		if sort.Desc {
			order = "DESC"
		}
		query = query.Order(fmt.Sprintf("%s %s", sort.Field, order))
	}

	return query
}

func ParseFiltersFromQuery(c *gin.Context, allowedFields []string) *QueryBuilder {
	qb := NewQueryBuilder(nil)

	allowedFieldsMap := make(map[string]bool)
	for _, field := range allowedFields {
		allowedFieldsMap[field] = true
	}

	for key, values := range c.Request.URL.Query() {
		if len(values) == 0 {
			continue
		}

		value := values[0]
		parts := strings.Split(key, "__")
		field := parts[0]

		if !allowedFieldsMap[field] {
			continue
		}

		if len(parts) == 1 {
			qb.Where(field, value)
			continue
		}

		operator := parts[1]
		switch operator {
		case "eq":
			qb.Where(field, value)
		case "ne":
			qb.WhereNot(field, value)
		case "gt":
			qb.WhereGreaterThan(field, value)
		case "lt":
			qb.WhereLessThan(field, value)
		case "gte":
			qb.WhereGreaterThanOrEqual(field, value)
		case "lte":
			qb.WhereLessThanOrEqual(field, value)
		case "like":
			qb.WhereLike(field, value)
		case "in":
			vals := strings.Split(value, ",")
			qb.WhereIn(field, vals)
		}
	}

	if sortBy := c.Query("sort_by"); sortBy != "" {
		if strings.HasPrefix(sortBy, "-") {
			qb.OrderByDesc(strings.TrimPrefix(sortBy, "-"))
		} else {
			qb.OrderBy(sortBy)
		}
	}

	return qb
}

func (qb *QueryBuilder) ApplyFilters(db *gorm.DB) *gorm.DB {
	qb.db = db
	return qb.Build()
}
