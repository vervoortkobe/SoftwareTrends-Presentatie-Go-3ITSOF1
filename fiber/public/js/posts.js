document.addEventListener("DOMContentLoaded", () => {
  const postForm = document.getElementById("postForm");
  if (postForm) {
    postForm.addEventListener("submit", async function (event) {
      event.preventDefault();

      const title = document.getElementById("title").value;
      const content = document.getElementById("content").value;

      try {
        const response = await fetch("/api/posts", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            title,
            content,
          }),
          credentials: "include",
        });

        if (response.ok) {
          document.getElementById("title").value = "";
          document.getElementById("content").value = "";
          loadPosts();
        } else {
          const data = await response.json();
          alert(data.error || "Failed to create post");
        }
      } catch (error) {
        console.error("Error creating post:", error);
        alert("An error occurred while creating the post");
      }
    });
  }
});

async function handleEditSubmit() {
  const postId = document.getElementById("editPostId").value;
  const title = document.getElementById("editTitle").value;
  const content = document.getElementById("editContent").value;

  try {
    const response = await fetch(`/api/posts/${postId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title,
        content,
      }),
      credentials: "include",
    });

    if (response.ok) {
      const editModal = bootstrap.Modal.getInstance(
        document.getElementById("editModal")
      );
      editModal.hide();
      loadPosts();
    } else {
      const data = await response.json();
      alert(data.error || "Failed to update post");
    }
  } catch (error) {
    console.error("Error updating post:", error);
    alert("An error occurred while updating the post");
  }
}

async function deletePost(postId) {
  if (!confirm("Are you sure you want to delete this post?")) {
    return;
  }

  try {
    const response = await fetch(`/api/posts/${postId}`, {
      method: "DELETE",
      credentials: "include",
    });

    if (response.ok) {
      loadPosts();
    } else {
      const data = await response.json();
      alert(data.error || "Failed to delete post");
    }
  } catch (error) {
    console.error("Error deleting post:", error);
    alert("An error occurred while deleting the post");
  }
}
