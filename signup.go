package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

//type PERSON struct {
//	Id    string
//	Name  string
//}

//func signupPage(res http.ResponseWriter, req *http.Request) {
func signupPage(c echo.Context) error {
	//if req.Method != "POST" {
	//	http.ServeFile(res, req, "signup.html")
	//	return
	//}
	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println(c.Path)
	fmt.Println("\nHere is signup\n\n")
	username := c.FormValue("username")
	password := c.FormValue("password")
	fmt.Println(username)
	fmt.Println(password, "\n\n")
	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Now add new user ^^ ")
		_, err = db.Exec("INSERT INTO users(username, password, logIn) VALUES(?, ?, ?)", username, password, 1)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"msg": "Servererror",
			})
		}
		rows, err := db.Query("SELECT id,logIn FROM users WHERE username=?", username)
		if err != nil {
			log.Panic(err)
		}
		var id int
		var logIn int
		for rows.Next() {
			rows.Scan(&id, &logIn)
		}
		fmt.Println("id = ", id)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg":     "success",
			"user_id": id,
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "Servererror",
		})
	default:
		return c.JSON(http.StatusMovedPermanently, map[string]interface{}{
			"msg": "redirect",
		})
		//return c.String(http.StatusMovedPermanently, "redirect, unable to create your account.")
		//http.Redirect(res, req, "/", 301)
	}
}
func loginPage(c echo.Context) error {

	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(c.Path)
	fmt.Println("\nHere is signup\n\n")
	username := c.FormValue("username")
	password := c.FormValue("password")
	fmt.Println(username)
	fmt.Println(password, "\n\n")

	rows, _ := db.Query("SELECT id,logIn FROM users WHERE username=?", username)
	var logIn int
	var id int
	for rows.Next() {
		rows.Scan(&id, &logIn)
		fmt.Println(id, logIn)
		//tweets = append(tweets, content)
	}
	logIn = 1
	_, err = db.Exec("UPDATE users SET logIn=? WHERE id=?", logIn, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "Servererror",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":     "success",
		"user_id": id,
	})
}

///////////////////////////////////////////////////////////////////////////////////////////

func insertTweet(c echo.Context) error {

	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//////////////////////////////////////// id of user

	fmt.Println("\nHere is new tweet\n\n")
	defer c.Request().Body.Close()
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idUser, _ := strconv.Atoi(id) //convert to int
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	////////////////////////////////////// content of tweet
	text := c.FormValue("text")
	file := c.FormValue("file")
	///////////////////////////////////////// inser content to data base
	fmt.Println("text - ", text, file)
	tnow := time.Now().Format("2006-01-02 15:04:05")
	res, _ := db.Exec("INSERT INTO tweets(user_id, content, tweet_date) VALUES(?, ?, ?)", idUser, text, tnow)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "Servererror",
		})
	}
	tweetID, _ := res.LastInsertId()
	fmt.Println("tweetId = ", tweetID)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":      "success",
		"tweet_id": tweetID,
	})
}

func deleteTweet(c echo.Context) error {

	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//////////////////////////////////////// id of user

	fmt.Println("\nHere is new tweet\n\n")
	defer c.Request().Body.Close()
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idt := c.Param("tweet")
	idUser, _ := strconv.Atoi(id)   //convert to int
	idTweet, _ := strconv.Atoi(idt) //convert to int
	fmt.Println("idsss ", idTweet, idUser)
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	///////////////////////////////////////// inser content to data base

	_, err = db.Exec("DELETE FROM tweets WHERE tweet_id=?", idTweet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "Servererror",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "success",
	})
}

func getSignupPage(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "success",
	})
}

func logout(c echo.Context) error {
	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//////////////////////////////////////// id of user
	fmt.Println("logout!!")
	//defer c.Request().Body.Close()
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idUser, _ := strconv.Atoi(id) //convert to int
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	//////////////////////////////////////
	var logIn int
	err := db.QueryRow("SELECT logIn FROM users WHERE id=?", idUser).Scan(&logIn)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", idUser)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		if logIn == 1 {
			logOut := 0 //UPDATE users SET logIn = 0 WHERE id = 6;
			_, err = db.Exec("UPDATE users SET logIn=? WHERE id=?", logOut, idUser)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"msg": "Servererror",
				})
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"msg":     "success",
				"user_id": idUser,
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "Servererror",
		})
	}
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"msg": "Servererror",
	})

}

func seeOtherPage(c echo.Context) error {
	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//////////////////////////////////
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idUser, _ := strconv.Atoi(id) //convert to int
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	//////////////////////////////
	fmt.Println("home page id_user = ", idUser)
	rows, err := db.Query("SELECT * FROM tweets WHERE user_id=? ORDER BY tweet_date DESC LIMIT 10", idUser)
	if err != nil {
		log.Panic(err)
	}

	var user_id int
	var tweet_id int
	var tweet_date string
	var content string
	var tweets []string
	var likes int
	var retweet string
	for rows.Next() {
		rows.Scan(&user_id, &content, &tweet_date, &tweet_id, &likes, &retweet)
		fmt.Println(user_id, content, tweet_date, tweet_id, likes, retweet)
		tweets = append(tweets, content)
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"msg":    "success",
		"tweets": tweets,
	})
}
func likeTweet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//////////////////////////////////////// id of user

	fmt.Println("\nHere is like tweet\n\n")
	defer c.Request().Body.Close()
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idTweet, _ := strconv.Atoi(id) //convert to int
	if idTweet == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	var likes int
	err := db.QueryRow("SELECT likes FROM tweets WHERE tweet_id=?", idTweet).Scan(&likes)
	fmt.Println("likes= ", likes, " for tweet:", idTweet)
	switch {
	case err == sql.ErrNoRows:
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "Servererror",
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "Servererror",
		})
	default:
		likes = likes + 1 //UPDATE users SET logIn = 0 WHERE id = 6;
		_, err = db.Exec("UPDATE tweets SET likes=? WHERE tweet_id=?", likes, idTweet)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"msg": "Servererror",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg":      "success",
			"tweet_id": idTweet,
			"likes":    likes,
		})
	}
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"msg": "Servererror",
	})
}
func follow(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idUser, _ := strconv.Atoi(id) //convert to int
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	id = c.Param("idfollow")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idfollow, _ := strconv.Atoi(id) //convert to int
	if idfollow == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	///////////////////////////////
	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("***", idUser, "  ", idfollow)
	rows, err := db.Query("SELECT * FROM followings WHERE id=? AND following=?", idUser, idfollow)
	fmt.Println("err-> ", err)
	var idd int
	var fodllowing int
	for rows.Next() {
		rows.Scan(&idd, &fodllowing)
		fmt.Println(idd, fodllowing)
	}
	if idd == 0 {
		fmt.Println("Now add new follower ^^ ")
		_, err = db.Exec("INSERT INTO followings(id, following) VALUES(?, ?)", idUser, idfollow)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"msg": "error",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "success",
		})
	}
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"msg": "success",
	})

}

func unfollow(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idUser, _ := strconv.Atoi(id) //convert to int
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	id = c.Param("idfollow")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idfollow, _ := strconv.Atoi(id) //convert to int
	if idfollow == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	///////////////////////////////
	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("***", idUser, "  ", idfollow)
	rows, err := db.Query("SELECT * FROM followings WHERE id=? AND following=?", idUser, idfollow)
	fmt.Println("err-> ", err)
	var idd int
	var fodllowing int
	for rows.Next() {
		rows.Scan(&idd, &fodllowing)
		fmt.Println(idd, fodllowing)
	}
	if idd != 0 {
		fmt.Println("Now add new follower ^^ ") //_, err = db.Exec("DELETE FROM tweets WHERE tweet_id=?", idTweet)
		_, err = db.Exec("DELETE FROM followings WHERE id=? AND following=?", idUser, idfollow)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"msg": "error",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "success",
		})
	}
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"msg": "error",
	})

}

func showTwitterFollowers(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idUser, _ := strconv.Atoi(id) //convert to int
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, _ := db.Query("SELECT following FROM followings WHERE id=?", idUser)
	//var a []int
	var a string
	var b string
	a = "("
	var fodllowing int
	for rows.Next() {
		rows.Scan(&fodllowing)
		fmt.Println(fodllowing)
		b = strconv.Itoa(fodllowing) + ","
		a = a + b
	}
	a = a + ")"
	a = strings.Replace(a, ",)", ")", -1)
	rows, err := db.Query("SELECT * FROM tweets WHERE user_id IN " + a + " ORDER BY tweet_date DESC") //db.Query(stmt, 4)
	if err != nil {
		log.Panic(err)
	}

	var user_id int
	var tweet_id int
	var tweet_date string
	var content string
	var tweets []string
	var tsends []int
	var likess []int
	var tids []int
	var likes int
	var retweet string
	for rows.Next() {
		rows.Scan(&user_id, &content, &tweet_date, &tweet_id, &likes, &retweet)
		fmt.Println(user_id, content, tweet_date, tweet_id, likes, retweet)
		tweets = append(tweets, content)
		tsends = append(tsends, user_id)
		tids = append(tids, tweet_id)
		likess = append(likess, likes)

	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"msg":      "success",
		"tweets":   tweets,
		"senders":  tsends,
		"likes":    likess,
		"tweet_id": tids,
	})

}
func editProfile(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idUser, _ := strconv.Atoi(id) //convert to int
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	newName := c.FormValue("newName")
	newPic := c.FormValue("newPic")
	if newName != "" {
		_, err = db.Exec("UPDATE users SET username=? WHERE id=?", newName, idUser)
	}
	if newPic != "" {
		_, err = db.Exec("UPDATE users SET logIn=? WHERE id=?", newPic, idUser)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":     "success",
		"newName": newName,
	})
}
func retweet(c echo.Context) error {
	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//////////////////////////////////////// id of user

	fmt.Println("\nHere is retweet\n\n")
	defer c.Request().Body.Close()
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	idt := c.Param("tweet")
	idUser, _ := strconv.Atoi(id)   //convert to int
	idTweet, _ := strconv.Atoi(idt) //convert to int
	fmt.Println("idsss ", idTweet, idUser)
	if idUser == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "error",
		})
	}
	////////////////////create string od UPDATE

	rows, _ := db.Query("SELECT retweet FROM tweets WHERE tweet_id=?", idTweet)
	//var a []int
	var b string
	var retweetids string
	for rows.Next() {
		rows.Scan(&retweetids)
		fmt.Println(retweetids)
	}
	b = id
	if retweetids != "" {
		b = "," + id
	}
	retweetids = retweetids + b
	fmt.Println("retweetids ", retweetids)
	///////////////////////////////////////////
	_, err = db.Exec("UPDATE tweets SET retweet=? WHERE tweet_id=?", retweetids, idTweet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "Servererror",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":     "success",
		"user_id": id,
	})
}

func main() {

	db, _ := sql.Open("sqlite3", "./twitter.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	e := echo.New()
	//e.Use(middleware.Logger()) SELECT * FROM tweets WHERE user_id=2 ORDER BY tweet_date DESC
	//e.Use(middleware.Recover())
	//var a []int
	//a = append(a, 1)

	e.Static("/", "")

	e.Use(middleware.Static(""))
	e.File("/FirstPage/", "FirstPage/twitter.html")
	e.POST("/signupPage", signupPage)
	e.GET("/signupPage", getSignupPage)
	e.POST("/loginPage", loginPage)
	e.PUT("/homepage/tweet/:id", insertTweet)
	e.DELETE("/homepage/tweet/:id/:tweet", deleteTweet)
	e.GET("/homepage/tweet/:id", likeTweet)

	e.GET("/homepage/retweet/:id/:tweet", retweet)

	e.GET("/logout/:id", logout)

	e.GET("/homepage/seeOtherPage/:id", seeOtherPage)

	e.GET("/homepage/follow/:id/:idfollow", follow)
	e.GET("/homepage/unfollow/:id/:idfollow", unfollow)

	e.GET("/homepage/:id", showTwitterFollowers)

	e.PUT("/homepage/edit/:id", editProfile)

	e.Logger.Fatal(e.Start(":8001"))

}
