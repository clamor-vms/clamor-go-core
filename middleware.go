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

package skaioskit

import (
    "errors"
    "net/http"
)

func PanicHandler(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var err error
        defer func() {
            r := recover()
            if r != nil {
                switch t := r.(type) {
                case string:
                    err = errors.New(t)
                case error:
                    err = t
                default:
                    err = errors.New("Unknown error")
                }
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        }()
        h.ServeHTTP(w, r)
    })
}
