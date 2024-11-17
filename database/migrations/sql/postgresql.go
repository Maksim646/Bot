package sql

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

const (
	uniqueViolation = "23505"
)

var (
	Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func IsUniqueViolation(err error) bool {
	var pqError *pq.Error
	if !errors.As(err, &pqError) {
		return false
	}

	return pqError.Code == uniqueViolation
}

func BuildQuery(query string, params []interface{}) string {
	for i := len(params) - 1; i >= 0; i-- {
		query = strings.ReplaceAll(query, "$"+strconv.Itoa(i+1), fmt.Sprint(params[i]))
	}

	return query
}
