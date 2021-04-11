package main

import (
                "log"
                "fmt"
                "net/http"
                "os"
                "strings"
                "database/sql"
                "encoding/json"
                "time"
                "io/ioutil"

                "github.com/dgrijalva/jwt-go"
                "golang.org/x/crypto/bcrypt"
                _ "github.com/go-sql-driver/mysql"
)
var db,err = sql.Open("mysql","securednsapi:rSiEN95k%6^NhKjcsScD7wsn@tcp(127.0.0.1:3306)/securedns")

type userStruct struct{
                uuid string
                password string
                macAddress string
                token string
}

func getJsonData(input string)[]byte{
                output,err := json.Marshal(input)
                if err!=nil{
                                return []byte("Error Occured")
                } else {
                                return output
                }

}

func home(w http.ResponseWriter, r *http.Request){
                w.Write(getJsonData("This is the api for the SecureDns Project.Vist /register for registration and /login for login."))
}

func verifyJwtToken(w http.ResponseWriter, r *http.Request)(string,string, error){
                bearToken := r.Header.Get("Authorization")
                tokenString := strings.Split(bearToken, " ")[1]

                token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                                                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
                                }
                                return []byte("10d61631988b57dc19e9c083d76ff45110d61631988b57dc19e9c083d76ff451"), nil
                })

                if _,ok := token.Claims.(jwt.Claims); (!ok || !token.Valid) && err!=nil{
                                return "","",err
                }


                claims, ok := token.Claims.(jwt.MapClaims)
                if ok && token.Valid{
                                uuid, ok := claims["uuid"].(string)
                                if !ok {
                                                return "","",err
                                }
                                has_mentor, ok := claims["has_mentor"].(string)
                                if !ok{
                                                return "","",err
                                }
                                return uuid,has_mentor,nil
                } else {
                                return "","",fmt.Errorf("Something bad happened with token")
                }

}

func userInfo (w http.ResponseWriter, r *http.Request){
                w.Header().Set("Access-Control-Allow-Origin",r.Header.Get("Origin"))
                w.Header().Set("Access-Control-Allow-Headers","authorization")

                uuid,has_mentor,err := verifyJwtToken(w,r)
                if err!=nil{
                                w.WriteHeader(401)
                                w.Write(getJsonData("Invalid Jwt Token. Please login again"))
                                return
                }

                if has_mentor=="false"{

                                type storage struct {
                                Uuid string `json:"uuid"`
                                Email string `json:"email"`
                                Created_at string `json:"created_at"`
                                Mentee []string `json:"mentee"`
                                Mac string `json:"mac"`
                }

                                details := new(storage)
                                query := "select uuid, men_email, created_at, mac  from mentor where uuid='" + uuid + "';"
                                getDetails,_:= db.Query(query)
                                log.Println(query)
                                defer getDetails.Close()
                                for getDetails.Next(){
                                                getDetails.Scan(&details.Uuid,&details.Email,&details.Created_at,&details.Mac)
                                }

                                query = "select email from mentee where men_uuid='"+ uuid +"';"
                                getMentee,_ := db.Query(query)
                                log.Println(query)
                                defer getMentee.Close()

                                for getMentee.Next(){
                                                var mentee string
                                                getMentee.Scan(&mentee)
                                                details.Mentee = append(details.Mentee, mentee)
                                }

                                tmp,_:= json.Marshal(details)
                                w.WriteHeader(200)
                                w.Write(tmp)
                                return
                }

}

func getLogs(w http.ResponseWriter, r *http.Request){
                w.Header().Set("Access-Control-Allow-Origin",r.Header.Get("Origin"))
                w.Header().Set("Access-Control-Allow-Headers","authorization")

                f, err := os.OpenFile("logs/analytics.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

                if err != nil {
                log.Fatalf("error opening file: %v", err)
                }
                defer f.Close()

                log.SetOutput(f)


                reqBody, err := ioutil.ReadAll(r.Body)
                if err!=nil{
                                panic(err)
                }
                log.Println(string(reqBody))

}


func createMentee(w http.ResponseWriter, r *http.Request){
                r.ParseForm()
                w.Header().Set("Access-Control-Allow-Origin",r.Header.Get("Origin"))
                w.Header().Set("Access-Control-Allow-Headers","authorization")
                if r.Method == "OPTIONS"{
                                w.Write([]byte("Ok"))
                                return
                }
                uuid,has_mentor,err := verifyJwtToken(w,r)
                if err!=nil{
                                w.WriteHeader(401)
                                w.Write(getJsonData("Invalid Jwt Token. Please login again"))
                                return
                }

                user := new(userStruct)
                if has_mentor == "false" && uuid!="" && len(r.PostForm)>4{
                                //check if the mentee exists already

                                query := "select uuid from mentee where men_uuid='"+uuid+"' and email='"+ r.PostForm["email"][0]+"';"
                                tmp := db.QueryRow(query)
                                log.Println(query)

                                tmp.Scan(&user.uuid)
                                if user.uuid==""{
                                                // Now that we confirmed user doesnt exists. Lets Register.
                                                // hash the password.
                                                hash := getPasswordHash(r.PostForm["password"][0])


                                                // check if using whitelist or blacklist

                                                if _,ok := r.PostForm["blacklist"]; ok{
                                                                log.Println(query)
                                                                query = "insert into mentee(uuid,men_uuid,email,name,password,age_grp,whitelist) values(uuid(),'"+ uuid +"','"+ r.PostForm["email"][0] +"','"+ r.PostForm["name"][0] +"','"+ hash +"','"+ r.PostForm["age_grp"][0] +"','"+ r.PostForm["blacklist"][0]+"')"
                                                } else {
                                                                query = "insert into mentee(uuid,men_uuid,email,name,password,age_grp,whitelist) values(uuid(),'"+ uuid +"','"+ r.PostForm["email"][0] +"','"+ r.PostForm["name"][0] +"','"+ hash +"','"+ r.PostForm["age_grp"][0] +"','"+ r.PostForm["whitelist"][0]+"')"
                                                                log.Println(query)
                                                }

                                                db.QueryRow(query)
                                                log.Println(query)

                                                w.WriteHeader(200)
                                                w.Write(getJsonData("Mentee registered successfully"))
                                                return
                                } else {
                                                w.WriteHeader(500)
                                                w.Write(getJsonData("Mentee already exists"))
                                                return
                                }

                }
}

func register(w http.ResponseWriter, r *http.Request){
                r.ParseForm()
                w.Header().Set("Access-Control-Allow-Origin",r.Header.Get("Origin"))

                if len(r.PostForm) < 2{
                                w.Write(getJsonData("Not Enough Values"))
                                return
                }

                user := new(userStruct)

                query := "select uuid from mentor where men_email='"+ r.PostForm["email"][0] + "';"
                existanceCheck := db.QueryRow(query)
                log.Print(query)
                existanceCheck.Scan(&user.uuid)


                if user.uuid == ""{
                                hash := getPasswordHash(r.PostForm["password"][0])

                                query = "insert into mentor (uuid,men_email,password,created_at) Values (uuid(),'" + r.PostForm["email"][0] + "','" + hash + "',now());"
                                db.QueryRow(query)
                                log.Println(query)

                                w.WriteHeader(200)
                                w.Write(getJsonData("User Registration Success"))
                } else {
                                w.WriteHeader(500)
                                w.Write(getJsonData("User Already in database; login at /login"))
                }
}

func login(w http.ResponseWriter, r *http.Request){
                r.ParseForm()
		log.Println(r.PostForm)
                w.Header().Set("Access-Control-Allow-Origin",r.Header.Get("Origin"))

                if len(r.PostForm) < 2{
                                w.Write(getJsonData("Not Enough Values"))
                                return
                }

                if r.PostForm["has_mentor"][0] =="false"{

                                user := new(userStruct)

                                query := "select uuid,password from mentor where men_email='" + r.PostForm["email"][0] + "';"
                                log.Println(query)
                                checkPassword,_ := db.Query(query)
                                log.Println(query)
                                defer checkPassword.Close()

                                for checkPassword.Next(){
                                                checkPassword.Scan(&user.uuid,&user.password)
                                }

                                if user.uuid == ""{
                                                w.WriteHeader(401)
                                                w.Write(getJsonData("Invalid Username or Password"))
                                                return
                                }

                                err=bcrypt.CompareHashAndPassword([]byte(user.password),[]byte(r.PostForm["password"][0]))

                                if err==nil{
                                                if _,ok := r.PostForm["macaddress"]; ok{
                                                        handleMacAddress(r.PostForm["macaddress"][0],r.PostForm["email"][0],1)
                                                }
                                                w.WriteHeader(200)
                                                jwt := createJWTToken(user.uuid,r.PostForm["has_mentor"][0])
                                                w.Write(getJsonData(jwt))
                                } else {
                                                w.WriteHeader(401)
                                                w.Write(getJsonData("Invalid Username or Password"))
                                                }
                } else {
                                user := new(userStruct)

                                //select m.uuid, m.password, men.uuid from mentee m, mentor men where men_email='xedatex702325@yncyjs.com' and m.email='mentee1@test.com';
                                query:= "select m.uuid, m.password from mentee m , mentor men where men_email='"+ r.PostForm["has_mentor"][0] +"' and m.email='"+ r.PostForm["email"][0]+"';"
                                checkPassword,_ := db.Query(query)
                                log.Println(query)
                                defer checkPassword.Close()

                                for checkPassword.Next(){
                                                checkPassword.Scan(&user.uuid, &user.password)
                                }

                                if user.uuid ==""{
                                                w.WriteHeader(401)
                                                w.Write(getJsonData("Invalid Username or Password"))
                                                return
                                }

                                err = bcrypt.CompareHashAndPassword([]byte(user.password),[]byte(r.PostForm["password"][0]))

                                if err!=nil{
                                } else {

                                                if _,ok := r.PostForm["macaddress"]; ok{
                                                        handleMacAddress(r.PostForm["macaddress"][0],r.PostForm["email"][0],0)
                                                }
                                                w.WriteHeader(200)
                                                jwt := createJWTToken(user.uuid,r.PostForm["has_mentor"][0])
                                                w.Write(getJsonData(jwt))
                                }

                }
}

func handleMacAddress(mac string, email string, is_mentor int){
// TODO Handle for multiple mac per user
                                if is_mentor == 1 {
                                //query := "select mac from mentor where men_email='" + email +"' and mac='" + mac  +"';"
                                //response := db.QueryRow(query)
                                //log.Println(query)
                                //response.Scan(&tmp)
                                //if tmp==""{
				query := "insert into mentor(mac) value('"+ mac  +"');"
                                db.Query(query)
                                log.Println(query)
                                //}
                } else {
                                //query := "select mac from mentee where email='" + email +"' and mac='" + mac  +"';"
                                //response := db.QueryRow(query)
                                //log.Println(query)
                                //response.Scan(&tmp)
                                //if tmp==""{
				query := "insert into mentee(mac) value('"+ mac  +"');"
                                db.Query(query)
                                log.Println(query)
                                //}
                }
}
func createJWTToken(uuid, has_mentor string) string{
                secret := []byte("10d61631988b57dc19e9c083d76ff45110d61631988b57dc19e9c083d76ff451")
                token := jwt.New(jwt.GetSigningMethod("HS256"))
                token.Claims = jwt.MapClaims{
                                "uuid":uuid,
                                "has_mentor": has_mentor,
                                "exp":time.Now().Add(time.Hour*12).Unix(),
                }
                tokenString, err := token.SignedString(secret)
                if err!=nil{
                                log.Println(err)
                                return ""
                }
                return tokenString
}


func getPasswordHash(password string) string{
                password = password
                bcryptHashBytes,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
                if err!=nil{
                                log.Println(err)
                }
                return string(bcryptHashBytes)
}

func main(){
                if err!=nil{
                                log.Println("Error Connecting to Database")
                }
                http.HandleFunc("/",home)
                http.HandleFunc("/register",register)
                http.HandleFunc("/login",login)
                http.HandleFunc("/createMentee",createMentee)
                http.HandleFunc("/userinfo",userInfo)
                http.HandleFunc("/log",getLogs)

                log.Println("Starting Listening on Port 8080...")
                err = http.ListenAndServe(":8080",nil)
                if err!=nil{
                                log.Fatal(err)
                }
}
