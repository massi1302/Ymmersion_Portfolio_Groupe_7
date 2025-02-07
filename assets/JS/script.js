const modal = document.getElementById("imageModal");
const modalImg = document.getElementById("modalImage");
const closeBtn = document.getElementsByClassName("close-modal")[0];

// Function to open modal
function openModal(img) {
    modal.style.display = "block";
    modalImg.src = img.src;
}

// Close modal when clicking (x)
closeBtn.onclick = function() {
    modal.style.display = "none";
}

// Close modal when clicking outside the image
modal.onclick = function(event) {
    if (event.target === modal) {
        modal.style.display = "none";
    }
}