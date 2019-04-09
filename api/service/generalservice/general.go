package generalservice

import (
    "net/http"
        "encoding/json"
    "avnos/api/model/usermodel"
        "avnos/api/db"
        "github.com/prologic/bitcask"
    "github.com/julienschmidt/httprouter"
    "encoding/binary"
    "github.com/satori/go.uuid"
    "strconv"
        )

func GenerateUniqueId() (int64) {
    u1, _ := uuid.NewV4()
    l1 := binary.BigEndian.Uint64(u1[:8])
    x := int64(l1 / 10000000000)
    return x
}

func TokenChecker(t string) bool{
    d := db.GetSession()
    defer d.Close()
    _, err := d.Get(t)
    if err != nil {
        return false
    }
    return true
}

func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    w.Header().Set("Content-Type", "application/json")
    decoder := json.NewDecoder(r.Body)
    var res usermodel.User
    err := decoder.Decode(&res)
    if err != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": err.Error()})
        return
    }
    res.EncryptPassword()
    if _, e := usermodel.GetUser(res.Username); e == nil {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": "Already Registered"})
        return
    }
    d := db.GetSession()
    defer d.Close()
    byteres, _ := usermodel.ConvertUser(res)
    d.Put(res.Email, byteres)
    d.Put(res.Username, byteres)
    json.NewEncoder(w).Encode(map[string]interface{}{"result":true, "obj": res})
}

func Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    w.Header().Set("Content-Type", "application/json")
    t := r.Header.Get("token")
    d := db.GetSession()
    defer d.Close()
    err := d.Delete(t)
    if err != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": err.Error()})
        return
    }
    json.NewEncoder(w).Encode(map[string]interface{}{"result":true, "msg": nil})
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    w.Header().Set("Content-Type", "application/json")
    decoder := json.NewDecoder(r.Body)
    var res usermodel.UserLogin
    err := decoder.Decode(&res)
    if err != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": err.Error()})
        return
    }
    res.EncryptPassword()
    
    d := db.GetSession()
    defer d.Close()
    val, errv := d.Get(res.Username)
    if errv == bitcask.ErrKeyNotFound {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": "Password Or Username Incorrect"})
        return
    }
    if errv != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": errv.Error()})
        return
    }
    
    u, errrev := usermodel.RevertUser(val)
    if errrev != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": errrev.Error()})
        return
    }
    if u.Password != res.Password {
        json.NewEncoder(w).Encode(map[string]interface{}{"result":false, "msg": "Password Or Username Incorrect"})
        return
    }
    t := strconv.Itoa(int(GenerateUniqueId()))
    d.Put(t, []byte(t))
    json.NewEncoder(w).Encode(map[string]interface{}{"result":true, "msg": "Success Login", "token": t})
}
