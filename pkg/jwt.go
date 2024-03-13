package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type Claims struct {
	UserId int
	jwt.StandardClaims
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	TokenUuid string
	UserId    int
	Exp       float64
}

func CreateToken(userId int) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 48).Unix() //expires after 30 min
	td.TokenUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 96 * 7).Unix()
	td.RefreshUuid = fmt.Sprintf("%s++%s", td.TokenUuid, strconv.Itoa(userId))

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = strconv.Itoa(userId)
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(""))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	td.RtExpires = time.Now().Add(time.Hour * 192 * 7).Unix()
	td.RefreshUuid = fmt.Sprintf("%s++%s", td.TokenUuid, strconv.Itoa(userId))

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = strconv.Itoa(userId)
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(""))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// get the token from the request body
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {

	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("SECRETJWT"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extract(token *jwt.Token) (*AccessDetails, error) {

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok2 := claims["access_uuid"].(string)
		userId, _ := claims["user_id"].(string)
		exp := claims["exp"].(float64)
		if ok2 == false {
			return nil, errors.New("unauthorized")
		} else {
			userid, _ := strconv.Atoi(userId)
			return &AccessDetails{
				TokenUuid: accessUuid,
				UserId:    userid,
				Exp:       exp,
			}, nil
		}
	}
	return nil, errors.New("something went wrong")
}

func ExtractToken(r *http.Request) (*AccessDetails, error) {
	token, err := verifyToken(r)
	if err != nil {
		fmt.Println("extrax", err)
		return nil, err
	}
	acc, err := extract(token)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

var jwtSecret []byte

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func ExtractTokenMiddleware(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
