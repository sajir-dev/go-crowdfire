package domain

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/lib/pq"

	"github.com/joho/godotenv"
	followcontract "github.com/sajir-dev/go-crowdfire/services/follow/contract"
	postscontract "github.com/sajir-dev/go-crowdfire/services/posts/contract"
	usercontract "github.com/sajir-dev/go-crowdfire/services/users/contract"
)

var dbInstance *sql.DB

// InitDB intialises the DB
func InitDB() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	// getting env variables
	dbHost := os.Getenv("DB_HOST")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbInstance, err = sql.Open("postgres", "postgres://"+dbUserName+":"+dbPassword+"@"+dbHost+"/"+dbName+"?sslmode=disable")
	if err != nil {
		return err
	}

	if err := dbInstance.Ping(); err != nil {
		return err
	}

	return nil
}

func CreateUser(req *usercontract.CreateUser) error {
	insertStmt := `INSERT INTO "users" ("name", "email", "password")VALUES ($1, $2, $3)`
	_, err := dbInstance.Exec(insertStmt, req.Name, req.Email, req.Password)
	return err
}

func Login(req *usercontract.LoginReq) (*usercontract.UserModel, error) {
	result := &usercontract.UserModel{}
	row := dbInstance.QueryRow(`SELECT * FROM "users" WHERE "email" = $1`, req.Email)
	if err := row.Scan(&result.Id, &result.Name, &result.Email, &result.Password); err != nil {
		return nil, err
	}

	if result.Email == req.Email && result.Password == req.Password {
		return result, nil
	}

	return nil, errors.New("invalid credentials")
}

func GetFollowers(req *followcontract.GetFollowersReq) (*followcontract.Following, error) {
	following := []string{}
	stmt := `SELECT "following" FROM "following" WHERE "userid" = $1::varchar`
	row := dbInstance.QueryRow(stmt, req.Userid)
	err := row.Scan(pq.Array(&following))
	if err != nil {
		return nil, err
	}
	return &followcontract.Following{
		Userids: following,
	}, nil
}

func Follow(req *followcontract.FollowReq) error {
	following := []string{}
	stmt := `SELECT "following" FROM "following" WHERE "userid" = $1::varchar`
	row := dbInstance.QueryRow(stmt, req.Userid)
	err := row.Scan(pq.Array(&following))
	if err != nil && err.Error() == "sql: no rows in result set" {
		following = []string{req.Following}
		stmt = `INSERT INTO "following" ("userid","following")VALUES ($1, $2)`
		_, err = dbInstance.Exec(stmt, req.Userid, pq.Array(following))
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	for _, id := range following {
		if id == req.Following {
			return errors.New("already following")
		}
	}

	following = append(following, req.Following)

	stmt = `UPDATE "following" SET "following"=$1 WHERE "userid" = $2::varchar`
	_, err = dbInstance.Exec(stmt, pq.Array(following), req.Userid)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePost(req *postscontract.UpdatePostReq) (*postscontract.PostModel, error) {
	var createdAt string
	var id string
	var createdBy string
	var content string

	stmt := `SELECT * from "posts" where id=$1`
	row := dbInstance.QueryRow(stmt, req.Id)
	if err := row.Scan(&createdAt, &id, &createdBy, &content); err != nil {
		return nil, err
	}

	if createdBy != req.Userid {
		return nil, errors.New("wrong post requested")
	}

	stmt = `UPDATE "posts" SET "content" = $1 WHERE id=$2`
	_, err := dbInstance.Exec(stmt, req.Content, req.Id)
	if err != nil {
		return nil, err
	}

	stmt = `SELECT * from "posts" where id=$1`
	row = dbInstance.QueryRow(stmt, req.Id)
	if err = row.Scan(&createdAt, &id, &createdBy, &content); err != nil {
		return nil, err
	}

	res := &postscontract.PostModel{
		Id:        id,
		CreatedBy: createdBy,
		Content:   content,
	}

	layout := "2006-01-02 15:04:05 -0700 MST"
	res.CreatedAt, _ = time.Parse(layout, createdAt)
	return res, nil
}

func GetPosts(req *postscontract.GetPostsReq) ([]postscontract.PostModel, error) {
	posts := []postscontract.PostModel{}
	stmt := `SELECT * from "posts" where "created_by" = $1`
	rows, err := dbInstance.Query(stmt, req.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var createdAt string
		var id string
		var createdBy string
		var content string
		if err = rows.Scan(&createdAt, &id, &createdBy, &content); err != nil {
			return nil, err
		}

		post := postscontract.PostModel{
			Id:        id,
			CreatedBy: createdBy,
			Content:   content,
		}
		layout := "2006-01-02 15:04:05 -0700 MST"
		post.CreatedAt, _ = time.Parse(layout, createdAt)
		posts = append(posts, post)
	}
	return posts, nil
}
