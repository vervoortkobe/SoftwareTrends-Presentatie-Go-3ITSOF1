async function loadUserData() {
  try {
    const response = await fetch("/api/user", {
      credentials: "include",
    });
    if (response.ok) {
      const data = await response.json();
      document.getElementById("usernameDisplay").textContent = data.username;
    } else {
      window.location.href = "/login";
    }
  } catch (error) {
    console.error("Error loading user data:", error);
    window.location.href = "/login";
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
                                <button class="btn btn-sm btn-outline-primary me-2" onclick="openEditModal(${
                                  post.id
                                }, '${post.title.replace(
                                /'/g,
                                "\\'"
                              )}', '${post.content.replace(/'/g, "\\'")}')">
                                    Edit
                                </button>
                                <button class="btn btn-sm btn-outline-danger" onclick="deletePost(${
                                  post.id
                                })">
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
  loadUserData();
  loadPosts();
});
