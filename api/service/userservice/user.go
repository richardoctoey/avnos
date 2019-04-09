package userservice

import (
    "net/http"
    "avnos/api/db"
        "github.com/julienschmidt/httprouter"
        "avnos/api/model/usermodel"
    "encoding/json"
    )

func ListUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Header().Set("Content-Type", "application/json")
    d := db.GetSession()
    defer d.Close()
    res := map[string]usermodel.User{}
    d.Scan("", func(k string) error {
        params := r.URL.Query()
        if _, ok := params["name"]; ok {
            if k == params["name"][0] {
                byteres, _ := d.Get(k)
                resn, _ := usermodel.RevertUser(byteres)
                if _, ok := res[resn.Username]; !ok {
                    res[resn.Username] = resn
                }
            }
        } else {
            byteres, _ := d.Get(k)
            resn, _ := usermodel.RevertUser(byteres)
            if len(resn.Username) != 0 {
                if _, ok := res[resn.Username]; !ok {
                    res[resn.Username] = resn
                }
            }
        }
        return nil
    })
    json.NewEncoder(w).Encode(map[string]interface{}{"result":true, "obj": res})
}

func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Header().Set("Content-Type", "application/json")
    d := db.GetSession()
    defer d.Close()
    k := ps.ByName("key")
    _, err := d.Get(k)
    if err != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": err.Error()})
        return
    }
    
    decoder := json.NewDecoder(r.Body)
    var res usermodel.User
    errd := decoder.Decode(&res)
    if errd != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": errd.Error()})
        return
    }
    res.EncryptPassword()
    byteres, _ := usermodel.ConvertUser(res)
    d.Put(res.Email, byteres)
    d.Put(res.Username, byteres)
    json.NewEncoder(w).Encode(map[string]interface{}{"result":true, "msg": "Success update"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Header().Set("Content-Type", "application/json")
    d := db.GetSession()
    defer d.Close()
    k := ps.ByName("key")
    d.Delete(k)
    json.NewEncoder(w).Encode(map[string]interface{}{"result":true, "msg": "Deleted"})
}
