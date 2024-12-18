package daos

import (
	"piotrek813/goth/db"
	"piotrek813/goth/models"
)

func FindAllPost() ([]models.Post, error) {
	db := db.GetDB()
	rows, err := db.Query("SELECT id, title, body, user_id FROM Posts")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.UserId); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	if posts == nil {
		posts = []models.Post{}
	}

	if err = rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func SavePost(post *models.Post) error {
	db := db.GetDB()

	stmt, err := db.Prepare("INSERT INTO Posts(title, body, user_id) VALUES(?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Body, post.UserId)

	if err != nil {
		return err
	}

	return nil
}

func DeletePost(id int) error {
	db := db.GetDB()

	stmt, err := db.Prepare("DELETE FROM Posts WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
