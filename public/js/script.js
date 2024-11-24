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
