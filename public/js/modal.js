function openEditModal(postId, title, content) {
  const editPostId = document.getElementById("editPostId");
  const editTitle = document.getElementById("editTitle");
  const editContent = document.getElementById("editContent");

  if (editPostId && editTitle && editContent) {
    editPostId.value = postId;
    editTitle.value = title;
    editContent.value = content;
    const editModal = new bootstrap.Modal(document.getElementById("editModal"));
    editModal.show();
  } else {
    console.error("Edit modal elements not found");
  }
}
