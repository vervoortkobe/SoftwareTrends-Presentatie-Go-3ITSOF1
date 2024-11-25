package handlers

import (
	"fiber/database"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func FetchAllPosts() ([]database.Post, error) {
	rows, err := database.DB.Query("SELECT p.id, p.title, p.content, u.username, p.user_id FROM posts p JOIN users u ON p.user_id = u.id ORDER BY p.id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []database.Post
	for rows.Next() {
		var post database.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CreatePost(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(http.StatusUnauthorized).SendString("You must be logged in to create a post")
	}

	title := c.FormValue("title")
	content := c.FormValue("content")

	_, err := database.DB.Exec("INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)", user.(map[string]interface{})["ID"], title, content)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error creating post")
	}

	log.Printf("Post created by user %s: %s\n", user.(map[string]interface{})["Username"], title)

	return c.Redirect("/")
}

func EditPost(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(http.StatusUnauthorized).SendString("You must be logged in to edit a post")
	}

	postID := c.Params("id")
	title := c.FormValue("title")
	content := c.FormValue("content")

	result, err := database.DB.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ? AND user_id = ?", title, content, postID, user.(map[string]interface{})["ID"])
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error updating post")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error checking affected rows")
	}

	if rowsAffected == 0 {
		return c.Status(http.StatusNotFound).SendString("Post not found or you do not have permission to edit this post")
	}

	log.Printf("Post edited by user %s: %s\n", user.(map[string]interface{})["Username"], title)

	return c.Redirect("/")
}

func DeletePost(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(http.StatusUnauthorized).SendString("You must be logged in to delete a post")
	}
	postID := c.Params("id")
	result, err := database.DB.Exec("DELETE FROM posts WHERE id = ? AND user_id = ?", postID, user.(map[string]interface{})["ID"])
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error deleting post")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error checking affected rows")
	}
	if rowsAffected == 0 {
		return c.Status(http.StatusNotFound).SendString("Post not found or you do not have permission to delete this post")
	}
	log.Printf("Post deleted by user %s: Post ID %s\n", user.(map[string]interface{})["Username"], postID)
	return c.SendStatus(http.StatusNoContent)
}
