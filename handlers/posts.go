package handlers

import (
	"fiber/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	currentUserId := c.Locals("userId").(int)

	rows, err := database.DB.Query(`
		SELECT p.id, p.title, p.content, p.user_id, u.username 
		FROM posts p 
		JOIN users u ON p.user_id = u.id 
		ORDER BY p.id DESC
	`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch posts",
		})
	}
	defer rows.Close()

	var posts []fiber.Map
	for rows.Next() {
		var id, userId int
		var title, content, username string
		if err := rows.Scan(&id, &title, &content, &userId, &username); err != nil {
			continue
		}
		posts = append(posts, fiber.Map{
			"id":       id,
			"title":    title,
			"content":  content,
			"userId":   userId,
			"username": username,
			"isOwner":  userId == currentUserId,
		})
	}

	return c.JSON(fiber.Map{
		"posts": posts,
	})
}

func CreatePost(c *fiber.Ctx) error {
	type PostInput struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	var input PostInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	userId := c.Locals("userId").(int)

	result, err := database.DB.Exec(`
		INSERT INTO posts (user_id, title, content)
		VALUES (?, ?, ?)
	`, userId, input.Title, input.Content)

	if err != nil {
		log.Printf("Failed to create post: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create post",
		})
	}

	id, _ := result.LastInsertId()
	log.Printf("Post created: User ID %d, Title: %s", userId, input.Title)
	return c.JSON(fiber.Map{
		"success": true,
		"id":      id,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}

	type PostUpdate struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	var input PostUpdate
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	userId := c.Locals("userId").(int)

	var postUserId int
	err = database.DB.QueryRow("SELECT user_id FROM posts WHERE id = ?", postId).Scan(&postUserId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	if postUserId != userId {
		return c.Status(403).JSON(fiber.Map{
			"error": "Not authorized to edit this post",
		})
	}

	_, err = database.DB.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ? AND user_id = ?",
		input.Title, input.Content, postId, userId)
	if err != nil {
		log.Printf("Failed to update post ID %d: %v", postId, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	}

	log.Printf("Post edited: ID %d, New Title: %s", postId, input.Title)
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func DeletePost(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}

	userId := c.Locals("userId").(int)

	var postUserId int
	err = database.DB.QueryRow("SELECT user_id FROM posts WHERE id = ?", postId).Scan(&postUserId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	if postUserId != userId {
		return c.Status(403).JSON(fiber.Map{
			"error": "Not authorized to delete this post",
		})
	}

	_, err = database.DB.Exec("DELETE FROM posts WHERE id = ? AND user_id = ?", postId, userId)
	if err != nil {
		log.Printf("Failed to delete post ID %d: %v", postId, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}

	log.Printf("Post deleted: ID %d", postId)
	return c.JSON(fiber.Map{
		"success": true,
	})
}
