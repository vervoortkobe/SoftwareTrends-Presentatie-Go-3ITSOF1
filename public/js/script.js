async function loadUserData() {
  try {
    const response = await fetch("/api/user", {
      credentials: "include",
    });
    if (response.ok) {
      const data = await response.json();
      document.getElementById("usernameDisplay").textContent = data.username;
    } else {
      console.error("Failed to load user data");
      document.getElementById("usernameDisplay").textContent = "User";
    }
  } catch (error) {
    console.error("Error loading user data:", error);
    document.getElementById("usernameDisplay").textContent = "User";
  }
}

async function loadPosts() {
  try {
    const response = await fetch("/api/posts", {
      credentials: "include",
    });
    if (response.ok) {
      const data = await response.json();
      const postList = document.getElementById("postList");

      if (data.posts && data.posts.length > 0) {
        postList.innerHTML = data.posts
          .map(
            (post) => `
                            <div class="list-group-item">
                                <h3 class="h5 mb-2">${post.title}</h3>
                                <p class="mb-2">${post.content}</p>
                                <small class="text-muted">Gepost door: ${
                                  post.username
                                }</small>
                                ${
                                  post.isOwner
                                    ? `
                                    <div class="mt-2">
                                        <button class="btn btn-sm btn-outline-danger" onclick="deletePost(${post.id})">
                                            Delete
                                        </button>
                                    </div>
                                `
                                    : ""
                                }
                            </div>
                        `
          )
          .join("");
      } else {
        postList.innerHTML =
          '<div class="list-group-item">Geen Posts beschikbaar.</div>';
      }
    } else {
      console.error("Failed to load posts");
      document.getElementById("postList").innerHTML =
        '<div class="list-group-item text-danger">Failed to load posts.</div>';
    }
  } catch (error) {
    console.error("Error loading posts:", error);
    document.getElementById("postList").innerHTML =
      '<div class="list-group-item text-danger">Error loading posts.</div>';
  }
}

document.addEventListener("DOMContentLoaded", () => {
  document
    .getElementById("postForm")
    .addEventListener("submit", async function (event) {
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
          // Clear form
          document.getElementById("title").value = "";
          document.getElementById("content").value = "";
          // Reload posts
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
});

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

function openEditModal(postId, title, content) {
  document.getElementById("editPostId").value = postId;
  document.getElementById("editTitle").value = title;
  document.getElementById("editContent").value = content;
  const editModal = new bootstrap.Modal(document.getElementById("editModal"));
  editModal.show();
}

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

document.addEventListener("DOMContentLoaded", () => {
  loadUserData();
  loadPosts();
});
