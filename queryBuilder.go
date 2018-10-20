/*
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.
    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package clamor

import (
    "reflect"

    "github.com/jinzhu/gorm"
)

type QueryFilter struct {
    Field string
    Value string
}

type QueryRequest struct {
    Limit uint
    Offset uint
    Include []string
    Filters []QueryFilter
}

func BuildQueryWithoutPagination(db *gorm.DB, query QueryRequest, obj interface{}) *gorm.DB {
    dbQuery := db.Model(obj);

    for _, filter := range query.Filters {
        field := reflect.ValueOf(obj).Elem().FieldByName(filter.Field)
        if field.IsValid() {
            field.SetString(filter.Value)
        } else {
            panic("Unknown Field: " + filter.Field)
        }
    }
    dbQuery = dbQuery.Where(obj)

    for _, include := range query.Include {
        dbQuery = dbQuery.Preload(include)
    }

    return dbQuery
}

func BuildQuery(db *gorm.DB, query QueryRequest, obj interface{}) *gorm.DB {
    dbQuery := BuildQueryWithoutPagination(db, query, obj)
    return dbQuery.Limit(buildQueryGetLimit(query)).Offset(buildQueryGetOffset(query))
}

func buildQueryGetLimit(query QueryRequest) uint {
    //TODO: Make the max query limit configurable somewhere.
    maxLimit := uint(1000)
    if query.Limit > maxLimit {
        return maxLimit
    } else if query.Limit == 0 {
        return maxLimit
    } else {
        return query.Limit
    }
}

func buildQueryGetOffset(query QueryRequest) uint {
    return query.Offset
}
