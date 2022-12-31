package services

import (
	"encoding/json"
	"fiff_golang_draft/database"
	"fiff_golang_draft/helper"
	"fiff_golang_draft/models"
	"fiff_golang_draft/module_socket"
	"fmt"
	"strconv"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/joho/godotenv"

	//"gopkg.in/square/go-jose.v2/jwt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func HandleServer() {
	//getClaims()
	r := gin.Default()
	// r.SetTrustedProxies([]string{"127.0.0.1"})
	// r.SetTrustedProxies([]string{"0.0.0.0"})
	// r.SetTrustedProxies([]string{"192.168.1.14"})
	// r.SetTrustedProxies([]string{"::1"})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	issuerURL, _ := url.Parse(os.Getenv("AUTH0_DOMAIN"))
	audience := os.Getenv("AUTH0_AUD")

	provider := jwks.NewCachingProvider(issuerURL, time.Duration(5*time.Minute))

	jwtValidator, _ := validator.New(provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)

	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)
	print("jwtMiddleware", jwtMiddleware)
	//r.Use(adapter.Wrap(jwtMiddleware.CheckJWT))

	var wrapCheckJWT = adapter.Wrap(jwtMiddleware.CheckJWT)

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000",
			"https://private-exchange-frontend-nextjs.vercel.app",
			"https://stellar-heliotrope-f5dece.netlify.app",
			"http://orange-exchange.xyz",
			"http://www.orange-exchange.xyz",
			"https://orange-exchange.xyz",
			"https://www.orange-exchange.xyz",
		},
		// AllowMethods:     []string{"PUT", "PATCH", "POST", "OPTIONS", "GET", "DELETE"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Headers", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "intro-ver" + strconv.Itoa(GetCurrentVersion()),
		})
	})

	r.GET("/config-page", func(c *gin.Context) {

		object1 := map[string]interface{}{
			"ver": GetCurrentVersion(),
			// "env1":               GetEnvTest(),
			"connString":         helper.GetConnStringObfuscate(),
			"testingMode":        helper.TestingMode,
			"pgdatabase_railway": helper.GetEnvRailway(),
			"is_prod_mode":       helper.IsProdMode(),
		}
		helper.GetConnStringObfuscate()
		// c.JSON(http.StatusOK, gin.H{object1})
		c.JSON(http.StatusOK, object1)
		// c.JSON(http.StatusOK, gin.H{
		// 	"hello": "intro-ver" + strconv.Itoa(GetCurrentVersion()),
		// })
	})

	r.GET("/ws", func(c *gin.Context) {
		module_socket.RegisterWebSocket(c)
		//wshandler(c.Writer, c.Request)
	})

	r.GET("/broadcastMessage", func(c *gin.Context) {
		module_socket.Broadcast("hello world")
		c.JSON(http.StatusOK, gin.H{
			"item": "sent",
		})
	})
	r.GET("/broadcastMessageWithInput/:id", func(c *gin.Context) {
		var message = c.Param("id")
		module_socket.Broadcast(message)
		c.JSON(http.StatusOK, gin.H{
			"item": "sent",
		})
	})

	r.GET("/hello", wrapCheckJWT, func(c *gin.Context) {
		var token = getToken(c)
		var subject = getSubject(token)
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.JSON(http.StatusOK, gin.H{
			"hello": "world - " + subject,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/order/getList", func(c *gin.Context) {
		var items = getOrderList("btcusd", 5)
		c.JSON(http.StatusOK, gin.H{
			"data": items,
		})
	})

	r.GET("/order/getListLong", func(c *gin.Context) {
		var items = getOrderList("btcusd", 10)
		c.JSON(http.StatusOK, gin.H{
			"data": items,
		})
	})

	r.POST("/order_submit", func(c *gin.Context) {

		var json models.OrderLite
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// var input = models.InputOrder{
		// 	RequestRate:  9999,
		// 	CurrencyPair: "btcusd",
		// 	//Uuid:         uuid,
		// 	Amount: 1,
		// 	IsBuy:  true,
		// }

		var input = models.InputOrder{
			RequestRate:  json.RequestRate,
			CurrencyPair: "btcusd",
			Amount:       json.Amount,
			IsBuy:        json.IsBuy,
		}
		fmt.Println("create order buy")
		fmt.Println(input)
		fmt.Println(json.RequestRate)
		fmt.Println(json.Amount)
		fmt.Println(json.IsBuy)
		ExecuteTransaction(input)

		output := fmt.Sprintf("amount: %.2f, request rate: %.2f, is buy: %t",
			json.Amount, json.RequestRate, json.IsBuy)
		c.JSON(http.StatusOK, gin.H{"status": output})
	})

	r.POST("/order_submit2", wrapCheckJWT, func(c *gin.Context) {

		var token = getToken(c)
		var subject = getSubject(token)

		var json models.OrderLite
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var input = models.InputOrder{
			RequestRate:  json.RequestRate,
			CurrencyPair: "btcusd",
			Amount:       json.Amount,
			IsBuy:        json.IsBuy,
			UserName:     subject,
		}
		InitParamDb()
		ExecuteTransaction(input)

		output := fmt.Sprintf("amount: %.2f, request rate: %.2f, is buy: %t",
			json.Amount, json.RequestRate, json.IsBuy)
		c.JSON(http.StatusOK, gin.H{"status": output})
	})

	{
		r.GET("/student/list", func(c *gin.Context) {
			var students = StudentGetList()
			str, err := json.Marshal(students)
			fmt.Println("get list")
			fmt.Println(string(str))
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": students})
		})

		r.GET("/student/getTotal", func(c *gin.Context) {
			var total = StudentGetTotal()
			c.JSON(http.StatusOK, gin.H{"total": total})
		})

		r.GET("/student/createOne", func(c *gin.Context) {
			StudentCreate()

			c.JSON(http.StatusOK, gin.H{
				"message": "create-item",
			})
		})
		r.PUT("/student", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "put",
			})
		})
		r.DELETE("/student", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete",
			})
		})
		r.GET("/student/:id", func(c *gin.Context) {
			var id = c.Param("id")
			c.JSON(http.StatusOK, gin.H{
				"message": "inquiry-" + id,
			})
		})

		r.POST("/student", func(c *gin.Context) {

			var json models.Login
			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if json.User != "manu" || json.Password != "123" {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		})
	}
	{
		r.GET("/order/initSample", func(c *gin.Context) {
			InitSample()
			c.JSON(http.StatusOK, gin.H{
				"data": "done",
			})
		})

		r.GET("/debug/db", func(c *gin.Context) {
			var output = database.DebugWebApi()
			c.JSON(http.StatusOK, gin.H{
				"message": output,
			})
		})
		r.GET("/debug/conn_string", func(c *gin.Context) {
			var connString = helper.GetConnString()
			c.JSON(http.StatusOK, gin.H{
				"message": connString,
			})
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.RunTLS(":8080", "./certs/localhost.crt", "./certs/localhost.key")
	// r.RunTLS(":8080", "./certs/localhost.pem", "./certs/localhost.key")
	// r.RunTLS(":8080", "./certs/cert.pem", "./certs/key.pem")
	// r.RunTLS(":8080", "./certs/server.crt", "./certs/server.key")

	//var w = autotls.Run(r, "orange-exchange.xyz")
	//log.Fatal(w)

}

func getToken(c *gin.Context) string {
	var d = c.Request.Header.Get("Authorization")
	var e = strings.Replace(d, "Bearer ", "", 1)
	return e
}

func getSubject(tokenString string) string {
	//secret := os.Getenv("AUTH0_CLIENT_SECRET")
	secret := "no_need"
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	fmt.Printf("", token)
	fmt.Printf("", err)

	//jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(secret), nil
	//})

	var p = claims["sub"]
	str := fmt.Sprintf("%v", p)
	return str

}
