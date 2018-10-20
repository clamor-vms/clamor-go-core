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
    "net/http"
    "encoding/json"
)

type ControllerResponse struct {
    Status int
    Body interface{}
}

//IController
type IController interface {
    //every controller should support these four methods... even if they just return 404
    Get(w http.ResponseWriter, r *http.Request) ControllerResponse
    Put(w http.ResponseWriter, r *http.Request) ControllerResponse
    Post(w http.ResponseWriter, r *http.Request) ControllerResponse
    Delete(w http.ResponseWriter, r *http.Request) ControllerResponse
}

//Controller wrapper to provide highlevel logic in a non dup'd way.
type ControllerProcessor struct {
    controller IController
}
func NewControllerProcessor(controller IController) *ControllerProcessor {
    return &ControllerProcessor{controller: controller}
}
func (p *ControllerProcessor) Logic(w http.ResponseWriter, r *http.Request) {
    resp := ControllerResponse{Status: http.StatusNotFound, Body: EmptyResponse{}}

    //route to the wrapped controller function based on request method
    switch r.Method {
        case "GET":
            resp = p.controller.Get(w, r)
            break
        case "POST":
            resp = p.controller.Post(w, r)
            break
        case "PUT":
            resp = p.controller.Put(w, r)
            break
        case "DELETE":
            resp = p.controller.Delete(w, r)
            break
    }

    p.writeJsonOutput(w, resp)
}
func (p *ControllerProcessor) writeJsonOutput(w http.ResponseWriter, resp ControllerResponse) {
    w.WriteHeader(resp.Status)
    w.Header().Set("Content-Type", "application/json")
    jsonResp, err := json.Marshal(resp.Body)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("{\"Success\": false,\"Message\": \"Cant Generate Json\"}"))
    } else {
        w.Write(jsonResp)
    }
}
