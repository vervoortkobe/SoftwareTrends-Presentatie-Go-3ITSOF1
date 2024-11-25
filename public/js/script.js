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
  loadUserData();
  loadPosts();
});

function openEditModal(postId, title, content) {
  document.getElementById("editPostId").value = postId;
  document.getElementById("editTitle").value = title;
  document.getElementById("editContent").value = content;
  document.getElementById("editModal").style.display = "block";
}

document.getElementById("closeModal").onclick = function () {
  document.getElementById("editModal").style.display = "none";
};

document.getElementById("discardEdit").onclick = function () {
  document.getElementById("editModal").style.display = "none";
};

window.onclick = function (event) {
  if (event.target == document.getElementById("editModal")) {
    document.getElementById("editModal").style.display = "none";
  }
};

document.getElementById("editForm").onsubmit = function (event) {
  event.preventDefault();
  const postId = document.getElementById("editPostId").value;
  const title = document.getElementById("editTitle").value;
  const content = document.getElementById("editContent").value;

  fetch(`/posts/${postId}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: new URLSearchParams({
      title: title,
      content: content,
    }),
  }).then((response) => {
    if (response.ok) {
      location.reload();
    } else {
      alert("Error editing post");
    }
  });
};

function deletePost(postId) {
  if (confirm("Weet je zeker dat je deze post wilt verwijderen?")) {
    fetch(`/posts/${postId}`, {
      method: "DELETE",
    }).then((response) => {
      if (response.ok) {
        document.getElementById(`post-${postId}`).remove();
      } else {
        alert("Error deleting post");
      }
    });
  }
}
