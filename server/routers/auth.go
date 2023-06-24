package routers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"suraj.com/refine/data"
	"suraj.com/refine/models"
)

var jwtkey = []byte(os.Getenv("JWT_SECRET"))

func getAuthRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", base)
	r.Get("/user", getCurrentUser)
	r.Post("/login", login)
	r.Post("/register", register)
	r.Delete("/{id}", deleteUser)

	return r
}

type Claims struct {
	UserID primitive.ObjectID
	jwt.StandardClaims
}

type BaseResponse struct {
	IsAuthenticated bool `json:"isAuthenticated"`
}

type RegisterResponse struct {
	Token      string             `json:"token"`
	InsertedId primitive.ObjectID `json:"insertedId"`
}

type DeleteResponse struct {
	DeleteCount int `json:"deleteCount"`
}

type UnauthorizedResponse struct {
	Message string `json:"message"`
}

func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	authorization, ok := r.Header["Authorization"]
	if !ok {
		http.Error(w, "Missing required Authorization in request headers", http.StatusUnauthorized)
		return
	}
	subStrs := strings.Split(authorization[0], " ")
	if len(subStrs) != 2 {
		http.Error(w, "Invlid Token", http.StatusUnauthorized)
		return
	}
	tokenStr := subStrs[1]
	if tokenStr == "" {
		http.Error(w, "Invlid Token", http.StatusUnauthorized)
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return t, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	if !token.Valid {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}

	uid := claims.UserID

	w.Write([]byte(uid.Hex()))
}

func login(w http.ResponseWriter, r *http.Request) {}

func register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user models.User
	user.FromJSON(&r.Body)

	// Check for duplicate
	if duplicate, err := data.UserFindByEmail(user.Email); duplicate != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, "Email already in use", http.StatusConflict)
		}
		return
	}

	// Hash the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Create user in db
	res, err := data.UserCreate(user.Email, string(hashedPass), user.UserName, user.DisplayName, user.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Set claims for JWT
	expiry := time.Now().Add(12 * time.Hour)
	claims := &Claims{
		UserID: primitive.NewObjectID(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
		},
	}

	// Create a new JWT
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtkey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := &RegisterResponse{
		Token:      tokenStr,
		InsertedId: res.InsertedID.(primitive.ObjectID),
	}
	json.NewEncoder(w).Encode(response)
}

func base(w http.ResponseWriter, r *http.Request) {}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	res, err := data.UserDeleteById(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := &DeleteResponse{
		DeleteCount: int(res.DeletedCount),
	}
	json.NewEncoder(w).Encode(response)
}
